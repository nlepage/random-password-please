// +build !js,!wasm

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

var (
	httpAddr = flag.String("http", defaultAddr(), "http listen address")
)

func main() {
	flag.Parse()

	if *counterFilePath != "" {
		var err error

		counterFile, err = os.OpenFile(*counterFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Failed to open counter file: %s", err)
		}
		counterBytes, err := ioutil.ReadAll(counterFile)
		if err != nil {
			log.Fatalf("Failed to read counter file: %s", err)
		}
		if len(counterBytes) > 0 {
			counter, err = strconv.ParseUint(string(bytes.TrimSpace(counterBytes)), 10, 64)
			if err != nil {
				log.Fatal("Failed to read counter value")
			}
		}
	}

	registerHandlers()

	// Ensure counter is saved on exit.
	go handleSignals()

	go generatePasswords()

	log.Print("Running at address ", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func saveCounter() {
	if counterFile == nil {
		return
	}

	counterFileLock.Lock()
	defer counterFileLock.Unlock()

	var err error

	if _, err = counterFile.Seek(0, 0); err == nil {
		if _, err = fmt.Fprint(counterFile, counter); err == nil {
			err = counterFile.Sync()
		}
	}
	if err != nil {
		// Complain, but doesn't seem worth bailing at this point.
		log.Print("Failed to write counter:", err)
	}
}

func handleSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill, os.Interrupt)
	<-sigChan
	saveCounter()
	os.Exit(0)
}

func defaultAddr() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}

	return ":8080"
}
