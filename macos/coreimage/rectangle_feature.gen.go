// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RectangleFeature] class.
var RectangleFeatureClass = _RectangleFeatureClass{objc.GetClass("CIRectangleFeature")}

type _RectangleFeatureClass struct {
	objc.Class
}

// An interface definition for the [RectangleFeature] class.
type IRectangleFeature interface {
	IFeature
	BottomRight() coregraphics.Point
	BottomLeft() coregraphics.Point
	TopRight() coregraphics.Point
	TopLeft() coregraphics.Point
}

// Information about a rectangular region detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirectanglefeature?language=objc
type RectangleFeature struct {
	Feature
}

func RectangleFeatureFrom(ptr unsafe.Pointer) RectangleFeature {
	return RectangleFeature{
		Feature: FeatureFrom(ptr),
	}
}

func (rc _RectangleFeatureClass) Alloc() RectangleFeature {
	rv := objc.Call[RectangleFeature](rc, objc.Sel("alloc"))
	return rv
}

func RectangleFeature_Alloc() RectangleFeature {
	return RectangleFeatureClass.Alloc()
}

func (rc _RectangleFeatureClass) New() RectangleFeature {
	rv := objc.Call[RectangleFeature](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRectangleFeature() RectangleFeature {
	return RectangleFeatureClass.New()
}

func (r_ RectangleFeature) Init() RectangleFeature {
	rv := objc.Call[RectangleFeature](r_, objc.Sel("init"))
	return rv
}

// The lower-right corner of the detected rectangle, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirectanglefeature/1437888-bottomright?language=objc
func (r_ RectangleFeature) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("bottomRight"))
	return rv
}

// The lower-left corner of the detected rectangle, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirectanglefeature/1437878-bottomleft?language=objc
func (r_ RectangleFeature) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("bottomLeft"))
	return rv
}

// The upper-right corner of the detected rectangle, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirectanglefeature/1438071-topright?language=objc
func (r_ RectangleFeature) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("topRight"))
	return rv
}

// The upper-left corner of the detected rectangle, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirectanglefeature/1437951-topleft?language=objc
func (r_ RectangleFeature) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("topLeft"))
	return rv
}