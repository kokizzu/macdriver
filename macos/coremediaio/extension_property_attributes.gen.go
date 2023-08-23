// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionPropertyAttributes] class.
var ExtensionPropertyAttributesClass = _ExtensionPropertyAttributesClass{objc.GetClass("CMIOExtensionPropertyAttributes")}

type _ExtensionPropertyAttributesClass struct {
	objc.Class
}

// An interface definition for the [ExtensionPropertyAttributes] class.
type IExtensionPropertyAttributes interface {
	objc.IObject
	IsReadOnly() bool
	ValidValues() []objc.Object
	MinValue() objc.Object
	MaxValue() objc.Object
}

// An object that describes the attributes of a property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes?language=objc
type ExtensionPropertyAttributes struct {
	objc.Object
}

func ExtensionPropertyAttributesFrom(ptr unsafe.Pointer) ExtensionPropertyAttributes {
	return ExtensionPropertyAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionPropertyAttributesClass) PropertyAttributesWithMinValueMaxValueValidValuesReadOnly(minValue objc.IObject, maxValue objc.IObject, validValues []objc.IObject, readOnly bool) ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](ec, objc.Sel("propertyAttributesWithMinValue:maxValue:validValues:readOnly:"), objc.Ptr(minValue), objc.Ptr(maxValue), validValues, readOnly)
	return rv
}

// Returns a new property attributes object with the specified configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915859-propertyattributeswithminvalue?language=objc
func ExtensionPropertyAttributes_PropertyAttributesWithMinValueMaxValueValidValuesReadOnly(minValue objc.IObject, maxValue objc.IObject, validValues []objc.IObject, readOnly bool) ExtensionPropertyAttributes {
	return ExtensionPropertyAttributesClass.PropertyAttributesWithMinValueMaxValueValidValuesReadOnly(minValue, maxValue, validValues, readOnly)
}

func (e_ ExtensionPropertyAttributes) InitWithMinValueMaxValueValidValuesReadOnly(minValue objc.IObject, maxValue objc.IObject, validValues []objc.IObject, readOnly bool) ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](e_, objc.Sel("initWithMinValue:maxValue:validValues:readOnly:"), objc.Ptr(minValue), objc.Ptr(maxValue), validValues, readOnly)
	return rv
}

// Creates a property attributes object with the specified configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915856-initwithminvalue?language=objc
func NewExtensionPropertyAttributesWithMinValueMaxValueValidValuesReadOnly(minValue objc.IObject, maxValue objc.IObject, validValues []objc.IObject, readOnly bool) ExtensionPropertyAttributes {
	instance := ExtensionPropertyAttributesClass.Alloc().InitWithMinValueMaxValueValidValuesReadOnly(minValue, maxValue, validValues, readOnly)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionPropertyAttributesClass) Alloc() ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionPropertyAttributes_Alloc() ExtensionPropertyAttributes {
	return ExtensionPropertyAttributesClass.Alloc()
}

func (ec _ExtensionPropertyAttributesClass) New() ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionPropertyAttributes() ExtensionPropertyAttributes {
	return ExtensionPropertyAttributesClass.New()
}

func (e_ ExtensionPropertyAttributes) Init() ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](e_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether a property is read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915860-readonly?language=objc
func (e_ ExtensionPropertyAttributes) IsReadOnly() bool {
	rv := objc.Call[bool](e_, objc.Sel("isReadOnly"))
	return rv
}

// An array of discrete values that this property supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915862-validvalues?language=objc
func (e_ ExtensionPropertyAttributes) ValidValues() []objc.Object {
	rv := objc.Call[[]objc.Object](e_, objc.Sel("validValues"))
	return rv
}

// The minimum value a property supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915858-minvalue?language=objc
func (e_ ExtensionPropertyAttributes) MinValue() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("minValue"))
	return rv
}

// A class property for a read-only property attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915861-readonlypropertyattribute?language=objc
func (ec _ExtensionPropertyAttributesClass) ReadOnlyPropertyAttribute() ExtensionPropertyAttributes {
	rv := objc.Call[ExtensionPropertyAttributes](ec, objc.Sel("readOnlyPropertyAttribute"))
	return rv
}

// A class property for a read-only property attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915861-readonlypropertyattribute?language=objc
func ExtensionPropertyAttributes_ReadOnlyPropertyAttribute() ExtensionPropertyAttributes {
	return ExtensionPropertyAttributesClass.ReadOnlyPropertyAttribute()
}

// The maximum value a property supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionpropertyattributes/3915857-maxvalue?language=objc
func (e_ ExtensionPropertyAttributes) MaxValue() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("maxValue"))
	return rv
}