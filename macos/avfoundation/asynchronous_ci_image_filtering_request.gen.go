// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coreimage"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [AsynchronousCIImageFilteringRequest] class.
var AsynchronousCIImageFilteringRequestClass = _AsynchronousCIImageFilteringRequestClass{objc.GetClass("AVAsynchronousCIImageFilteringRequest")}

type _AsynchronousCIImageFilteringRequestClass struct {
	objc.Class
}

// An interface definition for the [AsynchronousCIImageFilteringRequest] class.
type IAsynchronousCIImageFilteringRequest interface {
	objc.IObject
	FinishWithImageContext(filteredImage coreimage.IImage, context coreimage.IContext)
	FinishWithError(error foundation.IError)
	CompositionTime() coremedia.Time
	SourceImage() coreimage.Image
	RenderSize() coregraphics.Size
}

// An object that supprts using Core Image filters to process an individual video frame in a video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest?language=objc
type AsynchronousCIImageFilteringRequest struct {
	objc.Object
}

func AsynchronousCIImageFilteringRequestFrom(ptr unsafe.Pointer) AsynchronousCIImageFilteringRequest {
	return AsynchronousCIImageFilteringRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AsynchronousCIImageFilteringRequestClass) Alloc() AsynchronousCIImageFilteringRequest {
	rv := objc.Call[AsynchronousCIImageFilteringRequest](ac, objc.Sel("alloc"))
	return rv
}

func (ac _AsynchronousCIImageFilteringRequestClass) New() AsynchronousCIImageFilteringRequest {
	rv := objc.Call[AsynchronousCIImageFilteringRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsynchronousCIImageFilteringRequest() AsynchronousCIImageFilteringRequest {
	return AsynchronousCIImageFilteringRequestClass.New()
}

func (a_ AsynchronousCIImageFilteringRequest) Init() AsynchronousCIImageFilteringRequest {
	rv := objc.Call[AsynchronousCIImageFilteringRequest](a_, objc.Sel("init"))
	return rv
}

// Provides the filtered video frame image to AVFoundation for further processing or display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest/1389124-finishwithimage?language=objc
func (a_ AsynchronousCIImageFilteringRequest) FinishWithImageContext(filteredImage coreimage.IImage, context coreimage.IContext) {
	objc.Call[objc.Void](a_, objc.Sel("finishWithImage:context:"), filteredImage, context)
}

// Notifies AVFoundation that you cannot fulfill the image filtering request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest/1386608-finishwitherror?language=objc
func (a_ AsynchronousCIImageFilteringRequest) FinishWithError(error foundation.IError) {
	objc.Call[objc.Void](a_, objc.Sel("finishWithError:"), error)
}

// The time in the video composition corresponding to the frame being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest/1388240-compositiontime?language=objc
func (a_ AsynchronousCIImageFilteringRequest) CompositionTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("compositionTime"))
	return rv
}

// The current video frame image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest/1387577-sourceimage?language=objc
func (a_ AsynchronousCIImageFilteringRequest) SourceImage() coreimage.Image {
	rv := objc.Call[coreimage.Image](a_, objc.Sel("sourceImage"))
	return rv
}

// The width and height, in pixels, of the frame being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousciimagefilteringrequest/1387933-rendersize?language=objc
func (a_ AsynchronousCIImageFilteringRequest) RenderSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](a_, objc.Sel("renderSize"))
	return rv
}
