import { AddBlock } from "./game/index.js";

export const Reducer = (action, data) => {
	console.log(action, data);
	switch (action) {
		case "INITIALIZE":
			AddBlock(data);
			break;
	}
};
