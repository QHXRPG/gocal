package calcultools

import (
	"first_GoProject/niqe/typetrans"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
)

func Add2DFloat32(a [][]float32, b [][]float32) [][]float32 {
	if len(a) != len(b) && len(a[0]) != len(b[0]) {
		panic("两个数组大小不等")
	}
	c := typetrans.Make2DZeroFloat32(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func Add2DFloat32Int(a [][]float32, b int) [][]float32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			a[i][j] = a[i][j] + float32(b)
		}
	}
	return a
}

func Div2DFloat32(a [][]float32, b [][]float32) [][]float32 {
	if len(a) != len(b) && len(a[0]) != len(b[0]) {
		panic("两个数组大小不等")
	}
	c := typetrans.Make2DZeroFloat32(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			c[i][j] = a[i][j] / b[i][j]
		}
	}
	return c
}

func Div2DFloatFloat(a [][]float32, b float32) [][]float32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			a[i][j] = a[i][j] / b
		}
	}
	return a
}

func Sub2DFloat32(a [][]float32, b [][]float32) [][]float32 {
	if len(a) != len(b) && len(a[0]) != len(b[0]) {
		panic("两个数组大小不等")
	}
	c := typetrans.Make2DZeroFloat32(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

func Mul2DFloat32(matrix1 [][]float32, matrix2 [][]float32) [][]float32 {
	numRows1 := len(matrix1)
	numCols1 := len(matrix1[0])
	numRows2 := len(matrix2)
	numCols2 := len(matrix2[0])

	if numCols1 != numRows2 {
		// 矩阵无法相乘，返回空矩阵
		return [][]float32{}
	}

	result := make([][]float32, numRows1)
	for i := range result {
		result[i] = make([]float32, numCols2)
	}

	for i := 0; i < numRows1; i++ {
		for j := 0; j < numCols2; j++ {
			for k := 0; k < numCols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

func Sub1DFloat32Float32(a []float32, b float32) []float32 {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + b
	}
	return a
}

func Sub_2DFloat32_1DFloat(a [][]float32, b []float32) [][]float32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = a[i][j] - b[i]
		}
	}
	return a
}

func Square2DFloat32(a [][]float32) [][]float32 {
	b := typetrans.Make2DZeroFloat32(len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			b[i][j] = a[i][j] * a[i][j]
		}
	}
	return b
}

func Square1DFloat32(a []float32) []float32 {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
	}
	return a
}

func Abs2DFloat32(a [][]float32) [][]float32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] < 0 {
				a[i][j] = 0 - a[i][j]
			}
		}
	}
	return a
}

func Abs1DFloat32(a []float32) []float32 {
	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			a[i] = 0 - a[i]
		}
	}
	return a
}

func Sqrt2DFloat32(a [][]float32) [][]float32 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			a[i][j] = float32(math.Sqrt(float64(a[i][j])))
		}
	}
	return a
}

func Average1DFloat32(a []float32) float32 {
	var sum float32 = 0
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	sum = sum / float32(len(a))
	return sum
}

// ArgMin1DFloat32_ ArgMin1DFloat32 找出数组中最小的值以及索引
func ArgMin1DFloat32_(a []float32) (int, float32) {
	var temp int = 0
	var flog float32
	for i := 0; i < len(a); i++ {
		if i == 0 {
			flog = a[i]
		} else {
			if flog < a[i] {
			} else {
				flog = a[i]
				temp = i
			}
		}
	}
	return temp, flog
}

// Roll2DFloat 矩阵的平移
func Roll2DFloat(a [][]float32, shift int, axis int) [][]float32 {
	h, w := len(a), len(a[0])
	b := typetrans.Make2DZeroFloat32(h, w)
	var k int
	if axis == 1 { //列移动
		if shift >= 0 { // 右移
			for i := 0; i < h; i++ {
				k = shift
				for j := 0; j < w; j++ {
					if k != 0 {
						b[i][j] = a[i][w-k]
						k--
					} else {
						b[i][j] = a[i][j-shift]
					}
				}
			}
		} else { //左移
			shift = 0 - shift
			for i := 0; i < h; i++ {
				k = shift
				for j := 0; j < w; j++ {
					if k != 0 {
						b[i][w-k] = a[i][j]
						k--
					} else {
						b[i][j-shift] = a[i][j]
					}
				}
			}
		}
	}
	if axis == 0 { //行移动
		if shift >= 0 { // 下移
			k = shift
			for i := 0; i < h; i++ {
				if k != 0 {
					b[i] = a[h-k]
					k--
				} else {
					b[i] = a[i-shift]
				}
			}
		} else { //上移
			shift = 0 - shift
			k = shift
			for i := 0; i < h; i++ {
				if k != 0 {
					b[h-k] = a[i]
					k--
				} else {
					b[i-shift] = a[i]
				}
			}
		}
	}
	return b
}

