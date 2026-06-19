import { useState } from "react";
import styles from "./App.module.scss";
import Sidebar from "./components/Sidebar/Sidebar";
import TopBar from "./components/TopBar/TopBar";
import SceneList from "./components/SceneList/SceneList";
import Preview from "./components/Preview/Preview";
import NowPlaying from "./components/NowPlaying/NowPlaying";
import Controls from "./components/Controls/Controls";

function App() {
  const [activeNav, setActiveNav] = useState("stream-control");
  const [currentScene, setCurrentScene] = useState("");

  return (
    <div className={styles.layout}>
      <Sidebar activeNav={activeNav} onNavChange={setActiveNav} />
      <div className={styles.mainArea}>
        <TopBar />
        <main className={styles.content}>
          <div className={styles.colFull}>
            <Preview sceneName={currentScene || "Scene"} />
          </div>
          <SceneList onSceneChange={setCurrentScene} />
          <div>
            <NowPlaying sceneName={currentScene || "No scene selected"} />
            <div style={{ marginTop: 24 }}>
              <Controls />
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}

export default App;
