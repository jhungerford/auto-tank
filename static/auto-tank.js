document.addEventListener('keydown', sendKey);
document.addEventListener('keyup', sendKey);

function sendKey(event) {
  if (event.key.includes("Arrow")) {
    console.log(event.type, event.key);
  }
}