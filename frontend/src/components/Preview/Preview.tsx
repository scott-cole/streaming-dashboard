import { useEffect, useRef, useState } from "react";
import styles from "./Preview.module.scss";

interface PreviewProps {
  sceneName: string;
}

function Preview({ sceneName }: PreviewProps) {
  const [image, setImage] = useState<string | null>(null);
  const intervalRef = useRef<ReturnType<typeof setInterval> | null>(null);

  useEffect(() => {
    const fetchPreview = async () => {
      try {
        const res = await fetch(`/api/preview?scene=${encodeURIComponent(sceneName)}&width=640&height=360`);
        if (!res.ok) throw new Error("Preview failed");
        const data = await res.json();
        setImage(`data:image/png;base64,${data.image}`);
      } catch {
        setImage(null);
      }
    };

    fetchPreview();
    intervalRef.current = setInterval(fetchPreview, 2000);

    return () => {
      if (intervalRef.current) clearInterval(intervalRef.current);
    };
  }, [sceneName]);

  return (
    <div className={styles.container}>
      {image ? (
        <img src={image} alt={`Preview: ${sceneName}`} className={styles.image} />
      ) : (
        <div className={styles.placeholder}>
          <span className={styles.placeholderIcon}>◉</span>
          <span className={styles.placeholderText}>Preview Unavailable</span>
        </div>
      )}
      <div className={styles.label}>{sceneName}</div>
    </div>
  );
}

export default Preview;
