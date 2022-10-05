package main

import (
  _ "embed"
  "github.com/wailsapp/wails"
  "github.com/leaanthony/mewn"
)

func basic() string {
  return "World!"
}

func action(id string) string{
  return id + "Hello!"
}

//goembed frontend/build/static/js/main.js
var js string=mewn.String("./frontend/src/index.js")

//goembed frontend/build/static/css/main.css
var css string=mewn.String("./frontend/src/index.css")

func main() {

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "DockSetter",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Bind(action)
  app.Run()
}
