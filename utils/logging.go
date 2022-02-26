package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string)  {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// ログの出力先を指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 
	log.SetOutput(multiLogFile)
}