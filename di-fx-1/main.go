package main

import (
	"fmt"
	"log"

	"go.uber.org/fx"
)

type CustomString string

type MyAwesomeLogger struct {
	Prefix string
}

// LoggerInterface interfaces logging
type LoggerInterface interface {
	LogMe(str string) string
}

func (lg *MyAwesomeLogger) LogMe(data string) string {
	return fmt.Sprintf("%v: %v\n", lg.Prefix, data)
}

func ProvideString() CustomString {
	var str CustomString = "hello world"
	return str
}

func ProvideLogger() LoggerInterface {
	return &MyAwesomeLogger{Prefix: "Konichiwa"}
}

func RunMe(logger LoggerInterface, str CustomString) {
	logData := logger.LogMe("Hello Dependency Injection Works")
	fmt.Println(logData)
	fmt.Println("Received string from DI", str)
}

func main() {
	// logger := ProvideLogger()

	// str := ProvideString()
	// RunMe(logger, str)

	log.Println("Starting Application")
	fx.New(
		fx.Provide(ProvideString()),
		fx.Provide(ProvideLogger()),
		fx.Invoke(RunMe),
	).Run()

}
