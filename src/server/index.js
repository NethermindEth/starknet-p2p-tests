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