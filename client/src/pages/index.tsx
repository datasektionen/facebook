'use client'
import { useRouter } from 'next/navigation'

import RootLayout from '../components/layout';
import styles from '../styles/landing.module.css'

export default function Landing() {
  const router = useRouter()

  return (
    <RootLayout>
        <div className={styles.main}>
            <button className={styles.fun_button} onClick={()=>router.push('/stadschema')}>
                Click me!
            </button>
        </div>
    </RootLayout>
  )
}
