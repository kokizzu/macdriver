// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FontManager] class.
var FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}

type _FontManagerClass struct {
	objc.Class
}

// An interface definition for the [FontManager] class.
type IFontManager interface {
	objc.IObject
	SendAction() bool
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	LocalizedNameForFamilyFace(family string, faceKey string) string
	ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object
	WeightOfFont(fontObj IFont) int
	SetFontMenu(newMenu IMenu)
	TraitsOfFont(fontObj IFont) FontTraitMask
	RemoveFontTrait(sender objc.IObject)
	OrderFrontFontPanel(sender objc.IObject)
	AddFontTrait(sender objc.IObject)
	ModifyFont(sender objc.IObject)
	ConvertWeightOfFont(upFlag bool, fontObj IFont) Font
	OrderFrontStylesPanel(sender objc.IObject)
	SetSelectedAttributesIsMultiple(attributes map[string]objc.IObject, flag bool)
	FontMenu(create bool) Menu
	ConvertFontToSize(fontObj IFont, size float64) Font
	FontWithFamilyTraitsWeightSize(family string, traits FontTraitMask, weight int, size float64) Font
	SetSelectedFontIsMultiple(fontObj IFont, flag bool)
	FontPanel(create bool) FontPanel
	FontNamedHasTraits(fName string, someTraits FontTraitMask) bool
	ModifyFontViaPanel(sender objc.IObject)
	AvailableMembersOfFontFamily(fam string) [][]objc.Object
	CurrentFontAction() FontAction
	Target() objc.Object
	SetTarget(value objc.IObject)
	Action() objc.Selector
	SetAction(value objc.Selector)
	IsMultiple() bool
	AvailableFontFamilies() []string
	SelectedFont() Font
	IsEnabled() bool
	SetEnabled(value bool)
	AvailableFonts() []string
}

// The center of activity for the font-conversion system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager?language=objc
type FontManager struct {
	objc.Object
}

func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.Call[FontManager](fc, objc.Sel("alloc"))
	return rv
}

func FontManager_Alloc() FontManager {
	return FontManagerClass.Alloc()
}

func (fc _FontManagerClass) New() FontManager {
	rv := objc.Call[FontManager](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFontManager() FontManager {
	return FontManagerClass.New()
}

func (f_ FontManager) Init() FontManager {
	rv := objc.Call[FontManager](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether a responder handled the font manager’s action message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462386-sendaction?language=objc
func (f_ FontManager) SendAction() bool {
	rv := objc.Call[bool](f_, objc.Sel("sendAction"))
	return rv
}

// Converts font traits to a new traits mask value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462274-convertfonttraits?language=objc
func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.Call[FontTraitMask](f_, objc.Sel("convertFontTraits:"), traits)
	return rv
}

// Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the NSFont objects themselves). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462329-availablefontnameswithtraits?language=objc
func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.Call[[]string](f_, objc.Sel("availableFontNamesWithTraits:"), someTraits)
	return rv
}

// Returns a localized string with the name of the specified font family and face, if one exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462277-localizednameforfamily?language=objc
func (f_ FontManager) LocalizedNameForFamilyFace(family string, faceKey string) string {
	rv := objc.Call[string](f_, objc.Sel("localizedNameForFamily:face:"), family, faceKey)
	return rv
}

// Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462295-convertattributes?language=objc
func (f_ FontManager) ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](f_, objc.Sel("convertAttributes:"), attributes)
	return rv
}

// Returns an approximation of the specified font's weight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462351-weightoffont?language=objc
func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.Call[int](f_, objc.Sel("weightOfFont:"), objc.Ptr(fontObj))
	return rv
}

// Records the given menu as the application’s Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462381-setfontmenu?language=objc
func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.Call[objc.Void](f_, objc.Sel("setFontMenu:"), objc.Ptr(newMenu))
}

// Returns the traits of the given font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462374-traitsoffont?language=objc
func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.Call[FontTraitMask](f_, objc.Sel("traitsOfFont:"), objc.Ptr(fontObj))
	return rv
}

// Removes a trait from the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462276-removefonttrait?language=objc
func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("removeFontTrait:"), sender)
}

// Opens the Font panel, creating it if necessary, and displays that panel in front of the app's windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462384-orderfrontfontpanel?language=objc
func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("orderFrontFontPanel:"), sender)
}

// Adds a trait to the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462320-addfonttrait?language=objc
func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("addFontTrait:"), sender)
}

// Modifies a trait of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462353-modifyfont?language=objc
func (f_ FontManager) ModifyFont(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("modifyFont:"), sender)
}

// Returns a font object whose weight is greater or lesser than that of the given font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462321-convertweight?language=objc
func (f_ FontManager) ConvertWeightOfFont(upFlag bool, fontObj IFont) Font {
	rv := objc.Call[Font](f_, objc.Sel("convertWeight:ofFont:"), upFlag, objc.Ptr(fontObj))
	return rv
}

// Opens the Font Styles panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462392-orderfrontstylespanel?language=objc
func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("orderFrontStylesPanel:"), sender)
}

// Informs the Font panel that the specified font attributes changed for the selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462270-setselectedattributes?language=objc
func (f_ FontManager) SetSelectedAttributesIsMultiple(attributes map[string]objc.IObject, flag bool) {
	objc.Call[objc.Void](f_, objc.Sel("setSelectedAttributes:isMultiple:"), attributes, flag)
}

