package main

import "hackerman.ca/me/app"

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":5000")
}
