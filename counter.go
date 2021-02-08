package main

import (
	"flag"
	"os"
	"sync"
)

var (
	// Counts number of passwords generated.
	counter     uint64
	counterLock sync.Mutex // Overkill?

	// Optional file to load/save counter value.
	counterFilePath = flag.String("counter", "", "password counter file")
	counterFile     *os.File
	counterFileLock sync.Mutex
)
