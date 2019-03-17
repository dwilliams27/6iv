const PIXI = require("pixi.js");

function loadSpriteSheet(filename) {
  const sheetFile = require("../assets/" + filename);

  if(filepath in sheetFile) {
    const sheet = PIXI.BaseTexture.fromImage(sheetFile.filepath);
    var spriteGroups = [];
    delete sheetFile.filepath;
    for(var prop in sheet) {
      console.log(prop);
      if(sheet.hasOwnProperty(prop)) {
        var spriteArray = [];
      }
    }
  } else {
    console.error("filepath not defined for sheetFile")
  }
}

function makeTexture(sheet, data, options) {
  if(sameSize in options) {
    if("xOffset" in options && "yOffset" in options && "count" in options) {
      return new PIXI.Texture(sheet, new PIXI.Rectangle(data.x, data.y, data.width, data.height));
    } else {
      console.error("sameSize flag defined, but other needed options are missing");
      return null;
    }
  }
}

module.exports = {
  loadSpriteSheet: loadSpriteSheet,
  makeTexture: makeTexture
}