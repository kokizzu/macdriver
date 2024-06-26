// Code generated by DarwinKit. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [EmitterCell] class.
var EmitterCellClass = _EmitterCellClass{objc.GetClass("CAEmitterCell")}

type _EmitterCellClass struct {
	objc.Class
}

// An interface definition for the [EmitterCell] class.
type IEmitterCell interface {
	objc.IObject
	ShouldArchiveValueForKey(key string) bool
	SpinRange() float64
	SetSpinRange(value float64)
	GreenRange() float32
	SetGreenRange(value float32)
	MagnificationFilter() string
	SetMagnificationFilter(value string)
	XAcceleration() float64
	SetXAcceleration(value float64)
	AlphaRange() float32
	SetAlphaRange(value float32)
	ScaleSpeed() float64
	SetScaleSpeed(value float64)
	Contents() objc.Object
	SetContents(value objc.IObject)
	BirthRate() float32
	SetBirthRate(value float32)
	EmitterCells() []EmitterCell
	SetEmitterCells(value []IEmitterCell)
	ContentsScale() float64
	SetContentsScale(value float64)
	BlueRange() float32
	SetBlueRange(value float32)
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	ScaleRange() float64
	SetScaleRange(value float64)
	YAcceleration() float64
	SetYAcceleration(value float64)
	VelocityRange() float64
	SetVelocityRange(value float64)
	Name() string
	SetName(value string)
	Spin() float64
	SetSpin(value float64)
	EmissionLongitude() float64
	SetEmissionLongitude(value float64)
	Color() coregraphics.ColorRef
	SetColor(value coregraphics.ColorRef)
	BlueSpeed() float32
	SetBlueSpeed(value float32)
	RedRange() float32
	SetRedRange(value float32)
	ZAcceleration() float64
	SetZAcceleration(value float64)
	LifetimeRange() float32
	SetLifetimeRange(value float32)
	ContentsRect() coregraphics.Rect
	SetContentsRect(value coregraphics.Rect)
	Velocity() float64
	SetVelocity(value float64)
	MinificationFilter() string
	SetMinificationFilter(value string)
	IsEnabled() bool
	SetEnabled(value bool)
	EmissionLatitude() float64
	SetEmissionLatitude(value float64)
	Scale() float64
	SetScale(value float64)
	RedSpeed() float32
	SetRedSpeed(value float32)
	EmissionRange() float64
	SetEmissionRange(value float64)
	AlphaSpeed() float32
	SetAlphaSpeed(value float32)
	Lifetime() float32
	SetLifetime(value float32)
	Style() foundation.Dictionary
	SetStyle(value foundation.Dictionary)
	GreenSpeed() float32
	SetGreenSpeed(value float32)
}

// The definition of a particle emitted by a CAEmitterLayer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell?language=objc
type EmitterCell struct {
	objc.Object
}

