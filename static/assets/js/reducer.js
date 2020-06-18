import { game } from "./index.js";

export const Reducer = (action, data) => {
	console.log(action, data);
	switch (action) {
		case "INITIALIZE":
		case "MOVE":
			game.DrawSnake(data);
			break;
		case "GAME_OVER":
			game.GameOver();
			break;
	}
};
