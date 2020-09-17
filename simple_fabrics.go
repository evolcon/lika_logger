package lika_logger

func CreatePrintLogTarget(levels []string, categories []string) *TargetInterface {
	var target TargetInterface

	target = &PrintLogTarget{
		BaseLogTarget{
			Levels: levels, Categories: categories,
		},
	}

	return &target
}

func CreateFileLogTarget(FilePath string, levels []string, categories []string) *TargetInterface {
	var target TargetInterface

	target = &FileLogTarget{
		BaseLogTarget{
			Levels: levels, Categories: categories,
		},
		FilePath,
	}

	return &target
}
