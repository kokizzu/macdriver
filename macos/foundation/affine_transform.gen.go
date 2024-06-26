// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [AffineTransform] class.
var AffineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}

type _AffineTransformClass struct {
	objc.Class
}

// An interface definition for the [AffineTransform] class.
type IAffineTransform interface {
	objc.IObject
	PrependTransform(transform IAffineTransform)
	Concat()
	ScaleXByYBy(scaleX float64, scaleY float64)
	TransformBezierPath(path objc.IObject) objc.Object
	Invert()
	Set()
	RotateByDegrees(angle float64)
	TransformSize(aSize Size) Size
	RotateByRadians(angle float64)
	ScaleBy(scale float64)
	AppendTransform(transform IAffineTransform)
	TranslateXByYBy(deltaX float64, deltaY float64)
	TransformPoint(aPoint Point) Point
	TransformStruct() AffineTransformStruct
	SetTransformStruct(value AffineTransformStruct)
}

// A graphics coordinate transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform?language=objc
type AffineTransform struct {
	objc.Object
}

func AffineTransformFrom(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AffineTransform) Init() AffineTransform {
	rv := objc.Call[AffineTransform](a_, objc.Sel("init"))
	return rv
}

func (a_ AffineTransform) InitWithTransform(transform IAffineTransform) AffineTransform {
	rv := objc.Call[AffineTransform](a_, objc.Sel("initWithTransform:"), transform)
	return rv
}

// Initializes the receiver’s matrix using another transform object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1413399-initwithtransform?language=objc
func NewAffineTransformWithTransform(transform IAffineTransform) AffineTransform {
	instance := AffineTransformClass.Alloc().InitWithTransform(transform)
	instance.Autorelease()
	return instance
}

func (ac _AffineTransformClass) Alloc() AffineTransform {
	rv := objc.Call[AffineTransform](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AffineTransformClass) New() AffineTransform {
	rv := objc.Call[AffineTransform](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAffineTransform() AffineTransform {
	return AffineTransformClass.New()
}

// Prepends the specified matrix to the receiver’s matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1409793-prependtransform?language=objc
func (a_ AffineTransform) PrependTransform(transform IAffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("prependTransform:"), transform)
}

// Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1462135-concat?language=objc
func (a_ AffineTransform) Concat() {
	objc.Call[objc.Void](a_, objc.Sel("concat"))
}

// Applies scaling factors to each axis of the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1411806-scalexby?language=objc
func (a_ AffineTransform) ScaleXByYBy(scaleX float64, scaleY float64) {
	objc.Call[objc.Void](a_, objc.Sel("scaleXBy:yBy:"), scaleX, scaleY)
}

// Creates and returns a new NSBezierPath object with each point in the given path transformed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1462131-transformbezierpath?language=objc
func (a_ AffineTransform) TransformBezierPath(path objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("transformBezierPath:"), path)
	return rv
}

// Creates a new affine transform initialized to the identity matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1589057-transform?language=objc
func (ac _AffineTransformClass) Transform() AffineTransform {
	rv := objc.Call[AffineTransform](ac, objc.Sel("transform"))
	return rv
}

// Creates a new affine transform initialized to the identity matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1589057-transform?language=objc
func AffineTransform_Transform() AffineTransform {
	return AffineTransformClass.Transform()
}

// Replaces the receiver’s matrix with its inverse matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1412434-invert?language=objc
func (a_ AffineTransform) Invert() {
	objc.Call[objc.Void](a_, objc.Sel("invert"))
}

// Sets the current transformation matrix to the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1462133-set?language=objc
func (a_ AffineTransform) Set() {
	objc.Call[objc.Void](a_, objc.Sel("set"))
}

// Applies a rotation factor (measured in degrees) to the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1407401-rotatebydegrees?language=objc
func (a_ AffineTransform) RotateByDegrees(angle float64) {
	objc.Call[objc.Void](a_, objc.Sel("rotateByDegrees:"), angle)
}

// Applies the receiver’s transform to the specified size and returns the results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1416692-transformsize?language=objc
func (a_ AffineTransform) TransformSize(aSize Size) Size {
	rv := objc.Call[Size](a_, objc.Sel("transformSize:"), aSize)
	return rv
}

// Applies a rotation factor (measured in radians) to the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1411353-rotatebyradians?language=objc
func (a_ AffineTransform) RotateByRadians(angle float64) {
	objc.Call[objc.Void](a_, objc.Sel("rotateByRadians:"), angle)
}

// Applies the specified scaling factor along both x and y axes to the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1413649-scaleby?language=objc
func (a_ AffineTransform) ScaleBy(scale float64) {
	objc.Call[objc.Void](a_, objc.Sel("scaleBy:"), scale)
}

// Appends the specified matrix to the receiver’s matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1408404-appendtransform?language=objc
func (a_ AffineTransform) AppendTransform(transform IAffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("appendTransform:"), transform)
}

// Applies the specified translation factors to the receiver’s transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1417196-translatexby?language=objc
func (a_ AffineTransform) TranslateXByYBy(deltaX float64, deltaY float64) {
	objc.Call[objc.Void](a_, objc.Sel("translateXBy:yBy:"), deltaX, deltaY)
}

// Applies the receiver’s transform to the specified point and returns the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1411619-transformpoint?language=objc
func (a_ AffineTransform) TransformPoint(aPoint Point) Point {
	rv := objc.Call[Point](a_, objc.Sel("transformPoint:"), aPoint)
	return rv
}

// The matrix coefficients stored as the transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1414485-transformstruct?language=objc
func (a_ AffineTransform) TransformStruct() AffineTransformStruct {
	rv := objc.Call[AffineTransformStruct](a_, objc.Sel("transformStruct"))
	return rv
}

// The matrix coefficients stored as the transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsaffinetransform/1414485-transformstruct?language=objc
func (a_ AffineTransform) SetTransformStruct(value AffineTransformStruct) {
	objc.Call[objc.Void](a_, objc.Sel("setTransformStruct:"), value)
}
