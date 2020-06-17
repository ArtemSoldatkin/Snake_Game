import { Config } from "./config.js";
import { SendMsg } from "./websocket.js";

export class Game extends Config {
	constructor() {
		super();
		this.ctx = this.field.getContext("2d");
		this.addControl();
	}

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
			direction && SendMsg(direction);
		};
	};

	drawSnake = (snake) => {
		this.ctx.clearRect(0, 0, this.field.width, this.field.height);
		this.ctx.fillStyle = "white";
		snake.map((block) => {
			this.ctx.fillRect(block.x, block.y, block.width, block.height);
		});
	};

	Start = () => {
		this.uploadConfig();
	};
}
