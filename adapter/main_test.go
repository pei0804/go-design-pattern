package adpter

import "testing"

func TestAdpter(t *testing.T) {
	bc := NewPrintBannerComposition("A")
	if bc.getStrong() != "*A*" {
		t.Errorf("Expect get %s, actual %s", "*A*", bc.getStrong())
	}
	be := NewPrintBannerEmbedding("A")
	if be.getWeek() != "(A)" {
		t.Errorf("Expect get %s, actual %s", "(A)", be.getWeek())
	}
}