func EmitterCellFrom(ptr unsafe.Pointer) EmitterCell {
	return EmitterCell{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EmitterCellClass) EmitterCell() EmitterCell {
	rv := objc.Call[EmitterCell](ec, objc.Sel("emitterCell"))
	return rv
}

// Creates and returns an instance of CAEmitterCell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1584370-emittercell?language=objc
func EmitterCell_EmitterCell() EmitterCell {
	return EmitterCellClass.EmitterCell()
}

func (ec _EmitterCellClass) Alloc() EmitterCell {
	rv := objc.Call[EmitterCell](ec, objc.Sel("alloc"))
	return rv
}

func (ec _EmitterCellClass) New() EmitterCell {
	rv := objc.Call[EmitterCell](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEmitterCell() EmitterCell {
	return EmitterCellClass.New()
}

func (e_ EmitterCell) Init() EmitterCell {
	rv := objc.Call[EmitterCell](e_, objc.Sel("init"))
	return rv
}

// Returns the default value of the property with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521964-defaultvalueforkey?language=objc
func (ec _EmitterCellClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](ec, objc.Sel("defaultValueForKey:"), key)
	return rv
}

// Returns the default value of the property with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521964-defaultvalueforkey?language=objc
func EmitterCell_DefaultValueForKey(key string) objc.Object {
	return EmitterCellClass.DefaultValueForKey(key)
}

// Returns a Boolean value indicating whether the value for a given key should be archived. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522005-shouldarchivevalueforkey?language=objc
func (e_ EmitterCell) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Call[bool](e_, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}

// The amount by which the spin of the cell can vary over its lifetime. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522084-spinrange?language=objc
func (e_ EmitterCell) SpinRange() float64 {
	rv := objc.Call[float64](e_, objc.Sel("spinRange"))
	return rv
}

// The amount by which the spin of the cell can vary over its lifetime. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522084-spinrange?language=objc
func (e_ EmitterCell) SetSpinRange(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setSpinRange:"), value)
}

// The amount by which the green color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521867-greenrange?language=objc
func (e_ EmitterCell) GreenRange() float32 {
	rv := objc.Call[float32](e_, objc.Sel("greenRange"))
	return rv
}

// The amount by which the green color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521867-greenrange?language=objc
func (e_ EmitterCell) SetGreenRange(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setGreenRange:"), value)
}

// The filter used when increasing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522228-magnificationfilter?language=objc
func (e_ EmitterCell) MagnificationFilter() string {
	rv := objc.Call[string](e_, objc.Sel("magnificationFilter"))
	return rv
}

// The filter used when increasing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522228-magnificationfilter?language=objc
func (e_ EmitterCell) SetMagnificationFilter(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setMagnificationFilter:"), value)
}

// The x component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521879-xacceleration?language=objc
func (e_ EmitterCell) XAcceleration() float64 {
	rv := objc.Call[float64](e_, objc.Sel("xAcceleration"))
	return rv
}

// The x component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521879-xacceleration?language=objc
func (e_ EmitterCell) SetXAcceleration(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setXAcceleration:"), value)
}

// The amount by which the alpha component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522110-alpharange?language=objc
func (e_ EmitterCell) AlphaRange() float32 {
	rv := objc.Call[float32](e_, objc.Sel("alphaRange"))
	return rv
}

// The amount by which the alpha component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522110-alpharange?language=objc
func (e_ EmitterCell) SetAlphaRange(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setAlphaRange:"), value)
}

// The speed at which the scale changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522241-scalespeed?language=objc
func (e_ EmitterCell) ScaleSpeed() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scaleSpeed"))
	return rv
}

// The speed at which the scale changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522241-scalespeed?language=objc
func (e_ EmitterCell) SetScaleSpeed(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setScaleSpeed:"), value)
}

// An object that provides the contents of the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522109-contents?language=objc
func (e_ EmitterCell) Contents() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("contents"))
	return rv
}

// An object that provides the contents of the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522109-contents?language=objc
func (e_ EmitterCell) SetContents(value objc.IObject) {
	objc.Call[objc.Void](e_, objc.Sel("setContents:"), value)
}

// The number of emitted objects created every second. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522100-birthrate?language=objc
func (e_ EmitterCell) BirthRate() float32 {
	rv := objc.Call[float32](e_, objc.Sel("birthRate"))
	return rv
}

// The number of emitted objects created every second. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522100-birthrate?language=objc
func (e_ EmitterCell) SetBirthRate(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setBirthRate:"), value)
}

// An optional array containing the sub-cells of this cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521866-emittercells?language=objc
func (e_ EmitterCell) EmitterCells() []EmitterCell {
	rv := objc.Call[[]EmitterCell](e_, objc.Sel("emitterCells"))
	return rv
}

// An optional array containing the sub-cells of this cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521866-emittercells?language=objc
func (e_ EmitterCell) SetEmitterCells(value []IEmitterCell) {
	objc.Call[objc.Void](e_, objc.Sel("setEmitterCells:"), value)
}

// The scale factor of the cell contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522197-contentsscale?language=objc
func (e_ EmitterCell) ContentsScale() float64 {
	rv := objc.Call[float64](e_, objc.Sel("contentsScale"))
	return rv
}

// The scale factor of the cell contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522197-contentsscale?language=objc
func (e_ EmitterCell) SetContentsScale(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setContentsScale:"), value)
}

// The amount by which the blue color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522158-bluerange?language=objc
func (e_ EmitterCell) BlueRange() float32 {
	rv := objc.Call[float32](e_, objc.Sel("blueRange"))
	return rv
}

// The amount by which the blue color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522158-bluerange?language=objc
func (e_ EmitterCell) SetBlueRange(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setBlueRange:"), value)
}

// The bias factor used by the minification filter to determine the levels of detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521907-minificationfilterbias?language=objc
func (e_ EmitterCell) MinificationFilterBias() float32 {
	rv := objc.Call[float32](e_, objc.Sel("minificationFilterBias"))
	return rv
}

