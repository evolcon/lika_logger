package lika_logger

import "testing"
import "github.com/stretchr/testify/assert"

func TestLogger_AddTarget(t *testing.T) {

	var target TargetInterface

	target = &PrintLogTarget{}
	logger := Logger{}

	logger.AddTarget(&target)
	logger.AddTarget(&target)

	assert.Equal(t, len(logger.Targets), 2)
}
