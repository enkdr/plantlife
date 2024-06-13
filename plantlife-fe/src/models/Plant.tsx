import { useState, useEffect } from 'react';

interface Plant {
    id: number;
    name: string;
    description: string;
}

async function fetchPlants(): Promise<Plant[]> {
    const res = await fetch('/plants');
    const data = await res.json();
    return data;
}

function PlantList() {
    const [loading, setLoading] = useState<boolean>(true); // State for loading indicator
    const [plants, setPlants] = useState<Plant[]>([]);
    useEffect(() => {
        fetchPlants().then((data) => {
            setPlants(data)
            setLoading(false);
        });
    }, []);

    console.log(plants);

    if (loading) {
        return <div>Loading...</div>; // Render loading indicator
    }


    return (
        <div>
            {plants.map((plant) => (
                <div key={plant.id}>
                    <h2>{plant.name}</h2>
                    <p>{plant.description}</p>
                </div>
            ))}
        </div>

    );
}

export default PlantList;