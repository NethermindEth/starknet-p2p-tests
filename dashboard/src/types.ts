export interface TestRun {
  id: string;
  sourceNode: string;
  sourceVersion: string;
  targetNode: string;
  targetVersion: string;
  status: 'In Progress' | 'Passed' | 'Failed';
  startTime: string;
  endTime?: string;
  blocksProcessed: number;
  totalBlocks: number;
  avgBlockTime: string;
  errors: string[];
}