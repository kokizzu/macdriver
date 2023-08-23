// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNPad] class.
var NNPadClass = _NNPadClass{objc.GetClass("MPSNNPad")}

type _NNPadClass struct {
	objc.Class
}

// An interface definition for the [NNPad] class.
type INNPad interface {
	ICNNKernel
	PaddingSizeBefore() ImageCoordinate
	SetPaddingSizeBefore(value ImageCoordinate)
	PaddingSizeAfter() ImageCoordinate
	SetPaddingSizeAfter(value ImageCoordinate)
	FillValue() float64
	SetFillValue(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad?language=objc
type NNPad struct {
	CNNKernel
}

func NNPadFrom(ptr unsafe.Pointer) NNPad {
	return NNPad{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNPad) InitWithDevice(device metal.PDevice) NNPad {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNPad](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037429-initwithdevice?language=objc
func NewNNPadWithDevice(device metal.PDevice) NNPad {
	instance := NNPadClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNPadClass) Alloc() NNPad {
	rv := objc.Call[NNPad](nc, objc.Sel("alloc"))
	return rv
}

func NNPad_Alloc() NNPad {
	return NNPadClass.Alloc()
}

func (nc _NNPadClass) New() NNPad {
	rv := objc.Call[NNPad](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNPad() NNPad {
	return NNPadClass.New()
}

func (n_ NNPad) Init() NNPad {
	rv := objc.Call[NNPad](n_, objc.Sel("init"))
	return rv
}

func (n_ NNPad) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNPad {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNPad](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNPad_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNPad {
	instance := NNPadClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037433-paddingsizebefore?language=objc
func (n_ NNPad) PaddingSizeBefore() ImageCoordinate {
	rv := objc.Call[ImageCoordinate](n_, objc.Sel("paddingSizeBefore"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037433-paddingsizebefore?language=objc
func (n_ NNPad) SetPaddingSizeBefore(value ImageCoordinate) {
	objc.Call[objc.Void](n_, objc.Sel("setPaddingSizeBefore:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037432-paddingsizeafter?language=objc
func (n_ NNPad) PaddingSizeAfter() ImageCoordinate {
	rv := objc.Call[ImageCoordinate](n_, objc.Sel("paddingSizeAfter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037432-paddingsizeafter?language=objc
func (n_ NNPad) SetPaddingSizeAfter(value ImageCoordinate) {
	objc.Call[objc.Void](n_, objc.Sel("setPaddingSizeAfter:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037427-fillvalue?language=objc
func (n_ NNPad) FillValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("fillValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpad/3037427-fillvalue?language=objc
func (n_ NNPad) SetFillValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setFillValue:"), value)
}