package lika_logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

type FileLogTarget struct {
	BaseLogTarget
	FilePath string
}

func (t *FileLogTarget) Log(message interface{}, extraData map[string]interface{}, level string, category string) error {
	if err := createFileDir(t.FilePath); err != nil {
		log.Fatalf("error creating directory: %v", err)

		return err
	}

	f, err := openFile(t.FilePath)

	if err != nil {
		log.Fatalf("error opening file: %v", err)

		return err
	}

	defer f.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println(message)

	return nil
}

func openFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	return f, err
}

func createFileDir(filePath string) error {
	fileDir := filepath.Dir(filePath)

	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}
