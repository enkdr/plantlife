import React from 'react';
import './App.css';
import DataFetcher from './DataFetcher';

const App: React.FC = () => {
  return (
    <div className="App">
      <header className="App-header">
        <h1>React and TypeScript Data Fetching</h1>
      </header>
      <main>
        <DataFetcher />
      </main>
    </div>
  );
};

export default App;
