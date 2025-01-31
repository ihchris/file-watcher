package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// LogEvent logs file system events to the console.
func LogEvent(event fsnotify.Event) {
	switch {
	case event.Op&fsnotify.Write == fsnotify.Write:
		fmt.Printf("File modified: %s\n", event.Name)
	case event.Op&fsnotify.Create == fsnotify.Create:
		fmt.Printf("File created: %s\n", event.Name)
	case event.Op&fsnotify.Remove == fsnotify.Remove:
		fmt.Printf("File deleted: %s\n", event.Name)
	case event.Op&fsnotify.Rename == fsnotify.Rename:
		fmt.Printf("File renamed: %s\n", event.Name)
	default:
		fmt.Printf("Unknown event: %s\n", event.Name)
	}
}
