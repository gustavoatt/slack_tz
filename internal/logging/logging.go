package logging

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	loggerContextKey = "logger"
)

func SetupLogging(router *gin.Engine) {
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
	log.Logger = log.Output(output)

	router.Use(func(context *gin.Context) {
		t := time.Now()
		context.Set(loggerContextKey,
			log.Logger.With().Str("uri", context.Request.RequestURI).Logger())
		context.Next()

		// after request
		latency := time.Since(t)
		status := context.Writer.Status()
		log.Info().
			Str("method", context.Request.Method).
			Str("uri", context.Request.RequestURI).
			Int("status", status).
			Dur("latency", latency).
			Msg("Received Request")
	})
	router.Use(gin.Recovery())
}

func Logger(ctx *gin.Context) *zerolog.Logger {
	l := ctx.MustGet(loggerContextKey).(zerolog.Logger)
	return &l
}

func Info(ctx *gin.Context) *zerolog.Event {
	l := ctx.MustGet(loggerContextKey).(zerolog.Logger)
	return l.Info()
}

func Warn(ctx *gin.Context) *zerolog.Event {
	l := ctx.MustGet(loggerContextKey).(zerolog.Logger)
	return l.Warn()
}

func Error(ctx *gin.Context) *zerolog.Event {
	l := ctx.MustGet(loggerContextKey).(zerolog.Logger)
	return l.Error()
}
