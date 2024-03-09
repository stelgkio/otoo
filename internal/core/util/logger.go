package util

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	//	logger, _ := zap.NewDevelopment()
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	defer logger.Sync()
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(msg, info string) {
	l.logger.Info(msg,
		zap.String("time", getTime()),
		zap.String("info", info),
	)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Error(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Sync() {
	l.logger.Sync()
}

func getTime() string {
	return time.Now().Format(time.RFC3339Nano)
}
func (l *Logger) LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogError:     true,
		LogRequestID: true,
		HandleError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				l.logger.Info("request",
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
					zap.Duration("latency", v.Latency),
					zap.String("request_id", v.RequestID),
				)
			} else {
				l.logger.Error("request error",
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
					zap.Duration("latency", v.Latency),
					zap.String("request_id", v.RequestID),
					//zap.String("error", v.Error.Error()),
				)
			}
			return nil
		},
	})
}

//LogURI:      true,
//		LogStatus:   true,
//		LogError:    true,
//		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
//		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
//			if v.Error == nil {
//				logger.Info("request",
//					zap.String("URI", v.URI),
//					zap.Int("status", v.Status),
//				)
//			} else {
//				logger.Error("request error",
//					zap.String("URI", v.URI),
//					zap.Int("status", v.Status),
//					zap.Error(v.Error),
//				)
//			}
//			return nil
//		},
//	}))
//
