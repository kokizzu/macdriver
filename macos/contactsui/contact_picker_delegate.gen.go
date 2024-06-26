// Code generated by DarwinKit. DO NOT EDIT.

package contactsui

import (
	"github.com/progrium/darwinkit/macos/contacts"
	"github.com/progrium/darwinkit/objc"
)

// The methods that you implement to respond to contact-picker user events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate?language=objc
type PContactPickerDelegate interface {
	// optional
	ContactPickerWillClose(picker ContactPicker)
	HasContactPickerWillClose() bool

	// optional
	ContactPickerDidSelectContact(picker ContactPicker, contact contacts.Contact)
	HasContactPickerDidSelectContact() bool

	// optional
	ContactPickerDidClose(picker ContactPicker)
	HasContactPickerDidClose() bool

	// optional
	ContactPickerDidSelectContactProperty(picker ContactPicker, contactProperty contacts.ContactProperty)
	HasContactPickerDidSelectContactProperty() bool
}

// A delegate implementation builder for the [PContactPickerDelegate] protocol.
type ContactPickerDelegate struct {
	_ContactPickerWillClose                func(picker ContactPicker)
	_ContactPickerDidSelectContact         func(picker ContactPicker, contact contacts.Contact)
	_ContactPickerDidClose                 func(picker ContactPicker)
	_ContactPickerDidSelectContactProperty func(picker ContactPicker, contactProperty contacts.ContactProperty)
}

func (di *ContactPickerDelegate) HasContactPickerWillClose() bool {
	return di._ContactPickerWillClose != nil
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (di *ContactPickerDelegate) SetContactPickerWillClose(f func(picker ContactPicker)) {
	di._ContactPickerWillClose = f
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (di *ContactPickerDelegate) ContactPickerWillClose(picker ContactPicker) {
	di._ContactPickerWillClose(picker)
}
func (di *ContactPickerDelegate) HasContactPickerDidSelectContact() bool {
	return di._ContactPickerDidSelectContact != nil
}

// Called after a contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522595-contactpicker?language=objc
func (di *ContactPickerDelegate) SetContactPickerDidSelectContact(f func(picker ContactPicker, contact contacts.Contact)) {
	di._ContactPickerDidSelectContact = f
}

// Called after a contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522595-contactpicker?language=objc
func (di *ContactPickerDelegate) ContactPickerDidSelectContact(picker ContactPicker, contact contacts.Contact) {
	di._ContactPickerDidSelectContact(picker, contact)
}
func (di *ContactPickerDelegate) HasContactPickerDidClose() bool {
	return di._ContactPickerDidClose != nil
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (di *ContactPickerDelegate) SetContactPickerDidClose(f func(picker ContactPicker)) {
	di._ContactPickerDidClose = f
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (di *ContactPickerDelegate) ContactPickerDidClose(picker ContactPicker) {
	di._ContactPickerDidClose(picker)
}
func (di *ContactPickerDelegate) HasContactPickerDidSelectContactProperty() bool {
	return di._ContactPickerDidSelectContactProperty != nil
}

// Called when a property of the contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522593-contactpicker?language=objc
func (di *ContactPickerDelegate) SetContactPickerDidSelectContactProperty(f func(picker ContactPicker, contactProperty contacts.ContactProperty)) {
	di._ContactPickerDidSelectContactProperty = f
}

// Called when a property of the contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522593-contactpicker?language=objc
func (di *ContactPickerDelegate) ContactPickerDidSelectContactProperty(picker ContactPicker, contactProperty contacts.ContactProperty) {
	di._ContactPickerDidSelectContactProperty(picker, contactProperty)
}

// ensure impl type implements protocol interface
var _ PContactPickerDelegate = (*ContactPickerDelegateObject)(nil)

// A concrete type for the [PContactPickerDelegate] protocol.
type ContactPickerDelegateObject struct {
	objc.Object
}

func (c_ ContactPickerDelegateObject) HasContactPickerWillClose() bool {
	return c_.RespondsToSelector(objc.Sel("contactPickerWillClose:"))
}

// In macOS, called when the contact picker’s popover is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522594-contactpickerwillclose?language=objc
func (c_ ContactPickerDelegateObject) ContactPickerWillClose(picker ContactPicker) {
	objc.Call[objc.Void](c_, objc.Sel("contactPickerWillClose:"), picker)
}

func (c_ ContactPickerDelegateObject) HasContactPickerDidSelectContact() bool {
	return c_.RespondsToSelector(objc.Sel("contactPicker:didSelectContact:"))
}

// Called after a contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522595-contactpicker?language=objc
func (c_ ContactPickerDelegateObject) ContactPickerDidSelectContact(picker ContactPicker, contact contacts.Contact) {
	objc.Call[objc.Void](c_, objc.Sel("contactPicker:didSelectContact:"), picker, contact)
}

func (c_ ContactPickerDelegateObject) HasContactPickerDidClose() bool {
	return c_.RespondsToSelector(objc.Sel("contactPickerDidClose:"))
}

// In macOS, called when the contact picker’s popover has closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522584-contactpickerdidclose?language=objc
func (c_ ContactPickerDelegateObject) ContactPickerDidClose(picker ContactPicker) {
	objc.Call[objc.Void](c_, objc.Sel("contactPickerDidClose:"), picker)
}

func (c_ ContactPickerDelegateObject) HasContactPickerDidSelectContactProperty() bool {
	return c_.RespondsToSelector(objc.Sel("contactPicker:didSelectContactProperty:"))
}

// Called when a property of the contact has been selected by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactpickerdelegate/1522593-contactpicker?language=objc
func (c_ ContactPickerDelegateObject) ContactPickerDidSelectContactProperty(picker ContactPicker, contactProperty contacts.ContactProperty) {
	objc.Call[objc.Void](c_, objc.Sel("contactPicker:didSelectContactProperty:"), picker, contactProperty)
}
