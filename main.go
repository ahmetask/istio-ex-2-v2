package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999999Z",
	})
	logger := logrus.StandardLogger()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/b", func(c echo.Context) error {
		logger.Log(logrus.DebugLevel, "service-response-v2-2")
		return c.JSON(http.StatusOK, "service-response2-v2")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
