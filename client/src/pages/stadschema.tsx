import styles from '../styles/stadschema.module.css'

import WebSocketComponent from '../components/websockets'

export default function Stadschema() {
  return (
    <div className={styles.outer_container}>
          <div className={styles.login_container}>
              <span>Städkod</span>
              <span>{WebSocketComponent()}</span>
              <input type="text" />
              <button>GO!</button>
          </div>
    </div>
  )
}
