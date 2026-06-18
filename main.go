package main

import (
	"fmt"
	"os"
	"streaming-dashboard/obs"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <password> <command> [args]")
		fmt.Println("Commands: version, scenes, switch <scene name>")
		return
	}

	password := os.Args[1]
	command := os.Args[2]

	client, err := obs.New(password)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect()

	switch command {
	case "version":
		fmt.Println(client.Version())
	case "scenes":
		client.ListScenes()
	case "switch":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run . <password> switch <scene name>")
			return
		}
		client.SwitchScene(os.Args[3])
	default:
		fmt.Println("Unknown command:", command)
	}
}
