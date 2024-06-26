// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/macos/corefoundation"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A container that stores a sequence of GPU commands that you encode into it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer?language=objc
type PCommandBuffer interface {
	// optional
	PresentDrawable(drawable DrawableObject)
	HasPresentDrawable() bool

	// optional
	WaitUntilScheduled()
	HasWaitUntilScheduled() bool

	// optional
	Enqueue()
	HasEnqueue() bool

	// optional
	ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) ComputeCommandEncoderObject
	HasComputeCommandEncoderWithDispatchType() bool

	// optional
	ComputeCommandEncoder() ComputeCommandEncoderObject
	HasComputeCommandEncoder() bool

	// optional
	EncodeSignalEventValue(event EventObject, value uint64)
	HasEncodeSignalEventValue() bool

	// optional
	Commit()
	HasCommit() bool

	// optional
	PopDebugGroup()
	HasPopDebugGroup() bool

	// optional
	ComputeCommandEncoderWithDescriptor(computePassDescriptor ComputePassDescriptor) ComputeCommandEncoderObject
	HasComputeCommandEncoderWithDescriptor() bool

	// optional
	ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) ParallelRenderCommandEncoderObject
	HasParallelRenderCommandEncoderWithDescriptor() bool

	// optional
	EncodeWaitForEventValue(event EventObject, value uint64)
	HasEncodeWaitForEventValue() bool

	// optional
	BlitCommandEncoder() BlitCommandEncoderObject
	HasBlitCommandEncoder() bool

	// optional
	ResourceStateCommandEncoder() ResourceStateCommandEncoderObject
	HasResourceStateCommandEncoder() bool

	// optional
	RenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) RenderCommandEncoderObject
	HasRenderCommandEncoderWithDescriptor() bool

	// optional
	AddCompletedHandler(block CommandBufferHandler)
	HasAddCompletedHandler() bool

	// optional
	PresentDrawableAfterMinimumDuration(drawable DrawableObject, duration corefoundation.TimeInterval)
	HasPresentDrawableAfterMinimumDuration() bool

	// optional
	PresentDrawableAtTime(drawable DrawableObject, presentationTime corefoundation.TimeInterval)
	HasPresentDrawableAtTime() bool

	// optional
	ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor ResourceStatePassDescriptor) ResourceStateCommandEncoderObject
	HasResourceStateCommandEncoderWithDescriptor() bool

	// optional
	AccelerationStructureCommandEncoder() AccelerationStructureCommandEncoderObject
	HasAccelerationStructureCommandEncoder() bool

	// optional
	AddScheduledHandler(block CommandBufferHandler)
	HasAddScheduledHandler() bool

	// optional
	WaitUntilCompleted()
	HasWaitUntilCompleted() bool

	// optional
	PushDebugGroup(string_ string)
	HasPushDebugGroup() bool

	// optional
	BlitCommandEncoderWithDescriptor(blitPassDescriptor BlitPassDescriptor) BlitCommandEncoderObject
	HasBlitCommandEncoderWithDescriptor() bool

	// optional
	Logs() LogContainerObject
	HasLogs() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Status() CommandBufferStatus
	HasStatus() bool

	// optional
	ErrorOptions() CommandBufferErrorOption
	HasErrorOptions() bool

	// optional
	RetainedReferences() bool
	HasRetainedReferences() bool

	// optional
	GPUStartTime() corefoundation.TimeInterval
	HasGPUStartTime() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	KernelEndTime() corefoundation.TimeInterval
	HasKernelEndTime() bool

	// optional
	CommandQueue() CommandQueueObject
	HasCommandQueue() bool

	// optional
	KernelStartTime() corefoundation.TimeInterval
	HasKernelStartTime() bool

	// optional
	Error() foundation.Error
	HasError() bool

	// optional
	GPUEndTime() corefoundation.TimeInterval
	HasGPUEndTime() bool
}

// ensure impl type implements protocol interface
var _ PCommandBuffer = (*CommandBufferObject)(nil)

// A concrete type for the [PCommandBuffer] protocol.
type CommandBufferObject struct {
	objc.Object
}

func (c_ CommandBufferObject) HasPresentDrawable() bool {
	return c_.RespondsToSelector(objc.Sel("presentDrawable:"))
}

// Presents a drawable as early as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443029-presentdrawable?language=objc
func (c_ CommandBufferObject) PresentDrawable(drawable DrawableObject) {
	po0 := objc.WrapAsProtocol("MTLDrawable", drawable)
	objc.Call[objc.Void](c_, objc.Sel("presentDrawable:"), po0)
}

func (c_ CommandBufferObject) HasWaitUntilScheduled() bool {
	return c_.RespondsToSelector(objc.Sel("waitUntilScheduled"))
}

