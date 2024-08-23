package main

import (
	"fmt"
	"sync"
)

var (
	config map[string]string
	once   sync.Once
)

func main() {
	getConfig()
	getConfig()
	getConfig()
}
func getConfig() {
	fmt.Println("Getting config")
	once.Do(loadConfig)
	fmt.Printf("Config is %v\n", config)
}

func loadConfig() {
	config = make(map[string]string)
	config["key"] = "value"
	fmt.Println("Loading config")
}
