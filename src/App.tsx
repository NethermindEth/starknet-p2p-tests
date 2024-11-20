import React, { useEffect, useState } from 'react';
import { Activity } from 'lucide-react';
import type { TestRun } from './types';
import TestCard from './components/TestCard';
import TestDetails from './components/TestDetails';
import CompatibilityMatrix from './components/CompatibilityMatrix';

function App() {
  const [tests, setTests] = useState<TestRun[]>([]);
  const [selectedTest, setSelectedTest] = useState<TestRun | null>(null);

  useEffect(() => {
    const eventSource = new EventSource('/events');

    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data);
      
      if (data.type === 'initial') {
        setTests(data.data);
      } else if (data.type === 'newTest') {
        setTests(prev => [data.data, ...prev]);
      } else if (data.type === 'updateTest') {
        setTests(prev => prev.map(test => 
          test.id === data.data.id ? data.data : test
        ));
        if (selectedTest?.id === data.data.id) {
          setSelectedTest(data.data);
        }
      }
    };

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error);
      eventSource.close();
    };

    return () => eventSource.close();
  }, [selectedTest]);

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex items-center space-x-3">
            <Activity className="w-8 h-8 text-blue-500" />
            <h1 className="text-2xl font-bold text-gray-900">
              Starknet P2P Testing Dashboard
            </h1>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <CompatibilityMatrix tests={tests} />
        
        <div className="space-y-4">
          {tests.map(test => (
            <TestCard
              key={test.id}
              test={test}
              onClick={() => setSelectedTest(test)}
            />
          ))}
        </div>
      </main>

      {selectedTest && (
        <TestDetails
          test={selectedTest}
          onClose={() => setSelectedTest(null)}
        />
      )}
    </div>
  );
}

export default App;