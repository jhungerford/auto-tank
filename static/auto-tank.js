document.addEventListener('keydown', keyDown);
document.addEventListener('keyup', keyUp);

const http = new XMLHttpRequest();
const keys = {
  up: false,
  down: false,
  left: false,
  right: false
};
let directionStack = ['stop'];

function keyDown(event) {
  if (event.key.includes("Arrow")) {
    let direction = event.key.replace("Arrow", "").toLowerCase();
    // Only send the direction once when the key is pressed - down event is fired while the key is held.
    if (!keys[direction]) {
      directionStack.push(direction);
      keys[direction] = true;
      postMove(direction);
    }
  }
}

function keyUp(event) {
  if (event.key.includes("Arrow")) {
    let direction = event.key.replace("Arrow", "").toLowerCase();
    keys[direction] = false;

    let changeDirection = directionStack[directionStack.length - 1] === direction;

    directionStack = directionStack.filter(function (value) {
      return value !== direction;
    });

    if (changeDirection) {
      postMove(directionStack[directionStack.length - 1]);
    }
  }
}

function postMove(direction) {
  http.open("POST", "/move", true);
  http.send(direction);
}
