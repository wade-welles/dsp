package dsp

import (
	"fmt"
	"testing"
)

const (
	impulseResponse       = "impulse_response.dat"
	input15khzImx         = "input_15khz_imx.dat"
	input15khzRex         = "input_15khz_rex.dat"
	input15khzSignal      = "input_15khz_signal.dat"
	input20khzImx         = "input_20khz_imx.dat"
	input20khzRex         = "input_20khz_rex.dat"
	input20khzSignal      = "input_20khz_signal.dat"
	inputEcgImx           = "input_ecg_imx.dat"
	inputEcgRex           = "input_ecg_rex.dat"
	inputEcgSignal        = "input_ecg_signal.dat"
	output15khzMagnitude  = "output_magnitude.dat"
	outputConvolution     = "output_convolution.dat"
	outputDftImx          = "output_dft_imx.dat"
	outputDftRex          = "output_dft_rex.dat"
	outputEcgMagnitude    = "output_ecg_magnitude.dat"
	outputEcgPhase        = "output_ecg_phase.dat"
	outputFirstDifference = "output_first_difference.dat"
	outputIdft            = "output_idft.dat"
	outputIdftImx         = "output_idft_imx.dat"
	outputIdftRex         = "output_idft_rex.dat"
	output15khzImx        = "output_imx.dat"
	output15khzPhase      = "output_phase.dat"
	output15khzRex        = "output_rex.dat"
	outputRunningSum      = "output_running_sum.dat"
)

func loadTestFile(filename string) Sample {

	val, err := LoadSample(fmt.Sprintf("testdata/%s", filename))
	if err != nil {
		panic(fmt.Sprintf("unable to parse %s: %v", filename, err))
	}
	return val
}

var testSamples = map[string]Sample{
	impulseResponse:       loadTestFile(impulseResponse),
	input15khzImx:         loadTestFile(input15khzImx),
	input15khzRex:         loadTestFile(input15khzRex),
	input15khzSignal:      loadTestFile(input15khzSignal),
	input20khzImx:         loadTestFile(input20khzImx),
	input20khzRex:         loadTestFile(input20khzRex),
	input20khzSignal:      loadTestFile(input20khzSignal),
	inputEcgImx:           loadTestFile(inputEcgImx),
	inputEcgRex:           loadTestFile(inputEcgRex),
	inputEcgSignal:        loadTestFile(inputEcgSignal),
	output15khzMagnitude:  loadTestFile(output15khzMagnitude),
	outputConvolution:     loadTestFile(outputConvolution),
	outputDftImx:          loadTestFile(outputDftImx),
	outputDftRex:          loadTestFile(outputDftRex),
	outputEcgMagnitude:    loadTestFile(outputEcgMagnitude),
	outputEcgPhase:        loadTestFile(outputEcgPhase),
	outputFirstDifference: loadTestFile(outputFirstDifference),
	outputIdft:            loadTestFile(outputIdft),
	outputIdftImx:         loadTestFile(outputIdftImx),
	outputIdftRex:         loadTestFile(outputIdftRex),
	output15khzImx:        loadTestFile(output15khzImx),
	output15khzPhase:      loadTestFile(output15khzPhase),
	output15khzRex:        loadTestFile(output15khzRex),
	outputRunningSum:      loadTestFile(outputRunningSum),
}

func TestSampleMean(t *testing.T) {

	expected := Signal(0.037112)

	actual := testSamples[input15khzSignal].Mean()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleDeviation(t *testing.T) {
	expected := Signal(0.787502)

	actual := testSamples[input15khzSignal].Deviation()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleRunningSum(t *testing.T) {
	expected := testSamples[outputRunningSum]

	actual := testSamples[input15khzSignal].RunningSum()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleConvolution(t *testing.T) {

	expected := testSamples[outputConvolution]

	actual := testSamples[input15khzSignal].Convolution(testSamples[impulseResponse])

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleFirstDifference(t *testing.T) {

	expected := testSamples[outputFirstDifference]

	actual := testSamples[input15khzSignal].FirstDifference()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}

}

func TestSampleDFT(t *testing.T) {

	expected := DFT{
		testSamples[outputDftRex],
		testSamples[outputDftImx],
		len(testSamples[input15khzSignal]),
	}

	actual := testSamples[input15khzSignal].DFT()

	if !actual.Rex.Equals(expected.Rex) {
		t.Error("expected ", expected, " actual ", actual)
	}

	if !actual.Imx.Equals(expected.Imx) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleMagnitude(t *testing.T) {

	expected := testSamples[output15khzMagnitude]

	dft := testSamples[input15khzSignal].DFT()

	actual := dft.Magnitude()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestSampleInverse(t *testing.T) {

	expected := testSamples[outputIdft]

	dft := testSamples[input15khzSignal].DFT()

	actual := dft.Inverse()

	if !actual.Equals(expected) {
		t.Error("expected ", expected, " actual ", actual)
	}
}

func TestDFTPolar(t *testing.T) {

	expected := DFT{
		testSamples[outputDftRex],
		testSamples[outputDftImx],
		len(testSamples[input15khzSignal]),
	}

	dft := testSamples[input15khzSignal].DFT()

	actual := dft.Polar()

	if !actual.Rex.Equals(expected.Rex) {
		t.Error("expected ", expected.Rex, " actual ", actual.Rex)
	}

	if !actual.Imx.Equals(expected.Imx) {
		t.Error("expected ", expected.Imx, " actual ", actual.Imx)
	}
}
