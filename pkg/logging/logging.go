package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writers   []io.Writer
	logLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()

	if err != nil {
		return err
	}

	for _, write := range hook.Writers {
		write.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.logLevels
}

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{l}
}

var l *logrus.Entry

func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}
	log.SetFormatter(formatter)
	log.SetOutput(io.Discard)

	log.AddHook(&writerHook{
		Writers:   []io.Writer{os.Stdout},
		logLevels: logrus.AllLevels,
	})

	log.SetLevel(logrus.TraceLevel)
	l = logrus.NewEntry(log)
}
