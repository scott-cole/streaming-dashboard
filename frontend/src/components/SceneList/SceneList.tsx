import {useEffect, useState} from 'react';

function SceneList() {
  const [scenes, setScenes] = useState<string[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchScenes = async () => {
     try {
      const response = await fetch('/api/scenes')

      if (!response.ok) {
        throw new Error('Failed to fetch scenes')
      }

      const data = await response.json();
      setScenes(data.scenes)
     } 
     catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown Error')
     }
     finally {
      setLoading(false)
     }
    };

    fetchScenes()
  }, []);

  if (loading) return <p>Loading...</p>
  if (error) return <p>Error: {error}</p>

  return (
    <ul>
      {scenes.map((scene) => (
        <li key={scene}>{scene}</li>
      ))}
    </ul>
  );
}

export default SceneList;