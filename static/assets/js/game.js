import { Config } from "./config.js";
import { SendMsg } from "./websocket.js";

export class Game extends Config {
	constructor() {
		super();
		this.ctx = this.field.getContext("2d");
		this.addControl();
		this.gameBtn.addEventListener("click", this.startStop);
		this.isStarted = false;
	}

	// add start to Enter
	addControl = () => {
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
			if (!!direction && this.isStarted) SendMsg(direction);
		};
	};

	startStop = () => {
		this.isStarted = !this.isStarted;
		this.disableElements(".settings_controls", this.isStarted);
		this.gameBtn.innerText = this.isStarted ? "Pause" : "Start";
		SendMsg("START_PAUSE", this.isStarted);
	};

	DrawSnake = (snake) => {
		const fs = this.blockSize.value * this.fieldSize.value;
		this.ctx.clearRect(0, 0, fs, fs);
		this.ctx.fillStyle = "white";
		snake.map((block) => {
			this.ctx.fillRect(block.x, block.y, this.blockSize.value, this.blockSize.value);
		});
	};

	Start = () => {
		this.uploadConfig();
	};

	GameOver = () => {
		this.isStarted = false;
		this.gameBtn.innerText = "Start";
		alert("GAME_OVER");
		SendMsg("CONNECT");
	};
}
