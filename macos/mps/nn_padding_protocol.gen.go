// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"github.com/progrium/darwinkit/objc"
)

// The protocol that provides a description of how kernels should pad images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding?language=objc
type PNNPadding interface {
	// optional
	PaddingMethod() NNPaddingMethod
	HasPaddingMethod() bool

	// optional
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []Image, sourceStates []State, kernel Kernel, inDescriptor ImageDescriptor) ImageDescriptor
	HasDestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Inverse() objc.Object
	HasInverse() bool
}

// ensure impl type implements protocol interface
var _ PNNPadding = (*NNPaddingObject)(nil)

// A concrete type for the [PNNPadding] protocol.
type NNPaddingObject struct {
	objc.Object
}

func (n_ NNPaddingObject) HasPaddingMethod() bool {
	return n_.RespondsToSelector(objc.Sel("paddingMethod"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2866950-paddingmethod?language=objc
func (n_ NNPaddingObject) PaddingMethod() NNPaddingMethod {
	rv := objc.Call[NNPaddingMethod](n_, objc.Sel("paddingMethod"))
	return rv
}

func (n_ NNPaddingObject) HasDestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor() bool {
	return n_.RespondsToSelector(objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2867167-destinationimagedescriptorforsou?language=objc
func (n_ NNPaddingObject) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []Image, sourceStates []State, kernel Kernel, inDescriptor ImageDescriptor) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](n_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), sourceImages, sourceStates, kernel, inDescriptor)
	return rv
}

func (n_ NNPaddingObject) HasLabel() bool {
	return n_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2889870-label?language=objc
func (n_ NNPaddingObject) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}

func (n_ NNPaddingObject) HasInverse() bool {
	return n_.RespondsToSelector(objc.Sel("inverse"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2942456-inverse?language=objc
func (n_ NNPaddingObject) Inverse() objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("inverse"))
	return rv
}
