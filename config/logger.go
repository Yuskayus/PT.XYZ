package config

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Log adalah variabel global untuk logging
var Log = logrus.New()

func SetupLogger() {
	//Set format log
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	//Set output ke file jika diperlukan
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
		Log.Warn("Gagal membuka file log, menggunakan stdout sebagai default")
	}

	//Set level logging
	Log.SetLevel(logrus.InfoLevel)
}
