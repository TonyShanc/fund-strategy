package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWeekend(t *testing.T) {

	// Sunday
	assert.True(t, isWeekend(time.Date(2021, time.December, 26, 0, 0, 0, 0, time.UTC).Add(time.Hour * 8)))

	// Saturday
	assert.True(t, isWeekend(time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC).Add(time.Hour * 8)))

	// Friday
	assert.False(t, isWeekend(time.Date(2021, time.December, 24, 0, 0, 0, 0, time.UTC).Add(time.Hour * 8)))
}
