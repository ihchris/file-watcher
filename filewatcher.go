package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// WatchDirectory initializes a file watcher for the given directory.
func WatchDirectory(dir string) {
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
				// Log the event using the utility function
				LogEvent(event)
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
