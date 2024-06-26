// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a vignette filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette?language=objc
type PVignette interface {
	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool

	// optional
	SetIntensity(value float32)
	HasSetIntensity() bool

	// optional
	Intensity() float32
	HasIntensity() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PVignette = (*VignetteObject)(nil)

// A concrete type for the [PVignette] protocol.
type VignetteObject struct {
	objc.Object
}

func (v_ VignetteObject) HasSetRadius() bool {
	return v_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228828-radius?language=objc
func (v_ VignetteObject) SetRadius(value float32) {
	objc.Call[objc.Void](v_, objc.Sel("setRadius:"), value)
}

func (v_ VignetteObject) HasRadius() bool {
	return v_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228828-radius?language=objc
func (v_ VignetteObject) Radius() float32 {
	rv := objc.Call[float32](v_, objc.Sel("radius"))
	return rv
}

func (v_ VignetteObject) HasSetIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228827-intensity?language=objc
func (v_ VignetteObject) SetIntensity(value float32) {
	objc.Call[objc.Void](v_, objc.Sel("setIntensity:"), value)
}

func (v_ VignetteObject) HasIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228827-intensity?language=objc
func (v_ VignetteObject) Intensity() float32 {
	rv := objc.Call[float32](v_, objc.Sel("intensity"))
	return rv
}

func (v_ VignetteObject) HasSetInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228826-inputimage?language=objc
func (v_ VignetteObject) SetInputImage(value Image) {
	objc.Call[objc.Void](v_, objc.Sel("setInputImage:"), value)
}

func (v_ VignetteObject) HasInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228826-inputimage?language=objc
func (v_ VignetteObject) InputImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("inputImage"))
	return rv
}
