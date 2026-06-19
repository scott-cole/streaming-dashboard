import { useEffect, useState } from "react";
import Button from "../Button/Button";

function SceneList() {
  const [scenes, setScenes] = useState<string[]>([]);
  const [currentScene, setCurrentScene] = useState<string>("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchScenes = async () => {
      try {
        const res = await fetch("/api/scenes");
        if (!res.ok) {
          throw new Error("Failed to fetch scenes");
        }

        const currentRes = await fetch("/api/current-scene");
        if (!currentRes.ok) {
          throw new Error("Failed to fetch current scene");
        }

        const data = await res.json();
        const currentData = await currentRes.json();

        setScenes(data.scenes);
        setCurrentScene(currentData.scene);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Unknown Error");
      } finally {
        setLoading(false);
      }
    };

    fetchScenes();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  const switchScene = async (name: string) => {
    await fetch("/api/switch", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name }),
    });
    setCurrentScene(name);
  };

  return (
    <>
      <ul>
        {scenes.map((scene) => (
          <li key={scene}>
            <Button
              variant={scene === currentScene ? "primary" : "secondary"}
              onClick={() => switchScene(scene)}
            >
              {scene}
            </Button>
          </li>
        ))}
      </ul>
    </>
  );
}

export default SceneList;
