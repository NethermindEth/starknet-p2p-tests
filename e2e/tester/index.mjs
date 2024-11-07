import { RpcProvider } from "starknet";

if (!process.argv[2]) {
    console.error("Error: Node URL is required");
    process.exit(1);
}

const node = new RpcProvider({ nodeUrl: process.argv[2] });
const timeout = parseInt(process.argv[3], 10);
const targetBlockNumber = parseInt(process.argv[4], 10);

// Simple logging with timestamp
const log = (message) => console.log(`[${new Date().toISOString()}] ${message}`);

// Constants for configuration
const CONFIG = {
    CHECK_INTERVAL: 10_000,  // 10 seconds
    STALL_TIMEOUT: 5,        // 5 minutes
    NODE_READY_ATTEMPTS: 5,
    NODE_READY_INTERVAL: 5_000,
};

async function waitForNodeReady() {
    for (let attempt = 1; attempt <= CONFIG.NODE_READY_ATTEMPTS; attempt++) {
        try {
            const version = await node.getSpecVersion();
            log(`✓ Node ready (v${version})`);
            return true;
        } catch {
            log(`⧗ Waiting for node... (${attempt}/${CONFIG.NODE_READY_ATTEMPTS})`);
            await new Promise(r => setTimeout(r, CONFIG.NODE_READY_INTERVAL));
        }
    }
    throw new Error("Node failed to become ready");
}

async function getCurrentBlock() {
    try {
        return await node.getBlockLatestAccepted();
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
                }

                // Check completion or failure conditions
                if (currentBlock.block_number >= targetBlockNumber) {
                    log(`✓ Sync completed in ${elapsedMinutes.toFixed(1)}m`);
                    clearInterval(interval);
                    resolve();
                } else if ((Date.now() - lastBlockTime) / 1000 / 60 >= CONFIG.STALL_TIMEOUT) {
                    throw new Error(`Sync stalled for ${CONFIG.STALL_TIMEOUT}m`);
                } else if (Date.now() - startTime > timeout * 1000) {
                    throw new Error(`Sync timeout after ${timeout}s`);
                }
            } catch (error) {
                clearInterval(interval);
                reject(error);
            }
        }, CONFIG.CHECK_INTERVAL);
    });
}

// Main execution
syncNode()
    .then(() => process.exit(0))
    .catch(error => {
        log(`✗ Error: ${error.message}`);
        process.exit(1);
    });

