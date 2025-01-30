import fetch from 'node-fetch';
import { REPORTING_CONFIG } from './config.mjs';

export class TestReporter {
    constructor(sourceNode, targetNode, targetBlockNumber) {
        this.config = {
            ...REPORTING_CONFIG,
            sourceNode,
            targetNode,
            targetBlockNumber
        };
    }

    async reportProgress(blockNumber, startTime, errors = []) {
        if (!this.config.enabled) return;

        try {
            const status = this._getStatus(blockNumber, errors);

            const payload = {
                type: "updateTest",
                data: {
                    id: this.config.testId,
                    sourceNode: this.config.sourceNode.type,
                    sourceVersion: this.config.sourceNode.version,
                    targetNode: this.config.targetNode.type,
                    targetVersion: this.config.targetNode.version,
                    status: status,
                    startTime: new Date(startTime).toISOString(),
                    blocksProcessed: blockNumber,
                    totalBlocks: this.config.targetBlockNumber,
                    avgBlockTime: this._calculateAvgBlockTime(blockNumber, startTime),
                    errors
                }
            };

            if (status === "Passed" || status === "Failed") {
                payload.data.endTime = new Date().toISOString();
            }

            const response = await fetch(this.config.endpoint, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            return response.status;
        } catch (error) {
            console.warn(`Warning: Failed to report progress: ${error.message}`);
            return null;
        }
    }

    _getStatus(blockNumber, errors) {
        if (errors.length > 0) return "Failed";
        if (blockNumber >= this.config.targetBlockNumber) return "Passed";
        return "In Progress";
    }

    _calculateAvgBlockTime(blockNumber, startTime) {
        if (blockNumber <= 0) return "0.000";
        return ((Date.now() - startTime) / blockNumber / 1000).toFixed(3);
    }
} 