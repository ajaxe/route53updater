package logging

import (
	"log"
	"os"
)

// DBGLogger : Logs with prefix "Dbg"
var DBGLogger = log.New(os.Stdout, "[Dbg]", log.Ldate|log.Ltime|log.Lshortfile)
