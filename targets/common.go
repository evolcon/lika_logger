package targets

import (
	"github.com/thoas/go-funk"
)

type TargetInterface interface {
	Log(message interface{}, extraData map[string]interface{}, level string, category string) error
	CanLog(level string, category string) bool
}

type BaseLogTarget struct {
	Levels     []string
	Categories []string
}

func (t *BaseLogTarget) CanLog(level string, category string) bool {
	if len(t.Levels) > 0 && !funk.Contains(t.Levels, level) {
		return false
	}

	if len(t.Categories) > 0 && !funk.Contains(t.Categories, category) {
		return false
	}

	return true
}
