// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixCopyToImage] class.
var MatrixCopyToImageClass = _MatrixCopyToImageClass{objc.GetClass("MPSMatrixCopyToImage")}

type _MatrixCopyToImageClass struct {
	objc.Class
}

// An interface definition for the [MatrixCopyToImage] class.
type IMatrixCopyToImage interface {
	IKernel
	EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, destinationImages *foundation.Array)
	EncodeBatchToCommandBufferObjectSourceMatrixDestinationImages(commandBufferObject objc.IObject, sourceMatrix IMatrix, destinationImages *foundation.Array)
	EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, destinationImage IImage)
	EncodeToCommandBufferObjectSourceMatrixDestinationImage(commandBufferObject objc.IObject, sourceMatrix IMatrix, destinationImage IImage)
	DataLayout() DataLayout
	SourceMatrixBatchIndex() uint
	SetSourceMatrixBatchIndex(value uint)
	SourceMatrixOrigin() metal.Origin
	SetSourceMatrixOrigin(value metal.Origin)
}

// A kernel that copies matrix data to a Metal Performance Shaders image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage?language=objc
type MatrixCopyToImage struct {
	Kernel
}

func MatrixCopyToImageFrom(ptr unsafe.Pointer) MatrixCopyToImage {
	return MatrixCopyToImage{
		Kernel: KernelFrom(ptr),
	}
}

func (m_ MatrixCopyToImage) InitWithDeviceDataLayout(device metal.PDevice, dataLayout DataLayout) MatrixCopyToImage {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopyToImage](m_, objc.Sel("initWithDevice:dataLayout:"), po0, dataLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976460-initwithdevice?language=objc
func NewMatrixCopyToImageWithDeviceDataLayout(device metal.PDevice, dataLayout DataLayout) MatrixCopyToImage {
	instance := MatrixCopyToImageClass.Alloc().InitWithDeviceDataLayout(device, dataLayout)
	instance.Autorelease()
	return instance
}

func (mc _MatrixCopyToImageClass) Alloc() MatrixCopyToImage {
	rv := objc.Call[MatrixCopyToImage](mc, objc.Sel("alloc"))
	return rv
}

func MatrixCopyToImage_Alloc() MatrixCopyToImage {
	return MatrixCopyToImageClass.Alloc()
}

func (mc _MatrixCopyToImageClass) New() MatrixCopyToImage {
	rv := objc.Call[MatrixCopyToImage](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixCopyToImage() MatrixCopyToImage {
	return MatrixCopyToImageClass.New()
}

func (m_ MatrixCopyToImage) Init() MatrixCopyToImage {
	rv := objc.Call[MatrixCopyToImage](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixCopyToImage) InitWithDevice(device metal.PDevice) MatrixCopyToImage {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopyToImage](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixCopyToImageWithDevice(device metal.PDevice) MatrixCopyToImage {
	instance := MatrixCopyToImageClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixCopyToImage) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixCopyToImage {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopyToImage](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixCopyToImage_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixCopyToImage {
	instance := MatrixCopyToImageClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatchtocommandbuffer?language=objc
func (m_ MatrixCopyToImage) EncodeBatchToCommandBufferSourceMatrixDestinationImages(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, destinationImages *foundation.Array) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeBatchToCommandBuffer:sourceMatrix:destinationImages:"), po0, objc.Ptr(sourceMatrix), destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/3013770-encodebatchtocommandbuffer?language=objc
func (m_ MatrixCopyToImage) EncodeBatchToCommandBufferObjectSourceMatrixDestinationImages(commandBufferObject objc.IObject, sourceMatrix IMatrix, destinationImages *foundation.Array) {
	objc.Call[objc.Void](m_, objc.Sel("encodeBatchToCommandBuffer:sourceMatrix:destinationImages:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encodetocommandbuffer?language=objc
func (m_ MatrixCopyToImage) EncodeToCommandBufferSourceMatrixDestinationImage(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, destinationImage IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:destinationImage:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976458-encodetocommandbuffer?language=objc
func (m_ MatrixCopyToImage) EncodeToCommandBufferObjectSourceMatrixDestinationImage(commandBufferObject objc.IObject, sourceMatrix IMatrix, destinationImage IImage) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976457-datalayout?language=objc
func (m_ MatrixCopyToImage) DataLayout() DataLayout {
	rv := objc.Call[DataLayout](m_, objc.Sel("dataLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex?language=objc
func (m_ MatrixCopyToImage) SourceMatrixBatchIndex() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceMatrixBatchIndex"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976461-sourcematrixbatchindex?language=objc
func (m_ MatrixCopyToImage) SetSourceMatrixBatchIndex(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceMatrixBatchIndex:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin?language=objc
func (m_ MatrixCopyToImage) SourceMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("sourceMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopytoimage/2976462-sourcematrixorigin?language=objc
func (m_ MatrixCopyToImage) SetSourceMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceMatrixOrigin:"), value)
}