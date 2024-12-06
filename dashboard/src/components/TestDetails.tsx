import React from 'react';
import { X, CheckCircle2, XCircle, Clock, ArrowLeft } from 'lucide-react';
import { formatDistanceToNow, format } from 'date-fns';
import type { TestRun } from '../types';

interface TestDetailsProps {
  test: TestRun;
  onClose: () => void;
}

const statusColors = {
  'Passed': 'text-green-500',
  'Failed': 'text-red-500',
  'In Progress': 'text-blue-500'
};

export default function TestDetails({ test, onClose }: TestDetailsProps) {
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div className="bg-white rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div className="p-6">
          <div className="flex justify-between items-start">
            <h2 className="text-2xl font-bold text-gray-900">Test Details</h2>
            <button
              onClick={onClose}
              className="p-1 rounded-full hover:bg-gray-100 transition-colors"
            >
              <X className="w-6 h-6 text-gray-500" />
            </button>
          </div>

          <div className="mt-6 space-y-6">
            <div className="flex items-center justify-center space-x-4">
              <div className="text-center">
                <div className="font-semibold text-lg">{test.targetNode}</div>
                <div className="text-sm text-gray-500">{test.targetVersion}</div>
              </div>
              <ArrowLeft className="w-6 h-6 text-gray-400" />
              <div className="text-center">
                <div className="font-semibold text-lg">{test.sourceNode}</div>
                <div className="text-sm text-gray-500">{test.sourceVersion}</div>
              </div>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div className="bg-gray-50 p-4 rounded-lg">
                <div className="text-sm text-gray-500">Status</div>
                <div className={`font-semibold ${statusColors[test.status]} flex items-center space-x-2`}>
                  {test.status === 'Passed' && <CheckCircle2 className="w-5 h-5" />}
                  {test.status === 'Failed' && <XCircle className="w-5 h-5" />}
                  {test.status === 'In Progress' && <Clock className="w-5 h-5 animate-pulse" />}
                  <span>{test.status}</span>
                </div>
              </div>

              <div className="bg-gray-50 p-4 rounded-lg">
                <div className="text-sm text-gray-500">Started</div>
                <div className="font-semibold">
                  {format(new Date(test.startTime), 'PPp')}
                </div>
              </div>

              <div className="bg-gray-50 p-4 rounded-lg">
                <div className="text-sm text-gray-500">Blocks Processed</div>
                <div className="font-semibold">
                  {test.blocksProcessed.toLocaleString()} / {test.totalBlocks.toLocaleString()}
                </div>
              </div>

              <div className="bg-gray-50 p-4 rounded-lg">
                <div className="text-sm text-gray-500">Average Block Time</div>
                <div className="font-semibold">{test.avgBlockTime}s</div>
              </div>
            </div>

            {test.status === 'In Progress' && (
              <div>
                <div className="w-full bg-gray-200 rounded-full h-2">
                  <div 
                    className="bg-blue-500 h-2 rounded-full transition-all duration-500"
                    style={{ width: `${(test.blocksProcessed / test.totalBlocks) * 100}%` }}
                  />
                </div>
                <div className="mt-2 text-center text-sm text-gray-500">
                  {((test.blocksProcessed / test.totalBlocks) * 100).toFixed(1)}% Complete
                </div>
              </div>
            )}

            {test.errors.length > 0 && (
              <div>
                <h3 className="font-semibold text-red-500 mb-2">Errors</h3>
                <div className="bg-red-50 rounded-lg p-4">
                  <ul className="list-disc list-inside space-y-1">
                    {test.errors.map((error, index) => (
                      <li key={index} className="text-red-700">{error}</li>
                    ))}
                  </ul>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}