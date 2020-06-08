import { Block } from "./block.js";

export class Snake {
	constructor() {
		this.blocks = [];
	}

	AddBlock = (data) => {
		const block = new Block(data);
		this.blocks.push(block);
	};
}
