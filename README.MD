# File Watcher in Go

A simple Go program that monitors a directory for file changes (create, modify, delete, rename) and logs the events to the console. This can be used for automation, syncing, or triggering actions when files change.

---

## Features

- Watches a specified directory for file system events.
- Logs the following events:
  - File creation
  - File modification
  - File deletion
  - File renaming
- Easy to extend for custom actions (e.g., syncing, backups, notifications).

---

## Prerequisites

- Go installed on your machine (version 1.16 or higher recommended).
- Basic knowledge of Go and the command line.

---

## Installation

```bash
1. Clone this repository or download the `filewatcher.go` file:
   git clone https://github.com/yourusername/file-watcher-go.git
   cd file-watcher-go

2. Install the `fsnotify` package:
   go get github.com/fsnotify/fsnotify

3. Build and run the program:
   go run filewatcher.go