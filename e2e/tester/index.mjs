import { RpcProvider } from "starknet";

const node = new RpcProvider({ nodeUrl: process.argv[2] });
const timeout = parseInt(process.argv[3], 10);
const targetBlockNumber = parseInt(process.argv[4], 10);

function log(message) {
    process.stdout.write(`[${new Date().toISOString()}] ${message}\n`);
}

async function waitForNodeReady(maxAttempts = 30, interval = 10000) {
    for (let attempt = 0; attempt < maxAttempts; attempt++) {
        try {
            await node.getBlockLatestAccepted();
            log("✓ Node is ready");
            return true;
        } catch (error) {
            log(`⧗ Waiting for node... (${attempt + 1}/${maxAttempts})`);
            await new Promise(resolve => setTimeout(resolve, interval));
        }
    }
    log("✗ Node failed to become ready in time");
    return false;
}

async function syncNode() {
    log(`➜ Starting sync to target block ${targetBlockNumber}`);

    if (!await waitForNodeReady()) {
        process.exit(1);
    }

    const startTime = Date.now();
    let lastBlockTime = Date.now();
    let lastBlockNumber;

    try {
        const startBlock = await node.getBlockLatestAccepted();
        lastBlockNumber = startBlock.block_number;
        log(`➜ Initial block: ${lastBlockNumber}`);
    } catch (error) {
        log(`✗ Failed to get initial block: ${error.message}`);
        process.exit(1);
    }

    const checkSync = setInterval(async () => {
        try {
            const currentBlock = await node.getBlockLatestAccepted();
            const elapsedMinutes = (Date.now() - startTime) / 1000 / 60;
            const blocksSynced = currentBlock.block_number - lastBlockNumber;
            
            // Update progress
            if (currentBlock.block_number > lastBlockNumber) {
                const speed = blocksSynced / elapsedMinutes;
                log(`↑ Block ${currentBlock.block_number} | +${blocksSynced} blocks | ${elapsedMinutes.toFixed(1)}m | ${speed.toFixed(1)} blocks/min`);
                lastBlockNumber = currentBlock.block_number;
                lastBlockTime = Date.now();
            }

            // Check completion
            if (currentBlock.block_number >= targetBlockNumber) {
                const totalMinutes = (Date.now() - startTime) / 1000 / 60;
                log(`\n✓ Sync completed in ${totalMinutes.toFixed(1)} minutes`);
                clearInterval(checkSync);
                process.exit(0);
            }

            // Check timeouts
            const noProgressTime = (Date.now() - lastBlockTime) / 1000 / 60;
            if (noProgressTime >= 5) {
                log(`\n✗ Sync stalled - No progress for ${noProgressTime.toFixed(1)} minutes`);
                clearInterval(checkSync);
                process.exit(1);
            }

            if (Date.now() - startTime > timeout * 1000) {
                log(`\n✗ Sync timeout after ${timeout} seconds`);
                clearInterval(checkSync);
                process.exit(1);
            }
        } catch (error) {
            log(`! Error checking sync: ${error.message}`);
        }
    }, 10000);
}

syncNode().catch(error => {
    log(`✗ Fatal error: ${error.message}`);
    process.exit(1);
});
