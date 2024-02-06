'use client'
import { useRouter } from 'next/navigation'

import RootLayout from '../components/layout';
import styles from '../styles/landing.module.css'
import Header from '@/components/header';

export default function Landing() {
  const router = useRouter()

  const choices = ["Mötesrummet", "Catwalken", "Pant", "Mikros", "Dega"];
  const numbers = ["NEJ", "NEJ", "NEJ", "NEJ", "JA"]
  const choiceNums = choices.concat(numbers);
  const listItems = choiceNums.map((choiceNum, key) =>
    <li key={key}>
      {choiceNum}
    </li>);
    
  return (
    <RootLayout>
        <div className={styles.main}>
            <ul>{listItems}</ul>
            <br/>
            <button className={styles.fun_button} onClick={()=>router.push('/stadschema')}>
              Måndagsstäd!
            </button>
        </div>
    </RootLayout>
  )
}
