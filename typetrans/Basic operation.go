package typetrans

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"strconv"
)

func Make2DZeroFloat32(h int, w int) [][]float32 {
	slice := make([][]float32, h) // 创建一个长度为 h 的切片
	for i := 0; i < h; i++ {
		slice[i] = make([]float32, w) // 每个切片元素初始化为长度为 w 的 float32 切片
		for j := 0; j < w; j++ {
			slice[i][j] = 0 // 将每个元素赋值为 0.0
		}
	}
	return slice
}

func Make2DZeroFloat64(h int, w int) [][]float64 {
	slice := make([][]float64, h) // 创建一个长度为 h 的切片
	for i := 0; i < h; i++ {
		slice[i] = make([]float64, w) // 每个切片元素初始化为长度为 w 的 float32 切片
		for j := 0; j < w; j++ {
			slice[i][j] = 0 // 将每个元素赋值为 0.0
		}
	}
	return slice
}

func Uint8ToBytes(arr []uint8) []byte {
	return arr
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

func Float64ToFloat32Plus(a [][]float64) [][]float32 {
	b := Make2DZeroFloat32(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[i][j] = float32(a[i][j])
		}
	}
	return b
}
func Float32ToFloat64Plus(a [][]float32) [][]float64 {
	b := Make2DZeroFloat64(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[i][j] = float64(a[i][j])
		}
	}
	return b
}

func Uint8ToMat(array [][]uint8) gocv.Mat {
	w := len(array[0])
	h := len(array)
	array2 := SqueezeUint8(array)
	bytes := Uint8ToBytes(array2)
	mat, err2 := gocv.NewMatFromBytes(h, w, gocv.MatTypeCV8U, bytes)
	if err2 != nil {
		fmt.Println("转换出错：", err2)
		panic("转换出错")
	}
	return mat
}

// BytesToMat 1d的bytes转化为Mat
func BytesToMat(rows int, cols int, mt gocv.MatType, bytes []byte) gocv.Mat {
	mat, err2 := gocv.NewMatFromBytes(rows, cols, mt, bytes) // mt = MatTypeCV8UC1
	if err2 != nil {
		fmt.Println("转换出错：", err2)
		panic("转换出错")
	}
	return mat
}

// SqueezeFloat64 将2d数组降为1d
func SqueezeFloat64(a [][]float64) []float64 {
	b := make([]float64, len(a)*len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[j+len(a[0])*i] = a[i][j]
		}
	}
	return b
}

// UnSqueezeFloat32 将1d数组升到2d
func UnSqueezeFloat32(a []float32) [][]float32 {
	b := make([][]float32, len(a))
	for i := 0; i < len(a); i++ {
		c := []float32{a[i]}
		b[i] = c
	}
	return b
}

func SqueezeFloat32(a [][]float32) []float32 {
	b := make([]float32, len(a)*len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[j+len(a[0])*i] = a[i][j]
		}
	}
	return b
}

func SqueezeUint8(arr [][]uint8) []uint8 {
	var result []uint8
	for _, row := range arr {
		for _, val := range row {
			result = append(result, val)
		}
	}
	return result
}

// LoadCsvToFloat64 将Csv读取为Float64格式的2d矩阵
func LoadCsvToFloat64(path string) [][]float64 {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("读取模型出错：", err)
		panic("读取模型出错")
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("读取模型出错：", err)
		panic("读取模型出错")
	}
	floatData := make([][]float64, len(preData))
	// 遍历原始数据并进行转换
	for i, row := range preData {
		floatData[i] = make([]float64, len(row))
		for j, val := range row {
			// 将字符串转换为float64
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("无法转换为float64：", err)
				panic("无法转换为float64")
			}
			floatData[i][j] = num
		}
	}
	return floatData
}

func Uint8ToFloat32(array [][]uint8) [][]float32 {
	height := len(array)
	width := len(array[0])

	result := make([][]float32, height)
	for i := 0; i < height; i++ {
		result[i] = make([]float32, width)
		for j := 0; j < width; j++ {
			result[i][j] = float32(int(array[i][j]))
		}
	}

	return result
}
