package logger

import (
	"log"
)

type Clogger struct {
	verbose bool
}

func NewLogger(verbose bool) *Clogger {
	if !verbose {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	}
	return &Clogger{verbose: verbose}
}

func (c *Clogger) Info(format string, args ...interface{}) {
	if c.verbose {
		log.Printf(format, args)
	}
}

func (c *Clogger) Warn(format string, args ...interface{}) {
	if c.verbose {
		log.Printf(format, args)
	}
}

func (c *Clogger) Debug(format string, args ...interface{}) {
	if c.verbose {
		log.Printf(format, args)
	}
}

func (c *Clogger) Fatal(format string, args ...interface{}) {
	if c.verbose {
		log.Fatalf(format, args)
	}
}
