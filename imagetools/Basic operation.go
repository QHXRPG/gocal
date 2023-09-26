package imagetools

import (
	"encoding/json"
	"first_GoProject/niqe/calcultools"
	"first_GoProject/niqe/global"
	"first_GoProject/niqe/typetrans"
	"fmt"
	"gocv.io/x/gocv"
	"math"
)

// GetImageInfo 读取图像信息
func GetImageInfo(img gocv.Mat) (a, b, c int) {
	return img.Cols(), img.Rows(), img.Channels()
}

// ShowImage 小窗口显示图像数据Mat
func ShowImage(img gocv.Mat) {
	window := gocv.NewWindow("hello")
	imgk := img
	for {
		window.IMShow(imgk)
		window.WaitKey(1)
	}
}

// ReadImage 读取图像为Mat
func ReadImage(path string) gocv.Mat {
	img := gocv.IMRead(path, gocv.IMReadGrayScale) // 读取灰度图
	if img.Empty() {
		fmt.Println("无法读取图片")
		panic("无法读取图片")
	}
	return img
}

// Float64ToMat float64转化为Mat
func Float64ToMat(w int, h int, list []float64) gocv.Mat {
	bytes, err1 := json.Marshal(list)
	if err1 != nil {
		fmt.Println("转换出错：", err1)
		panic("转换出错")
	}
	mat, err2 := gocv.NewMatFromBytes(h, w, gocv.MatTypeCV64F, bytes)
	if err2 != nil {
		fmt.Println("转换出错：", err2)
		panic("转换出错")
	}
	return mat
}

func Float64ToFloat32(a []float64) []float32 {
	b := make([]float32, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = float32(a[i])
	}
	return b
}

// CheckImageSize 检查图片是否过小
func CheckImageSize(mat gocv.Mat, patchSize int) {
	wight, high, _ := GetImageInfo(mat)
	if !(wight > patchSize*2+1) {
		panic("图像宽过小")
	}
	if !(high > patchSize*2+1) {
		panic("图像高过小")
	}
}

// GrayImageMatToArray 灰度图转2d矩阵
func GrayImageMatToArray(width int, height int, img gocv.Mat) [][]uint8 {
	// 初始化二维切片
	matrix := make([][]uint8, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([]uint8, width)
		for j := 0; j < width; j++ {
			matrix[i][j] = img.GetVecbAt(i, j)[0]
		}
	}
	return matrix
}

// CropImage 图像裁剪
func CropImage(img [][]uint8, hoffset int, woffset int) [][]uint8 {
	height := len(img)
	width := len(img[0])
	if hoffset > 0 {
		img = img[:height-hoffset]
	}
	if woffset > 0 {
		for i := 0; i < height-hoffset; i++ {
			img[i] = img[i][:width-woffset]
		}
	}
	return img
}

func GenGaussWindow(lw int, sigma float32) []float32 {
	k := 2*lw + 1
	var sum float32 = 1
	weights := make([]float32, k)
	for i := 0; i < k; i++ {
		weights[i] = 0
	}
	weights[lw] = 1
	sigma = sigma * sigma
	for ii := 1; ii <= lw; ii++ {
		tmp := float32(math.Exp(float64(-0.5 * float32(ii*ii) / sigma)))
		weights[lw+ii] = tmp
		weights[lw-ii] = tmp
		sum = sum + tmp*2
	}
	for ii := 0; ii < 2*lw+1; ii++ {
		weights[ii] = weights[ii] / sum
	}
	return weights
}

func Correlate1dFloat32(matrix1 [][]float32, avgWindow []float32, axis int) [][]float32 {
	var sum float32 = 0
	Matrix1Height := len(matrix1)
	Matrix1Width := len(matrix1[0])
	lenAvgwindow := len(avgWindow)
	matrix2 := typetrans.Make2DZeroFloat32(Matrix1Height, Matrix1Width)
	if axis == 1 {
		for i := 0; i < Matrix1Height; i++ {
			for j := 0; j < Matrix1Width; j++ {
				for k := -lenAvgwindow / 2; k < lenAvgwindow/2+1; k++ {
					if j+k >= 0 && j+k < Matrix1Width {
						sum = sum + matrix1[i][j+k]*avgWindow[k+lenAvgwindow/2]
					}
				}
				matrix2[i][j] = sum
				sum = 0
			}
		}
	}
	sum = 0
	if axis == 0 {
		for i := 0; i < Matrix1Width; i++ {
			for j := 0; j < Matrix1Height; j++ {
				for k := -lenAvgwindow / 2; k < lenAvgwindow/2+1; k++ {
					if j+k >= 0 && j+k < Matrix1Height {
						sum = sum + matrix1[j+k][i]*avgWindow[k+lenAvgwindow/2]
					}
				}
				matrix2[j][i] = sum
				sum = 0
			}
		}
	}
	return matrix2
}

func ImageMSCNTransform(img [][]float32) ([][]float32, [][]float32, [][]float32) {
	avgWindow := GenGaussWindow(3, float32(7)/float32(6))
	h := len(img)
	w := len(img[0])
	muImage := typetrans.Make2DZeroFloat32(h, w)
	varImage := typetrans.Make2DZeroFloat32(h, w)
	muImage = Correlate1dFloat32(img, avgWindow, 0)
	muImage = Correlate1dFloat32(muImage, avgWindow, 1)
	varImage = Correlate1dFloat32(img, avgWindow, 0)
	varImage = Correlate1dFloat32(varImage, avgWindow, 1)
	varImage = calcultools.Sqrt2DFloat32(calcultools.Abs2DFloat32(calcultools.Sub2DFloat32(varImage, calcultools.Square2DFloat32(muImage))))
	result := calcultools.Div2DFloat32(calcultools.Sub2DFloat32(img, muImage), calcultools.Add2DFloat32Int(varImage, 1))
	return result, varImage, muImage
}

