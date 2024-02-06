import { useRouter } from "next/navigation";
import * as React from "react";
import styles from '../styles/header.module.css'
import logo from '../../public/METAdorerna.gif'

export default function Header() {
    const router = useRouter();

    return (
        <header className={styles.header}>
            <div className={styles.headerDiv} onClick={() => router.push("/stadschema")}>
                <img className={styles.img} src={logo.src} />
                <h1>METADORERNA</h1>
            </div>
        </header>
    );
}