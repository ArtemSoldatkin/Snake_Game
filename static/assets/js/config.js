import { SendMsg } from "./websocket.js";

export class Config {
	body = document.getElementsByTagName("body")[0];
	// game field
	field = document.getElementById("game_field");
	// result
	result = document.getElementById("results_value");
	// settings inputs
	blockSize = document.getElementById("block_size");
	fieldSize = document.getElementById("field_size");
	gameSpeed = document.getElementById("game_speed");
	// settings button
	settingSBtn = document.getElementById("send_config");
	// game button
	gameBtn = document.getElementById("game_button");
	// game over modal
	gameOverModal = document.getElementById("game_over");
	gameOverResult = document.getElementById("game_over_result");
	gameOverCloseBtn = document.getElementById("close_button");

	constructor() {
		this.disableElements("input", true);
		this.disableElements("button", true);
		this.settingSBtn.addEventListener("click", this.setSettings);
		this.addInputValidate();
		this.gameOverCloseBtn.addEventListener("click", this.closeModal);
		this.body.addEventListener("click", this.clickOutsideModal);
	}

	clickOutsideModal = (e) => {
		e.target === this.gameOverModal && this.closeModal();
	};

	closeModal = () => {
		this.gameOverModal.close();
	};

	addInputValidate = () => {
		const inputs = document.getElementsByTagName("input");
		for (let i = 0; i < inputs.length; ++i) inputs[i].addEventListener("blur", this.validateInputs);
	};

	validateInputs = (e) => {
		let { value, min, max } = e.target;
		+value > +max && (value = max);
		+value < +min && (value = min);
		e.target.value = value;
	};

	disableElements = (selector, value) => {
		const elements = document.querySelectorAll(selector);
		for (let i = 0; i < elements.length; ++i) elements[i].disabled = value;
	};

	setSettings = () => {
		const fieldSize = this.blockSize.value * this.fieldSize.value;
		this.field.width = fieldSize;
		this.field.height = fieldSize;
		SendMsg("SET_GAME_SETTINGS", {
			field_size: parseInt(this.fieldSize.value),
			block_side: parseInt(this.blockSize.value),
			speed: parseInt(this.gameSpeed.value),
		});
	};

	uploadConfig = async () => {
		const CONFIG = await fetch("../../config.json").then((data) => data.json());
		const bs = CONFIG["block_side"];
		const fs = CONFIG["field_size"];
		const fieldSize = fs * bs;
		this.field.width = fieldSize;
		this.field.height = fieldSize;
		this.blockSize.value = bs;
		this.fieldSize.value = fs;
		this.gameSpeed.value = CONFIG["speed"];
		this.gameSpeed.max = CONFIG["max_speed"];
		this.disableElements("input", false);
		this.disableElements("button", false);
	};
}
