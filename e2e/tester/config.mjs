export const SYNC_CONFIG = {
    CHECK_INTERVAL: 10_000,  // 10 seconds
    STALL_TIMEOUT: 5,        // 1 minute
    NODE_READY_ATTEMPTS: 5,
    NODE_READY_INTERVAL: 5_000,
};

export const REPORTING_CONFIG = {
    enabled: true,
    endpoint: 'https://starknet-p2p-testing-dashboard.onrender.com/update',
    testId: Date.now().toString(),
};

export const NODE_TYPES = {
    JUNO: 'Juno',
    PATHFINDER: 'Pathfinder',
    UNKNOWN: 'Unknown'
}; 