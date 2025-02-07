import { ArrowRight, CheckCircle2, XCircle, Clock } from 'lucide-react';
import { formatDistanceToNow, differenceInSeconds, intervalToDuration } from 'date-fns';
import type { TestRun } from '../types';

interface TestCardProps {
  test: TestRun;
  onClick: () => void;
}

const statusIcons = {
  'Passed': <CheckCircle2 className="w-5 h-5 text-green-500" />,
  'Failed': <XCircle className="w-5 h-5 text-red-500" />,
  'In Progress': <Clock className="w-5 h-5 text-blue-500 animate-pulse" />
};

const statusColors = {
  'Passed': 'bg-green-50 border-green-200',
  'Failed': 'bg-red-50 border-red-200',
  'In Progress': 'bg-blue-50 border-blue-200'
};

export default function TestCard({ test, onClick }: TestCardProps) {
  const calculateSyncSpeed = () => {
    const elapsedSeconds = differenceInSeconds(new Date(), new Date(test.startTime));
    if (elapsedSeconds === 0) return 0;
    return Math.round(test.blocksProcessed / elapsedSeconds);
  };

  const formatSyncSummary = () => {
    if (!test.endTime) {
      return `${test.totalBlocks.toLocaleString()} blocks in N/A`;
    }

    const endTimeDate = new Date(test.endTime);
    const duration = intervalToDuration({
      start: new Date(test.startTime),
      end: endTimeDate
    });

    const formatTime = () => {
      const parts = [];
      if (duration.days) parts.push(`${duration.days}d`);
      if (duration.hours) parts.push(`${duration.hours}h`);
      if (duration.minutes) parts.push(`${duration.minutes}m`);
      if (duration.seconds) parts.push(`${duration.seconds}s`);
      if (parts.length === 0) parts.push('<1s');
      return parts.join(' ');
    };

    return `${test.totalBlocks.toLocaleString()} blocks in ${formatTime()}`;
  };

  return (
    <div 
      onClick={onClick}
      className={`p-4 rounded-lg border ${statusColors[test.status]} cursor-pointer 
        transition-all duration-200 hover:shadow-md`}
    >
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <span className="font-semibold text-gray-700">{test.targetNode}</span>
          <span className="text-sm text-gray-500">{test.targetVersion}</span>
          <ArrowRight className="w-4 h-4 text-gray-400 rotate-180" />
          <span className="font-semibold text-gray-700">{test.sourceNode}</span>
          <span className="text-sm text-gray-500">{test.sourceVersion}</span>
        </div>
        <div className="flex items-center space-x-2">
          {statusIcons[test.status]}
          {(test.status === 'Passed' || test.status === 'Failed') ? (
            <span className="text-sm font-medium text-gray-600">{formatSyncSummary()}</span>
          ) : (
            <span className="text-sm font-medium text-gray-600">
              {formatDistanceToNow(new Date(test.startTime), { addSuffix: true })}
            </span>
          )}
        </div>
      </div>
      
      {test.status === 'In Progress' && (
        <div className="mt-3">
          <div className="w-full bg-gray-200 rounded-full h-2">
            <div 
              className="bg-blue-500 h-2 rounded-full transition-all duration-500"
              style={{ width: `${(test.blocksProcessed / test.totalBlocks) * 100}%` }}
            />
          </div>
          <div className="mt-1 flex justify-between items-center text-sm text-gray-500">
            <span>{test.blocksProcessed.toLocaleString()} / {test.totalBlocks.toLocaleString()} blocks</span>
            <span>{calculateSyncSpeed()} blocks/sec</span>
          </div>
        </div>
      )}
    </div>
  );
}