// Hstack2DFloat32 两个矩阵水平方向堆叠
func Hstack2DFloat32(a [][]float32, b [][]float32) [][]float32 {
	for i := 0; i < len(a); i++ {
		a[i] = append(a[i], b[i]...)
	}
	return a
}

// Mean2DFloat32 计算每行/列的平均值
func Mean2DFloat32(a [][]float32, axis int) []float32 {
	var result []float32
	var sum float32 = 0
	if axis != 0 && axis != 1 {
		panic("请检查axis输入的值")
	}
	if axis == 1 {
		result = make([]float32, len(a))
		sum = 0
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				sum = sum + a[i][j]
			}
			result[i] = sum / float32(len(a[0]))
			sum = 0
		}
	}
	if axis == 0 {
		result = make([]float32, len(a[0]))
		sum = 0
		for i := 0; i < len(a[0]); i++ {
			for j := 0; j < len(a); j++ {
				sum = sum + a[j][i]
			}
			result[i] = sum / float32(len(a))
			sum = 0
		}
	}
	return result
}

// TransposeFloat32 Transpose 转置
func TransposeFloat32(matrix [][]float32) [][]float32 {
	rows := len(matrix)
	cols := len(matrix[0])
	// 创建一个新的矩阵，行列互换
	transposed := make([][]float32, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]float32, rows)
	}
	// 将原矩阵的元素复制到新矩阵的对应位置
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

// Float32ToNumMat 2d的float32转换为 mat.Dense 类型
func Float32ToNumMat(data [][]float32) *mat.Dense {
	// 转换为 mat.Dense 类型
	rows, cols := len(data), len(data[0])
	flatData := make([]float64, rows*cols)
	for i, row := range data {
		for j, val := range row {
			flatData[i*cols+j] = float64(val)
		}
	}
	dense := mat.NewDense(rows, cols, flatData)
	return dense
}

func NumMatToFloat32(dense mat.Dense) [][]float32 {
	rows, cols := dense.Dims()
	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			data[i][j] = dense.At(i, j)
		}
	}
	result := typetrans.Float64ToFloat32Plus(data)
	return result
}

func CovarianceFloat32(a [][]float32) [][]float32 {
	aMean := Mean2DFloat32(a, 1)
	aTmp := Sub_2DFloat32_1DFloat(a, aMean)
	dense := Float32ToNumMat(aTmp)
	denseT := mat.Transpose{Matrix: dense}

	h, _ := dense.Dims()
	_, w := denseT.Dims()

	covMat := mat.NewDense(h, w, nil)
	covMat.Mul(dense, denseT)
	a = Div2DFloatFloat(NumMatToFloat32(*covMat), float32(len(aTmp[0])-1))
	return a
}

// NoZeroReciprocal 对非零元素求倒数
func NoZeroReciprocal(a []float64) []float64 {
	for i := 0; i < len(a); i++ {
		if a[i] > 0.00001 || a[i] < -0.00001 {
			a[i] = 1 / a[i]
		}
	}
	return a
}

func SVDMatrixProcess(h int, c []float64) [][]float64 {
	a := typetrans.Make2DZeroFloat64(h, len(c))
	for i := 0; i < len(c); i++ {
		a[i][i] = c[i]
	}
	return a
}

/*
PinvFloat32
SVD分解求伪逆
原理和公式：1. SVD分解得到的矩阵:U和V是正交阵，S是对角阵
 2. 正交阵的逆=转置
 3. 对角阵的逆=非零元素求倒

Step1: 求解A的SVD分解

	[U,S,V] = svd(A); % A = U*S*V'

Step2: 将S中的非零元素求倒

	T=S;
	T(find(S~=0)) = 1./S(find(S~=0));

Step3: 求invA  svdInvA = V * T' * U';
*/
func PinvFloat32(data [][]float32) [][]float32 {
	h, w := len(data), len(data[0])
	svdInvA := mat.NewDense(w, h, nil)
	a := mat.NewDense(h, w, typetrans.SqueezeFloat64(typetrans.Float32ToFloat64Plus(data))) //原始矩阵
	//Apinv := mat.NewDense(h, w, nil)                                    //伪逆矩阵

	var svd mat.SVD
	ok := svd.Factorize(a, mat.SVDFull)
	if !ok {
		log.Fatal("failed to factorize A")
	}

	const rcond = 1e-15
	rank := svd.Rank(rcond)
	if rank == 0 {
		log.Fatal("zero rank system")
	}

	s, u, vt := svd.GetValues()

	//sDense := mat.NewDense(h, w, SqueezeFloat64(SVDMatrixProcess(h, s)))
	tDense := mat.NewDense(h, w, NoZeroReciprocal(typetrans.SqueezeFloat64(SVDMatrixProcess(h, s))))
	uDense := mat.NewDense(h, h, u)
	vtDense := mat.NewDense(w, w, vt)

	svdInvA.Mul(vtDense, tDense.T())
	svdInvA.Mul(svdInvA, uDense.T())

	result := NumMatToFloat32(*svdInvA)
	return result
}
