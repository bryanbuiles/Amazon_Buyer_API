package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO:", 0)
	errorLogger = log.New(os.Stdout, "", log.Lshortfile|log.Ltime) // Lshortfile final file name element and line number
	// Ltime display the time
)

// Info To display Info in stdout
func Info(s interface{}) {
	infoLogger.Println(s)
}

// Error log to display the error in stdout
func Error(s interface{}) {
	errorLogger.Println(s)
}
