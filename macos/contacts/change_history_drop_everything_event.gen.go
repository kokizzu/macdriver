// Code generated by DarwinKit. DO NOT EDIT.

package contacts

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ChangeHistoryDropEverythingEvent] class.
var ChangeHistoryDropEverythingEventClass = _ChangeHistoryDropEverythingEventClass{objc.GetClass("CNChangeHistoryDropEverythingEvent")}

type _ChangeHistoryDropEverythingEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryDropEverythingEvent] class.
type IChangeHistoryDropEverythingEvent interface {
	IChangeHistoryEvent
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistorydropeverythingevent?language=objc
type ChangeHistoryDropEverythingEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryDropEverythingEventFrom(ptr unsafe.Pointer) ChangeHistoryDropEverythingEvent {
	return ChangeHistoryDropEverythingEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryDropEverythingEventClass) Alloc() ChangeHistoryDropEverythingEvent {
	rv := objc.Call[ChangeHistoryDropEverythingEvent](cc, objc.Sel("alloc"))
	return rv
}

func (cc _ChangeHistoryDropEverythingEventClass) New() ChangeHistoryDropEverythingEvent {
	rv := objc.Call[ChangeHistoryDropEverythingEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryDropEverythingEvent() ChangeHistoryDropEverythingEvent {
	return ChangeHistoryDropEverythingEventClass.New()
}

func (c_ ChangeHistoryDropEverythingEvent) Init() ChangeHistoryDropEverythingEvent {
	rv := objc.Call[ChangeHistoryDropEverythingEvent](c_, objc.Sel("init"))
	return rv
}
