package main

import "github.com/victorhugonf/go-expert/07-APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
