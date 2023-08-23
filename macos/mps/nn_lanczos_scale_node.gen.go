// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNLanczosScaleNode] class.
var NNLanczosScaleNodeClass = _NNLanczosScaleNodeClass{objc.GetClass("MPSNNLanczosScaleNode")}

type _NNLanczosScaleNodeClass struct {
	objc.Class
}

// An interface definition for the [NNLanczosScaleNode] class.
type INNLanczosScaleNode interface {
	INNScaleNode
}

// A representation of a Lanczos resampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlanczosscalenode?language=objc
type NNLanczosScaleNode struct {
	NNScaleNode
}

func NNLanczosScaleNodeFrom(ptr unsafe.Pointer) NNLanczosScaleNode {
	return NNLanczosScaleNode{
		NNScaleNode: NNScaleNodeFrom(ptr),
	}
}

func (nc _NNLanczosScaleNodeClass) Alloc() NNLanczosScaleNode {
	rv := objc.Call[NNLanczosScaleNode](nc, objc.Sel("alloc"))
	return rv
}

func NNLanczosScaleNode_Alloc() NNLanczosScaleNode {
	return NNLanczosScaleNodeClass.Alloc()
}

func (nc _NNLanczosScaleNodeClass) New() NNLanczosScaleNode {
	rv := objc.Call[NNLanczosScaleNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNLanczosScaleNode() NNLanczosScaleNode {
	return NNLanczosScaleNodeClass.New()
}

func (n_ NNLanczosScaleNode) Init() NNLanczosScaleNode {
	rv := objc.Call[NNLanczosScaleNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNLanczosScaleNodeClass) NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNLanczosScaleNode {
	rv := objc.Call[NNLanczosScaleNode](nc, objc.Sel("nodeWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource?language=objc
func NNLanczosScaleNode_NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNLanczosScaleNode {
	return NNLanczosScaleNodeClass.NodeWithSourceOutputSize(sourceNode, size)
}

func (n_ NNLanczosScaleNode) InitWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNLanczosScaleNode {
	rv := objc.Call[NNLanczosScaleNode](n_, objc.Sel("initWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource?language=objc
func NewNNLanczosScaleNodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNLanczosScaleNode {
	instance := NNLanczosScaleNodeClass.Alloc().InitWithSourceOutputSize(sourceNode, size)
	instance.Autorelease()
	return instance
}