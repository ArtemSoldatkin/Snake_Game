/*import { DrawSnake } from "./game/index.js";
import { GameOver } from "./config.js";*/

export const Reducer = (action, data) => {
	console.log(action, data);
	switch (action) {
		case "INITIALIZE":
			//DrawSnake(data);
			break;
		case "MOVE":
			//DrawSnake(data);
			break;
		case "GAME_OVER":
			//GameOver();
			break;
	}
};