// The bias factor used by the minification filter to determine the levels of detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521907-minificationfilterbias?language=objc
func (e_ EmitterCell) SetMinificationFilterBias(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setMinificationFilterBias:"), value)
}

// Specifies the range over which the scale value can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521915-scalerange?language=objc
func (e_ EmitterCell) ScaleRange() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scaleRange"))
	return rv
}

// Specifies the range over which the scale value can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521915-scalerange?language=objc
func (e_ EmitterCell) SetScaleRange(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setScaleRange:"), value)
}

// The y component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522077-yacceleration?language=objc
func (e_ EmitterCell) YAcceleration() float64 {
	rv := objc.Call[float64](e_, objc.Sel("yAcceleration"))
	return rv
}

// The y component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522077-yacceleration?language=objc
func (e_ EmitterCell) SetYAcceleration(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setYAcceleration:"), value)
}

// The amount by which the velocity of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522330-velocityrange?language=objc
func (e_ EmitterCell) VelocityRange() float64 {
	rv := objc.Call[float64](e_, objc.Sel("velocityRange"))
	return rv
}

// The amount by which the velocity of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522330-velocityrange?language=objc
func (e_ EmitterCell) SetVelocityRange(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setVelocityRange:"), value)
}

// The name of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521909-name?language=objc
func (e_ EmitterCell) Name() string {
	rv := objc.Call[string](e_, objc.Sel("name"))
	return rv
}

// The name of the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521909-name?language=objc
func (e_ EmitterCell) SetName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setName:"), value)
}

// The rotational velocity, measured in radians per second, to apply to the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522361-spin?language=objc
func (e_ EmitterCell) Spin() float64 {
	rv := objc.Call[float64](e_, objc.Sel("spin"))
	return rv
}

// The rotational velocity, measured in radians per second, to apply to the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522361-spin?language=objc
func (e_ EmitterCell) SetSpin(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setSpin:"), value)
}

// The longitudinal orientation of the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522013-emissionlongitude?language=objc
func (e_ EmitterCell) EmissionLongitude() float64 {
	rv := objc.Call[float64](e_, objc.Sel("emissionLongitude"))
	return rv
}

// The longitudinal orientation of the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522013-emissionlongitude?language=objc
func (e_ EmitterCell) SetEmissionLongitude(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEmissionLongitude:"), value)
}

// The color of each emitted object. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522322-color?language=objc
func (e_ EmitterCell) Color() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](e_, objc.Sel("color"))
	return rv
}

// The color of each emitted object. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522322-color?language=objc
func (e_ EmitterCell) SetColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](e_, objc.Sel("setColor:"), value)
}

// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522082-bluespeed?language=objc
func (e_ EmitterCell) BlueSpeed() float32 {
	rv := objc.Call[float32](e_, objc.Sel("blueSpeed"))
	return rv
}

// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522082-bluespeed?language=objc
func (e_ EmitterCell) SetBlueSpeed(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setBlueSpeed:"), value)
}

// The amount by which the red color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522176-redrange?language=objc
func (e_ EmitterCell) RedRange() float32 {
	rv := objc.Call[float32](e_, objc.Sel("redRange"))
	return rv
}

// The amount by which the red color component of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522176-redrange?language=objc
func (e_ EmitterCell) SetRedRange(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setRedRange:"), value)
}

// The z component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522298-zacceleration?language=objc
func (e_ EmitterCell) ZAcceleration() float64 {
	rv := objc.Call[float64](e_, objc.Sel("zAcceleration"))
	return rv
}

// The z component of an acceleration vector applied to cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522298-zacceleration?language=objc
func (e_ EmitterCell) SetZAcceleration(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setZAcceleration:"), value)
}

// The mean value by which the [quartzcore/caemittercell/lifetime] of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522101-lifetimerange?language=objc
func (e_ EmitterCell) LifetimeRange() float32 {
	rv := objc.Call[float32](e_, objc.Sel("lifetimeRange"))
	return rv
}

// The mean value by which the [quartzcore/caemittercell/lifetime] of the cell can vary. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522101-lifetimerange?language=objc
func (e_ EmitterCell) SetLifetimeRange(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setLifetimeRange:"), value)
}

// A rectangle (in the unit coordinate space) that specifies the portion of [quartzcore/caemittercell/contents] that the receiver should draw. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522124-contentsrect?language=objc
func (e_ EmitterCell) ContentsRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](e_, objc.Sel("contentsRect"))
	return rv
}

