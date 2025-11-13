package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* temporary testing until the rework */
func main() {
	flag.Parse()
	todos := Todos{}
	args := os.Args[1:] // skip program name

	if len(args) == 0 {
		fmt.Println("Usage: todo [add|list|delete|toggle]")
		return
	}

	command := args[0]

	switch command {
	case "add":
		desc := strings.Join(args[1:], "")
		todos.Load()
		if *completed {

		}
		todos.Add(desc)
		todos.Save("storage.json") // testing saving

	case "list":
		todos.Load()
		todos.List()

	case "delete":
		id, _ := strconv.Atoi(args[1])
		todos.Load()
		todos.Delete(id)
		todos.Save("storage.json") // testing saving

	case "toggle":
		id, _ := strconv.Atoi(args[1])
		todos.Load()
		todos.Toggle(id)
		todos.Save("storage.json") // testing saving

	default:
		fmt.Println("Unknown command")
	}
}
