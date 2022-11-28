package discount

import (
	"testing"
)

func TestDiscount_One(t *testing.T) {
	result := discount(10000, 10)
	if result != 1000 {
		t.Errorf("test failed")
	}
}
