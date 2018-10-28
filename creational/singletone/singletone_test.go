package singletone

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to Singleton not nil")
	}

	count := counter1.AddOne()

	if count != 1 {
		t.Error("Count after first calling AddOne() must be 1")
	}

	counter2 := GetInstance()
	count = counter2.AddOne()
	if count != 2 {
		t.Error("Count value must be 2")
	}
}
