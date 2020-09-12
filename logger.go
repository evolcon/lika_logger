package lika_logger

type Logger struct {
	FlushInterval int
	Targets       []TargetInterface
}

func (l *Logger) Log(message interface{}, extraData map[string]interface{}, level string, category string) {
	for _, target := range l.Targets {
		if target.canLog(level, category) {
			target.Log(message, extraData, level, category)
		}
	}
}

func (l *Logger) AddTarget(target *TargetInterface) {
	l.Targets = append(l.Targets, *target)
}
