import { RpcProvider } from "starknet";

const node = new RpcProvider({ nodeUrl: process.argv[2] });
const timeout = parseInt(process.argv[3], 10);
const targetBlockNumber = parseInt(process.argv[4], 10);

function log(message) {
    console.log(message);
    if (typeof process !== 'undefined' && process.stdout && process.stdout.isTTY) {
        process.stdout.write('');  // Force flush
    }
}

async function waitForNodeReady(maxAttempts = 30, interval = 10000) {
    for (let attempt = 0; attempt < maxAttempts; attempt++) {
        try {
            await node.getBlockLatestAccepted();
            log("Node is ready.");
            return true;
        } catch (error) {
            log(`Waiting for node to be ready... (Attempt ${attempt + 1}/${maxAttempts})`);
            await new Promise(resolve => setTimeout(resolve, interval));
        }
    }
    log("Node failed to become ready in time.");
    return false;
}

async function syncNode() {
    log(`Syncing to target block: ${targetBlockNumber}`);

    const nodeReady = await waitForNodeReady();
    if (!nodeReady) {
        process.exit(1);
    }

    const startTime = Date.now();

    const timer = setInterval(async () => {
        try {
            const currentBlock = await node.getBlockLatestAccepted();
            log(`Current syncing block: ${currentBlock.block_number}`);
            
            if (currentBlock.block_number >= targetBlockNumber) {
                log(`Syncing node has reached or surpassed the target block ${targetBlockNumber}.`);
                log("Sync successful. Stopping checks.");
                clearInterval(timer);
                process.exit(0);
            }

            if (Date.now() - startTime > timeout * 1000) {
                log(`Timeout of ${timeout} seconds reached. Sync unsuccessful.`);
                clearInterval(timer);
                process.exit(1);
            }
        } catch (error) {
            log(`Error during sync check: ${error.message}`);
            // Instead of exiting immediately, we'll continue the loop
        }
    }, 10000);
}

syncNode().catch(error => {
    log(`Unhandled error occurred: ${error.message}`);
    process.exit(1);
});
