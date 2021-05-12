package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	ft "github.com/x-cray/logrus-prefixed-formatter"

	"go-boiler-plate/factory"
	"go-boiler-plate/router"
)

var version = "0.0.0"

func main() {
	const Port = 8000
	l := logrus.New()
	l.Level = logrus.DebugLevel
	l.Formatter = &ft.TextFormatter{
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	f := factory.NewFactory("")

	l.Infof("Running server version: %s", version)
	n := negroni.New()
	n.UseHandler(router.NewCustomRouter(f, l))
	n.Run(fmt.Sprintf(":%d", Port))
}
