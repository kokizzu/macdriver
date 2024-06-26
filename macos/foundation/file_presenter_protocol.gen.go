// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/darwinkit/objc"
)

// The interface a file coordinator uses to inform an object presenting a file about changes to that file made elsewhere in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter?language=objc
type PFilePresenter interface {
	// optional
	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler func(errorOrNil Error))
	HasAccommodatePresentedItemDeletionWithCompletionHandler() bool

	// optional
	PresentedItemDidChangeUbiquityAttributes(attributes Set)
	HasPresentedItemDidChangeUbiquityAttributes() bool

	// optional
	PresentedSubitemAtURLDidResolveConflictVersion(url URL, version FileVersion)
	HasPresentedSubitemAtURLDidResolveConflictVersion() bool

	// optional
	PresentedItemDidResolveConflictVersion(version FileVersion)
	HasPresentedItemDidResolveConflictVersion() bool

	// optional
	AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url URL, completionHandler func(errorOrNil Error))
	HasAccommodatePresentedSubitemDeletionAtURLCompletionHandler() bool

	// optional
	PresentedSubitemDidChangeAtURL(url URL)
	HasPresentedSubitemDidChangeAtURL() bool

	// optional
	PresentedSubitemAtURLDidGainVersion(url URL, version FileVersion)
	HasPresentedSubitemAtURLDidGainVersion() bool

	// optional
	PresentedSubitemAtURLDidLoseVersion(url URL, version FileVersion)
	HasPresentedSubitemAtURLDidLoseVersion() bool

	// optional
	PresentedItemDidGainVersion(version FileVersion)
	HasPresentedItemDidGainVersion() bool

	// optional
	PresentedSubitemAtURLDidMoveToURL(oldURL URL, newURL URL)
	HasPresentedSubitemAtURLDidMoveToURL() bool

	// optional
	SavePresentedItemChangesWithCompletionHandler(completionHandler func(errorOrNil Error))
	HasSavePresentedItemChangesWithCompletionHandler() bool

	// optional
	PresentedItemDidLoseVersion(version FileVersion)
	HasPresentedItemDidLoseVersion() bool

	// optional
	PresentedItemDidMoveToURL(newURL URL)
	HasPresentedItemDidMoveToURL() bool

	// optional
	PresentedSubitemDidAppearAtURL(url URL)
	HasPresentedSubitemDidAppearAtURL() bool

	// optional
	PresentedItemDidChange()
	HasPresentedItemDidChange() bool

	// optional
	RelinquishPresentedItemToWriter(writer func(arg0 func()))
	HasRelinquishPresentedItemToWriter() bool

	// optional
	RelinquishPresentedItemToReader(reader func(arg0 func()))
	HasRelinquishPresentedItemToReader() bool

	// optional
	ObservedPresentedItemUbiquityAttributes() Set
	HasObservedPresentedItemUbiquityAttributes() bool

	// optional
	PresentedItemURL() URL
	HasPresentedItemURL() bool

	// optional
	PrimaryPresentedItemURL() URL
	HasPrimaryPresentedItemURL() bool

	// optional
	PresentedItemOperationQueue() OperationQueue
	HasPresentedItemOperationQueue() bool
}

// ensure impl type implements protocol interface
var _ PFilePresenter = (*FilePresenterObject)(nil)

// A concrete type for the [PFilePresenter] protocol.
type FilePresenterObject struct {
	objc.Object
}

func (f_ FilePresenterObject) HasAccommodatePresentedItemDeletionWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"))
}

// Tells your object that its presented item is about to be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414732-accommodatepresenteditemdeletion?language=objc
func (f_ FilePresenterObject) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), completionHandler)
}

func (f_ FilePresenterObject) HasPresentedItemDidChangeUbiquityAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidChangeUbiquityAttributes:"))
}

// Tells your object that the file or file package's ubiquity attributes have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/2909021-presenteditemdidchangeubiquityat?language=objc
func (f_ FilePresenterObject) PresentedItemDidChangeUbiquityAttributes(attributes Set) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), attributes)
}