// Blocks the current thread until the command queue schedules the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443036-waituntilscheduled?language=objc
func (c_ CommandBufferObject) WaitUntilScheduled() {
	objc.Call[objc.Void](c_, objc.Sel("waitUntilScheduled"))
}

func (c_ CommandBufferObject) HasEnqueue() bool {
	return c_.RespondsToSelector(objc.Sel("enqueue"))
}

// Reserves the next available place for the command buffer in its command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443019-enqueue?language=objc
func (c_ CommandBufferObject) Enqueue() {
	objc.Call[objc.Void](c_, objc.Sel("enqueue"))
}

func (c_ CommandBufferObject) HasComputeCommandEncoderWithDispatchType() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoderWithDispatchType:"))
}

// Creates a compute command encoder with a dispatch type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966541-computecommandencoderwithdispatc?language=objc
func (c_ CommandBufferObject) ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) ComputeCommandEncoderObject {
	rv := objc.Call[ComputeCommandEncoderObject](c_, objc.Sel("computeCommandEncoderWithDispatchType:"), dispatchType)
	return rv
}

func (c_ CommandBufferObject) HasComputeCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoder"))
}

// Creates a compute command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443044-computecommandencoder?language=objc
func (c_ CommandBufferObject) ComputeCommandEncoder() ComputeCommandEncoderObject {
	rv := objc.Call[ComputeCommandEncoderObject](c_, objc.Sel("computeCommandEncoder"))
	return rv
}

func (c_ CommandBufferObject) HasEncodeSignalEventValue() bool {
	return c_.RespondsToSelector(objc.Sel("encodeSignalEvent:value:"))
}

// Encodes a command into the command buffer that pauses the GPU from running subsequent passes until the event equals or exceeds a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966542-encodesignalevent?language=objc
func (c_ CommandBufferObject) EncodeSignalEventValue(event EventObject, value uint64) {
	po0 := objc.WrapAsProtocol("MTLEvent", event)
	objc.Call[objc.Void](c_, objc.Sel("encodeSignalEvent:value:"), po0, value)
}

func (c_ CommandBufferObject) HasCommit() bool {
	return c_.RespondsToSelector(objc.Sel("commit"))
}

// Submits the command buffer to run on the GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443003-commit?language=objc
func (c_ CommandBufferObject) Commit() {
	objc.Call[objc.Void](c_, objc.Sel("commit"))
}

func (c_ CommandBufferObject) HasPopDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("popDebugGroup"))
}

// Marks the end of a debug group and, if applicable, restores the previous group from a stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2869549-popdebuggroup?language=objc
func (c_ CommandBufferObject) PopDebugGroup() {
	objc.Call[objc.Void](c_, objc.Sel("popDebugGroup"))
}

func (c_ CommandBufferObject) HasComputeCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoderWithDescriptor:"))
}

// Creates a compute command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3564432-computecommandencoderwithdescrip?language=objc
func (c_ CommandBufferObject) ComputeCommandEncoderWithDescriptor(computePassDescriptor ComputePassDescriptor) ComputeCommandEncoderObject {
	rv := objc.Call[ComputeCommandEncoderObject](c_, objc.Sel("computeCommandEncoderWithDescriptor:"), computePassDescriptor)
	return rv
}

func (c_ CommandBufferObject) HasParallelRenderCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("parallelRenderCommandEncoderWithDescriptor:"))
}

// Creates a parallel render command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443009-parallelrendercommandencoderwith?language=objc
func (c_ CommandBufferObject) ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) ParallelRenderCommandEncoderObject {
	rv := objc.Call[ParallelRenderCommandEncoderObject](c_, objc.Sel("parallelRenderCommandEncoderWithDescriptor:"), renderPassDescriptor)
	return rv
}

func (c_ CommandBufferObject) HasEncodeWaitForEventValue() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWaitForEvent:value:"))
}

// Encodes a command into the command buffer that pauses the GPU from running subsequent passes until the event equals or exceeds a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966543-encodewaitforevent?language=objc
func (c_ CommandBufferObject) EncodeWaitForEventValue(event EventObject, value uint64) {
	po0 := objc.WrapAsProtocol("MTLEvent", event)
	objc.Call[objc.Void](c_, objc.Sel("encodeWaitForEvent:value:"), po0, value)
}

func (c_ CommandBufferObject) HasBlitCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("blitCommandEncoder"))
}

// Creates a block information transfer (blit) encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443001-blitcommandencoder?language=objc
func (c_ CommandBufferObject) BlitCommandEncoder() BlitCommandEncoderObject {
	rv := objc.Call[BlitCommandEncoderObject](c_, objc.Sel("blitCommandEncoder"))
	return rv
}

func (c_ CommandBufferObject) HasResourceStateCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("resourceStateCommandEncoder"))
}

