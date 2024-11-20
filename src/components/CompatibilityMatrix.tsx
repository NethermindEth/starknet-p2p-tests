import React from 'react';
import { CheckCircle2, XCircle } from 'lucide-react';
import type { TestRun } from '../types';

interface CompatibilityMatrixProps {
  tests: TestRun[];
}

export default function CompatibilityMatrix({ tests }: CompatibilityMatrixProps) {
  const nodes = ['Pathfinder', 'Juno'];
  
  const getLatestTestResult = (target: string, source: string) => {
    const latestTest = tests.find(test => 
      test.sourceNode === source && 
      test.targetNode === target &&
      test.status !== 'In Progress'
    );
    return {
      passed: latestTest?.status === 'Passed',
      sourceVersion: latestTest?.sourceVersion,
      targetVersion: latestTest?.targetVersion
    };
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
      <h2 className="text-lg font-semibold text-gray-900 mb-4">Compatibility Matrix</h2>
      <div className="relative overflow-x-auto">
        <table className="w-full text-left">
          <thead>
            <tr>
              <th className="px-4 py-2">Syncing to ↓</th>
              {nodes.map(node => (
                <th key={node} className="px-4 py-2 font-medium text-gray-700">
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
                      <div className="flex flex-col items-center">
                        {result.passed ? (
                          <CheckCircle2 className="w-6 h-6 text-green-500" />
                        ) : (
                          <XCircle className="w-6 h-6 text-red-500" />
                        )}
                        {(result.sourceVersion || result.targetVersion) && (
                          <div className="text-xs text-gray-500 mt-1">
                            {result.targetVersion} ← {result.sourceVersion}
                          </div>
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