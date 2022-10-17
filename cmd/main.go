package main

import "github.com/alexm24/ticket-booking/internal/app"

func main() {
	const configPath = "configs/config.yaml"
	app.App(configPath)
}
