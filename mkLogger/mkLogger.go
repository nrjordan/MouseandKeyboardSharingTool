package mkLogger

import (
	"log"
	"os"
	"sync"
)

type mkLogger struct {
	*log.Logger
	filename string
}

var mLogger *mkLogger
var once sync.Once

func GetInstance() *mkLogger {
	once.Do(func() {
		mLogger = createLogger("mkLogger.log")
	})
	return mLogger
}

func createLogger(fname string) *mkLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &mkLogger{
		filename: fname,
		Logger: log.New(file, "MK ", log.Lshortfile),
	}
}