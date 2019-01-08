package dsp

import (
	"errors"
	"math"
)

// DFT Discrete Fourier Transform
type dft struct {
	rex Sample
	imx Sample
}

const pi2 = math.Pi * 2

// NewDFT Initializes a new DFT
func NewDFT(rex Sample, imx Sample) (*dft, error) {

	if len(rex) != len(imx) {
		return nil, errors.New("invalid argument")
	}

	return &dft{rex, imx}, nil
}

// Equals compares two DFT instances
func (dft *dft) Equals(other *dft) bool {
	return dft.rex.Equals(other.rex) && dft.imx.Equals(other.imx)
}

func (dft *dft) length() int {
	return len(dft.rex)
}

func (dft *dft) size() int {
	return dft.length() * 2
}

// Magnitude returns the magnitude of the samples
func (dft *dft) Magnitude() Sample {

	len := dft.length()

	output := make(Sample, len)

	for i := 0; i < len; i++ {
		val := math.Sqrt(math.Pow(float64(dft.rex[i]), 2) + math.Pow(float64(dft.imx[i]), 2))
		output[i] = Signal(val)
	}

	return output
}

// Inverse return the inverse sample of the discrete fourier transform
func (dft *dft) Inverse() Sample {

	length := dft.length()
	size := dft.size()
	fsize := float64(size)

	output := make(Sample, size)

	for i := 0; i < length; i++ {
		dft.rex[i] = dft.rex[i] / Signal(length)
		dft.imx[i] = -dft.imx[i] / Signal(length)
	}

	dft.rex[0] = dft.rex[0] / 2
	dft.imx[0] = -dft.imx[0] / 2

	for i := 0; i < size; i++ {
		for j := 0; j < length; j++ {
			val := float64(i * j)
			output[i] += dft.rex[j] * Signal(math.Cos(pi2*val/fsize))
			output[i] += dft.imx[j] * Signal(math.Sin(pi2*val/fsize))
		}
	}

	return output
}

// Polar converts from Rect to Polar
func (dft *dft) Polar() (*dft, error) {

	len := dft.length()

	mag := make(Sample, len)
	phase := make(Sample, len)

	for i := 0; i < len; i++ {

		mag[i] = Signal(math.Sqrt(math.Pow(float64(dft.rex[i]), 2) + math.Pow(float64(dft.imx[i]), 2)))

		if dft.rex[i].Equals(0) {
			dft.rex[i] = Signal(math.Pow(10, -20))
			phase[i] = Signal(math.Atan(float64(dft.imx[i]) / float64(dft.rex[i])))
		}

		if dft.rex[i] < 0 && dft.imx[i] < 0 {
			phase[i] -= Signal(math.Pi)
		}
		if dft.rex[i] < 0 && dft.imx[i] >= 0 {
			phase[i] += Signal(math.Pi)
		}
	}

	return NewDFT(mag, phase)
}
