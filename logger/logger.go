package logger

import (
	"log"
)

/**
 * @var		type	CLogge
 * @global
 */
type CLogger log.Logger

/**
 * InitLogger.
 *
 * @author	Unknown
 * @since	v0.0.1
 * @version	v1.0.0	Sunday, January 5th, 2020.
 * @global
 * @return	void
 */
func (l *CLogger) InitLogger() {
	CLogger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

/**
 * @var		l	*log.Logge
 * @global
 */
func (l *CLogger) LInfo(v interface{}) {
	log.Print(v)
}

/**
 * @var		l	*log.Logge
 * @global
 */
func (l *CLogger) LInfof(format string, v ...interface{}) {
	log.Printf(format, v)
}
