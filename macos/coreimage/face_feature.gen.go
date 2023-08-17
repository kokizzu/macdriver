// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FaceFeature] class.
var FaceFeatureClass = _FaceFeatureClass{objc.GetClass("CIFaceFeature")}

type _FaceFeatureClass struct {
	objc.Class
}

// An interface definition for the [FaceFeature] class.
type IFaceFeature interface {
	IFeature
	FaceAngle() float64
	RightEyeClosed() bool
	TrackingID() int
	HasTrackingFrameCount() bool
	HasLeftEyePosition() bool
	HasMouthPosition() bool
	RightEyePosition() coregraphics.Point
	HasSmile() bool
	HasFaceAngle() bool
	HasTrackingID() bool
	TrackingFrameCount() int
	LeftEyeClosed() bool
	LeftEyePosition() coregraphics.Point
	HasRightEyePosition() bool
	MouthPosition() coregraphics.Point
}

// Information about a face detected in a still or video image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature?language=objc
type FaceFeature struct {
	Feature
}

func FaceFeatureFrom(ptr unsafe.Pointer) FaceFeature {
	return FaceFeature{
		Feature: FeatureFrom(ptr),
	}
}

func (fc _FaceFeatureClass) Alloc() FaceFeature {
	rv := objc.Call[FaceFeature](fc, objc.Sel("alloc"))
	return rv
}

func FaceFeature_Alloc() FaceFeature {
	return FaceFeatureClass.Alloc()
}

func (fc _FaceFeatureClass) New() FaceFeature {
	rv := objc.Call[FaceFeature](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFaceFeature() FaceFeature {
	return FaceFeatureClass.New()
}

func (f_ FaceFeature) Init() FaceFeature {
	rv := objc.Call[FaceFeature](f_, objc.Sel("init"))
	return rv
}

// The rotation of the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437689-faceangle?language=objc
func (f_ FaceFeature) FaceAngle() float64 {
	rv := objc.Call[float64](f_, objc.Sel("faceAngle"))
	return rv
}

// A Boolean value that indicates whether a closed right eye is detected in the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437615-righteyeclosed?language=objc
func (f_ FaceFeature) RightEyeClosed() bool {
	rv := objc.Call[bool](f_, objc.Sel("rightEyeClosed"))
	return rv
}

// The tracking identifier of the face object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437709-trackingid?language=objc
func (f_ FaceFeature) TrackingID() int {
	rv := objc.Call[int](f_, objc.Sel("trackingID"))
	return rv
}

// A Boolean value that indicates the face object has a tracking frame count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437731-hastrackingframecount?language=objc
func (f_ FaceFeature) HasTrackingFrameCount() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasTrackingFrameCount"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s left eye. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437900-haslefteyeposition?language=objc
func (f_ FaceFeature) HasLeftEyePosition() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasLeftEyePosition"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s mouth. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437976-hasmouthposition?language=objc
func (f_ FaceFeature) HasMouthPosition() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasMouthPosition"))
	return rv
}

// The coordinates of the right eye, in image coordinates [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1438213-righteyeposition?language=objc
func (f_ FaceFeature) RightEyePosition() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("rightEyePosition"))
	return rv
}

// A Boolean value that indicates whether a smile is detected in the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437882-hassmile?language=objc
func (f_ FaceFeature) HasSmile() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasSmile"))
	return rv
}

// A Boolean value that indicates whether information about face rotation is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1438165-hasfaceangle?language=objc
func (f_ FaceFeature) HasFaceAngle() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasFaceAngle"))
	return rv
}

// A Boolean value that indicates whether the face object has a tracking ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437683-hastrackingid?language=objc
func (f_ FaceFeature) HasTrackingID() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasTrackingID"))
	return rv
}

// The tracking frame count of the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437953-trackingframecount?language=objc
func (f_ FaceFeature) TrackingFrameCount() int {
	rv := objc.Call[int](f_, objc.Sel("trackingFrameCount"))
	return rv
}

// A Boolean value that indicates whether a closed left eye is detected in the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437630-lefteyeclosed?language=objc
func (f_ FaceFeature) LeftEyeClosed() bool {
	rv := objc.Call[bool](f_, objc.Sel("leftEyeClosed"))
	return rv
}

// The coordinates of the left eye, in image coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1437923-lefteyeposition?language=objc
func (f_ FaceFeature) LeftEyePosition() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("leftEyePosition"))
	return rv
}

// A Boolean value that indicates whether the detector found the face’s right eye. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1438076-hasrighteyeposition?language=objc
func (f_ FaceFeature) HasRightEyePosition() bool {
	rv := objc.Call[bool](f_, objc.Sel("hasRightEyePosition"))
	return rv
}

// The coordinates of the mouth, in image coordinates [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifacefeature/1438020-mouthposition?language=objc
func (f_ FaceFeature) MouthPosition() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("mouthPosition"))
	return rv
}