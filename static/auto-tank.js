document.addEventListener('keydown', keyDown);
document.addEventListener('keyup', keyUp);

const http = new XMLHttpRequest();
const keys = {
  up: false,
  down: false,
  left: false,
  right: false
};

function keyDown(event) {
  if (event.key.includes("Arrow")) {
    let direction = event.key.replace("Arrow", "").toLowerCase();
    // Only send the direction once when the key is pressed - down event is fired while the key is held.
    if (!keys[direction]) {
      keys[direction] = true;
      http.open("POST", "/move", true);
      http.send(direction);
    }
  }
}

function keyUp(event) {
  if (event.key.includes("Arrow")) {
    let direction = event.key.replace("Arrow", "").toLowerCase();
    keys[direction] = false;
  }
}
