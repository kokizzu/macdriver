// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PlayerInterstitialEventMonitor] class.
var PlayerInterstitialEventMonitorClass = _PlayerInterstitialEventMonitorClass{objc.GetClass("AVPlayerInterstitialEventMonitor")}

type _PlayerInterstitialEventMonitorClass struct {
	objc.Class
}

// An interface definition for the [PlayerInterstitialEventMonitor] class.
type IPlayerInterstitialEventMonitor interface {
	objc.IObject
	InterstitialPlayer() QueuePlayer
	CurrentEvent() PlayerInterstitialEvent
	PrimaryPlayer() Player
	Events() []PlayerInterstitialEvent
}

// An object that monitors the scheduling and progress of interstitial events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor?language=objc
type PlayerInterstitialEventMonitor struct {
	objc.Object
}

func PlayerInterstitialEventMonitorFrom(ptr unsafe.Pointer) PlayerInterstitialEventMonitor {
	return PlayerInterstitialEventMonitor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PlayerInterstitialEventMonitor) InitWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventMonitor {
	rv := objc.Call[PlayerInterstitialEventMonitor](p_, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	return rv
}

// Creates an observer with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800565-initwithprimaryplayer?language=objc
func NewPlayerInterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventMonitor {
	instance := PlayerInterstitialEventMonitorClass.Alloc().InitWithPrimaryPlayer(primaryPlayer)
	instance.Autorelease()
	return instance
}

func (pc _PlayerInterstitialEventMonitorClass) InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventMonitor {
	rv := objc.Call[PlayerInterstitialEventMonitor](pc, objc.Sel("interstitialEventMonitorWithPrimaryPlayer:"), primaryPlayer)
	return rv
}

// A convenience initializer that creates an observer with a player item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800566-interstitialeventmonitorwithprim?language=objc
func PlayerInterstitialEventMonitor_InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IPlayer) PlayerInterstitialEventMonitor {
	return PlayerInterstitialEventMonitorClass.InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer)
}

func (pc _PlayerInterstitialEventMonitorClass) Alloc() PlayerInterstitialEventMonitor {
	rv := objc.Call[PlayerInterstitialEventMonitor](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PlayerInterstitialEventMonitorClass) New() PlayerInterstitialEventMonitor {
	rv := objc.Call[PlayerInterstitialEventMonitor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerInterstitialEventMonitor() PlayerInterstitialEventMonitor {
	return PlayerInterstitialEventMonitorClass.New()
}

func (p_ PlayerInterstitialEventMonitor) Init() PlayerInterstitialEventMonitor {
	rv := objc.Call[PlayerInterstitialEventMonitor](p_, objc.Sel("init"))
	return rv
}

// An object that plays interstitial content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800567-interstitialplayer?language=objc
func (p_ PlayerInterstitialEventMonitor) InterstitialPlayer() QueuePlayer {
	rv := objc.Call[QueuePlayer](p_, objc.Sel("interstitialPlayer"))
	return rv
}

// The current interstitial event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800563-currentevent?language=objc
func (p_ PlayerInterstitialEventMonitor) CurrentEvent() PlayerInterstitialEvent {
	rv := objc.Call[PlayerInterstitialEvent](p_, objc.Sel("currentEvent"))
	return rv
}

// An object that plays primary content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800568-primaryplayer?language=objc
func (p_ PlayerInterstitialEventMonitor) PrimaryPlayer() Player {
	rv := objc.Call[Player](p_, objc.Sel("primaryPlayer"))
	return rv
}

// The schedule of interstitial events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/3800564-events?language=objc
func (p_ PlayerInterstitialEventMonitor) Events() []PlayerInterstitialEvent {
	rv := objc.Call[[]PlayerInterstitialEvent](p_, objc.Sel("events"))
	return rv
}
