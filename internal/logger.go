package internal

import (
	"log"
	"os"
)

func SetupLogger() {
	logFile, err := os.OpenFile("./fuckregex.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Logger initialized")
}
