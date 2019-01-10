package dsp

import "testing"

func TestSignalDeviation(t *testing.T) {

	expected := Signal(5)

	signal := Signal(25)

	actual := signal.Deviation()

	if expected != actual {
		t.Error("expected ", expected, " actual ", actual)
	}
}
