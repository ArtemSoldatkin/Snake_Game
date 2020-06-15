import { SendMsg } from "./websocket.js";

// Initialize fields
const field = document.getElementById("game_field");
// deprecated
const widthSize = document.getElementById("width_size");
const heightSize = document.getElementById("height_size");

const fieldSize = document.getElementById("field_size");
const configButton = document.getElementById("send_config");
const startPauseButton = document.getElementById("start_pause_game");

let isStarted = false;

startPauseButton.addEventListener("click", () => {
	setIsStarted();
});

configButton.addEventListener("click", () => {
	const width = widthSize.value;
	const height = heightSize.value;
	if (!width || !height) {
		alert("Wrong parameters");
		return;
	}
	field.width = width;
	field.height = height;
	SendMsg("SET_GAME_SETTINGS", { width, height });
});

// Load JSON config
(async () => {
	const CONFIG = await fetch("../../config.json").then((data) => data.json());
	const blockSize = CONFIG["block_side"];
	const fieldMult = CONFIG["field_size"];
	const fs = fieldMult * blockSize;
	fieldSize.value = fieldMult;
	field.width = fs;
	field.height = fs;
	fieldSize.value = fs;
})();

const setIsStarted = (newIsStarted) => {
	isStarted = newIsStarted === undefined ? !isStarted : newIsStarted;
	startPauseButton.innerText = isStarted ? "Pause" : "Start";
	widthSize.disabled = isStarted;
	heightSize.disabled = isStarted;
	configButton.disabled = isStarted;
	newIsStarted === undefined && SendMsg("START_PAUSE", isStarted);
};

export const GameOver = () => {
	alert("GAME OVER");
	setIsStarted(false);
};
