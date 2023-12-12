package main

import "github.com/rgoncalvesrr/fullcycle-api/configs"

func main() {
	cfg := configs.NewConfig()

	println(cfg.DBDriver)
}
