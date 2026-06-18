import styles from './App.module.scss'
import SceneList from './components/SceneList/SceneList'

function App() {
  return (
    <div className={styles.app}>
      <h1>Streaming Dashboard</h1>
      <SceneList />
    </div>
  )
}

export default App
