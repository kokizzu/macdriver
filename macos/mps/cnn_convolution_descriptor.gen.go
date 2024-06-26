// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNConvolutionDescriptor] class.
var CNNConvolutionDescriptorClass = _CNNConvolutionDescriptorClass{objc.GetClass("MPSCNNConvolutionDescriptor")}

type _CNNConvolutionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionDescriptor] class.
type ICNNConvolutionDescriptor interface {
	objc.IObject
	SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean *float32, variance *float32, gamma *float32, beta *float32, epsilon float32)
	EncodeWithCoder(aCoder foundation.ICoder)
	KernelWidth() uint
	SetKernelWidth(value uint)
	OutputFeatureChannels() uint
	SetOutputFeatureChannels(value uint)
	Neuron() CNNNeuron
	SetNeuron(value ICNNNeuron)
	InputFeatureChannels() uint
	SetInputFeatureChannels(value uint)
	StrideInPixelsY() uint
	SetStrideInPixelsY(value uint)
	StrideInPixelsX() uint
	SetStrideInPixelsX(value uint)
	Groups() uint
	SetGroups(value uint)
	KernelHeight() uint
	SetKernelHeight(value uint)
	DilationRateX() uint
	SetDilationRateX(value uint)
	FusedNeuronDescriptor() NNNeuronDescriptor
	SetFusedNeuronDescriptor(value INNNeuronDescriptor)
	DilationRateY() uint
	SetDilationRateY(value uint)
}

// A description of the attributes of a convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor?language=objc
type CNNConvolutionDescriptor struct {
	objc.Object
}

func CNNConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNConvolutionDescriptor {
	return CNNConvolutionDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](cc, objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648813-cnnconvolutiondescriptorwithkern?language=objc
func CNNConvolutionDescriptor_CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNConvolutionDescriptor {
	return CNNConvolutionDescriptorClass.CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
}

func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter ICNNNeuron) CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](cc, objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return rv
}

// Creates a convolution descriptor with an optional neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648876-cnnconvolutiondescriptorwithkern?language=objc
func CNNConvolutionDescriptor_CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter ICNNNeuron) CNNConvolutionDescriptor {
	return CNNConvolutionDescriptorClass.CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
}

func (cc _CNNConvolutionDescriptorClass) Alloc() CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNConvolutionDescriptorClass) New() CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionDescriptor() CNNConvolutionDescriptor {
	return CNNConvolutionDescriptorClass.New()
}

func (c_ CNNConvolutionDescriptor) Init() CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867057-setbatchnormalizationparametersf?language=objc
func (c_ CNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean *float32, variance *float32, gamma *float32, beta *float32, epsilon float32) {
	objc.Call[objc.Void](c_, objc.Sel("setBatchNormalizationParametersForInferenceWithMean:variance:gamma:beta:epsilon:"), mean, variance, gamma, beta, epsilon)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2866972-encodewithcoder?language=objc
func (c_ CNNConvolutionDescriptor) EncodeWithCoder(aCoder foundation.ICoder) {
	objc.Call[objc.Void](c_, objc.Sel("encodeWithCoder:"), aCoder)
}

// The width of the kernel window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth?language=objc
func (c_ CNNConvolutionDescriptor) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

// The width of the kernel window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth?language=objc
func (c_ CNNConvolutionDescriptor) SetKernelWidth(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelWidth:"), value)
}

// The number of feature channels per pixel in the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels?language=objc
func (c_ CNNConvolutionDescriptor) OutputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("outputFeatureChannels"))
	return rv
}

// The number of feature channels per pixel in the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels?language=objc
func (c_ CNNConvolutionDescriptor) SetOutputFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setOutputFeatureChannels:"), value)
}

// The neuron filter to be applied as part of the convolution operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron?language=objc
func (c_ CNNConvolutionDescriptor) Neuron() CNNNeuron {
	rv := objc.Call[CNNNeuron](c_, objc.Sel("neuron"))
	return rv
}

// The neuron filter to be applied as part of the convolution operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron?language=objc
func (c_ CNNConvolutionDescriptor) SetNeuron(value ICNNNeuron) {
	objc.Call[objc.Void](c_, objc.Sel("setNeuron:"), value)
}

// The number of feature channels per pixel in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels?language=objc
func (c_ CNNConvolutionDescriptor) InputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("inputFeatureChannels"))
	return rv
}

// The number of feature channels per pixel in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels?language=objc
func (c_ CNNConvolutionDescriptor) SetInputFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setInputFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding?language=objc
func (cc _CNNConvolutionDescriptorClass) SupportsSecureCoding() bool {
	rv := objc.Call[bool](cc, objc.Sel("supportsSecureCoding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding?language=objc
func CNNConvolutionDescriptor_SupportsSecureCoding() bool {
	return CNNConvolutionDescriptorClass.SupportsSecureCoding()
}

// The output stride (downsampling factor) in the y dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy?language=objc
func (c_ CNNConvolutionDescriptor) StrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsY"))
	return rv
}

// The output stride (downsampling factor) in the y dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy?language=objc
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInPixelsY:"), value)
}

// The output stride (downsampling factor) in the x dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx?language=objc
func (c_ CNNConvolutionDescriptor) StrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsX"))
	return rv
}

// The output stride (downsampling factor) in the x dimension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx?language=objc
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInPixelsX:"), value)
}

// The number of groups that the input and output channels are divided into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups?language=objc
func (c_ CNNConvolutionDescriptor) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}

// The number of groups that the input and output channels are divided into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups?language=objc
func (c_ CNNConvolutionDescriptor) SetGroups(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setGroups:"), value)
}

// The height of the kernel window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight?language=objc
func (c_ CNNConvolutionDescriptor) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

// The height of the kernel window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight?language=objc
func (c_ CNNConvolutionDescriptor) SetKernelHeight(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex?language=objc
func (c_ CNNConvolutionDescriptor) DilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex?language=objc
func (c_ CNNConvolutionDescriptor) SetDilationRateX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor?language=objc
func (c_ CNNConvolutionDescriptor) FusedNeuronDescriptor() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](c_, objc.Sel("fusedNeuronDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor?language=objc
func (c_ CNNConvolutionDescriptor) SetFusedNeuronDescriptor(value INNNeuronDescriptor) {
	objc.Call[objc.Void](c_, objc.Sel("setFusedNeuronDescriptor:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey?language=objc
func (c_ CNNConvolutionDescriptor) DilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey?language=objc
func (c_ CNNConvolutionDescriptor) SetDilationRateY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateY:"), value)
}
