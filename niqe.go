package main

import (
	cal "first_GoProject/niqe/calcultools"
	"first_GoProject/niqe/global"
	imgs "first_GoProject/niqe/imagetools"
	"first_GoProject/niqe/typetrans"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"math"
)

func main() {
	start := 0.2
	end := 10.0
	step := 0.001

	for i := start; i < end-step; i += step {
		global.GammaRange = append(global.GammaRange, i)
	}
	fmt.Println(len(global.GammaRange))
	var a []float64
	var b []float64
	var c []float64

	for _, val := range global.GammaRange {
		a = append(a, math.Pow(math.Gamma(2.0/val), 2))
		b = append(b, math.Gamma(1.0/val))
		c = append(c, math.Gamma(3.0/val))
	}

	for i := range global.GammaRange {
		global.PrecGammas = append(global.PrecGammas, a[i]/(b[i]*c[i]))
	}
	patchSize := 96
	//
	popMu := typetrans.Float64ToFloat32Plus(typetrans.LoadCsvToFloat64("/Users/qiuhaoxuan/GolandProjects/first_GoProject/niqe/data/pop_mu.csv"))
	popCov := typetrans.Float64ToFloat32Plus(typetrans.LoadCsvToFloat64("/Users/qiuhaoxuan/GolandProjects/first_GoProject/niqe/data/pop_cov.csv"))

	img := imgs.ReadImage("/Users/qiuhaoxuan/Desktop/截图/iShot_2023-09-26_18.18.15.jpg")
	h, w, _ := imgs.GetImageInfo(img)
	hoffset := h % patchSize
	woffset := w % patchSize

	matrix := imgs.GrayImageMatToArray(img.Cols(), img.Rows(), img)
	matrix = imgs.CropImage(matrix, hoffset, woffset)
	img = typetrans.Uint8ToMat(matrix)
	img2 := gocv.NewMat()

	gocv.Resize(img, &img2, image.Point{X: int(math.Floor(float64(h) / 2)), Y: int(math.Floor(float64(w) / 2))}, 0, 0, gocv.InterpolationArea)

	matrix2 := typetrans.Uint8ToFloat32(imgs.GrayImageMatToArray(img2.Cols(), img2.Rows(), img2)) //cols表示宽度，rows表示高度
	matrix1 := typetrans.Uint8ToFloat32(matrix)

	factor := math.Pow(float64(512)/float64((len(matrix1)+len(matrix1))/2), 2)

	mscn1, _, _ := imgs.ImageMSCNTransform(matrix1)
	mscn2, _, _ := imgs.ImageMSCNTransform(matrix2)

	featsLvl1 := imgs.ExtractOnPatches(mscn1, patchSize)
	featsLvl2 := imgs.ExtractOnPatches(mscn2, patchSize/2)

	feats := cal.Hstack2DFloat32(featsLvl1, featsLvl2)

	sampleMu := cal.Mean2DFloat32(feats, 0)
	sampleMu2 := typetrans.UnSqueezeFloat32(sampleMu)
	sampleCov := cal.CovarianceFloat32(cal.TransposeFloat32(feats))

	X := cal.Sub2DFloat32(sampleMu2, popMu)

	covMat := cal.Div2DFloatFloat(cal.Add2DFloat32(popCov, sampleCov), 2)
	pinvMat := cal.PinvFloat32(covMat)

	x := cal.TransposeFloat32(X)
	step1 := cal.Mul2DFloat32(x, pinvMat)
	step2 := typetrans.SqueezeFloat32(step1)
	xx := typetrans.SqueezeFloat32(X)
	var result float32 = 0
	for i := 0; i < len(xx); i++ {
		result = result + step2[i]*xx[i]
	}
	if result < 0 {
		result = 0 - result
	}
	fmt.Println(math.Sqrt(float64(result)) * float64(factor))
}
