import { SendMsg } from "./websocket.js";

// Initialize fields
const field = document.getElementById("game_field");
const widthSize = document.getElementById("width_size");
const heightSize = document.getElementById("height_size");
const configButton = document.getElementById("send_config");

configButton.addEventListener("click", () => {
	const width = widthSize.value;
	const height = heightSize.value;
	if (!width || !height) {
		alert("Wrong parameters");
		return;
	}
	SendMsg("SET_GAME_SETTINGS", { width, height });
});

// Load JSON config
(async () => {
	const CONFIG = await fetch("../../config.json").then((data) => data.json());
	//const CONFIG = await data.json();
	field.style.width = CONFIG["width"];
	field.style.height = CONFIG["height"];
})();
