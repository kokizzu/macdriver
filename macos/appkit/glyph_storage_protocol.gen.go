// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that a glyph storage object must implement to interact properly with NSGlyphGenerator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage?language=objc
type PGlyphStorage interface {
	// optional
	AttributedString() foundation.AttributedString
	HasAttributedString() bool

	// optional
	InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint)
	HasInsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex() bool

	// optional
	LayoutOptions() uint
	HasLayoutOptions() bool

	// optional
	SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
	HasSetIntAttributeValueForGlyphAtIndex() bool
}

// ensure impl type implements protocol interface
var _ PGlyphStorage = (*GlyphStorageObject)(nil)

// A concrete type for the [PGlyphStorage] protocol.
type GlyphStorageObject struct {
	objc.Object
}

func (g_ GlyphStorageObject) HasAttributedString() bool {
	return g_.RespondsToSelector(objc.Sel("attributedString"))
}

// Returns the text storage object from which the NSGlyphGenerator object procures characters for glyph generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425147-attributedstring?language=objc
func (g_ GlyphStorageObject) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](g_, objc.Sel("attributedString"))
	return rv
}

func (g_ GlyphStorageObject) HasInsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex() bool {
	return g_.RespondsToSelector(objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"))
}

// Inserts the given glyphs into the glyph cache and maps them to the specified characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425153-insertglyphs?language=objc
func (g_ GlyphStorageObject) InsertGlyphsLengthForStartingGlyphAtIndexCharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	objc.Call[objc.Void](g_, objc.Sel("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
}

func (g_ GlyphStorageObject) HasLayoutOptions() bool {
	return g_.RespondsToSelector(objc.Sel("layoutOptions"))
}

// Returns the current layout options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425149-layoutoptions?language=objc
func (g_ GlyphStorageObject) LayoutOptions() uint {
	rv := objc.Call[uint](g_, objc.Sel("layoutOptions"))
	return rv
}

func (g_ GlyphStorageObject) HasSetIntAttributeValueForGlyphAtIndex() bool {
	return g_.RespondsToSelector(objc.Sel("setIntAttribute:value:forGlyphAtIndex:"))
}

// Sets a custom attribute value for a given glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphstorage/1425141-setintattribute?language=objc
func (g_ GlyphStorageObject) SetIntAttributeValueForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.Call[objc.Void](g_, objc.Sel("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}
