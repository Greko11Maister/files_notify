package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

// main
func main() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				fmt.Println(event)
				//fmt.Printf(event.Op.String())

				file, errorF := os.Open(event.Name) // Para acceso de lectura.

				if errorF != nil {
					//log.Fatal(err)
					fmt.Println(errorF)
				}
				fi, err := file.Stat()
				file.Close()
				if err != nil {
					fmt.Println(err)
					//return
				}

				fmt.Println(fi)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("files"); err != nil {
		fmt.Println("ERROR 2", err)
	}

	<-done
}
