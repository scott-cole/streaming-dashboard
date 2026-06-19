import { useState } from "react";
import styles from "./Controls.module.scss";

type OutputAction = "stream" | "record";

function Controls() {
  const [streaming, setStreaming] = useState(false);
  const [recording, setRecording] = useState(false);
  const [loading, setLoading] = useState<OutputAction | null>(null);

  const toggle = async (action: OutputAction) => {
    setLoading(action);
    try {
      const isActive = action === "stream" ? streaming : recording;
      const endpoint = `/${action}/${isActive ? "stop" : "start"}`;
      const res = await fetch(`/api${endpoint}`, { method: "POST" });
      if (!res.ok) throw new Error("Failed");

      if (action === "stream") setStreaming(!streaming);
      else setRecording(!recording);
    } catch {
      // could surface error to user
    } finally {
      setLoading(null);
    }
  };

  return (
    <div className={styles.controls}>
      <button
        className={`${styles.output} ${streaming ? styles.active : ""}`}
        onClick={() => toggle("stream")}
        disabled={loading === "stream"}
      >
        <span className={styles.outputDot} />
        <span className={styles.outputLabel}>
          {loading === "stream" ? "..." : streaming ? "Stop Stream" : "Start Stream"}
        </span>
      </button>

      <button
        className={`${styles.output} ${styles.record} ${recording ? styles.active : ""}`}
        onClick={() => toggle("record")}
        disabled={loading === "record"}
      >
        <span className={styles.outputDot} />
        <span className={styles.outputLabel}>
          {loading === "record" ? "..." : recording ? "Stop Recording" : "Start Recording"}
        </span>
      </button>
    </div>
  );
}

export default Controls;
