package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

/* Helper */
func isPortInUse(port int) bool {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false // Port is closed or unreachable
	}
	defer func() { _ = conn.Close() }()
	return true // Port is in use
}

/* Main */
var shellRunning bool

func parseCli(command []string) error {
	if len(command) == 0 {
		return nil
	}

	switch command[0] {
	case "-h", "--help":
		printHelp()
	case "shell":
		if shellRunning {
			return errors.New("instance of shell already running")
		}

		runShell()
		shellRunning = true
	case "start":
		if isPortInUse(8080) {
			fmt.Println("Server could not be started, something already running in port 8080")
			return nil
		}
		return startServer()
	default:
		fmt.Println("Unknown Command:", command[0])
		fmt.Println("Run --help, -h for more info.")
	}

	return nil
}

func printHelp() {
	fmt.Println("Usage: ushort start | ushort <commands>")
	fmt.Println("\nCommands: ")
	fmt.Printf("  %-20s\tStart localhost server at port 8080.\n", "start")
	fmt.Printf("  %-20s\tStart an instance of shell.\n", "shell")
	fmt.Printf("  %-20s\tDisplay this output.\n", "-h, --help")
}

func runShell() {
	// TODO: implement interactive shell
	fmt.Println("shell not implemented yet")
}
