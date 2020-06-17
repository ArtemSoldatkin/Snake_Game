import { Reducer } from "./reducer.js";

const socket = new WebSocket("ws://localhost:5000/echo");

export const SendMsg = (action, data = null) => socket.send(JSON.stringify({ action, data }));

socket.onopen = () => {
	console.info("Connected to server");
	SendMsg("CONNECT");
};

socket.onmessage = (e) => {
	const msgData = JSON.parse(e.data);
	const { action, data } = msgData;
	Reducer(action, data);
};

socket.onclose = () => console.warn("Connection is lost");

socket.onerror = (err) => console.error(err);