// Returns the menu that’s connected to the font conversion system, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462337-fontmenu?language=objc
func (f_ FontManager) FontMenu(create bool) Menu {
	rv := objc.Call[Menu](f_, objc.Sel("fontMenu:"), create)
	return rv
}

// Sets the class that creates the shared Font panel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462388-setfontpanelfactory?language=objc
func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.IClass) {
	objc.Call[objc.Void](fc, objc.Sel("setFontPanelFactory:"), objc.Ptr(factoryId))
}

// Sets the class that creates the shared Font panel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462388-setfontpanelfactory?language=objc
func FontManager_SetFontPanelFactory(factoryId objc.IClass) {
	FontManagerClass.SetFontPanelFactory(factoryId)
}

// Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462378-convertfont?language=objc
func (f_ FontManager) ConvertFontToSize(fontObj IFont, size float64) Font {
	rv := objc.Call[Font](f_, objc.Sel("convertFont:toSize:"), objc.Ptr(fontObj), size)
	return rv
}

// Attempts to load a font with the specified characteristics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462332-fontwithfamily?language=objc
func (f_ FontManager) FontWithFamilyTraitsWeightSize(family string, traits FontTraitMask, weight int, size float64) Font {
	rv := objc.Call[Font](f_, objc.Sel("fontWithFamily:traits:weight:size:"), family, traits, weight, size)
	return rv
}

// Records the specified font as the currently selected font and updates the Font panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462398-setselectedfont?language=objc
func (f_ FontManager) SetSelectedFontIsMultiple(fontObj IFont, flag bool) {
	objc.Call[objc.Void](f_, objc.Sel("setSelectedFont:isMultiple:"), objc.Ptr(fontObj), flag)
}

// Returns the application’s shared Font panel object, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462283-fontpanel?language=objc
func (f_ FontManager) FontPanel(create bool) FontPanel {
	rv := objc.Call[FontPanel](f_, objc.Sel("fontPanel:"), create)
	return rv
}

// Sets the class that creates the shared font manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462402-setfontmanagerfactory?language=objc
func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.IClass) {
	objc.Call[objc.Void](fc, objc.Sel("setFontManagerFactory:"), objc.Ptr(factoryId))
}

// Sets the class that creates the shared font manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462402-setfontmanagerfactory?language=objc
func FontManager_SetFontManagerFactory(factoryId objc.IClass) {
	FontManagerClass.SetFontManagerFactory(factoryId)
}

// Indicates whether the given font has all the specified traits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462327-fontnamed?language=objc
func (f_ FontManager) FontNamedHasTraits(fName string, someTraits FontTraitMask) bool {
	rv := objc.Call[bool](f_, objc.Sel("fontNamed:hasTraits:"), fName, someTraits)
	return rv
}

// Modifies a font trait using input from the Font panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462355-modifyfontviapanel?language=objc
func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("modifyFontViaPanel:"), sender)
}

// Returns an array with one entry for each available member of a font family. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462316-availablemembersoffontfamily?language=objc
func (f_ FontManager) AvailableMembersOfFontFamily(fam string) [][]objc.Object {
	rv := objc.Call[[][]objc.Object](f_, objc.Sel("availableMembersOfFontFamily:"), fam)
	return rv
}

// The current font conversion action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462362-currentfontaction?language=objc
func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.Call[FontAction](f_, objc.Sel("currentFontAction"))
	return rv
}

// The object that receives action messages related to the font manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462380-target?language=objc
func (f_ FontManager) Target() objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("target"))
	return rv
}

// The object that receives action messages related to the font manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462380-target?language=objc
func (f_ FontManager) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setTarget:"), value)
}

// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462349-action?language=objc
func (f_ FontManager) Action() objc.Selector {
	rv := objc.Call[objc.Selector](f_, objc.Sel("action"))
	return rv
}

// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462349-action?language=objc
func (f_ FontManager) SetAction(value objc.Selector) {
	objc.Call[objc.Void](f_, objc.Sel("setAction:"), value)
}

// A Boolean value that indicates whether the currently selected font has multiple fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462376-multiple?language=objc
func (f_ FontManager) IsMultiple() bool {
	rv := objc.Call[bool](f_, objc.Sel("isMultiple"))
	return rv
}

// Returns the shared instance of the font manager for the application, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462360-sharedfontmanager?language=objc
func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.Call[FontManager](fc, objc.Sel("sharedFontManager"))
	return rv
}

// Returns the shared instance of the font manager for the application, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462360-sharedfontmanager?language=objc
func FontManager_SharedFontManager() FontManager {
	return FontManagerClass.SharedFontManager()
}

// The names of the font families available in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462323-availablefontfamilies?language=objc
func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.Call[[]string](f_, objc.Sel("availableFontFamilies"))
	return rv
}

// The currently selected font object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462268-selectedfont?language=objc
func (f_ FontManager) SelectedFont() Font {
	rv := objc.Call[Font](f_, objc.Sel("selectedFont"))
	return rv
}

// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462341-enabled?language=objc
func (f_ FontManager) IsEnabled() bool {
	rv := objc.Call[bool](f_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462341-enabled?language=objc
func (f_ FontManager) SetEnabled(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setEnabled:"), value)
}

// The names of the fonts available in the system (not the NSFont objects themselves). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontmanager/1462372-availablefonts?language=objc
func (f_ FontManager) AvailableFonts() []string {
	rv := objc.Call[[]string](f_, objc.Sel("availableFonts"))
	return rv
}