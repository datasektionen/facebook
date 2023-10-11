import React, { useState, useEffect } from 'react';

const WebSocketComponent = () => {

    const [socket_state, set_socket_state] = useState<string[]>([]);

    useEffect(() => {
        const socket = new WebSocket('ws://localhost:5001/websocket'); // Replace with your WebSocket server URL

        socket.onopen = (event) => {
            console.log('WebSocket connection opened:', event);
        };

        socket.onmessage = (event) => {
            console.log('Received message from server:', event.data);
            set_socket_state((prevSocketState) => [event.data.toString(), ...prevSocketState]);

        };

        socket.onclose = (event) => {
            console.log('WebSocket connection closed:', event);
        };

        return () => {
            // Clean up the WebSocket connection when the component unmounts
            socket.close();
        };
    }, []);

    const display_sockets = () => {
        const socket_display_arr: JSX.Element[] = []
        socket_state.forEach(element => {
            socket_display_arr.push(<p>{element}</p>)
        }); 
        return socket_display_arr
    }

    return (
        <div>
            <p>{display_sockets()}</p>
        </div>
    );
};

export default WebSocketComponent;