// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ManagedObject] class.
var ManagedObjectClass = _ManagedObjectClass{objc.GetClass("NSManagedObject")}

type _ManagedObjectClass struct {
	objc.Class
}

// An interface definition for the [ManagedObject] class.
type IManagedObject interface {
	objc.IObject
	WillTurnIntoFault()
	AwakeFromFetch()
	ValueForKey(key string) objc.Object
	WillChangeValueForKey(key string)
	ChangedValues() map[string]objc.Object
	PrepareForDeletion()
	InitWithEntityInsertIntoManagedObjectContext(entity IEntityDescription, context IManagedObjectContext) ManagedObject
	ValidateForDelete(error unsafe.Pointer) bool
	ValidateValueForKeyError(value unsafe.Pointer, key string, error unsafe.Pointer) bool
	DidAccessValueForKey(key string)
	HasFaultForRelationshipNamed(key string) bool
	WillSave()
	SetObservationInfo(inObservationInfo unsafe.Pointer)
	DidTurnIntoFault()
	AwakeFromInsert()
	DidChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind foundation.KeyValueSetMutationKind, inObjects foundation.ISet)
	ObservationInfo() unsafe.Pointer
	ObjectIDsForRelationshipNamed(key string) []ManagedObjectID
	WillAccessValueForKey(key string)
	SetPrimitiveValueForKey(value objc.IObject, key string)
	WillChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind foundation.KeyValueSetMutationKind, inObjects foundation.ISet)
	ValidateForInsert(error unsafe.Pointer) bool
	PrimitiveValueForKey(key string) objc.Object
	AwakeFromSnapshotEvents(flags SnapshotEventType)
	CommittedValuesForKeys(keys []string) map[string]objc.Object
	DidChangeValueForKey(key string)
	ChangedValuesForCurrentEvent() map[string]objc.Object
	SetValueForKey(value objc.IObject, key string)
	DidSave()
	ValidateForUpdate(error unsafe.Pointer) bool
	HasPersistentChangedValues() bool
	FaultingState() uint
	HasChanges() bool
	ManagedObjectContext() ManagedObjectContext
	IsInserted() bool
	IsDeleted() bool
	ObjectID() ManagedObjectID
	IsFault() bool
	IsUpdated() bool
}

// The base class that all Core Data model objects inherit from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject?language=objc
type ManagedObject struct {
	objc.Object
}

func ManagedObjectFrom(ptr unsafe.Pointer) ManagedObject {
	return ManagedObject{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ ManagedObject) InitWithContext(moc IManagedObjectContext) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("initWithContext:"), moc)
	return rv
}

// Initializes a managed object subclass and inserts it into the specified managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640602-initwithcontext?language=objc
func NewManagedObjectWithContext(moc IManagedObjectContext) ManagedObject {
	instance := ManagedObjectClass.Alloc().InitWithContext(moc)
	instance.Autorelease()
	return instance
}

func (mc _ManagedObjectClass) Alloc() ManagedObject {
	rv := objc.Call[ManagedObject](mc, objc.Sel("alloc"))
	return rv
}

func (mc _ManagedObjectClass) New() ManagedObject {
	rv := objc.Call[ManagedObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewManagedObject() ManagedObject {
	return ManagedObjectClass.New()
}

func (m_ ManagedObject) Init() ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("init"))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object before converting it to a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506537-willturnintofault?language=objc
func (m_ ManagedObject) WillTurnIntoFault() {
	objc.Call[objc.Void](m_, objc.Sel("willTurnIntoFault"))
}

// Provides an opportunity to add code into the life cycle of the managed object when fufilling it from a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506424-awakefromfetch?language=objc
func (m_ ManagedObject) AwakeFromFetch() {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromFetch"))
}

// Returns the value for the property specified by key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506613-valueforkey?language=objc
func (m_ ManagedObject) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("valueForKey:"), key)
	return rv
}

// Provides an opportunity to respond when a value of a given property is about to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506229-willchangevalueforkey?language=objc
func (m_ ManagedObject) WillChangeValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("willChangeValueForKey:"), key)
}

// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506775-changedvalues?language=objc
func (m_ ManagedObject) ChangedValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("changedValues"))
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object before deleting it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506674-preparefordeletion?language=objc
func (m_ ManagedObject) PrepareForDeletion() {
	objc.Call[objc.Void](m_, objc.Sel("prepareForDeletion"))
}

// Initializes a managed object from an entity description and inserts it into the specified managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506357-initwithentity?language=objc
func (m_ ManagedObject) InitWithEntityInsertIntoManagedObjectContext(entity IEntityDescription, context IManagedObjectContext) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("initWithEntity:insertIntoManagedObjectContext:"), entity, context)
	return rv
}

