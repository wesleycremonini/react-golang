package main

import (
	"runtime/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	err := run()
	if err != nil {
		zap.L().Fatal(err.Error(), zap.ByteString("debug_stack", debug.Stack()))
	}
}

type application struct {}

func run() error {
	logger, err := newLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	app := &application{}

	zap.L().Info("server started")
	err = app.server().ListenAndServe()
	zap.L().Info("server stopped")
	if err != nil {
		return err
	}

	return nil
}

func newLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.FunctionKey = "func"
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig.LevelKey = zapcore.OmitKey
	config.EncoderConfig.LineEnding = "\n\n"
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}
