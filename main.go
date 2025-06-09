package main

import "fmt"

const TODO_DIR_ENV_VAR = "TODO_DIR"

var todoDir = "~/.todo"
var configFile = "~/.todo/config"

func main() {
	fmt.Println("godo")
}
