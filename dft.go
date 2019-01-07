package dsp

import "math"

// DFT Discrete Fourier Transform
type DFT struct {
	Rex Sample
	Imx Sample
}

// Len the len of the samples
func (dft DFT) Len() int {
	return int(math.Max(float64(len(dft.Rex)), float64(len(dft.Imx))))
}

// Magnitude returns the magnitude of the samples
func (dft DFT) Magnitude() Sample {

	len := dft.Len()

	output := make(Sample, len)

	for i := 0; i < len; i++ {
		output[i] = Signal(math.Sqrt(math.Pow(float64(dft.Rex[i]), 2) + math.Pow(float64(dft.Imx[i]), 2)))
	}
	return output
}
