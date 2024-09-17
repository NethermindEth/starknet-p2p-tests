import { RpcProvider } from "starknet";

const base = new RpcProvider({ nodeUrl: process.argv[2] });
const syncing = new RpcProvider({ nodeUrl: process.argv[3] });

function log(message) {
    console.log(message);
    if (typeof process !== 'undefined' && process.stdout && process.stdout.isTTY) {
        process.stdout.write('');  // Force flush
    }
}

async function syncNode(targetBlockNumber = 10) {
    log(`Syncing to target block: ${targetBlockNumber}`);

    const timer = setInterval(async () => {
        try {
            const syncingBlock = await syncing.getBlockLatestAccepted();
            log(`Current syncing block: ${syncingBlock.block_number}`);
            
            if (syncingBlock.block_number >= targetBlockNumber) {
                log(`Syncing node has reached or surpassed the target block ${targetBlockNumber}.`);
                log("Sync successful. Stopping checks.");
                clearInterval(timer);
                process.exit(0);
            }
        } catch (error) {
            log(`Error during sync check: ${error.message}`);
            clearInterval(timer);
            process.exit(1);
        }
    }, 10000);

    // Set a timeout of 3 hours
    setTimeout(() => {
        log("Stopping automatic checks after 3h. Marking as failure due to timeout.");
        clearInterval(timer);
        process.exit(1);
    }, 3 * 60 * 60 * 1000);
}

// Call syncNode with the target block number (default is 1000)
syncNode().catch(error => {
    log(`Error occurred: ${error.message}`);
    process.exit(1);
});
