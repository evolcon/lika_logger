package lika_logger

import "fmt"
import funk "github.com/thoas/go-funk"

type TargetInterface interface {
	Log(message interface{}, extraData map[string]interface{}, level string, category string)
	canLog(level string, category string) bool
}

type BaseLogTarget struct {
	Levels     []string
	Categories []string
}

func (t *BaseLogTarget) canLog(level string, category string) bool {
	if len(t.Levels) > 0 && !funk.Contains(t.Levels, level) {
		return false
	}

	if len(t.Categories) > 0 && !funk.Contains(t.Categories, category) {
		return false
	}

	return true
}

type PrintLogTarget struct {
	BaseLogTarget
}

func (l *PrintLogTarget) Log(message interface{}, extraData map[string]interface{}, level string, category string) {
	fmt.Printf("Call log message in PrintLogTarget \n message: %v \n extraData %s  \n level %v \n category %v \n",
		message, extraData, level, category,
	)
}
