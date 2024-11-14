import { RpcProvider } from "starknet";
import fetch from 'node-fetch';

if (!process.argv[2] || !process.argv[3]) {
    console.error("Error: Source and Target Node URLs are required");
    process.exit(1);
}

const sourceNode = new RpcProvider({ nodeUrl: process.argv[2] });
const targetNode = new RpcProvider({ nodeUrl: process.argv[3] });
const timeout = parseInt(process.argv[4], 10);
const targetBlockNumber = parseInt(process.argv[5], 10);
const sourceType = process.argv[6] || 'Unknown';
const targetType = process.argv[7] || 'Unknown';

// Simple logging with timestamp
const log = (message) => console.log(`[${new Date().toISOString()}] ${message}`);

// Constants for configuration
const CONFIG = {
    CHECK_INTERVAL: 10_000,  // 10 seconds
    STALL_TIMEOUT: 1,        // 1 minute
    NODE_READY_ATTEMPTS: 5,
    NODE_READY_INTERVAL: 5_000,
};

const REPORTING_CONFIG = {
    enabled: true,
    endpoint: 'https://starknet-p2p-testing-dashboard.onrender.com/update',
    testId: Date.now().toString(),
    sourceNode: process.argv[2] || 'Juno',
    sourceVersion: process.argv[6] || 'Unknown',
    targetNode: process.argv[3] || 'Juno',
    targetVersion: process.argv[8] || 'Unknown',
};

async function detectNodeTypeAndVersion(provider, url) {
    // Try Juno version endpoint first
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                method: "juno_version",
                jsonrpc: "2.0",
                id: 0
            })
        });
        
        const data = await response.json();
        if (data.result) {
            return { type: 'Juno', version: data.result };
        }
    } catch (error) {
        // Silently fail and try Pathfinder
    }

    // Try Pathfinder version endpoint
    try {
        const response = await fetch(`${url}/rpc/pathfinder/v0_1`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                method: "pathfinder_version",
                jsonrpc: "2.0",
                id: 0
            })
        });
        
        const data = await response.json();
        if (data.result) {
            return { type: 'Pathfinder', version: data.result };
        }
    } catch (error) {
        // Silently fail
    }

    // If both fail, return unknown
    return { type: 'Unknown', version: 'Unknown' };
}

async function waitForNodeReady() {
    for (let attempt = 1; attempt <= CONFIG.NODE_READY_ATTEMPTS; attempt++) {
        try {
            // Detect node types and versions
            const sourceInfo = await detectNodeTypeAndVersion(sourceNode, process.argv[2]);
            const targetInfo = await detectNodeTypeAndVersion(targetNode, process.argv[3]);
            
            REPORTING_CONFIG.sourceNode = sourceInfo.type;
            REPORTING_CONFIG.targetNode = targetInfo.type;
            REPORTING_CONFIG.sourceVersion = sourceInfo.version;
            REPORTING_CONFIG.targetVersion = targetInfo.version;
            
            log(`✓ Source node ready (${REPORTING_CONFIG.sourceNode} ${REPORTING_CONFIG.sourceVersion})`);
            log(`✓ Target node ready (${REPORTING_CONFIG.targetNode} ${REPORTING_CONFIG.targetVersion})`);
            return true;
        } catch (error) {
            log(`⧗ Waiting for nodes... (${attempt}/${CONFIG.NODE_READY_ATTEMPTS})`);
            await new Promise(r => setTimeout(r, CONFIG.NODE_READY_INTERVAL));
        }
    }
    throw new Error("One or both nodes failed to become ready");
}

async function getCurrentBlock() {
    try {
        return await targetNode.getBlockLatestAccepted();
    } catch (error) {
        if (error.message.includes("There are no blocks")) {
            log("➜ Starting from genesis");
            return { block_number: 0 };
        }
        throw error;
    }
}

async function syncNode() {
    const startTime = Date.now();
    let lastBlockTime = startTime;
    let lastBlockNumber = (await getCurrentBlock()).block_number;
    
    log(`➜ Starting sync to block ${targetBlockNumber} from ${lastBlockNumber}`);
    await waitForNodeReady();

    return new Promise((resolve, reject) => {
        const interval = setInterval(async () => {
            try {
                const currentBlock = await getCurrentBlock();
                const elapsedMinutes = (Date.now() - startTime) / 1000 / 60;
                const blocksSynced = currentBlock.block_number - lastBlockNumber;

                // Progress update with completion percentage
                if (currentBlock.block_number > lastBlockNumber) {
                    const speed = blocksSynced / elapsedMinutes;
                    const progress = ((currentBlock.block_number / targetBlockNumber) * 100).toFixed(1);
                    log(`↑ Block ${currentBlock.block_number} | ${progress}% | ${speed.toFixed(1)} blocks/min`);
                    lastBlockNumber = currentBlock.block_number;
                    lastBlockTime = Date.now();
                    
                    // Report progress
                    await reportProgress(currentBlock.block_number, startTime);
                }

                // Check completion or failure conditions
                if (currentBlock.block_number >= targetBlockNumber) {
                    log(`✓ Sync completed in ${elapsedMinutes.toFixed(1)}m`);
                    await reportProgress(currentBlock.block_number, startTime);
                    clearInterval(interval);
                    resolve();
                } else if ((Date.now() - lastBlockTime) / 1000 / 60 >= CONFIG.STALL_TIMEOUT) {
                    throw new Error(`Sync stalled for ${CONFIG.STALL_TIMEOUT}m`);
                } else if (Date.now() - startTime > timeout * 1000) {
                    throw new Error(`Sync timeout after ${timeout}s`);
                }
            } catch (error) {
                clearInterval(interval);
                await reportProgress(lastBlockNumber, startTime, [error.message]);
                reject(error);
            }
        }, CONFIG.CHECK_INTERVAL);
    });
}

// Modify reportProgress to use REPORTING_CONFIG.targetVersion directly
async function reportProgress(blockNumber, startTime, errors = []) {
    if (!REPORTING_CONFIG.enabled) return;

    try {
        const payload = {
            type: "updateTest",
            data: {
                id: REPORTING_CONFIG.testId,
                sourceNode: REPORTING_CONFIG.sourceNode,
                sourceVersion: REPORTING_CONFIG.sourceVersion,
                targetNode: REPORTING_CONFIG.targetNode,
                targetVersion: REPORTING_CONFIG.targetVersion, // Use config directly
                status: errors.length > 0 ? "Failed" : 
                       blockNumber >= targetBlockNumber ? "Passed" : 
                       "In Progress",
                startTime: new Date(startTime).toISOString(),
                blocksProcessed: blockNumber,
                totalBlocks: targetBlockNumber,
                avgBlockTime: blockNumber > 0 ? 
                    ((Date.now() - startTime) / blockNumber / 1000).toFixed(3) : 
                    "0.000",
                errors: errors
            }
        };

        const response = await fetch(REPORTING_CONFIG.endpoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        log(`Progress report response: ${response.status}`);
    } catch (error) {
        log(`Warning: Failed to report progress: ${error.message}`);
    }
}

// Main execution
syncNode()
    .then(() => process.exit(0))
    .catch(error => {
        log(`✗ Error: ${error.message}`);
        process.exit(1);
    });

