import { useRouter } from "next/navigation";
import * as React from "react";
import styles from '../styles/header.module.css'
//import logo from "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Facebook_Logo_%282019%29.png/240px-Facebook_Logo_%282019%29.png"

export default function Header() {
    const router = useRouter();

    return (
        <header className={styles.header}>
            <button onClick={() => router.push("/stadschema")}>
                <img className={styles.img} src={"https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Facebook_Logo_%282019%29.png/240px-Facebook_Logo_%282019%29.png"} />
                <h1>METADORERNA</h1>
            </button>
        </header>
    );
}