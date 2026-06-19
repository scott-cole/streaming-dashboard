import styles from "./NowPlaying.module.scss";

interface NowPlayingProps {
  sceneName: string;
}

function NowPlaying({ sceneName }: NowPlayingProps) {
  return (
    <div className={styles.card}>
      <div className={styles.header}>
        <span className={styles.label}>Current Scene</span>
        <div className={styles.liveBadge}>
          <span className={styles.liveDot} />
          Active
        </div>
      </div>
      <div className={styles.sceneName}>{sceneName}</div>
      <div className={styles.meta}>
        <span>● Scene</span>
      </div>
    </div>
  );
}

export default NowPlaying;
