package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCashContext_GetMoney(t *testing.T) {
	context := NewCashContext(HalfOff)
	assert.Equal(t, 50.0, context.GetMoney(100))

	otherContext := NewCashContext(HundredDiscountTwenty)
	assert.Equal(t, 80.0, otherContext.GetMoney(100))

	normalContext := NewCashContext(0)
	assert.Equal(t, 100.0, normalContext.GetMoney(100))
}