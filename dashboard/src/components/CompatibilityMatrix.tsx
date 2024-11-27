import React from 'react';
import { CheckCircle2, XCircle, Ban } from 'lucide-react';
import type { TestRun } from '../types';

interface CompatibilityMatrixProps {
  tests: TestRun[];
}

export default function CompatibilityMatrix({ tests }: CompatibilityMatrixProps) {
  const nodes = ['Pathfinder', 'Juno', 'Madara', 'Papyrus'];
  const availableNodes = ['Pathfinder', 'Juno'];
  
  const getLatestTestResult = (target: string, source: string) => {
    if (!availableNodes.includes(target) || !availableNodes.includes(source)) {
      return undefined;
    }

    const latestTest = tests.find(test => 
      test.sourceNode === source && 
      test.targetNode === target &&
      test.status !== 'In Progress'
    );
    return latestTest?.status === 'Passed';
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
      <h2 className="text-lg font-semibold text-gray-900 mb-4">P2P Sync Compatibility</h2>
      <div className="relative overflow-x-auto">
        <table className="w-full">
          <thead>
            <tr>
              <th className="px-4 py-2 w-32">
                <div className="text-sm text-gray-500">
                  To ↓ / From →
                </div>
              </th>
              {nodes.map(node => (
                <th key={node} className="px-4 py-2 font-medium text-gray-700 text-center">
                  {node}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {nodes.map(targetNode => (
              <tr key={targetNode}>
                <td className="px-4 py-2 font-medium text-gray-700">{targetNode}</td>
                {nodes.map(sourceNode => {
                  const result = getLatestTestResult(targetNode, sourceNode);
                  return (
                    <td key={`${targetNode}-${sourceNode}`} className="px-4 py-2">
                      <div className="flex justify-center">
                        {result === undefined ? (
                          <Ban className="w-6 h-6 text-gray-300" />
                        ) : result ? (
                          <CheckCircle2 className="w-6 h-6 text-green-500" />
                        ) : (
                          <XCircle className="w-6 h-6 text-red-500" />
                        )}
                      </div>
                    </td>
                  );
                })}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
} 