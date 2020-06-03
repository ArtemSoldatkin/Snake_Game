import { SendMsg } from "../websocket.js";

document.onkeyup = (e) => {
	let direction;
	switch (e.key) {
		case "ArrowUp":
			direction = "UP";
			break;
		case "ArrowDown":
			direction = "DOWN";
			break;
		case "ArrowRight":
			direction = "RIGHT";
			break;
		case "ArrowLeft":
			direction = "LEFT";
			break;
	}
	direction && SendMsg(direction);
};
