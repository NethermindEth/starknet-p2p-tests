import { RpcProvider } from "starknet";
import { SYNC_CONFIG } from './config.mjs';
import { TestReporter } from './reporter.mjs';
import { NodeDetector } from './nodeDetector.mjs';

if (!process.argv[2] || !process.argv[3]) {
    console.error("Error: Source and Target Node URLs are required");
    process.exit(1);
}

const sourceNode = new RpcProvider({ nodeUrl: process.argv[2] });
const targetNode = new RpcProvider({ nodeUrl: process.argv[3] });
const timeout = parseInt(process.argv[4], 10);
const targetBlockNumber = parseInt(process.argv[5], 10);

// Simple logging with timestamp
const log = (message) => console.log(`[${new Date().toISOString()}] ${message}`);

async function waitForNodeReady() {
    for (let attempt = 1; attempt <= SYNC_CONFIG.NODE_READY_ATTEMPTS; attempt++) {
        try {
            const sourceInfo = await NodeDetector.detect(process.argv[2]);
            const targetInfo = await NodeDetector.detect(process.argv[3]);
            
            log(`✓ Source node ready (${sourceInfo.type} ${sourceInfo.version})`);
            log(`✓ Target node ready (${targetInfo.type} ${targetInfo.version})`);
            
            return { sourceInfo, targetInfo };
        } catch (error) {
            log(`⧗ Waiting for nodes... (${attempt}/${SYNC_CONFIG.NODE_READY_ATTEMPTS})`);
            await new Promise(r => setTimeout(r, SYNC_CONFIG.NODE_READY_INTERVAL));
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
    const { sourceInfo, targetInfo } = await waitForNodeReady();
    
    const reporter = new TestReporter(sourceInfo, targetInfo, targetBlockNumber);

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
                    await reporter.reportProgress(currentBlock.block_number, startTime);
                }

                // Check completion or failure conditions
                if (currentBlock.block_number >= targetBlockNumber) {
                    log(`✓ Sync completed in ${elapsedMinutes.toFixed(1)}m`);
                    await reporter.reportProgress(currentBlock.block_number, startTime);
                    clearInterval(interval);
                    resolve();
                } else if ((Date.now() - lastBlockTime) / 1000 / 60 >= SYNC_CONFIG.STALL_TIMEOUT) {
                    throw new Error(`Sync stalled for ${SYNC_CONFIG.STALL_TIMEOUT}m`);
                } else if (Date.now() - startTime > timeout * 1000) {
                    throw new Error(`Sync timeout after ${timeout}s`);
                }
            } catch (error) {
                clearInterval(interval);
                await reporter.reportProgress(lastBlockNumber, startTime, [error.message]);
                reject(error);
            }
        }, SYNC_CONFIG.CHECK_INTERVAL);
    });
}

// Main execution
syncNode()
    .then(() => process.exit(0))
    .catch(error => {
        log(`✗ Error: ${error.message}`);
        process.exit(1);
    });

