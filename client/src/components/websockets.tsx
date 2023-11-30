import React, { useState, useEffect } from 'react';
import styles from '../styles/websockets.module.css';

interface Task {
    name: string;
    description: string;
    done: boolean;
}

interface Stadschema {
    [key: string]: Task[];
}

const stadschema: Stadschema = {
    Mötesrummet: [
        {
            name: 'Skrivbord',
            description: 'Torka av skrivbordet',
            done: false,
        },
        {
            name: 'Fönsterbräda',
            description: 'Torka av fönsterbrädan',
            done: false,
        },
        {
            name: 'Fönster',
            description: 'Torka av fönstret',
            done: false,
        },
        {
            name: 'Golv',
            description: 'Sopa och torka golvet',
            done: false,
        },
        {
            name: 'Skräp',
            description: 'Töm papperskorgen',
            done: false,
        },
    ],
};

export default function WebSocketComponent({ code }: { code: string }) {
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const [data, setData] = useState<Stadschema | null>(stadschema);

    useEffect(() => {
        const socket = new WebSocket(`ws://localhost:5001/websocket?code=${code}`); // Replace with your WebSocket server URL

        socket.onopen = (event) => {
            console.log('WebSocket connection opened:', event);
        };

        socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log('Received message from server:', data);
        };

        socket.onclose = (event) => {
            console.log('WebSocket connection closed:', event);
        };

        setSocket(socket);

        return () => {
            // Clean up the WebSocket connection when the component unmounts
            socket.close();
        };
    }, [code]);

    function sendData() {
        if (socket) {
            const data = 'Hello, world!';
            socket.send(data);
            console.log(`Sent message to server: ${data}`);
        }
    }

    async function onTaskChanged(room: string, task: Task, done: boolean) {
        setData((data) => {
            data![room].find((t) => t.name === task.name)!.done = done;
            return { ...data };
        });

        const body = {
            room,
            task: task.name,
            done,
        };

        try {
            await fetch(`http://localhost:5001/update_schedule?key=sokior`, {
                method: 'PUT',
                body: JSON.stringify(body),
            });
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className={styles.container}>
            {data &&
                Object.keys(data).map((room) => (
                    <div
                        className={styles.room}
                        key={room}>
                        <h1>{room}</h1>
                        {data[room].map((task) => (
                            <div
                                key={room + task.name}
                                className={styles.task}>
                                <input
                                    type="checkbox"
                                    checked={task.done}
                                    onChange={() => onTaskChanged(room, task, !task.done)}
                                    className={styles.input}
                                />
                                <div>
                                    <h2>{task.name}</h2>
                                    <p>{task.description}</p>
                                </div>
                            </div>
                        ))}
                    </div>
                ))}
        </div>
    );
}
