import React, { useEffect, useState } from 'react';

interface Plant {
  name: string;
}

const App: React.FC = () => {
  const [plants, setPlants] = useState<Plant[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPlants = async () => {
      try {
        const response = await fetch('http://localhost:3000/');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data: Plant[] = await response.json();
        setPlants(data);
        setLoading(false);
      } catch (error) {
        setError((error as Error).message);
        setLoading(false);
      }
    };

    fetchPlants();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <h1>Plants</h1>
      <ul>
        {plants.map((plant, index) => (
          <li key={index}>{plant.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default App;
