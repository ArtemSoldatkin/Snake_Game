import { SendMsg } from "./websocket.js";

// Initialize fields
const field = document.getElementById("game_field");
const widthSize = document.getElementById("width_size");
const heightSize = document.getElementById("height_size");
const configButton = document.getElementById("send_config");
const startPauseButton = document.getElementById("start_pause_game");

let isStarted = false;

startPauseButton.addEventListener("click", () => {
	isStarted = !isStarted;
	startPauseButton.innerText = isStarted ? "Pause" : "Start";
	SendMsg("START_PAUSE", isStarted);
});

configButton.addEventListener("click", () => {
	const width = widthSize.value;
	const height = heightSize.value;
	if (!width || !height) {
		alert("Wrong parameters");
		return;
	}
	field.style.width = width;
	field.style.height = height;
	SendMsg("SET_GAME_SETTINGS", { width, height });
});

// Load JSON config
(async () => {
	const CONFIG = await fetch("../../config.json").then((data) => data.json());
	//const CONFIG = await data.json();
	field.style.width = CONFIG["width"];
	field.style.height = CONFIG["height"];
})();
