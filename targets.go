package lika_logger

import (
	"fmt"
)
import funk "github.com/thoas/go-funk"

type TargetInterface interface {
	Log(message interface{}, extraData map[string]interface{}, level string, category string) error
	canLog(level string, category string) bool
	SetLevels(levels []string)
	SetCategories(categories []string)
}

type BaseLogTarget struct {
	Levels     []string
	Categories []string
}

func (t *BaseLogTarget) SetLevels(levels []string) {
	t.Levels = levels
}

func (t *BaseLogTarget) SetCategories(categories []string) {
	t.Categories = categories
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

func (l *PrintLogTarget) Log(message interface{}, extraData map[string]interface{}, level string, category string) error {
	fmt.Printf("Call log message in PrintLogTarget \n message: %v \n extraData %s  \n level %v \n category %v \n",
		message, extraData, level, category,
	)

	return nil
}

func CreateTarget(driver string, levels []string, categories []string) *TargetInterface {
	var target TargetInterface

	switch driver {
	case "printer":
		target = &PrintLogTarget{}
	default:
		target = &PrintLogTarget{}
	}

	target.SetCategories(categories)
	target.SetLevels(levels)

	return &target
}
