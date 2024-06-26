// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [AttributedStringMarkdownParsingOptions] class.
var AttributedStringMarkdownParsingOptionsClass = _AttributedStringMarkdownParsingOptionsClass{objc.GetClass("NSAttributedStringMarkdownParsingOptions")}

type _AttributedStringMarkdownParsingOptionsClass struct {
	objc.Class
}

// An interface definition for the [AttributedStringMarkdownParsingOptions] class.
type IAttributedStringMarkdownParsingOptions interface {
	objc.IObject
	LanguageCode() string
	SetLanguageCode(value string)
	FailurePolicy() AttributedStringMarkdownParsingFailurePolicy
	SetFailurePolicy(value AttributedStringMarkdownParsingFailurePolicy)
	InterpretedSyntax() AttributedStringMarkdownInterpretedSyntax
	SetInterpretedSyntax(value AttributedStringMarkdownInterpretedSyntax)
	AllowsExtendedAttributes() bool
	SetAllowsExtendedAttributes(value bool)
}

// Options that affect the parsing of Markdown content into an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions?language=objc
type AttributedStringMarkdownParsingOptions struct {
	objc.Object
}

func AttributedStringMarkdownParsingOptionsFrom(ptr unsafe.Pointer) AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AttributedStringMarkdownParsingOptions) Init() AttributedStringMarkdownParsingOptions {
	rv := objc.Call[AttributedStringMarkdownParsingOptions](a_, objc.Sel("init"))
	return rv
}

func (ac _AttributedStringMarkdownParsingOptionsClass) Alloc() AttributedStringMarkdownParsingOptions {
	rv := objc.Call[AttributedStringMarkdownParsingOptions](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AttributedStringMarkdownParsingOptionsClass) New() AttributedStringMarkdownParsingOptions {
	rv := objc.Call[AttributedStringMarkdownParsingOptions](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAttributedStringMarkdownParsingOptions() AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptionsClass.New()
}

// The BCP-47 language code for this document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746889-languagecode?language=objc
func (a_ AttributedStringMarkdownParsingOptions) LanguageCode() string {
	rv := objc.Call[string](a_, objc.Sel("languageCode"))
	return rv
}

// The BCP-47 language code for this document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746889-languagecode?language=objc
func (a_ AttributedStringMarkdownParsingOptions) SetLanguageCode(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLanguageCode:"), value)
}

// The policy for handling a parsing failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746887-failurepolicy?language=objc
func (a_ AttributedStringMarkdownParsingOptions) FailurePolicy() AttributedStringMarkdownParsingFailurePolicy {
	rv := objc.Call[AttributedStringMarkdownParsingFailurePolicy](a_, objc.Sel("failurePolicy"))
	return rv
}

// The policy for handling a parsing failure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746887-failurepolicy?language=objc
func (a_ AttributedStringMarkdownParsingOptions) SetFailurePolicy(value AttributedStringMarkdownParsingFailurePolicy) {
	objc.Call[objc.Void](a_, objc.Sel("setFailurePolicy:"), value)
}

// The syntax for intepreting a Markdown string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3787562-interpretedsyntax?language=objc
func (a_ AttributedStringMarkdownParsingOptions) InterpretedSyntax() AttributedStringMarkdownInterpretedSyntax {
	rv := objc.Call[AttributedStringMarkdownInterpretedSyntax](a_, objc.Sel("interpretedSyntax"))
	return rv
}

// The syntax for intepreting a Markdown string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3787562-interpretedsyntax?language=objc
func (a_ AttributedStringMarkdownParsingOptions) SetInterpretedSyntax(value AttributedStringMarkdownInterpretedSyntax) {
	objc.Call[objc.Void](a_, objc.Sel("setInterpretedSyntax:"), value)
}

// A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746886-allowsextendedattributes?language=objc
func (a_ AttributedStringMarkdownParsingOptions) AllowsExtendedAttributes() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowsExtendedAttributes"))
	return rv
}

// A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstringmarkdownparsingoptions/3746886-allowsextendedattributes?language=objc
func (a_ AttributedStringMarkdownParsingOptions) SetAllowsExtendedAttributes(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAllowsExtendedAttributes:"), value)
}
