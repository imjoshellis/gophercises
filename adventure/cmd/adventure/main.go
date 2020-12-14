package main

import (
	"adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the app on")
	file := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := adventure.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story)
	fmt.Printf("Starting the server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), h))
}
