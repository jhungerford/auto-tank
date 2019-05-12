keyMap = {
	"ArrowUp": "up",
	"ArrowDown": "down",
	"ArrowLeft": "left",
	"ArrowRight": "right"
};

keys = {
	"up": false,
	"down": false,
	"left": false,
	"right": false
};

document.body.addEventListener('keydown', function(e) {
	let first = !keys[keyMap[e.key]];

	keys[keyMap[e.key]] = true;

	if (first) {
		console.log("Down", keys);
	}
});

document.body.addEventListener('keyup', function(e) {
	keys[keyMap[e.key]] = false;
}

