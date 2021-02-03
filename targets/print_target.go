package targets

import "fmt"

type PrintLogTarget struct {
	BaseLogTarget
}

func (l *PrintLogTarget) Log(message interface{}, extraData map[string]interface{}, level string, category string) error {
	fmt.Printf("Call log message in PrintLogTarget \n message: %v \n extraData %s  \n level %v \n category %v \n",
		message, extraData, level, category,
	)

	return nil
}
