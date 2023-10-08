import styles from '../styles/stadschema.module.css'

export default function Stadschema() {
  return (
    <div className={styles.outer_container}>
          <div className={styles.login_container}>
              <span>Städkod</span>
              <input type="text" />
              <button>GO!</button>
          </div>
    </div>
  )
}
