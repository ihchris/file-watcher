package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Directory to watch
	dir := "./watchdir"

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
		fmt.Printf("Created directory: %s\n", dir)
	}

	// Create a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	// Add the directory to the watcher
	err = watcher.Add(dir)
	if err != nil {
		log.Fatalf("Failed to add directory to watcher: %v", err)
	}
	fmt.Printf("Watching directory: %s\n", dir)

	// Start listening for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Log the event
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("File modified: %s\n", event.Name)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("File created: %s\n", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Printf("File deleted: %s\n", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Printf("File renamed: %s\n", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Keep the program running
	fmt.Println("File watcher is running. Press Ctrl+C to exit.")
	select {}
}