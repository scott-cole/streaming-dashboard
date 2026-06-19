import { useEffect, useState } from "react";
import styles from "./TopBar.module.scss";

type StreamStatus = "offline" | "live" | "recording";

interface StatusData {
  streaming: boolean;
  recording: boolean;
  scene: string;
}

function TopBar() {
  const [status, setStatus] = useState<StreamStatus>("offline");
  const [statusText, setStatusText] = useState("Offline");
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    const fetchStatus = async () => {
      try {
        const res = await fetch("/api/status");
        if (!res.ok) throw new Error("Not connected");
        const data: StatusData = await res.json();
        setConnected(true);

        if (data.streaming) {
          setStatus("live");
          setStatusText(`LIVE — ${data.scene}`);
        } else if (data.recording) {
          setStatus("recording");
          setStatusText(`REC — ${data.scene}`);
        } else {
          setStatus("offline");
          setStatusText("Offline");
        }
      } catch {
        setConnected(false);
        setStatus("offline");
        setStatusText("Disconnected");
      }
    };

    fetchStatus();
    const interval = setInterval(fetchStatus, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <header className={styles.topbar}>
      <div className={styles.search}>
        <span className={styles.searchIcon}>⌕</span>
        <input type="text" placeholder="Search..." className={styles.searchInput} />
      </div>

      <div className={styles.right}>
        <div className={`${styles.status} ${styles[status]}`}>
          <span className={styles.statusDot} />
          <span className={styles.statusText}>{statusText}</span>
        </div>

        <button className={styles.goLive}>
          Go Live
        </button>

        <div className={styles.profileIcon}>SC</div>
      </div>
    </header>
  );
}

export default TopBar;