// A rectangle (in the unit coordinate space) that specifies the portion of [quartzcore/caemittercell/contents] that the receiver should draw. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522124-contentsrect?language=objc
func (e_ EmitterCell) SetContentsRect(value coregraphics.Rect) {
	objc.Call[objc.Void](e_, objc.Sel("setContentsRect:"), value)
}

// The initial velocity of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521837-velocity?language=objc
func (e_ EmitterCell) Velocity() float64 {
	rv := objc.Call[float64](e_, objc.Sel("velocity"))
	return rv
}

// The initial velocity of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521837-velocity?language=objc
func (e_ EmitterCell) SetVelocity(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setVelocity:"), value)
}

// The filter used when reducing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522222-minificationfilter?language=objc
func (e_ EmitterCell) MinificationFilter() string {
	rv := objc.Call[string](e_, objc.Sel("minificationFilter"))
	return rv
}

// The filter used when reducing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522222-minificationfilter?language=objc
func (e_ EmitterCell) SetMinificationFilter(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setMinificationFilter:"), value)
}

// A Boolean value indicating whether or not cells from this emitter are rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521831-enabled?language=objc
func (e_ EmitterCell) IsEnabled() bool {
	rv := objc.Call[bool](e_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value indicating whether or not cells from this emitter are rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521831-enabled?language=objc
func (e_ EmitterCell) SetEnabled(value bool) {
	objc.Call[objc.Void](e_, objc.Sel("setEnabled:"), value)
}

// The latitudinal orientation of the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521857-emissionlatitude?language=objc
func (e_ EmitterCell) EmissionLatitude() float64 {
	rv := objc.Call[float64](e_, objc.Sel("emissionLatitude"))
	return rv
}

// The latitudinal orientation of the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521857-emissionlatitude?language=objc
func (e_ EmitterCell) SetEmissionLatitude(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEmissionLatitude:"), value)
}

// Specifies the scale factor applied to the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522287-scale?language=objc
func (e_ EmitterCell) Scale() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scale"))
	return rv
}

// Specifies the scale factor applied to the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522287-scale?language=objc
func (e_ EmitterCell) SetScale(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setScale:"), value)
}

// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521859-redspeed?language=objc
func (e_ EmitterCell) RedSpeed() float32 {
	rv := objc.Call[float32](e_, objc.Sel("redSpeed"))
	return rv
}

// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521859-redspeed?language=objc
func (e_ EmitterCell) SetRedSpeed(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setRedSpeed:"), value)
}

// The angle, in radians, defining a cone around the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521847-emissionrange?language=objc
func (e_ EmitterCell) EmissionRange() float64 {
	rv := objc.Call[float64](e_, objc.Sel("emissionRange"))
	return rv
}

// The angle, in radians, defining a cone around the emission angle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521847-emissionrange?language=objc
func (e_ EmitterCell) SetEmissionRange(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEmissionRange:"), value)
}

// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522120-alphaspeed?language=objc
func (e_ EmitterCell) AlphaSpeed() float32 {
	rv := objc.Call[float32](e_, objc.Sel("alphaSpeed"))
	return rv
}

// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522120-alphaspeed?language=objc
func (e_ EmitterCell) SetAlphaSpeed(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setAlphaSpeed:"), value)
}

// The lifetime of the cell, in seconds. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522075-lifetime?language=objc
func (e_ EmitterCell) Lifetime() float32 {
	rv := objc.Call[float32](e_, objc.Sel("lifetime"))
	return rv
}

// The lifetime of the cell, in seconds. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1522075-lifetime?language=objc
func (e_ EmitterCell) SetLifetime(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setLifetime:"), value)
}

// An optional dictionary containing additional style values that are not explicitly defined by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521925-style?language=objc
func (e_ EmitterCell) Style() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](e_, objc.Sel("style"))
	return rv
}

// An optional dictionary containing additional style values that are not explicitly defined by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521925-style?language=objc
func (e_ EmitterCell) SetStyle(value foundation.Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setStyle:"), value)
}

// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521946-greenspeed?language=objc
func (e_ EmitterCell) GreenSpeed() float32 {
	rv := objc.Call[float32](e_, objc.Sel("greenSpeed"))
	return rv
}

// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caemittercell/1521946-greenspeed?language=objc
func (e_ EmitterCell) SetGreenSpeed(value float32) {
	objc.Call[objc.Void](e_, objc.Sel("setGreenSpeed:"), value)
}
