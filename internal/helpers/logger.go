package helpers

import (
    "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
    Logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    Logger.SetLevel(logrus.InfoLevel)
}
