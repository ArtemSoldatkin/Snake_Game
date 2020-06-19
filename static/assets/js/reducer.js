import { game } from "./index.js";

export const Reducer = (action, data) => {
	switch (action) {
		case "INITIALIZE":
		case "MOVE":
			const { snake, food } = data;
			game.DrawSnake(snake, food);
			break;
		case "GAME_OVER":
			game.GameOver();
			break;
	}
};
