package logx

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var GlobalLogger zerolog.Logger

func Init(config Config) {
	var w io.Writer
	switch config.Env {
	case "test", "dev", "development":
		w = &zerolog.ConsoleWriter{
			Out:         os.Stdout,
			TimeFormat:  time.RFC3339,
			FormatLevel: nil, // use default
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("%s", i)
			},
			FormatFieldName: func(i interface{}) string {
				return fmt.Sprintf("[%s=", i)
			},
			FormatFieldValue: func(i interface{}) string {
				return fmt.Sprintf("%s]", i)
			},
		}
	case "prod", "production":
		w = os.Stdout
	}

	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		fmt.Printf("failed to parse log level in config file, err:%v", err)
		os.Exit(1)
	}

	GlobalLogger = zerolog.New(w).With().Str("service", config.ServiceName).Timestamp().Caller().Logger().Level(level)
}

func Fatalf(format string, v ...interface{}) {
	GlobalLogger.Fatal().Msgf(format, v...)
}

func FatalIfErr(err error) {
	if err != nil {
		GlobalLogger.Fatal().Msg(err.Error())
	}
}

func FatalIfErrf(err error, format string, v ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(format, v...)
		GlobalLogger.Fatal().Msgf(msg+", err:%v", err)
	}
}

func Errorf(format string, v ...interface{}) {
	GlobalLogger.Error().Msgf(format, v...)
}

func Infof(format string, v ...interface{}) {
	GlobalLogger.Info().Msgf(format, v...)
}

func Debug(format string, v ...interface{}) {
	GlobalLogger.Debug().Msgf(format, v...)
}
