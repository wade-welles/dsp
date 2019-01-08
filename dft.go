package dsp

import "math"

// DFT Discrete Fourier Transform
type DFT struct {
	Rex Sample
	Imx Sample
	Len int
}

const pi2 = math.Pi * 2

func (dft DFT) length() int {
	i, j := len(dft.Rex), len(dft.Imx)

	if j > i {
		return j
	}
	return i
}

// Magnitude returns the magnitude of the samples
func (dft DFT) Magnitude() Sample {

	len := dft.length()

	output := make(Sample, len)

	for i := 0; i < len; i++ {
		val := math.Sqrt(math.Pow(float64(dft.Rex[i]), 2) + math.Pow(float64(dft.Imx[i]), 2))
		output[i] = Signal(val)
	}

	return output
}

// Inverse return the inverse sample of the discrete fourier transform
func (dft DFT) Inverse() Sample {

	size := dft.length()

	output := make(Sample, dft.Len)

	for i := 0; i < size; i++ {
		dft.Rex[i] = dft.Rex[i] / Signal(size)
		dft.Imx[i] = -dft.Imx[i] / Signal(size)
	}

	dft.Rex[0] = dft.Rex[0] / 2
	dft.Imx[0] = -dft.Imx[0] / 2

	for i := 0; i < dft.Len; i++ {
		for j := 0; j < size; j++ {
			val := float64(i * j)
			len := float64(dft.Len)
			output[i] += dft.Rex[j] * Signal(math.Cos(pi2*val/len))
			output[i] += dft.Imx[j] * Signal(math.Sin(pi2*val/len))
		}
	}

	return output
}

// Polar converts from Rect to Polar
func (dft DFT) Polar() DFT {

	len := dft.length()

	mag := make(Sample, len)
	phase := make(Sample, len)

	for i := 0; i < len; i++ {

		mag[i] = Signal(math.Sqrt(math.Pow(float64(dft.Rex[i]), 2) + math.Pow(float64(dft.Imx[i]), 2)))

		if dft.Rex[i].Equals(0) {
			dft.Rex[i] = Signal(math.Pow(10, -20))
			phase[i] = Signal(math.Atan(float64(dft.Imx[i]) / float64(dft.Rex[i])))
		}

		if dft.Rex[i] < 0 && dft.Imx[i] < 0 {
			phase[i] -= Signal(math.Pi)
		}
		if dft.Rex[i] < 0 && dft.Imx[i] >= 0 {
			phase[i] += Signal(math.Pi)
		}
	}

	return DFT{mag, phase, dft.Len}
}
