package main

import "githug.com/guilhermeayusso/api-goexpert/pkg/config"

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(cfg.DBDriver)
	println(cfg.DBHost)
	println(cfg.DBPort)
	println(cfg.DBUser)
	println(cfg.DBPassword)
	println(cfg.DBName)
	println(cfg.WebServerPort)
	println(cfg.JWTSecret)
	println(cfg.JWTExperesIn)
}
