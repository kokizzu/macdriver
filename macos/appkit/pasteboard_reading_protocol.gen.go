// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that defines the interface for initializing an object from a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreading?language=objc
type PPasteboardReading interface {
	// optional
	InitWithPasteboardPropertyListOfType(propertyList objc.Object, type_ PasteboardType) objc.Object
	HasInitWithPasteboardPropertyListOfType() bool
}

// ensure impl type implements protocol interface
var _ PPasteboardReading = (*PasteboardReadingObject)(nil)

// A concrete type for the [PPasteboardReading] protocol.
type PasteboardReadingObject struct {
	objc.Object
}

func (p_ PasteboardReadingObject) HasInitWithPasteboardPropertyListOfType() bool {
	return p_.RespondsToSelector(objc.Sel("initWithPasteboardPropertyList:ofType:"))
}

// Initializes an instance with a property list object and a type string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreading/1528252-initwithpasteboardpropertylist?language=objc
func (p_ PasteboardReadingObject) InitWithPasteboardPropertyListOfType(propertyList objc.Object, type_ PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, type_)
	return rv
}