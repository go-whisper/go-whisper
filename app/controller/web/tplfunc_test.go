package web

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testNow time.Time

func testInit() {
	testNow = time.Date(2021, 4, 1, 23, 59, 59, 0, time.Local)
}

func TestTimeYear(t *testing.T) {
	testInit()
	assert.Equal(t, TimeYear(testNow), 2021)
}

func TestTimeMonth(t *testing.T) {
	testInit()
	assert.Equal(t, TimeMonth(testNow), "April")
}

func TestTimeDay(t *testing.T) {
	testInit()
	assert.Equal(t, TimeDay(testNow), 1)
}
