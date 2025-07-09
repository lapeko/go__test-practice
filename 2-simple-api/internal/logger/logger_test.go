package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	require.NotNil(t, Log)
	assert.Equal(t, Log.GetLevel(), logrus.DebugLevel)
	assert.IsType(t, Log.Out, os.Stdout)
	_, ok := Log.Formatter.(*logrus.JSONFormatter)
	assert.True(t, ok)
}