// Determines whether the managed object can be deleted in its current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506195-validatefordelete?language=objc
func (m_ ManagedObject) ValidateForDelete(error unsafe.Pointer) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForDelete:"), error)
	return rv
}

// Validates a property value for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506776-validatevalue?language=objc
func (m_ ManagedObject) ValidateValueForKeyError(value unsafe.Pointer, key string, error unsafe.Pointer) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateValue:forKey:error:"), value, key, error)
	return rv
}

// Provides support for key-value observing access notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506865-didaccessvalueforkey?language=objc
func (m_ ManagedObject) DidAccessValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("didAccessValueForKey:"), key)
}

// Returns a Boolean value that indicates whether the relationship for a given key is a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506973-hasfaultforrelationshipnamed?language=objc
func (m_ ManagedObject) HasFaultForRelationshipNamed(key string) bool {
	rv := objc.Call[bool](m_, objc.Sel("hasFaultForRelationshipNamed:"), key)
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object before saving it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506209-willsave?language=objc
func (m_ ManagedObject) WillSave() {
	objc.Call[objc.Void](m_, objc.Sel("willSave"))
}

// Sets the observation info of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506535-setobservationinfo?language=objc
func (m_ ManagedObject) SetObservationInfo(inObservationInfo unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("setObservationInfo:"), inObservationInfo)
}

// Provides an opportunity to add code into the life cycle of the managed object after converting it to a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506470-didturnintofault?language=objc
func (m_ ManagedObject) DidTurnIntoFault() {
	objc.Call[objc.Void](m_, objc.Sel("didTurnIntoFault"))
}

// Provides an opportunity to add code into the life cycle of the managed object when initially creating it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506548-awakefrominsert?language=objc
func (m_ ManagedObject) AwakeFromInsert() {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromInsert"))
}

// Provides an opportunity to respond when a change was made to a specified to-many relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506936-didchangevalueforkey?language=objc
func (m_ ManagedObject) DidChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind foundation.KeyValueSetMutationKind, inObjects foundation.ISet) {
	objc.Call[objc.Void](m_, objc.Sel("didChangeValueForKey:withSetMutation:usingObjects:"), inKey, inMutationKind, inObjects)
}

// Returns the observation info of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506210-observationinfo?language=objc
func (m_ ManagedObject) ObservationInfo() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](m_, objc.Sel("observationInfo"))
	return rv
}

// Returns the object IDs for all of the managed objects that are in the named relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506201-objectidsforrelationshipnamed?language=objc
func (m_ ManagedObject) ObjectIDsForRelationshipNamed(key string) []ManagedObjectID {
	rv := objc.Call[[]ManagedObjectID](m_, objc.Sel("objectIDsForRelationshipNamed:"), key)
	return rv
}

// Provides support for key-value observing access notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1507001-willaccessvalueforkey?language=objc
func (m_ ManagedObject) WillAccessValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("willAccessValueForKey:"), key)
}

// Sets the value of a given property in the managed object's private internal storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506960-setprimitivevalue?language=objc
func (m_ ManagedObject) SetPrimitiveValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](m_, objc.Sel("setPrimitiveValue:forKey:"), value, key)
}

// Provides an opportunity to respond when a change is about to be made to a specified to-many relationship. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506801-willchangevalueforkey?language=objc
func (m_ ManagedObject) WillChangeValueForKeyWithSetMutationUsingObjects(inKey string, inMutationKind foundation.KeyValueSetMutationKind, inObjects foundation.ISet) {
	objc.Call[objc.Void](m_, objc.Sel("willChangeValueForKey:withSetMutation:usingObjects:"), inKey, inMutationKind, inObjects)
}

// Determines whether the managed object can be inserted in its current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506683-validateforinsert?language=objc
func (m_ ManagedObject) ValidateForInsert(error unsafe.Pointer) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForInsert:"), error)
	return rv
}

// Returns the value for the specified property from the managed object’s private internal storage . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506728-primitivevalueforkey?language=objc
func (m_ ManagedObject) PrimitiveValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("primitiveValueForKey:"), key)
	return rv
}

// Provides an opportunity to add code into the life cycle of the managed object when fulfilling it from a snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506861-awakefromsnapshotevents?language=objc
func (m_ ManagedObject) AwakeFromSnapshotEvents(flags SnapshotEventType) {
	objc.Call[objc.Void](m_, objc.Sel("awakeFromSnapshotEvents:"), flags)
}

