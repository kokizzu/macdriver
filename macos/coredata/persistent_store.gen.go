// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentStore] class.
var PersistentStoreClass = _PersistentStoreClass{objc.GetClass("NSPersistentStore")}

type _PersistentStoreClass struct {
	objc.Class
}

// An interface definition for the [PersistentStore] class.
type IPersistentStore interface {
	objc.IObject
	WillRemoveFromPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator)
	DidAddToPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator)
	LoadMetadata(error foundation.IError) bool
	IsReadOnly() bool
	SetReadOnly(value bool)
	PersistentStoreCoordinator() PersistentStoreCoordinator
	Options() foundation.Dictionary
	Metadata() map[string]objc.Object
	SetMetadata(value map[string]objc.IObject)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	CoreSpotlightExporter() CoreDataCoreSpotlightDelegate
	ConfigurationName() string
	Type() string
	Identifier() string
	SetIdentifier(value string)
}

// The abstract base class for all Core Data persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore?language=objc
type PersistentStore struct {
	objc.Object
}

func PersistentStoreFrom(ptr unsafe.Pointer) PersistentStore {
	return PersistentStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentStore) InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), objc.Ptr(root), name, objc.Ptr(url), options)
	return rv
}

// Returns a store initialized with the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506232-initwithpersistentstorecoordinat?language=objc
func PersistentStore_InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options foundation.Dictionary) PersistentStore {
	return PersistentStoreClass.Alloc().InitWithPersistentStoreCoordinatorConfigurationNameURLOptions(root, name, url, options)
}

func (pc _PersistentStoreClass) Alloc() PersistentStore {
	rv := objc.Call[PersistentStore](pc, objc.Sel("alloc"))
	return rv
}

func PersistentStore_Alloc() PersistentStore {
	return PersistentStoreClass.Alloc()
}

func (pc _PersistentStoreClass) New() PersistentStore {
	rv := objc.Call[PersistentStore](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStore() PersistentStore {
	return PersistentStoreClass.New()
}

func (p_ PersistentStore) Init() PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("init"))
	return rv
}

// Invoked before the persistent store is removed from the persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506731-willremovefrompersistentstorecoo?language=objc
func (p_ PersistentStore) WillRemoveFromPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator) {
	objc.Call[objc.Void](p_, objc.Sel("willRemoveFromPersistentStoreCoordinator:"), objc.Ptr(coordinator))
}

// Sets the metadata for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506824-setmetadata?language=objc
func (pc _PersistentStoreClass) SetMetadataForPersistentStoreWithURLError(metadata map[string]objc.IObject, url foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](pc, objc.Sel("setMetadata:forPersistentStoreWithURL:error:"), metadata, objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Sets the metadata for the store at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506824-setmetadata?language=objc
func PersistentStore_SetMetadataForPersistentStoreWithURLError(metadata map[string]objc.IObject, url foundation.IURL, error foundation.IError) bool {
	return PersistentStoreClass.SetMetadataForPersistentStoreWithURLError(metadata, url, error)
}

// Returns the metadata from the persistent store at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506741-metadataforpersistentstorewithur?language=objc
func (pc _PersistentStoreClass) MetadataForPersistentStoreWithURLError(url foundation.IURL, error foundation.IError) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](pc, objc.Sel("metadataForPersistentStoreWithURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns the metadata from the persistent store at the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506741-metadataforpersistentstorewithur?language=objc
func PersistentStore_MetadataForPersistentStoreWithURLError(url foundation.IURL, error foundation.IError) map[string]objc.Object {
	return PersistentStoreClass.MetadataForPersistentStoreWithURLError(url, error)
}

// Invoked after the persistent store has been added to the persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506873-didaddtopersistentstorecoordinat?language=objc
func (p_ PersistentStore) DidAddToPersistentStoreCoordinator(coordinator IPersistentStoreCoordinator) {
	objc.Call[objc.Void](p_, objc.Sel("didAddToPersistentStoreCoordinator:"), objc.Ptr(coordinator))
}

// Instructs the persistent store to load its metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506273-loadmetadata?language=objc
func (p_ PersistentStore) LoadMetadata(error foundation.IError) bool {
	rv := objc.Call[bool](p_, objc.Sel("loadMetadata:"), objc.Ptr(error))
	return rv
}

// Returns the migration manager class for this store class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506361-migrationmanagerclass?language=objc
func (pc _PersistentStoreClass) MigrationManagerClass() objc.Class {
	rv := objc.Call[objc.Class](pc, objc.Sel("migrationManagerClass"))
	return rv
}

// Returns the migration manager class for this store class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506361-migrationmanagerclass?language=objc
func PersistentStore_MigrationManagerClass() objc.Class {
	return PersistentStoreClass.MigrationManagerClass()
}

// A Boolean value that indicates whether the persistent store is read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506183-readonly?language=objc
func (p_ PersistentStore) IsReadOnly() bool {
	rv := objc.Call[bool](p_, objc.Sel("isReadOnly"))
	return rv
}

// A Boolean value that indicates whether the persistent store is read-only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506183-readonly?language=objc
func (p_ PersistentStore) SetReadOnly(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setReadOnly:"), value)
}

// The persistent store coordinator that loads the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506226-persistentstorecoordinator?language=objc
func (p_ PersistentStore) PersistentStoreCoordinator() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](p_, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The options that Core Data uses to create the store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506821-options?language=objc
func (p_ PersistentStore) Options() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("options"))
	return rv
}

// The metadata for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506564-metadata?language=objc
func (p_ PersistentStore) Metadata() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("metadata"))
	return rv
}

// The metadata for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506564-metadata?language=objc
func (p_ PersistentStore) SetMetadata(value map[string]objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setMetadata:"), value)
}

// The URL for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506700-url?language=objc
func (p_ PersistentStore) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// The URL for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506700-url?language=objc
func (p_ PersistentStore) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), objc.Ptr(value))
}

// The spotlight exporter associated with this persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/2897191-corespotlightexporter?language=objc
func (p_ PersistentStore) CoreSpotlightExporter() CoreDataCoreSpotlightDelegate {
	rv := objc.Call[CoreDataCoreSpotlightDelegate](p_, objc.Sel("coreSpotlightExporter"))
	return rv
}

// The name of the managed object model configuration that creates the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506620-configurationname?language=objc
func (p_ PersistentStore) ConfigurationName() string {
	rv := objc.Call[string](p_, objc.Sel("configurationName"))
	return rv
}

// The type string of the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506250-type?language=objc
func (p_ PersistentStore) Type() string {
	rv := objc.Call[string](p_, objc.Sel("type"))
	return rv
}

// The unique identifier for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506215-identifier?language=objc
func (p_ PersistentStore) Identifier() string {
	rv := objc.Call[string](p_, objc.Sel("identifier"))
	return rv
}

// The unique identifier for the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/1506215-identifier?language=objc
func (p_ PersistentStore) SetIdentifier(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setIdentifier:"), value)
}