// Creates a resource state command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3043915-resourcestatecommandencoder?language=objc
func (c_ CommandBufferObject) ResourceStateCommandEncoder() ResourceStateCommandEncoderObject {
	rv := objc.Call[ResourceStateCommandEncoderObject](c_, objc.Sel("resourceStateCommandEncoder"))
	return rv
}

func (c_ CommandBufferObject) HasRenderCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("renderCommandEncoderWithDescriptor:"))
}

// Creates a render command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442999-rendercommandencoderwithdescript?language=objc
func (c_ CommandBufferObject) RenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) RenderCommandEncoderObject {
	rv := objc.Call[RenderCommandEncoderObject](c_, objc.Sel("renderCommandEncoderWithDescriptor:"), renderPassDescriptor)
	return rv
}

func (c_ CommandBufferObject) HasAddCompletedHandler() bool {
	return c_.RespondsToSelector(objc.Sel("addCompletedHandler:"))
}

// Registers a completion handler the GPU device calls immediately after the GPU finishes running the commands in the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442997-addcompletedhandler?language=objc
func (c_ CommandBufferObject) AddCompletedHandler(block CommandBufferHandler) {
	objc.Call[objc.Void](c_, objc.Sel("addCompletedHandler:"), block)
}

func (c_ CommandBufferObject) HasPresentDrawableAfterMinimumDuration() bool {
	return c_.RespondsToSelector(objc.Sel("presentDrawable:afterMinimumDuration:"))
}

// Presents a drawable after the system presents the previous drawable for an amount of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2806849-presentdrawable?language=objc
func (c_ CommandBufferObject) PresentDrawableAfterMinimumDuration(drawable DrawableObject, duration corefoundation.TimeInterval) {
	po0 := objc.WrapAsProtocol("MTLDrawable", drawable)
	objc.Call[objc.Void](c_, objc.Sel("presentDrawable:afterMinimumDuration:"), po0, duration)
}

func (c_ CommandBufferObject) HasPresentDrawableAtTime() bool {
	return c_.RespondsToSelector(objc.Sel("presentDrawable:atTime:"))
}

// Presents a drawable at a specific time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442989-presentdrawable?language=objc
func (c_ CommandBufferObject) PresentDrawableAtTime(drawable DrawableObject, presentationTime corefoundation.TimeInterval) {
	po0 := objc.WrapAsProtocol("MTLDrawable", drawable)
	objc.Call[objc.Void](c_, objc.Sel("presentDrawable:atTime:"), po0, presentationTime)
}

func (c_ CommandBufferObject) HasResourceStateCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("resourceStateCommandEncoderWithDescriptor:"))
}

// Creates a resource state command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3566536-resourcestatecommandencoderwithd?language=objc
func (c_ CommandBufferObject) ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor ResourceStatePassDescriptor) ResourceStateCommandEncoderObject {
	rv := objc.Call[ResourceStateCommandEncoderObject](c_, objc.Sel("resourceStateCommandEncoderWithDescriptor:"), resourceStatePassDescriptor)
	return rv
}

func (c_ CommandBufferObject) HasAccelerationStructureCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("accelerationStructureCommandEncoder"))
}

// Creates a ray-tracing acceleration structure command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553937-accelerationstructurecommandenco?language=objc
func (c_ CommandBufferObject) AccelerationStructureCommandEncoder() AccelerationStructureCommandEncoderObject {
	rv := objc.Call[AccelerationStructureCommandEncoderObject](c_, objc.Sel("accelerationStructureCommandEncoder"))
	return rv
}

func (c_ CommandBufferObject) HasAddScheduledHandler() bool {
	return c_.RespondsToSelector(objc.Sel("addScheduledHandler:"))
}

// Registers a completion handler the GPU device calls immediately after it schedules the command buffer to run on the GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442991-addscheduledhandler?language=objc
func (c_ CommandBufferObject) AddScheduledHandler(block CommandBufferHandler) {
	objc.Call[objc.Void](c_, objc.Sel("addScheduledHandler:"), block)
}

func (c_ CommandBufferObject) HasWaitUntilCompleted() bool {
	return c_.RespondsToSelector(objc.Sel("waitUntilCompleted"))
}

// Blocks the current thread until the GPU finishes executing the command buffer and all of its completion handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443039-waituntilcompleted?language=objc
func (c_ CommandBufferObject) WaitUntilCompleted() {
	objc.Call[objc.Void](c_, objc.Sel("waitUntilCompleted"))
}

func (c_ CommandBufferObject) HasPushDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("pushDebugGroup:"))
}

// Marks the beginning of a debug group and gives it an identifying label, which temporarily replaces the previous group, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2869550-pushdebuggroup?language=objc
func (c_ CommandBufferObject) PushDebugGroup(string_ string) {
	objc.Call[objc.Void](c_, objc.Sel("pushDebugGroup:"), string_)
}

func (c_ CommandBufferObject) HasBlitCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("blitCommandEncoderWithDescriptor:"))
}

