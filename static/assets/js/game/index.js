const field = document.getElementById("game_field");
const ctx = field.getContext("2d");

export const DrawSnake = (snake) => {
	ctx.clearRect(0, 0, field.width, field.height);
	ctx.fillStyle = "white";
	snake.map((block) => {
		ctx.fillRect(block.x, block.y, block.width, block.height);
	});
};
