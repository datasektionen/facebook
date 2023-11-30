import styles from '../styles/stadschema.module.css';

import WebSocketComponent from '../components/websockets';
import { FormEvent } from 'react';
import { useRouter } from 'next/router';

export default function Stadschema() {
    const router = useRouter();
    const code = router.query.code as string;
    /* 
    function connectSocket(form: FormEvent<HTMLFormElement>) {
        // Prevent url params from being added to url
        form.preventDefault();
        router.push(`/stadschema?code=${form.currentTarget.code.value}`);
    } */

    // If we don't have a code, show the code input form
    if (!code) {
        return (
            <div className={styles.outer_container}>
                <form className={styles.login_container}>
                    <label>Städkod</label>
                    <input
                        type="text"
                        name="code"
                    />
                    <button type="submit">GO!</button>
                </form>
            </div>
        );
    }

    // If we have a code, show the WebSocketComponent
    return <WebSocketComponent code={code} />;
}
