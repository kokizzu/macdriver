// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that an outline view calls to retrieve data and information about it from the data source delegate, and—optionally—to update data values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource?language=objc
type POutlineViewDataSource interface {
	// optional
	OutlineViewSortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	HasOutlineViewSortDescriptorsDidChange() bool

	// optional
	OutlineViewChildOfItem(outlineView OutlineView, index int, item objc.Object) objc.Object
	HasOutlineViewChildOfItem() bool

	// optional
	OutlineViewPersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.Object
	HasOutlineViewPersistentObjectForItem() bool

	// optional
	OutlineViewValidateDropProposedItemProposedChildIndex(outlineView OutlineView, info DraggingInfoObject, item objc.Object, index int) DragOperation
	HasOutlineViewValidateDropProposedItemProposedChildIndex() bool

	// optional
	OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	HasOutlineViewDraggingSessionWillBeginAtPointForItems() bool

	// optional
	OutlineViewAcceptDropItemChildIndex(outlineView OutlineView, info DraggingInfoObject, item objc.Object, index int) bool
	HasOutlineViewAcceptDropItemChildIndex() bool

	// optional
	OutlineViewDraggingSessionEndedAtPointOperation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	HasOutlineViewDraggingSessionEndedAtPointOperation() bool

	// optional
	OutlineViewSetObjectValueForTableColumnByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object)
	HasOutlineViewSetObjectValueForTableColumnByItem() bool

	// optional
	OutlineViewPasteboardWriterForItem(outlineView OutlineView, item objc.Object) PasteboardWritingObject
	HasOutlineViewPasteboardWriterForItem() bool

	// optional
	OutlineViewNumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int
	HasOutlineViewNumberOfChildrenOfItem() bool

	// optional
	OutlineViewObjectValueForTableColumnByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.Object
	HasOutlineViewObjectValueForTableColumnByItem() bool

	// optional
	OutlineViewUpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoObject)
	HasOutlineViewUpdateDraggingItemsForDrag() bool

	// optional
	OutlineViewItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.Object
	HasOutlineViewItemForPersistentObject() bool

	// optional
	OutlineViewIsItemExpandable(outlineView OutlineView, item objc.Object) bool
	HasOutlineViewIsItemExpandable() bool
}

// ensure impl type implements protocol interface
var _ POutlineViewDataSource = (*OutlineViewDataSourceObject)(nil)

// A concrete type for the [POutlineViewDataSource] protocol.
type OutlineViewDataSourceObject struct {
	objc.Object
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewSortDescriptorsDidChange() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:sortDescriptorsDidChange:"))
}

// Invoked by an outline view to notify the data source that the descriptors changed and the data may need to be resorted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1535892-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewSortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:sortDescriptorsDidChange:"), outlineView, oldDescriptors)
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewChildOfItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:child:ofItem:"))
}

// Returns the child item at the specified index of a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1528977-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewChildOfItem(outlineView OutlineView, index int, item objc.Object) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("outlineView:child:ofItem:"), outlineView, index, item)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewPersistentObjectForItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:persistentObjectForItem:"))
}

// Invoked by outlineView to return an archived object for item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1532545-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewPersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("outlineView:persistentObjectForItem:"), outlineView, item)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewValidateDropProposedItemProposedChildIndex() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:validateDrop:proposedItem:proposedChildIndex:"))
}

// Used by an outline view to determine a valid drop target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1533597-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewValidateDropProposedItemProposedChildIndex(outlineView OutlineView, info DraggingInfoObject, item objc.Object, index int) DragOperation {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.Call[DragOperation](o_, objc.Sel("outlineView:validateDrop:proposedItem:proposedChildIndex:"), outlineView, po1, item, index)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewDraggingSessionWillBeginAtPointForItems() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:draggingSession:willBeginAtPoint:forItems:"))
}

// Implement this method know when the given dragging session is about to begin and potentially modify the dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1535142-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:draggingSession:willBeginAtPoint:forItems:"), outlineView, session, screenPoint, draggedItems)
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewAcceptDropItemChildIndex() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:acceptDrop:item:childIndex:"))
}

// Returns a Boolean value that indicates whether a drop operation was successful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1529572-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewAcceptDropItemChildIndex(outlineView OutlineView, info DraggingInfoObject, item objc.Object, index int) bool {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.Call[bool](o_, objc.Sel("outlineView:acceptDrop:item:childIndex:"), outlineView, po1, item, index)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewDraggingSessionEndedAtPointOperation() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:draggingSession:endedAtPoint:operation:"))
}

// Implement this method to know when the given dragging session has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1532073-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewDraggingSessionEndedAtPointOperation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:draggingSession:endedAtPoint:operation:"), outlineView, session, screenPoint, operation)
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewSetObjectValueForTableColumnByItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:setObjectValue:forTableColumn:byItem:"))
}

// Set the data object for a given item in a given column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1534817-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewSetObjectValueForTableColumnByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:setObjectValue:forTableColumn:byItem:"), outlineView, object, tableColumn, item)
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewPasteboardWriterForItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:pasteboardWriterForItem:"))
}

// Implement this method to enable the table to be an NSDraggingSource that supports dragging multiple items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1525837-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewPasteboardWriterForItem(outlineView OutlineView, item objc.Object) PasteboardWritingObject {
	rv := objc.Call[PasteboardWritingObject](o_, objc.Sel("outlineView:pasteboardWriterForItem:"), outlineView, item)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewNumberOfChildrenOfItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:numberOfChildrenOfItem:"))
}

// Returns the number of child items encompassed by a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1535549-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewNumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int {
	rv := objc.Call[int](o_, objc.Sel("outlineView:numberOfChildrenOfItem:"), outlineView, item)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewObjectValueForTableColumnByItem() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:objectValueForTableColumn:byItem:"))
}

// Invoked by outlineView to return the data object associated with the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1531606-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewObjectValueForTableColumnByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("outlineView:objectValueForTableColumn:byItem:"), outlineView, tableColumn, item)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewUpdateDraggingItemsForDrag() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:updateDraggingItemsForDrag:"))
}

// Implement this method to enable the table to update dragging items as they are dragged over the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1534424-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewUpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoObject) {
	po1 := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.Call[objc.Void](o_, objc.Sel("outlineView:updateDraggingItemsForDrag:"), outlineView, po1)
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewItemForPersistentObject() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:itemForPersistentObject:"))
}

// Invoked by outlineView to return the item for the archived object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1533602-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("outlineView:itemForPersistentObject:"), outlineView, object)
	return rv
}

func (o_ OutlineViewDataSourceObject) HasOutlineViewIsItemExpandable() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:isItemExpandable:"))
}

// Returns a Boolean value that indicates whether the a given item is expandable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1535198-outlineview?language=objc
func (o_ OutlineViewDataSourceObject) OutlineViewIsItemExpandable(outlineView OutlineView, item objc.Object) bool {
	rv := objc.Call[bool](o_, objc.Sel("outlineView:isItemExpandable:"), outlineView, item)
	return rv
}
