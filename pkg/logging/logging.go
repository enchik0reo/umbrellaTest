package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fName := path.Base(f.File)
			function = fmt.Sprintf("%s()", f.Function)
			file = fmt.Sprintf("%s:%d", fName, f.Line)
			return function, file
		},
		FullTimestamp: true,
	}

	if err := os.MkdirAll("logs", 0744); err != nil {
		panic(err)
	}

	fl, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&Hooker{
		Writer:    []io.Writer{fl, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	e = logrus.NewEntry(l)
}

var e *logrus.Entry

type Lgr struct {
	*logrus.Entry
}

func InitLogger() *Lgr {
	return &Lgr{
		Entry: e,
	}
}

type Hooker struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (h *Hooker) Levels() []logrus.Level {
	return h.LogLevels
}

func (h *Hooker) Fire(entry *logrus.Entry) error {
	l, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range h.Writer {
		_, err = w.Write([]byte(l))
	}
	return err
}
