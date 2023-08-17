// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Preferences] class.
var PreferencesClass = _PreferencesClass{objc.GetClass("WKPreferences")}

type _PreferencesClass struct {
	objc.Class
}

// An interface definition for the [Preferences] class.
type IPreferences interface {
	objc.IObject
	MinimumFontSize() float64
	SetMinimumFontSize(value float64)
	IsTextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)
}

// An object that encapsulates the standard behaviors to apply to websites. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences?language=objc
type Preferences struct {
	objc.Object
}

func PreferencesFrom(ptr unsafe.Pointer) Preferences {
	return Preferences{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PreferencesClass) Alloc() Preferences {
	rv := objc.Call[Preferences](pc, objc.Sel("alloc"))
	return rv
}

func Preferences_Alloc() Preferences {
	return PreferencesClass.Alloc()
}

func (pc _PreferencesClass) New() Preferences {
	rv := objc.Call[Preferences](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPreferences() Preferences {
	return PreferencesClass.New()
}

func (p_ Preferences) Init() Preferences {
	rv := objc.Call[Preferences](p_, objc.Sel("init"))
	return rv
}

// The minimum font size, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc
func (p_ Preferences) MinimumFontSize() float64 {
	rv := objc.Call[float64](p_, objc.Sel("minimumFontSize"))
	return rv
}

// The minimum font size, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc
func (p_ Preferences) SetMinimumFontSize(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMinimumFontSize:"), value)
}

// A Boolean value that indicates whether to allow people to select or otherwise interact with text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc
func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isTextInteractionEnabled"))
	return rv
}

// A Boolean value that indicates whether to allow people to select or otherwise interact with text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc
func (p_ Preferences) SetTextInteractionEnabled(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setTextInteractionEnabled:"), value)
}

// A Boolean value that indicates whether JavaScript can open windows without user interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc
func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.Call[bool](p_, objc.Sel("javaScriptCanOpenWindowsAutomatically"))
	return rv
}

// A Boolean value that indicates whether JavaScript can open windows without user interaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc
func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setJavaScriptCanOpenWindowsAutomatically:"), value)
}

// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc
func (p_ Preferences) TabFocusesLinks() bool {
	rv := objc.Call[bool](p_, objc.Sel("tabFocusesLinks"))
	return rv
}

// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc
func (p_ Preferences) SetTabFocusesLinks(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setTabFocusesLinks:"), value)
}

// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc
func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isFraudulentWebsiteWarningEnabled"))
	return rv
}

// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc
func (p_ Preferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setFraudulentWebsiteWarningEnabled:"), value)
}