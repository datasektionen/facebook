import React, { useState, useEffect } from 'react';
import styles from '../styles/websockets.module.css';

// interface Stadschema {
//     [key: string]: Task[];
// }

// interface Task {
//     name: string;
//     description: string;
//     done: boolean;
// }

// const stadschema: Stadschema = {
//     Mötesrummet: [
//         {
//             name: 'Skrivbord',
//             description: 'Torka av skrivbordet',
//             done: false,
//         },
//         {
//             name: 'Fönsterbräda',
//             description: 'Torka av fönsterbrädan',
//             done: false,
//         },
//         {
//             name: 'Fönster',
//             description: 'Torka av fönstret',
//             done: false,
//         },
//         {
//             name: 'Golv',
//             description: 'Sopa och torka golvet',
//             done: false,
//         },
//         {
//             name: 'Skräp',
//             description: 'Töm papperskorgen',
//             done: false,
//         },
//     ],
// };

interface Stadschema {
    Category: string;
    SwedishValues: SwedishValue[];
    EnglishValues: EnglishValue[];
}

interface SwedishValue {
    ChecklistItemID: string;
    Prompt: string;
}

interface EnglishValue {
    ChecklistItemID: string;
    Prompt: string;
}

const stadschema: Stadschema[] = [
    {
        Category: "Hallen & Toaletter",
        SwedishValues: [
            {
                Prompt: "Plocka bort stolar och bord som använts i entrén",
                ChecklistItemID: "A1",
            },
            {
                Prompt: "Ställ tillbaka saker ordentligt i städskrubben",
                ChecklistItemID: "A2",
            },
            {
                Prompt: "!Se till att brandcentralen inte blockeras!",
                ChecklistItemID: "A3",
            },
            {
                Prompt: "!Återställ dörröppnaren så att den funkar dagen efter. Båda knapparna på apparaten ska vara tryckta uppåt för att återställa!",
                ChecklistItemID: "A4",
            },
            {
                Prompt: "!Se till att Entrédörren är stängd och låst!",
                ChecklistItemID: "A5",
            },
            {
                Prompt: "Se till att det inte ser förjävligt ut",
                ChecklistItemID: "A6",
            },
        ],
        EnglishValues: [
            {
                Prompt: "Remove chairs and tables that have been used at the entrance",
                ChecklistItemID: "A1",
            },
            {
                Prompt: "Put things back properly in the cleaning closet",
                ChecklistItemID: "A2",
            },
            {
                Prompt: "!Ensure that the fire center is not blocked!",
                ChecklistItemID: "A3",
            },
            {
                Prompt: "!Reset the door opener so that it works the day after. Both buttons on the device should be pushed upward to reset!",
                ChecklistItemID: "A4",
            },
            {
                Prompt: "!Make sure that the entrance door is closed and locked!",
                ChecklistItemID: "A5",
            },
            {
                Prompt: "Make sure it doesn't look terrible",
                ChecklistItemID: "A6",
            },
        ],
    }
];


export default function WebSocketComponent({ code }: { code: string }) {
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const [data, setData] = useState<Stadschema[] | null>(stadschema);

    useEffect(() => {
        const socket = new WebSocket(`ws://localhost:5001/websocket?code=${code}`); // Replace with your WebSocket server URL

        socket.onopen = (event) => {
            console.log('WebSocket connection opened:', event);
        };

        socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log('Received message from server:', data);

            if (data.parsedChecklistItems) {
                console.log("Updateing data to: ", data.parsedChecklistItems)
                setData(data.parsedChecklistItems[0].swedish_values);
            }

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
            await fetch(`http://localhost:5001/api/update_schedule?key=sokior`, {
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
