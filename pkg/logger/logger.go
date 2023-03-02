// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Info(msg string, filelds ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

type LogLevel = zapcore.Level

const (
	InfoLevel  LogLevel = zap.InfoLevel
	WarnLevel  LogLevel = zap.WarnLevel
	ErrorLevel LogLevel = zap.ErrorLevel
	PanicLevel LogLevel = zap.PanicLevel
	FatalLevel LogLevel = zap.FatalLevel
	DebugLevel LogLevel = zap.DebugLevel
)

type Field = zap.Field

type Logger struct {
	logger   *zap.Logger
	logLevel LogLevel
}

type TeeOption struct {
	W        io.Writer
	logLevel LogLevel
}

func NewLoggerWithTee(options []TeeOption) (ILogger, error) {
	if len(options) == 0 {
		return nil, fmt.Errorf("options is nil")
	}

	var (
		zapCore zapcore.Core
		cores   []zapcore.Core
	)

	cfg := zap.NewProductionConfig()

	for _, option := range options {
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(option.W),
			zapcore.Level(option.logLevel),
		)
		cores = append(cores, core)
	}

	zapCore = zapcore.NewTee(cores...)

	logger := &Logger{
		logger: zap.New(zapCore),
	}

	return logger, nil
}

func NewLogger(logLevel LogLevel, writer io.Writer) (ILogger, error) {
	cfg := zap.NewProductionConfig()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(logLevel),
	)

	logger := &Logger{
		logger:   zap.New(core),
		logLevel: logLevel,
	}

	return logger, nil
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}