// Returns an initialized fetch request with the entity this subclass represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640605-fetchrequest?language=objc
func (mc _ManagedObjectClass) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](mc, objc.Sel("fetchRequest"))
	return rv
}

// Returns an initialized fetch request with the entity this subclass represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640605-fetchrequest?language=objc
func ManagedObject_FetchRequest() FetchRequest {
	return ManagedObjectClass.FetchRequest()
}

// Returns a dictionary of the most recent fetched or saved values of the managed object for the properties of the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506771-committedvaluesforkeys?language=objc
func (m_ ManagedObject) CommittedValuesForKeys(keys []string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("committedValuesForKeys:"), keys)
	return rv
}

// Provides an opportunity to respond when a value of a given property has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506976-didchangevalueforkey?language=objc
func (m_ ManagedObject) DidChangeValueForKey(key string) {
	objc.Call[objc.Void](m_, objc.Sel("didChangeValueForKey:"), key)
}

// Returns a dictionary containing the keys and new values of persistent properties with changes since the last fetching or saving of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506472-changedvaluesforcurrentevent?language=objc
func (m_ ManagedObject) ChangedValuesForCurrentEvent() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("changedValuesForCurrentEvent"))
	return rv
}

// Sets the specified property of the managed object to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506397-setvalue?language=objc
func (m_ ManagedObject) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](m_, objc.Sel("setValue:forKey:"), value, key)
}

// Returns the entity description that is associated with this subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640588-entity?language=objc
func (mc _ManagedObjectClass) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](mc, objc.Sel("entity"))
	return rv
}

// Returns the entity description that is associated with this subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1640588-entity?language=objc
func ManagedObject_Entity() EntityDescription {
	return ManagedObjectClass.Entity()
}

// Provides an opportunity to add code into the life cycle of the managed object after the managed object’s context completes a save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506585-didsave?language=objc
func (m_ ManagedObject) DidSave() {
	objc.Call[objc.Void](m_, objc.Sel("didSave"))
}

// Determines whether the managed object's current state is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506998-validateforupdate?language=objc
func (m_ ManagedObject) ValidateForUpdate(error unsafe.Pointer) bool {
	rv := objc.Call[bool](m_, objc.Sel("validateForUpdate:"), error)
	return rv
}

// A Boolean value that indicates whether the managed object has persistent changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506240-haspersistentchangedvalues?language=objc
func (m_ ManagedObject) HasPersistentChangedValues() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasPersistentChangedValues"))
	return rv
}

// A Boolean value that indicates whether to mark instances of the class as having changes when an unmodeled property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506727-contextshouldignoreunmodeledprop?language=objc
func (mc _ManagedObjectClass) ContextShouldIgnoreUnmodeledPropertyChanges() bool {
	rv := objc.Call[bool](mc, objc.Sel("contextShouldIgnoreUnmodeledPropertyChanges"))
	return rv
}

// A Boolean value that indicates whether to mark instances of the class as having changes when an unmodeled property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506727-contextshouldignoreunmodeledprop?language=objc
func ManagedObject_ContextShouldIgnoreUnmodeledPropertyChanges() bool {
	return ManagedObjectClass.ContextShouldIgnoreUnmodeledPropertyChanges()
}

// The faulting state of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506720-faultingstate?language=objc
func (m_ ManagedObject) FaultingState() uint {
	rv := objc.Call[uint](m_, objc.Sel("faultingState"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted, has been deleted, or has unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506654-haschanges?language=objc
func (m_ ManagedObject) HasChanges() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasChanges"))
	return rv
}

// The managed object context with which the managed object is registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506677-managedobjectcontext?language=objc
func (m_ ManagedObject) ManagedObjectContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("managedObjectContext"))
	return rv
}

// A Boolean value that indicates whether the managed object has been inserted in a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506281-inserted?language=objc
func (m_ ManagedObject) IsInserted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isInserted"))
	return rv
}

// A Boolean value that indicates whether the managed object will be deleted during the next save. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506681-deleted?language=objc
func (m_ ManagedObject) IsDeleted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isDeleted"))
	return rv
}

// The object ID of the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506848-objectid?language=objc
func (m_ ManagedObject) ObjectID() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](m_, objc.Sel("objectID"))
	return rv
}

// A Boolean value that indicates whether the managed object is a fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506837-fault?language=objc
func (m_ ManagedObject) IsFault() bool {
	rv := objc.Call[bool](m_, objc.Sel("isFault"))
	return rv
}

// A Boolean value that indicates whether the managed object has unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobject/1506867-updated?language=objc
func (m_ ManagedObject) IsUpdated() bool {
	rv := objc.Call[bool](m_, objc.Sel("isUpdated"))
	return rv
}