func (f_ FilePresenterObject) HasPresentedSubitemAtURLDidResolveConflictVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemAtURL:didResolveConflictVersion:"))
}

// Tells the delegate that the item inside the presented directory had a version conflict resolved by an outside entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1416913-presentedsubitematurl?language=objc
func (f_ FilePresenterObject) PresentedSubitemAtURLDidResolveConflictVersion(url URL, version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemAtURL:didResolveConflictVersion:"), url, version)
}

func (f_ FilePresenterObject) HasPresentedItemDidResolveConflictVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidResolveConflictVersion:"))
}

// Tells the delegate that some other entity resolved a version conflict for the presenter’s file or file package. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1418445-presenteditemdidresolveconflictv?language=objc
func (f_ FilePresenterObject) PresentedItemDidResolveConflictVersion(version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidResolveConflictVersion:"), version)
}

func (f_ FilePresenterObject) HasAccommodatePresentedSubitemDeletionAtURLCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("accommodatePresentedSubitemDeletionAtURL:completionHandler:"))
}

// Tells the delegate that some entity wants to delete an item that is inside of a presented directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415657-accommodatepresentedsubitemdelet?language=objc
func (f_ FilePresenterObject) AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url URL, completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("accommodatePresentedSubitemDeletionAtURL:completionHandler:"), url, completionHandler)
}

func (f_ FilePresenterObject) HasPresentedSubitemDidChangeAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemDidChangeAtURL:"))
}

// Tells the delegate that the contents or attributes of the specified item changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1411135-presentedsubitemdidchangeaturl?language=objc
func (f_ FilePresenterObject) PresentedSubitemDidChangeAtURL(url URL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemDidChangeAtURL:"), url)
}

func (f_ FilePresenterObject) HasPresentedSubitemAtURLDidGainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemAtURL:didGainVersion:"))
}

// Tells the delegate that the item inside the presented directory gained a new version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415472-presentedsubitematurl?language=objc
func (f_ FilePresenterObject) PresentedSubitemAtURLDidGainVersion(url URL, version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemAtURL:didGainVersion:"), url, version)
}

func (f_ FilePresenterObject) HasPresentedSubitemAtURLDidLoseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemAtURL:didLoseVersion:"))
}

// Tells the delegate that the item inside the presented directory lost an existing version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1413957-presentedsubitematurl?language=objc
func (f_ FilePresenterObject) PresentedSubitemAtURLDidLoseVersion(url URL, version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemAtURL:didLoseVersion:"), url, version)
}

func (f_ FilePresenterObject) HasPresentedItemDidGainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidGainVersion:"))
}

// Tells the delegate that a new version of the file or file package was added. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415018-presenteditemdidgainversion?language=objc
func (f_ FilePresenterObject) PresentedItemDidGainVersion(version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidGainVersion:"), version)
}

func (f_ FilePresenterObject) HasPresentedSubitemAtURLDidMoveToURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemAtURL:didMoveToURL:"))
}

// Tells the delegate that an item in the presented directory moved to a new location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1409465-presentedsubitematurl?language=objc
func (f_ FilePresenterObject) PresentedSubitemAtURLDidMoveToURL(oldURL URL, newURL URL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemAtURL:didMoveToURL:"), oldURL, newURL)
}

func (f_ FilePresenterObject) HasSavePresentedItemChangesWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("savePresentedItemChangesWithCompletionHandler:"))
}

// Tells your object to save any unsaved changes for the presented item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414407-savepresenteditemchangeswithcomp?language=objc
func (f_ FilePresenterObject) SavePresentedItemChangesWithCompletionHandler(completionHandler func(errorOrNil Error)) {
	objc.Call[objc.Void](f_, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), completionHandler)
}

func (f_ FilePresenterObject) HasPresentedItemDidLoseVersion() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidLoseVersion:"))
}

// Tells the delegate that a version of the file or file package was removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1417258-presenteditemdidloseversion?language=objc
func (f_ FilePresenterObject) PresentedItemDidLoseVersion(version FileVersion) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidLoseVersion:"), version)
}

