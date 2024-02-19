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

interface Tasks {
    ChecklistJSON: ChecklistJSON,
    Comments: string,
    Key: string,
    Overseers: string,
}

interface ChecklistJSON {
    A1: boolean,
    A2: boolean,
    A3: boolean,
    A4: boolean,
    A5: boolean,
    A6: boolean,
    B1: boolean,
    B2: boolean,
    B3: boolean,
    B4: boolean,
    B5: boolean,
    B6: boolean,
    B7: boolean,
    B8: boolean,
    B9: boolean,
    B10: boolean,
    C1: boolean,
    C2: boolean,
    C3: boolean,
    C4: boolean,
    C5: boolean,
    C6: boolean,
    C7: boolean,
    C8: boolean,
    C9: boolean,
    C10: boolean,
    C11: boolean,
    C12: boolean,
    C13: boolean,
    D1: boolean,
    D2: boolean,
    D3: boolean,
    D4: boolean,
    D5: boolean,
    E1: boolean,
    E2: boolean,
    E3: boolean,
    E4: boolean,
    E5: boolean,
    F1: boolean,
    F2: boolean
}

interface Stadschema {
    category: string;
    swedish_values: SwedishValue[];
    english_values: EnglishValue[];
}

interface SwedishValue {
    checklistItemID: string;
    prompt: string;
}

interface EnglishValue {
    checklistItemID: string;
    prompt: string;
}

const stadschema: Stadschema[] = [
    {
        category: "Hallen & Toaletter",
        swedish_values: [
            {
                prompt: "Plocka bort stolar och bord som använts i entrén",
                ChecklistItemID: "A1",
            },
            {
                prompt: "Ställ tillbaka saker ordentligt i städskrubben",
                ChecklistItemID: "A2",
            },
            {
                prompt: "!Se till att brandcentralen inte blockeras!",
                ChecklistItemID: "A3",
            },
            {
                prompt: "!Återställ dörröppnaren så att den funkar dagen efter. Båda knapparna på apparaten ska vara tryckta uppåt för att återställa!",
                ChecklistItemID: "A4",
            },
            {
                prompt: "!Se till att Entrédörren är stängd och låst!",
                ChecklistItemID: "A5",
            },
            {
                prompt: "Se till att det inte ser förjävligt ut",
                ChecklistItemID: "A6",
            },
        ],
        english_values: [
            {
                prompt: "Remove chairs and tables that have been used at the entrance",
                ChecklistItemID: "A1",
            },
            {
                prompt: "Put things back properly in the cleaning closet",
                ChecklistItemID: "A2",
            },
            {
                prompt: "!Ensure that the fire center is not blocked!",
                ChecklistItemID: "A3",
            },
            {
                prompt: "!Reset the door opener so that it works the day after. Both buttons on the device should be pushed upward to reset!",
                ChecklistItemID: "A4",
            },
            {
                prompt: "!Make sure that the entrance door is closed and locked!",
                ChecklistItemID: "A5",
            },
            {
                prompt: "Make sure it doesn't look terrible",
                ChecklistItemID: "A6",
            },
        ],
    }
];



export default function WebSocketComponent({ code }: { code: string }) {
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const [data, setData] = useState<Stadschema[] | null>(stadschema);
    const [tasks, setTasks] = useState<Tasks | null>(null);

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
                setData(data.parsedChecklistItems);
            }

            if (data.scheduleItems){
                setTasks(data.scheduleItems[0])
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
            await fetch(`http://localhost:5001/api/update_schedule?key=${code}`, {
                method: 'PUT',
                body: JSON.stringify(body),
            });
        } catch (error) {
            console.error(error);
        }
    }

    const language = "swedish_values"

    const out_log = (room)=>{
        console.log("CATEGORY: ", room.category)
        return(null)
    }

    return (
        <div className={styles.container}>
            {
                data && data.map(room => (
                    // room != null? out_log(room):
                    <div
                        className={styles.room}
                        key={room.category}>
                        <h1>{room.category}</h1>
                        {room[language].map((task)=>(
                            <div
                                key={task.prompt}
                                className={styles.task}>
                                <input 
                                    type="checkbox" 
                                    checked={
                                        tasks?
                                        tasks.ChecklistJSON[task.checklistItemID]
                                        :
                                        false
                                    }
                                    onChange={onTaskChanged()}
                                />
                            </div>
                        ))}
                    </div>
                ))
            }
        </div>

        // <div className={styles.container}>
        //     {data &&
        //         Object.keys(data).map((room) => (
        //             <div
        //                 className={styles.room}
        //                 key={room}>
        //                 <h1>{room}</h1>
        //                 {data[room].map((task) => (
        //                     <div
        //                         key={room + task.name}
        //                         className={styles.task}>
        //                         <input
        //                             type="checkbox"
        //                             checked={task.done}
        //                             onChange={() => onTaskChanged(room, task, !task.done)}
        //                             className={styles.input}
        //                         />
        //                         <div>
        //                             <h2>{task.name}</h2>
        //                             <p>{task.description}</p>
        //                         </div>
        //                     </div>
        //                 ))}
        //             </div>
        //         ))}
        // </div>
    );
}