// Creates a block information transfer (blit) encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3564431-blitcommandencoderwithdescriptor?language=objc
func (c_ CommandBufferObject) BlitCommandEncoderWithDescriptor(blitPassDescriptor BlitPassDescriptor) BlitCommandEncoderObject {
	rv := objc.Call[BlitCommandEncoderObject](c_, objc.Sel("blitCommandEncoderWithDescriptor:"), blitPassDescriptor)
	return rv
}

func (c_ CommandBufferObject) HasLogs() bool {
	return c_.RespondsToSelector(objc.Sel("logs"))
}

// The messages the command buffer records as the GPU runs its commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553939-logs?language=objc
func (c_ CommandBufferObject) Logs() LogContainerObject {
	rv := objc.Call[LogContainerObject](c_, objc.Sel("logs"))
	return rv
}

func (c_ CommandBufferObject) HasSetLabel() bool {
	return c_.RespondsToSelector(objc.Sel("setLabel:"))
}

// An optional name that can help you identify the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443031-label?language=objc
func (c_ CommandBufferObject) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

func (c_ CommandBufferObject) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// An optional name that can help you identify the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443031-label?language=objc
func (c_ CommandBufferObject) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CommandBufferObject) HasStatus() bool {
	return c_.RespondsToSelector(objc.Sel("status"))
}

// The command buffer’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443048-status?language=objc
func (c_ CommandBufferObject) Status() CommandBufferStatus {
	rv := objc.Call[CommandBufferStatus](c_, objc.Sel("status"))
	return rv
}

func (c_ CommandBufferObject) HasErrorOptions() bool {
	return c_.RespondsToSelector(objc.Sel("errorOptions"))
}

// Settings that determine which information the command buffer records about execution errors, and how it does it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553938-erroroptions?language=objc
func (c_ CommandBufferObject) ErrorOptions() CommandBufferErrorOption {
	rv := objc.Call[CommandBufferErrorOption](c_, objc.Sel("errorOptions"))
	return rv
}

func (c_ CommandBufferObject) HasRetainedReferences() bool {
	return c_.RespondsToSelector(objc.Sel("retainedReferences"))
}

// A Boolean value that indicates whether the command buffer maintains strong references to the resources it uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443011-retainedreferences?language=objc
func (c_ CommandBufferObject) RetainedReferences() bool {
	rv := objc.Call[bool](c_, objc.Sel("retainedReferences"))
	return rv
}

func (c_ CommandBufferObject) HasGPUStartTime() bool {
	return c_.RespondsToSelector(objc.Sel("GPUStartTime"))
}

// The host time, in seconds, when the GPU starts command buffer execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639924-gpustarttime?language=objc
func (c_ CommandBufferObject) GPUStartTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("GPUStartTime"))
	return rv
}

func (c_ CommandBufferObject) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The GPU device that indirectly owns the command buffer because you create it from a command queue the device also owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442995-device?language=objc
func (c_ CommandBufferObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](c_, objc.Sel("device"))
	return rv
}

func (c_ CommandBufferObject) HasKernelEndTime() bool {
	return c_.RespondsToSelector(objc.Sel("kernelEndTime"))
}

// The host time, in seconds, when the CPU finishes scheduling the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1640027-kernelendtime?language=objc
func (c_ CommandBufferObject) KernelEndTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("kernelEndTime"))
	return rv
}

func (c_ CommandBufferObject) HasCommandQueue() bool {
	return c_.RespondsToSelector(objc.Sel("commandQueue"))
}

// The command queue that creates the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443015-commandqueue?language=objc
func (c_ CommandBufferObject) CommandQueue() CommandQueueObject {
	rv := objc.Call[CommandQueueObject](c_, objc.Sel("commandQueue"))
	return rv
}

func (c_ CommandBufferObject) HasKernelStartTime() bool {
	return c_.RespondsToSelector(objc.Sel("kernelStartTime"))
}

// The host time, in seconds, when the CPU begins to schedule the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639925-kernelstarttime?language=objc
func (c_ CommandBufferObject) KernelStartTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("kernelStartTime"))
	return rv
}

func (c_ CommandBufferObject) HasError() bool {
	return c_.RespondsToSelector(objc.Sel("error"))
}

// A description of an error when the GPU encounters an issue as it runs the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443040-error?language=objc
func (c_ CommandBufferObject) Error() foundation.Error {
	rv := objc.Call[foundation.Error](c_, objc.Sel("error"))
	return rv
}

func (c_ CommandBufferObject) HasGPUEndTime() bool {
	return c_.RespondsToSelector(objc.Sel("GPUEndTime"))
}

// The host time, in seconds, when the GPU finishes execution of the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639926-gpuendtime?language=objc
func (c_ CommandBufferObject) GPUEndTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("GPUEndTime"))
	return rv
}
