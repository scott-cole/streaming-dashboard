import styles from "./Sidebar.module.scss";

const navItems = [
  { label: "Dashboard", icon: "⊞", id: "dashboard" },
  { label: "Stream Control", icon: "▶", id: "stream-control" },
  { label: "Settings", icon: "⚙", id: "settings" },
];

interface SidebarProps {
  activeNav: string;
  onNavChange: (id: string) => void;
}

function Sidebar({ activeNav, onNavChange }: SidebarProps) {
  return (
    <aside className={styles.sidebar}>
      <div className={styles.logo}>
        <span className={styles.logoIcon}>◆</span>
        <span className={styles.logoText}>Streaming Dashboard</span>
      </div>

      <nav className={styles.nav}>
        {navItems.map((item) => (
          <button
            key={item.id}
            className={`${styles.navItem} ${activeNav === item.id ? styles.navItemActive : ""}`}
            onClick={() => onNavChange(item.id)}
          >
            <span className={styles.navIcon}>{item.icon}</span>
            <span>{item.label}</span>
          </button>
        ))}
      </nav>

      <div className={styles.section}>
        <div className={styles.sectionLabel}>Integrations</div>
        <div className={styles.integration}>
          <span className={styles.dot} />
          <span>OBS</span>
        </div>
        <div className={`${styles.integration} ${styles.disabled}`}>
          <span className={`${styles.dot} ${styles.dotInactive}`} />
          <span>Discord</span>
        </div>
        <div className={`${styles.integration} ${styles.disabled}`}>
          <span className={`${styles.dot} ${styles.dotInactive}`} />
          <span>Twitch</span>
        </div>
      </div>

      <div className={styles.profile}>
        <div className={styles.avatar}>SC</div>
        <div className={styles.profileInfo}>
          <div className={styles.profileName}>Scott Cole</div>
          <div className={styles.profileRole}>Streamer</div>
        </div>
      </div>
    </aside>
  );
}

export default Sidebar;