func AggdFeatures(imdata [][]float32) (float32, float32, float32, float32, float32, float32) {
	imdata2, imdata3 := typetrans.SqueezeFloat32(imdata), typetrans.SqueezeFloat32(imdata)
	leftIndex := make([]int, 0)
	rightIndex := make([]int, 0)
	for i := 0; i < len(imdata2); i++ {
		if imdata2[i] < 0 {
			leftIndex = append(leftIndex, i)
		}
		if imdata2[i] >= 0 {
		}
		rightIndex = append(rightIndex, i)
	}
	imdata2 = calcultools.Square1DFloat32(imdata2)
	leftData := make([]float32, 0)
	rightData := make([]float32, 0)
	for i := 0; i < len(leftIndex); i++ {
		leftData = append(leftData, imdata2[leftIndex[i]])
	}
	for i := 0; i < len(rightIndex); i++ {
		rightData = append(rightData, imdata2[rightIndex[i]])
	}
	var leftMeanSqrt float32 = 0
	var rightMeanSqrt float32 = 0
	var gammaHat float32
	var rHat float32
	if len(leftData) > 0 {
		leftMeanSqrt = float32(math.Sqrt(float64(calcultools.Average1DFloat32(leftData))))
	}
	if len(rightData) > 0 {
		rightMeanSqrt = float32(math.Sqrt(float64(calcultools.Average1DFloat32(rightData))))
	}
	if rightMeanSqrt != 0 {
		gammaHat = leftMeanSqrt / rightMeanSqrt
	} else {
		gammaHat = 9999999
	}
	imdata2Mean := calcultools.Average1DFloat32(imdata2)
	if imdata2Mean != 0 {
		rHat = calcultools.Average1DFloat32(calcultools.Square1DFloat32(calcultools.Abs1DFloat32(imdata3)))
	} else {
		rHat = 9999999
	}

	rHatNorm := rHat * float32(((math.Pow(float64(gammaHat), 3)+1)*(float64(gammaHat)+1))/math.Pow(math.Pow(float64(gammaHat), 2)+1, 2))
	pos, _ := calcultools.ArgMin1DFloat32_(calcultools.Square1DFloat32(calcultools.Sub1DFloat32Float32(Float64ToFloat32(global.PrecGammas), rHatNorm)))
	Alpha := global.GammaRange[pos]

	gam1 := math.Gamma(1 / Alpha)
	gam2 := math.Gamma(1 / Alpha)
	gam3 := math.Gamma(1 / Alpha)

	aggDratio := math.Sqrt(gam1) / math.Sqrt(gam3)
	bl := float32(aggDratio) * leftMeanSqrt
	br := float32(aggDratio) * rightMeanSqrt

	N := (br - bl) * float32(gam2/gam1)

	return float32(Alpha), N, bl, br, leftMeanSqrt, rightMeanSqrt
}

func PiredProduct(img [][]float32) ([][]float32, [][]float32, [][]float32, [][]float32) {
	shift1 := calcultools.Roll2DFloat(img, 1, 1)
	shift2 := calcultools.Roll2DFloat(img, 1, 0)
	shift3 := calcultools.Roll2DFloat(calcultools.Roll2DFloat(img, 1, 0), 1, 1)
	shift4 := calcultools.Roll2DFloat(calcultools.Roll2DFloat(img, 1, 0), -1, 1)
	Himg := calcultools.Mul2DFloat32(shift1, img)
	Vimg := calcultools.Mul2DFloat32(shift2, img)
	D1Img := calcultools.Mul2DFloat32(shift3, img)
	D2Img := calcultools.Mul2DFloat32(shift4, img)
	return Himg, Vimg, D1Img, D2Img
}

func NiqeExtractSubbandFeats(mscncoefs [][]float32) []float32 {
	alpha_m, _, bl, br, _, _ := AggdFeatures(mscncoefs)
	pps1, pps2, pps3, pps4 := PiredProduct(mscncoefs)
	alpha1, N1, bl1, br1, _, _ := AggdFeatures(pps1)
	alpha2, N2, bl2, br2, _, _ := AggdFeatures(pps2)
	alpha3, N3, bl3, _, _, _ := AggdFeatures(pps3)
	alpha4, N4, bl4, _, _, _ := AggdFeatures(pps4)
	result := []float32{alpha_m, (bl + br) / 2.0, alpha1, N1, bl1, br1, alpha2, N2, bl2, br2, alpha3, N3, bl3, bl3, alpha4, N4, bl4, bl4}
	return result

}

func ExtractOnPatches(img [][]float32, patchSize int) [][]float32 {
	index := 0
	h, w := len(img), len(img[0])
	patches := make([][][]float32, h/patchSize*(w/patchSize))
	for j := 0; j < h-patchSize+1; j = j + patchSize {
		for i := 0; i < w-patchSize+1; i = i + patchSize {
			for n := 0; n < patchSize; n++ {
				patches[index] = append(patches[index], img[j : j+patchSize][n][i:i+patchSize])
			}
			index++
		}
	}
	patchFeatures := make([][]float32, len(patches))
	for i := 0; i < len(patches); i++ {
		patchFeatures[i] = NiqeExtractSubbandFeats(patches[i])
	}
	return patchFeatures
}