func (f_ FilePresenterObject) HasPresentedItemDidMoveToURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidMoveToURL:"))
}

// Tells your object that the presented item moved or was renamed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1417861-presenteditemdidmovetourl?language=objc
func (f_ FilePresenterObject) PresentedItemDidMoveToURL(newURL URL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidMoveToURL:"), newURL)
}

func (f_ FilePresenterObject) HasPresentedSubitemDidAppearAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedSubitemDidAppearAtURL:"))
}

// Tells the delegate that an item was added to the presented directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1408642-presentedsubitemdidappearaturl?language=objc
func (f_ FilePresenterObject) PresentedSubitemDidAppearAtURL(url URL) {
	objc.Call[objc.Void](f_, objc.Sel("presentedSubitemDidAppearAtURL:"), url)
}

func (f_ FilePresenterObject) HasPresentedItemDidChange() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemDidChange"))
}

// Tells your object that the presented item’s contents or attributes changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1416103-presenteditemdidchange?language=objc
func (f_ FilePresenterObject) PresentedItemDidChange() {
	objc.Call[objc.Void](f_, objc.Sel("presentedItemDidChange"))
}

func (f_ FilePresenterObject) HasRelinquishPresentedItemToWriter() bool {
	return f_.RespondsToSelector(objc.Sel("relinquishPresentedItemToWriter:"))
}

// Notifies your object that another object or process wants to write to the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1413688-relinquishpresenteditemtowriter?language=objc
func (f_ FilePresenterObject) RelinquishPresentedItemToWriter(writer func(arg0 func())) {
	objc.Call[objc.Void](f_, objc.Sel("relinquishPresentedItemToWriter:"), writer)
}

func (f_ FilePresenterObject) HasRelinquishPresentedItemToReader() bool {
	return f_.RespondsToSelector(objc.Sel("relinquishPresentedItemToReader:"))
}

// Notifies your object that another object or process wants to read the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1410743-relinquishpresenteditemtoreader?language=objc
func (f_ FilePresenterObject) RelinquishPresentedItemToReader(reader func(arg0 func())) {
	objc.Call[objc.Void](f_, objc.Sel("relinquishPresentedItemToReader:"), reader)
}

func (f_ FilePresenterObject) HasObservedPresentedItemUbiquityAttributes() bool {
	return f_.RespondsToSelector(objc.Sel("observedPresentedItemUbiquityAttributes"))
}

// A list of ubiquity attributes used to generate and send notifications whenever an attribute in the list changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/2909022-observedpresenteditemubiquityatt?language=objc
func (f_ FilePresenterObject) ObservedPresentedItemUbiquityAttributes() Set {
	rv := objc.Call[Set](f_, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return rv
}

func (f_ FilePresenterObject) HasPresentedItemURL() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemURL"))
}

// The URL of the presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1414861-presenteditemurl?language=objc
func (f_ FilePresenterObject) PresentedItemURL() URL {
	rv := objc.Call[URL](f_, objc.Sel("presentedItemURL"))
	return rv
}

func (f_ FilePresenterObject) HasPrimaryPresentedItemURL() bool {
	return f_.RespondsToSelector(objc.Sel("primaryPresentedItemURL"))
}

// The URL of a secondary item’s primary presented file or directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415415-primarypresenteditemurl?language=objc
func (f_ FilePresenterObject) PrimaryPresentedItemURL() URL {
	rv := objc.Call[URL](f_, objc.Sel("primaryPresentedItemURL"))
	return rv
}

func (f_ FilePresenterObject) HasPresentedItemOperationQueue() bool {
	return f_.RespondsToSelector(objc.Sel("presentedItemOperationQueue"))
}

// The operation queue in which to execute presenter-related messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilepresenter/1415250-presenteditemoperationqueue?language=objc
func (f_ FilePresenterObject) PresentedItemOperationQueue() OperationQueue {
	rv := objc.Call[OperationQueue](f_, objc.Sel("presentedItemOperationQueue"))
	return rv
}
