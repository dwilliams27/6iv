const PIXI = require("pixi.js");

function loadImage(filename) {
  PIXI.loader
    .add(filename)
    .load(createImage(filename));
  return filename;
}

function createImage(filename) {
  const sprite = new PIXI.Sprite(
    PIXI.loader.resources[filename].texture
  );
}