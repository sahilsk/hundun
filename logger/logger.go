package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type Clogger struct {
	verbose bool
}

func NewLogger(verbose bool) *Clogger {
	if verbose {
		//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	}
	return &Clogger{verbose: verbose}
}

func (c *Clogger) Info(format string, args ...interface{}) {
	if c.verbose {
		_, file, no, _ := runtime.Caller(1)
		appendedString := fmt.Sprintf("[INFO] [%s/%s:%d] %s", filepath.Base(filepath.Dir(file)), filepath.Base(file), no, format)
		log.Printf(appendedString, args)
	}
}

func (c *Clogger) Warn(format string, args ...interface{}) {
	if c.verbose {
		_, file, no, _ := runtime.Caller(1)
		appendedString := fmt.Sprintf("[WARN] [%s/%s:%d] %s", filepath.Base(filepath.Dir(file)), filepath.Base(file), no, format)
		log.Printf(appendedString, args)
	}
}

func (c *Clogger) Debug(format string, args ...interface{}) {
	if c.verbose {
		_, file, no, _ := runtime.Caller(1)
		appendedString := fmt.Sprintf("[DEBUG] [%s/%s:%d] %s", filepath.Base(filepath.Dir(file)), filepath.Base(file), no, format)
		log.Printf(appendedString, args)
	}
}

func (c *Clogger) Fatal(format string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	appendedString := fmt.Sprintf("[FATAL] [%s/%s:%d] %s", filepath.Base(filepath.Dir(file)), filepath.Base(file), no, format)
	log.Fatalf(appendedString, args)

}
