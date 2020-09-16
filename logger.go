package lika_logger

import funk "github.com/thoas/go-funk"

type Logger struct {
	FlushInterval int
	Targets       []TargetInterface
}

func (l *Logger) Log(message interface{}, extraData map[string]interface{}, level string, category string, except []int) {
	for k, target := range l.Targets {
		if !funk.Contains(except, k) && target.canLog(level, category) {
			err := target.Log(message, extraData, level, category)

			if err != nil {
				except = append(except, k)
				logErrorMessage := map[string]interface{}{
					"errorMessage": "Error on trying to log message",
					"error":        err,
					"message":      message,
				}

				l.Log(logErrorMessage, extraData, level, category, except)
			}
		}
	}
}

func (l *Logger) AddTarget(target *TargetInterface) {
	l.Targets = append(l.Targets, *target)
}
