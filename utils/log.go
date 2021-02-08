package utils

import (
	"github.com/getsentry/sentry-go"
	"log"
)

//CreateLog ...
func CreateLog(error string) {
	dsn := GetEnv("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	})

	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln("sentry.Init: %s", err)
	}

	sentry.CaptureMessage(error)
}
