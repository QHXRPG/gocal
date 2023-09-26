# gocal
# 基于go的一些算法实现

## calcultools
### Add2DFloat32 二维矩阵加法
### Add2DFloat32Int 二维矩阵与整型加法
### Div2DFloat32 二维矩阵除法
### Div2DFloatFloat 二维矩阵与浮点型除法
### Sub2DFloat32 二维矩阵减法
### Mul2DFloat32 矩阵乘法
### Sub1DFloat32Float32 二维矩阵与浮点型减法
### Sub_2DFloat32_1DFloat 二维矩阵与一维矩阵减法
### Square2DFloat32 二维矩阵平方
### Square1DFloat32 一维矩阵平方
### Abs2DFloat32 二维矩阵绝对值
### Abs1DFloat32 一维矩阵绝对值
### Sqrt2DFloat32 二维矩阵开平方根
### Average1DFloat32 一维矩阵求平均值
### ArgMin1DFloat32_ 找出数组中最小的值以及索引
### Roll2DFloat 矩阵的平移
### Hstack2DFloat32 两个矩阵水平方向堆叠
### Mean2DFloat32 计算每行/列的平均值
### TransposeFloat32 Transpose 转置
### Float32ToNumMat 2d的float32转换为 mat.Dense 类型
### NumMatToFloat32 mat.Dense转float32
### CovarianceFloat32 求协方差矩阵
### NoZeroReciprocal 对非零元素求倒数
### SVDMatrixProcess
### PinvFloat32 SVD分解求伪逆矩阵

## imagetools
### GetImageInfo 读取图像信息
### ShowImage 小窗口显示图像数据Mat
### ReadImage 读取图像为Mat
### Float64ToMat float64转化为Mat
### CheckImageSize 检查图片是否过小
### GrayImageMatToArray 灰度图转2d矩阵
### CropImage 图像裁剪
### GenGaussWindow 定义高斯模糊
### Correlate1dFloat32 一维卷积
### ImageMSCNTransform 图像MSCN变换
### AggdFeatures
### PiredProduct
### NiqeExtractSubbandFeats
### ExtractOnPatches

## typetrans
### Make2DZeroFloat32创建二维矩阵
### Make2DZeroFloat64
### Uint8ToBytes uint8转bytes
### Float64ToFloat32 数组float64转float32
### Float64ToFloat32Plus 二维矩阵float64转float32
### Float32ToFloat64Plus二维矩阵float32转float64
### Uint8ToMat uint8转gocv.Mat
### BytesToMat 1d的bytes转化为Mat
### SqueezeFloat64 将2d数组降为1d
### UnSqueezeFloat32 将1d数组升到2d
### SqueezeFloat32 将2d数组降为1d
### SqueezeUint8 将2d数组降为1d
### LoadCsvToFloat64 将Csv读取为Float64格式的2d矩阵
### Uint8ToFloat32 uint8转float32


