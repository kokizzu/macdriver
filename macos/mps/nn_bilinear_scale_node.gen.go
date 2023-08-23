// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNBilinearScaleNode] class.
var NNBilinearScaleNodeClass = _NNBilinearScaleNodeClass{objc.GetClass("MPSNNBilinearScaleNode")}

type _NNBilinearScaleNodeClass struct {
	objc.Class
}

// An interface definition for the [NNBilinearScaleNode] class.
type INNBilinearScaleNode interface {
	INNScaleNode
}

// A representation of a bilinear resampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbilinearscalenode?language=objc
type NNBilinearScaleNode struct {
	NNScaleNode
}

func NNBilinearScaleNodeFrom(ptr unsafe.Pointer) NNBilinearScaleNode {
	return NNBilinearScaleNode{
		NNScaleNode: NNScaleNodeFrom(ptr),
	}
}

func (nc _NNBilinearScaleNodeClass) Alloc() NNBilinearScaleNode {
	rv := objc.Call[NNBilinearScaleNode](nc, objc.Sel("alloc"))
	return rv
}

func NNBilinearScaleNode_Alloc() NNBilinearScaleNode {
	return NNBilinearScaleNodeClass.Alloc()
}

func (nc _NNBilinearScaleNodeClass) New() NNBilinearScaleNode {
	rv := objc.Call[NNBilinearScaleNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNBilinearScaleNode() NNBilinearScaleNode {
	return NNBilinearScaleNodeClass.New()
}

func (n_ NNBilinearScaleNode) Init() NNBilinearScaleNode {
	rv := objc.Call[NNBilinearScaleNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNBilinearScaleNodeClass) NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNBilinearScaleNode {
	rv := objc.Call[NNBilinearScaleNode](nc, objc.Sel("nodeWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource?language=objc
func NNBilinearScaleNode_NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNBilinearScaleNode {
	return NNBilinearScaleNodeClass.NodeWithSourceOutputSize(sourceNode, size)
}

func (n_ NNBilinearScaleNode) InitWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNBilinearScaleNode {
	rv := objc.Call[NNBilinearScaleNode](n_, objc.Sel("initWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource?language=objc
func NewNNBilinearScaleNodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNBilinearScaleNode {
	instance := NNBilinearScaleNodeClass.Alloc().InitWithSourceOutputSize(sourceNode, size)
	instance.Autorelease()
	return instance
}