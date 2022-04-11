package logging

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
)

type sequenceHook struct{}

var sequence = new(uint64)

func (hook sequenceHook) Fire(entry *logrus.Entry) error {
	entry.Data["sequence"] = atomic.AddUint64(sequence, 1)
	return nil
}

func (hook sequenceHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}
