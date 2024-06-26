// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableVideoCompositionInstruction] class.
var MutableVideoCompositionInstructionClass = _MutableVideoCompositionInstructionClass{objc.GetClass("AVMutableVideoCompositionInstruction")}

type _MutableVideoCompositionInstructionClass struct {
	objc.Class
}

// An interface definition for the [MutableVideoCompositionInstruction] class.
type IMutableVideoCompositionInstruction interface {
	IVideoCompositionInstruction
	SetEnablePostProcessing(value bool)
	SetLayerInstructions(value []IVideoCompositionLayerInstruction)
	SetBackgroundColor(value coregraphics.ColorRef)
	SetRequiredSourceSampleDataTrackIDs(value []foundation.INumber)
	SetTimeRange(value coremedia.TimeRange)
}

// A mutable video composition instruction subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction?language=objc
type MutableVideoCompositionInstruction struct {
	VideoCompositionInstruction
}

func MutableVideoCompositionInstructionFrom(ptr unsafe.Pointer) MutableVideoCompositionInstruction {
	return MutableVideoCompositionInstruction{
		VideoCompositionInstruction: VideoCompositionInstructionFrom(ptr),
	}
}

func (mc _MutableVideoCompositionInstructionClass) VideoCompositionInstruction() MutableVideoCompositionInstruction {
	rv := objc.Call[MutableVideoCompositionInstruction](mc, objc.Sel("videoCompositionInstruction"))
	return rv
}

// Returns a new mutable video composition instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/1519701-videocompositioninstruction?language=objc
func MutableVideoCompositionInstruction_VideoCompositionInstruction() MutableVideoCompositionInstruction {
	return MutableVideoCompositionInstructionClass.VideoCompositionInstruction()
}

func (mc _MutableVideoCompositionInstructionClass) Alloc() MutableVideoCompositionInstruction {
	rv := objc.Call[MutableVideoCompositionInstruction](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableVideoCompositionInstructionClass) New() MutableVideoCompositionInstruction {
	rv := objc.Call[MutableVideoCompositionInstruction](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableVideoCompositionInstruction() MutableVideoCompositionInstruction {
	return MutableVideoCompositionInstructionClass.New()
}

func (m_ MutableVideoCompositionInstruction) Init() MutableVideoCompositionInstruction {
	rv := objc.Call[MutableVideoCompositionInstruction](m_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the instruction requires post processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/1385876-enablepostprocessing?language=objc
func (m_ MutableVideoCompositionInstruction) SetEnablePostProcessing(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setEnablePostProcessing:"), value)
}

// Instructions that specify how to layer and compose video frames from source tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/1388912-layerinstructions?language=objc
func (m_ MutableVideoCompositionInstruction) SetLayerInstructions(value []IVideoCompositionLayerInstruction) {
	objc.Call[objc.Void](m_, objc.Sel("setLayerInstructions:"), value)
}

// The background color of the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/1390236-backgroundcolor?language=objc
func (m_ MutableVideoCompositionInstruction) SetBackgroundColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](m_, objc.Sel("setBackgroundColor:"), value)
}

// The track identifiers of source sample data that the compositor requires to compose frames for the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/3750317-requiredsourcesampledatatrackids?language=objc
func (m_ MutableVideoCompositionInstruction) SetRequiredSourceSampleDataTrackIDs(value []foundation.INumber) {
	objc.Call[objc.Void](m_, objc.Sel("setRequiredSourceSampleDataTrackIDs:"), value)
}

// The time range to which the instruction applies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositioninstruction/1390418-timerange?language=objc
func (m_ MutableVideoCompositionInstruction) SetTimeRange(value coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setTimeRange:"), value)
}
