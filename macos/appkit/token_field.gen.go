// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TokenField] class.
var TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}

type _TokenFieldClass struct {
	objc.Class
}

// An interface definition for the [TokenField] class.
type ITokenField interface {
	ITextField
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
}

// A text field that converts text into visually distinct tokens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield?language=objc
type TokenField struct {
	TextField
}

func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: TextFieldFrom(ptr),
	}
}

func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TokenFieldClass) New() TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTokenField() TokenField {
	return TokenFieldClass.New()
}

func (t_ TokenField) Init() TokenField {
	rv := objc.Call[TokenField](t_, objc.Sel("init"))
	return rv
}

func (tc _TokenFieldClass) LabelWithString(stringValue string) TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("labelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func TokenField_LabelWithString(stringValue string) TokenField {
	return TokenFieldClass.LabelWithString(stringValue)
}

func (tc _TokenFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func TokenField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TokenField {
	return TokenFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (tc _TokenFieldClass) WrappingLabelWithString(stringValue string) TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a multiline static label with selectable text that uses the system default font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func TokenField_WrappingLabelWithString(stringValue string) TokenField {
	return TokenFieldClass.WrappingLabelWithString(stringValue)
}

func (tc _TokenFieldClass) TextFieldWithString(stringValue string) TokenField {
	rv := objc.Call[TokenField](tc, objc.Sel("textFieldWithString:"), stringValue)
	return rv
}

// Initializes a single-line editable text field for user input using the system default font and standard visual appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func TokenField_TextFieldWithString(stringValue string) TokenField {
	return TokenFieldClass.TextFieldWithString(stringValue)
}

func (t_ TokenField) InitWithFrame(frameRect foundation.Rect) TokenField {
	rv := objc.Call[TokenField](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewTokenFieldWithFrame(frameRect foundation.Rect) TokenField {
	instance := TokenFieldClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the default completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1529243-defaultcompletiondelay?language=objc
func (tc _TokenFieldClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](tc, objc.Sel("defaultCompletionDelay"))
	return rv
}

// Returns the default completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1529243-defaultcompletiondelay?language=objc
func TokenField_DefaultCompletionDelay() foundation.TimeInterval {
	return TokenFieldClass.DefaultCompletionDelay()
}

// The token style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1534077-tokenstyle?language=objc
func (t_ TokenField) TokenStyle() TokenStyle {
	rv := objc.Call[TokenStyle](t_, objc.Sel("tokenStyle"))
	return rv
}

// The token style of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1534077-tokenstyle?language=objc
func (t_ TokenField) SetTokenStyle(value TokenStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setTokenStyle:"), value)
}

// Returns the default tokenizing character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1530131-defaulttokenizingcharacterset?language=objc
func (tc _TokenFieldClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](tc, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}

// Returns the default tokenizing character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1530131-defaulttokenizingcharacterset?language=objc
func TokenField_DefaultTokenizingCharacterSet() foundation.CharacterSet {
	return TokenFieldClass.DefaultTokenizingCharacterSet()
}

// The receiver’s completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1535823-completiondelay?language=objc
func (t_ TokenField) CompletionDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](t_, objc.Sel("completionDelay"))
	return rv
}

// The receiver’s completion delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1535823-completiondelay?language=objc
func (t_ TokenField) SetCompletionDelay(value foundation.TimeInterval) {
	objc.Call[objc.Void](t_, objc.Sel("setCompletionDelay:"), value)
}

// The recevier’s tokenizing character set to characterSet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1534230-tokenizingcharacterset?language=objc
func (t_ TokenField) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Call[foundation.CharacterSet](t_, objc.Sel("tokenizingCharacterSet"))
	return rv
}

// The recevier’s tokenizing character set to characterSet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstokenfield/1534230-tokenizingcharacterset?language=objc
func (t_ TokenField) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	objc.Call[objc.Void](t_, objc.Sel("setTokenizingCharacterSet:"), value)
}
