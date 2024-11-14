import fetch from 'node-fetch';
import { NODE_TYPES } from './config.mjs';

export class NodeDetector {
    static async detect(url) {
        const junoVersion = await this._detectJuno(url);
        if (junoVersion) {
            return { type: NODE_TYPES.JUNO, version: junoVersion };
        }

        const pathfinderVersion = await this._detectPathfinder(url);
        if (pathfinderVersion) {
            return { type: NODE_TYPES.PATHFINDER, version: pathfinderVersion };
        }

        return { type: NODE_TYPES.UNKNOWN, version: 'Unknown' };
    }

    static async _detectJuno(url) {
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
            return data.result || null;
        } catch {
            return null;
        }
    }

    static async _detectPathfinder(url) {
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
            return data.result || null;
        } catch {
            return null;
        }
    }
} 