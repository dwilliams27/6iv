const xPad = 17, yPad = 17;
console.log("dad2")
//Create a Pixi Application
let app = new PIXI.Application();
app.renderer.view.style.position = "absolute";
app.renderer.view.style.display = "block";
app.renderer.autoResize = true;
app.renderer.resize(window.innerWidth - xPad, window.innerHeight - yPad);
//Add the canvas that Pixi automatically created for you to the HTML document
document.body.appendChild(app.view);