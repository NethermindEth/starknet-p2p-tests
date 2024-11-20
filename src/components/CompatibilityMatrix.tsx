import React from 'react';
import { CheckCircle2, XCircle } from 'lucide-react';
import type { TestRun } from '../types';

interface CompatibilityMatrixProps {
  tests: TestRun[];
}

export default function CompatibilityMatrix({ tests }: CompatibilityMatrixProps) {
  const nodes = ['Juno', 'Pathfinder'];
  
  const getLatestTestResult = (source: string, target: string) => {
    const latestTest = tests.find(test => 
      test.sourceNode === source && 
      test.targetNode === target &&
      test.status !== 'In Progress'
    );
    return latestTest?.status === 'Passed';
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
      <h2 className="text-lg font-semibold text-gray-900 mb-4">Compatibility Matrix</h2>
      <div className="relative overflow-x-auto">
        <table className="w-full text-left">
          <thead>
            <tr>
              <th className="px-4 py-2"></th>
              {nodes.map(node => (
                <th key={node} className="px-4 py-2 font-medium text-gray-700">
                  {node}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {nodes.map(sourceNode => (
              <tr key={sourceNode}>
                <td className="px-4 py-2 font-medium text-gray-700">{sourceNode}</td>
                {nodes.map(targetNode => (
                  <td key={`${sourceNode}-${targetNode}`} className="px-4 py-2">
                    {getLatestTestResult(sourceNode, targetNode) ? (
                      <CheckCircle2 className="w-6 h-6 text-green-500" />
                    ) : (
                      <XCircle className="w-6 h-6 text-red-500" />
                    )}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
} 