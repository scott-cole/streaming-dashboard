import { useEffect, useState } from "react";
import styles from "./SceneList.module.scss";

interface SceneListProps {
  onSceneChange: (scene: string) => void;
}

function SceneList({ onSceneChange }: SceneListProps) {
  const [scenes, setScenes] = useState<string[]>([]);
  const [currentScene, setCurrentScene] = useState<string>("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchScenes = async () => {
      try {
        const res = await fetch("/api/scenes");
        if (!res.ok) throw new Error("Failed to fetch scenes");

        const currentRes = await fetch("/api/current-scene");
        if (!currentRes.ok) throw new Error("Failed to fetch current scene");

        const data = await res.json();
        const currentData = await currentRes.json();

        setScenes(data.scenes);
        setCurrentScene(currentData.scene);
        onSceneChange(currentData.scene);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Unknown Error");
      } finally {
        setLoading(false);
      }
    };

    fetchScenes();
  }, [onSceneChange]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  const switchScene = async (name: string) => {
    await fetch("/api/switch", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name }),
    });
    setCurrentScene(name);
    onSceneChange(name);
  };

  return (
    <div>
      <div style={{
        fontSize: "var(--font-xs)",
        fontWeight: 600,
        color: "var(--text-muted)",
        textTransform: "uppercase",
        letterSpacing: "0.05em",
        marginBottom: 12,
      }}>
        Scenes
      </div>
      <div className={styles.list}>
        {scenes.map((scene) => (
          <button
            key={scene}
            className={`${styles.scene} ${scene === currentScene ? styles.active : ""}`}
            onClick={() => switchScene(scene)}
          >
            <span className={`${styles.indicator} ${scene === currentScene ? styles.indicatorActive : ""}`} />
            {scene}
          </button>
        ))}
      </div>
    </div>
  );
}

export default SceneList;
