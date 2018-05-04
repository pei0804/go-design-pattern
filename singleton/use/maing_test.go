package singleton

import (
	"testing"

	"github.com/pei0804/go-design-pattern/singleton"
)

func TestSingleton(t *testing.T) {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expect instance to equal, but not equal.\n")
	}
}
