import express from 'express';
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const app = express();
const PORT = 3322;

const clients = new Set();
const testRuns = new Map();

// Serve static files from the dist directory
app.use(express.static(join(__dirname, '../../dist')));

app.use(express.json());

// Enable CORS
app.use((req, res, next) => {
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
  next();
});

// SSE endpoint
app.get('/events', (req, res) => {
  res.setHeader('Content-Type', 'text/event-stream');
  res.setHeader('Cache-Control', 'no-cache');
  res.setHeader('Connection', 'keep-alive');

  // Send initial data
  const initialData = Array.from(testRuns.values());
  res.write(`data: ${JSON.stringify({ type: 'initial', data: initialData })}\n\n`);

  clients.add(res);

  req.on('close', () => clients.delete(res));
});

// Serve index.html for all routes
app.get('*', (req, res) => {
  res.sendFile(join(__dirname, '../../dist/index.html'));
});

function broadcastUpdate(data) {
  clients.forEach(client => {
    client.write(`data: ${JSON.stringify(data)}\n\n`);
  });
}

// Simulate new test runs
function generateTestRun() {
  const nodes = ['Juno', 'Pathfinder'];
  const versions = {
    Juno: ['v0.8.0', 'v0.7.9', 'v0.7.8'],
    Pathfinder: ['v0.9.1', 'v0.9.0', 'v0.8.9']
  };
  
  const sourceNode = nodes[Math.floor(Math.random() * nodes.length)];
  const targetNode = nodes.find(n => n !== sourceNode);
  
  const sourceVersion = versions[sourceNode][Math.floor(Math.random() * versions[sourceNode].length)];
  const targetVersion = versions[targetNode][Math.floor(Math.random() * versions[targetNode].length)];
  
  const status = Math.random() > 0.3 ? 'In Progress' : Math.random() > 0.5 ? 'Passed' : 'Failed';
  
  const testRun = {
    id: Date.now().toString(),
    sourceNode,
    sourceVersion,
    targetNode,
    targetVersion,
    status,
    startTime: new Date().toISOString(),
    blocksProcessed: status === 'In Progress' ? Math.floor(Math.random() * 1000) : Math.floor(Math.random() * 10000),
    totalBlocks: 10000,
    avgBlockTime: (Math.random() * 0.5 + 0.1).toFixed(3),
    errors: status === 'Failed' ? ['Block hash mismatch at 1337', 'State root mismatch at 1338'] : []
  };

  testRuns.set(testRun.id, testRun);
  broadcastUpdate({ type: 'newTest', data: testRun });
  
  // Update progress if in progress
  if (status === 'In Progress') {
    const interval = setInterval(() => {
      const test = testRuns.get(testRun.id);
      if (test && test.status === 'In Progress') {
        test.blocksProcessed += Math.floor(Math.random() * 100);
        if (test.blocksProcessed >= test.totalBlocks) {
          test.status = Math.random() > 0.8 ? 'Failed' : 'Passed';
          test.endTime = new Date().toISOString();
          clearInterval(interval);
        }
        testRuns.set(test.id, test);
        broadcastUpdate({ type: 'updateTest', data: test });
      }
    }, 2000);
  }
}

// Generate new test every 10-30 seconds
//setInterval(generateTestRun, Math.random() * 20000 + 10000);

// Initial test runs
for (let i = 0; i < 5; i++) {
  //generateTestRun();
}

// Add this new endpoint after the existing endpoints
app.post('/update', (req, res) => {
  const update = req.body;
  
  if (update.type === 'newTest') {
    testRuns.set(update.data.id, update.data);
  } else if (update.type === 'updateTest') {
    testRuns.set(update.data.id, update.data);
  }
  
  broadcastUpdate(update);
  res.sendStatus(200);
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});