// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging?language=objc
type PFontChanging interface {
	// optional
	ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask
	HasValidModesForFontPanel() bool

	// optional
	ChangeFont(sender FontManager)
	HasChangeFont() bool
}

// ensure impl type implements protocol interface
var _ PFontChanging = (*FontChangingObject)(nil)

// A concrete type for the [PFontChanging] protocol.
type FontChangingObject struct {
	objc.Object
}

func (f_ FontChangingObject) HasValidModesForFontPanel() bool {
	return f_.RespondsToSelector(objc.Sel("validModesForFontPanel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging/3005181-validmodesforfontpanel?language=objc
func (f_ FontChangingObject) ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask {
	rv := objc.Call[FontPanelModeMask](f_, objc.Sel("validModesForFontPanel:"), objc.Ptr(fontPanel))
	return rv
}

func (f_ FontChangingObject) HasChangeFont() bool {
	return f_.RespondsToSelector(objc.Sel("changeFont:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging/3005180-changefont?language=objc
func (f_ FontChangingObject) ChangeFont(sender FontManager) {
	objc.Call[objc.Void](f_, objc.Sel("changeFont:"), objc.Ptr(sender))
}