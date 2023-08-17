// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchableIndex] class.
var SearchableIndexClass = _SearchableIndexClass{objc.GetClass("CSSearchableIndex")}

type _SearchableIndexClass struct {
	objc.Class
}

// An interface definition for the [SearchableIndex] class.
type ISearchableIndex interface {
	objc.IObject
	FetchLastClientStateWithCompletionHandler(completionHandler func(clientState []byte, error foundation.Error))
	DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler func(error foundation.Error))
	BeginIndexBatch()
	DeleteAllSearchableItemsWithCompletionHandler(completionHandler func(error foundation.Error))
	IndexSearchableItemsCompletionHandler(items []ISearchableItem, completionHandler func(error foundation.Error))
	EndIndexBatchWithClientStateCompletionHandler(clientState []byte, completionHandler func(error foundation.Error))
	DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler func(error foundation.Error))
	IndexDelegate() SearchableIndexDelegateWrapper
	SetIndexDelegate(value PSearchableIndexDelegate)
	SetIndexDelegateObject(valueObject objc.IObject)
}

// The on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex?language=objc
type SearchableIndex struct {
	objc.Object
}

func SearchableIndexFrom(ptr unsafe.Pointer) SearchableIndex {
	return SearchableIndex{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SearchableIndexClass) DefaultSearchableIndex() SearchableIndex {
	rv := objc.Call[SearchableIndex](sc, objc.Sel("defaultSearchableIndex"))
	return rv
}

// Returns the default on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620341-defaultsearchableindex?language=objc
func SearchableIndex_DefaultSearchableIndex() SearchableIndex {
	return SearchableIndexClass.DefaultSearchableIndex()
}

func (s_ SearchableIndex) InitWithName(name string) SearchableIndex {
	rv := objc.Call[SearchableIndex](s_, objc.Sel("initWithName:"), name)
	return rv
}

// Returns an on-device index with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620347-initwithname?language=objc
func SearchableIndex_InitWithName(name string) SearchableIndex {
	return SearchableIndexClass.Alloc().InitWithName(name)
}

func (sc _SearchableIndexClass) Alloc() SearchableIndex {
	rv := objc.Call[SearchableIndex](sc, objc.Sel("alloc"))
	return rv
}

func SearchableIndex_Alloc() SearchableIndex {
	return SearchableIndexClass.Alloc()
}

func (sc _SearchableIndexClass) New() SearchableIndex {
	rv := objc.Call[SearchableIndex](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchableIndex() SearchableIndex {
	return SearchableIndexClass.New()
}

func (s_ SearchableIndex) Init() SearchableIndex {
	rv := objc.Call[SearchableIndex](s_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether indexing is available on the current device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620350-isindexingavailable?language=objc
func (sc _SearchableIndexClass) IsIndexingAvailable() bool {
	rv := objc.Call[bool](sc, objc.Sel("isIndexingAvailable"))
	return rv
}

// Returns a Boolean value that indicates whether indexing is available on the current device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620350-isindexingavailable?language=objc
func SearchableIndex_IsIndexingAvailable() bool {
	return SearchableIndexClass.IsIndexingAvailable()
}

// Gets the app’s most recently stored state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620346-fetchlastclientstatewithcompleti?language=objc
func (s_ SearchableIndex) FetchLastClientStateWithCompletionHandler(completionHandler func(clientState []byte, error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("fetchLastClientStateWithCompletionHandler:"), completionHandler)
}

// Removes from the index all items with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620337-deletesearchableitemswithidentif?language=objc
func (s_ SearchableIndex) DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("deleteSearchableItemsWithIdentifiers:completionHandler:"), identifiers, completionHandler)
}

// Begins a batch of updates to an index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620331-beginindexbatch?language=objc
func (s_ SearchableIndex) BeginIndexBatch() {
	objc.Call[objc.Void](s_, objc.Sel("beginIndexBatch"))
}

// Deletes all searchable items from the index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620342-deleteallsearchableitemswithcomp?language=objc
func (s_ SearchableIndex) DeleteAllSearchableItemsWithCompletionHandler(completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("deleteAllSearchableItemsWithCompletionHandler:"), completionHandler)
}

// Adds or updates items in the index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620333-indexsearchableitems?language=objc
func (s_ SearchableIndex) IndexSearchableItemsCompletionHandler(items []ISearchableItem, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("indexSearchableItems:completionHandler:"), items, completionHandler)
}

// Ends a batch of index updates and stores the specified state information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620344-endindexbatchwithclientstate?language=objc
func (s_ SearchableIndex) EndIndexBatchWithClientStateCompletionHandler(clientState []byte, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("endIndexBatchWithClientState:completionHandler:"), clientState, completionHandler)
}

// Removes from the index all searchable items associated with the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620351-deletesearchableitemswithdomaini?language=objc
func (s_ SearchableIndex) DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler func(error foundation.Error)) {
	objc.Call[objc.Void](s_, objc.Sel("deleteSearchableItemsWithDomainIdentifiers:completionHandler:"), domainIdentifiers, completionHandler)
}

// The delegate object that can handle index-management tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620354-indexdelegate?language=objc
func (s_ SearchableIndex) IndexDelegate() SearchableIndexDelegateWrapper {
	rv := objc.Call[SearchableIndexDelegateWrapper](s_, objc.Sel("indexDelegate"))
	return rv
}

// The delegate object that can handle index-management tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620354-indexdelegate?language=objc
func (s_ SearchableIndex) SetIndexDelegate(value PSearchableIndexDelegate) {
	po0 := objc.WrapAsProtocol("CSSearchableIndexDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setIndexDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setIndexDelegate:"), po0)
}

// The delegate object that can handle index-management tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableindex/1620354-indexdelegate?language=objc
func (s_ SearchableIndex) SetIndexDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setIndexDelegate:"), objc.Ptr(valueObject))
}