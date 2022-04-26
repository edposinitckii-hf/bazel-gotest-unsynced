package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	l, err := NewLogger()
	assert.NoError(t, err)
	assert.IsType(t, &zap.Logger{}, l)
}
