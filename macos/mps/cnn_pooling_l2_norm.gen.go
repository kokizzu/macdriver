// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingL2Norm] class.
var CNNPoolingL2NormClass = _CNNPoolingL2NormClass{objc.GetClass("MPSCNNPoolingL2Norm")}

type _CNNPoolingL2NormClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingL2Norm] class.
type ICNNPoolingL2Norm interface {
	ICNNPooling
}

// An L2-norm pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2norm?language=objc
type CNNPoolingL2Norm struct {
	CNNPooling
}

func CNNPoolingL2NormFrom(ptr unsafe.Pointer) CNNPoolingL2Norm {
	return CNNPoolingL2Norm{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}

func (c_ CNNPoolingL2Norm) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2Norm {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2Norm](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes an L2-norm pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2norm/2875162-initwithdevice?language=objc
func NewCNNPoolingL2NormWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2Norm {
	instance := CNNPoolingL2NormClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingL2NormClass) Alloc() CNNPoolingL2Norm {
	rv := objc.Call[CNNPoolingL2Norm](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingL2Norm_Alloc() CNNPoolingL2Norm {
	return CNNPoolingL2NormClass.Alloc()
}

func (cc _CNNPoolingL2NormClass) New() CNNPoolingL2Norm {
	rv := objc.Call[CNNPoolingL2Norm](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingL2Norm() CNNPoolingL2Norm {
	return CNNPoolingL2NormClass.New()
}

func (c_ CNNPoolingL2Norm) Init() CNNPoolingL2Norm {
	rv := objc.Call[CNNPoolingL2Norm](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingL2Norm) InitWithDevice(device metal.PDevice) CNNPoolingL2Norm {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2Norm](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNPoolingL2NormWithDevice(device metal.PDevice) CNNPoolingL2Norm {
	instance := CNNPoolingL2NormClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingL2Norm) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingL2Norm {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2Norm](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingL2Norm_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingL2Norm {
	instance := CNNPoolingL2NormClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}