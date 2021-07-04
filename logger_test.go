package lika_logger

import (
	"./targets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger_AddTarget(t *testing.T) {

	var target targets.TargetInterface

	target = &targets.PrintLogTarget{}
	logger := Logger{}

	logger.AddTarget(&target)
	logger.AddTarget(&target)

	assert.Equal(t, len(logger.Targets), 2)
}
