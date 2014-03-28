package mc

import (
	"log"
	"os"
	"strings"
)

var (
	dlog  = log.New(os.Stderr, "mc: ", log.Lmicroseconds|log.Lshortfile)
	debug = strings.Contains(os.Getenv("STTRACE"), "mc")
)
