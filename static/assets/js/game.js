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

	addControl = () => {
		document.onkeyup = (e) => {
			let direction;
			switch (e.key) {
				case "w":
				case "ArrowUp":
					direction = "UP";
					break;
				case "s":
				case "ArrowDown":
					direction = "DOWN";
					break;
				case "d":
				case "ArrowRight":
					direction = "RIGHT";
					break;
				case "a":
				case "ArrowLeft":
					direction = "LEFT";
					break;
				case " ":
				case "Enter":
					this.gameBtn.click();
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

	DrawSnake = (snake, food) => {
		const fs = this.blockSize.value * this.fieldSize.value;
		this.ctx.clearRect(0, 0, fs, fs);
		this.ctx.fillStyle = "white";
		snake.map((block) => {
			this.ctx.fillRect(block.x, block.y, this.blockSize.value - 1, this.blockSize.value - 1);
		});
		food?.map((block) => {
			this.ctx.fillRect(block.x, block.y, this.blockSize.value - 1, this.blockSize.value - 1);
		});
		snake.length > 1 && snake.length != this.result.innerText && (this.result.innerText = snake.length - 1);
	};

	Start = () => {
		this.uploadConfig();
	};

	GameOver = () => {
		this.isStarted = false;
		this.disableElements(".settings_controls", this.isStarted);
		this.gameBtn.innerText = "Start";
		this.gameOverResult.innerText = this.result.innerText;
		this.gameOverModal.showModal();
		SendMsg("CONNECT");
	};
}
