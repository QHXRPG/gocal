# gocal
基于Go语言的一些算法实现

## calcultools
提供各种矩阵和数组的基本运算功能。

### 矩阵运算
- **Add2DFloat32**: 二维矩阵加法
- **Add2DFloat32Int**: 二维矩阵与整型加法
- **Div2DFloat32**: 二维矩阵除法
- **Div2DFloatFloat**: 二维矩阵与浮点型除法
- **Sub2DFloat32**: 二维矩阵减法
- **Mul2DFloat32**: 矩阵乘法
- **Sub1DFloat32Float32**: 二维矩阵与浮点型减法
- **Sub_2DFloat32_1DFloat**: 二维矩阵与一维矩阵减法
- **Square2DFloat32**: 二维矩阵平方
- **Square1DFloat32**: 一维矩阵平方
- **Abs2DFloat32**: 二维矩阵绝对值
- **Abs1DFloat32**: 一维矩阵绝对值
- **Sqrt2DFloat32**: 二维矩阵开平方根

### 数学运算
- **Average1DFloat32**: 一维矩阵求平均值
- **ArgMin1DFloat32_**: 找出数组中最小的值以及索引
- **CovarianceFloat32**: 求协方差矩阵
- **NoZeroReciprocal**: 对非零元素求倒数

### 矩阵变换
- **Roll2DFloat**: 矩阵的平移
- **Hstack2DFloat32**: 两个矩阵水平方向堆叠
- **Mean2DFloat32**: 计算每行/列的平均值
- **TransposeFloat32**: 矩阵转置
- **Float32ToNumMat**: 2D的float32转换为mat.Dense类型
- **NumMatToFloat32**: mat.Dense转float32
- **SVDMatrixProcess**: SVD矩阵处理
- **PinvFloat32**: SVD分解求伪逆矩阵

## imagetools
提供图像处理相关的功能。

### 图像读取与显示
- **GetImageInfo**: 读取图像信息
- **ShowImage**: 小窗口显示图像数据Mat
- **ReadImage**: 读取图像为Mat

### 图像转换
- **Float64ToMat**: float64转化为Mat
- **GrayImageMatToArray**: 灰度图转2D矩阵

### 图像处理
- **CheckImageSize**: 检查图片是否过小
- **CropImage**: 图像裁剪
- **GenGaussWindow**: 定义高斯模糊
- **Correlate1dFloat32**: 一维卷积
- **ImageMSCNTransform**: 图像MSCN变换

### 特征提取
- **AggdFeatures**: 提取AGGD特征
- **PiredProduct**: 提取配对乘积
- **NiqeExtractSubbandFeats**: 提取NIQE子带特征
- **ExtractOnPatches**: 在图像块上提取特征

## typetrans
提供类型转换相关的功能。

### 矩阵创建
- **Make2DZeroFloat32**: 创建二维零矩阵
- **Make2DZeroFloat64**: 创建二维零矩阵

### 类型转换
- **Uint8ToBytes**: uint8转bytes
- **Float64ToFloat32**: 数组float64转float32
- **Float64ToFloat32Plus**: 二维矩阵float64转float32
- **Float32ToFloat64Plus**: 二维矩阵float32转float64
- **Uint8ToMat**: uint8转gocv.Mat
- **BytesToMat**: 1D的bytes转化为Mat
- **Uint8ToFloat32**: uint8转float32

### 数组维度变换
- **SqueezeFloat64**: 将2D数组降为1D
- **SqueezeFloat32**: 将2D数组降为1D
- **SqueezeUint8**: 将2D数组降为1D
- **UnSqueezeFloat32**: 将1D数组升到2D

### CSV读取
- **LoadCsvToFloat64**: 将CSV读取为Float64格式的2D矩阵

## 使用示例

以下是一些简单的使用示例，展示如何使用`gocal`库中的部分功能。

