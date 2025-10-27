package logger

import (
	"log"
	"os"
)

// Init configures the standard logger to output to stdout with date, time, and file name.
func Init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
