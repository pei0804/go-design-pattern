package prototype

import "testing"

func TestPrototype(t *testing.T) {
	messageBox := MessageBox{"A"}
	manager := Manager{}
	manager.Register("A", &messageBox)

	cloned := manager.Create("A")

	if messageBox.Use() != cloned.Use() {
		t.Errorf("Expect instance to equal, but not equal.")
	}

	underline := Underline{"B"}
	manager.Register("B", &underline)
	cloned = manager.Create("B")
	if underline.Use() != cloned.Use() {
		t.Errorf("Expect instance to equal, but not equal.")
	}
}
