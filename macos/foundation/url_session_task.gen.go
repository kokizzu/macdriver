// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionTask] class.
var URLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}

type _URLSessionTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objc.IObject
	Suspend()
	Resume()
	Cancel()
	CountOfBytesExpectedToSend() int64
	Priority() float64
	SetPriority(value float64)
	Error() Error
	State() URLSessionTaskState
	TaskIdentifier() uint
	CountOfBytesClientExpectsToReceive() int64
	SetCountOfBytesClientExpectsToReceive(value int64)
	CountOfBytesSent() int64
	TaskDescription() string
	SetTaskDescription(value string)
	PrefersIncrementalDelivery() bool
	SetPrefersIncrementalDelivery(value bool)
	EarliestBeginDate() Date
	SetEarliestBeginDate(value IDate)
	CountOfBytesExpectedToReceive() int64
	Delegate() URLSessionTaskDelegateWrapper
	SetDelegate(value PURLSessionTaskDelegate)
	SetDelegateObject(valueObject objc.IObject)
	CountOfBytesReceived() int64
	Progress() Progress
	Response() URLResponse
	CountOfBytesClientExpectsToSend() int64
	SetCountOfBytesClientExpectsToSend(value int64)
	CurrentRequest() URLRequest
	OriginalRequest() URLRequest
}

// A task, like downloading a specific resource, performed in a URL session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask?language=objc
type URLSessionTask struct {
	objc.Object
}

func URLSessionTaskFrom(ptr unsafe.Pointer) URLSessionTask {
	return URLSessionTask{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Call[URLSessionTask](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionTask_Alloc() URLSessionTask {
	return URLSessionTaskClass.Alloc()
}

func (uc _URLSessionTaskClass) New() URLSessionTask {
	rv := objc.Call[URLSessionTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionTask() URLSessionTask {
	return URLSessionTaskClass.New()
}

func (u_ URLSessionTask) Init() URLSessionTask {
	rv := objc.Call[URLSessionTask](u_, objc.Sel("init"))
	return rv
}

// Temporarily suspends a task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411565-suspend?language=objc
func (u_ URLSessionTask) Suspend() {
	objc.Call[objc.Void](u_, objc.Sel("suspend"))
}

// Resumes the task, if it is suspended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411121-resume?language=objc
func (u_ URLSessionTask) Resume() {
	objc.Call[objc.Void](u_, objc.Sel("resume"))
}

// Cancels the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411591-cancel?language=objc
func (u_ URLSessionTask) Cancel() {
	objc.Call[objc.Void](u_, objc.Sel("cancel"))
}

// The number of bytes that the task expects to send in the request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411534-countofbytesexpectedtosend?language=objc
func (u_ URLSessionTask) CountOfBytesExpectedToSend() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesExpectedToSend"))
	return rv
}

// The relative priority at which you’d like a host to handle the task, specified as a floating point value between 0.0 (lowest priority) and 1.0 (highest priority). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1410569-priority?language=objc
func (u_ URLSessionTask) Priority() float64 {
	rv := objc.Call[float64](u_, objc.Sel("priority"))
	return rv
}

// The relative priority at which you’d like a host to handle the task, specified as a floating point value between 0.0 (lowest priority) and 1.0 (highest priority). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1410569-priority?language=objc
func (u_ URLSessionTask) SetPriority(value float64) {
	objc.Call[objc.Void](u_, objc.Sel("setPriority:"), value)
}

// An error object that indicates why the task failed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1408145-error?language=objc
func (u_ URLSessionTask) Error() Error {
	rv := objc.Call[Error](u_, objc.Sel("error"))
	return rv
}

// The current state of the task—active, suspended, in the process of being canceled, or completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1409888-state?language=objc
func (u_ URLSessionTask) State() URLSessionTaskState {
	rv := objc.Call[URLSessionTaskState](u_, objc.Sel("state"))
	return rv
}

// An identifier uniquely identifying the task within a given session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411231-taskidentifier?language=objc
func (u_ URLSessionTask) TaskIdentifier() uint {
	rv := objc.Call[uint](u_, objc.Sel("taskIdentifier"))
	return rv
}

// A best-guess upper bound on the number of bytes the client expects to receive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873414-countofbytesclientexpectstorecei?language=objc
func (u_ URLSessionTask) CountOfBytesClientExpectsToReceive() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesClientExpectsToReceive"))
	return rv
}

// A best-guess upper bound on the number of bytes the client expects to receive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873414-countofbytesclientexpectstorecei?language=objc
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToReceive(value int64) {
	objc.Call[objc.Void](u_, objc.Sel("setCountOfBytesClientExpectsToReceive:"), value)
}

// The number of bytes that the task has sent to the server in the request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1410444-countofbytessent?language=objc
func (u_ URLSessionTask) CountOfBytesSent() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesSent"))
	return rv
}

// An app-provided string value for the current task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1409798-taskdescription?language=objc
func (u_ URLSessionTask) TaskDescription() string {
	rv := objc.Call[string](u_, objc.Sel("taskDescription"))
	return rv
}

// An app-provided string value for the current task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1409798-taskdescription?language=objc
func (u_ URLSessionTask) SetTaskDescription(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setTaskDescription:"), value)
}

// A Boolean value that determines whether to deliver a partial response body in increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/3735881-prefersincrementaldelivery?language=objc
func (u_ URLSessionTask) PrefersIncrementalDelivery() bool {
	rv := objc.Call[bool](u_, objc.Sel("prefersIncrementalDelivery"))
	return rv
}

// A Boolean value that determines whether to deliver a partial response body in increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/3735881-prefersincrementaldelivery?language=objc
func (u_ URLSessionTask) SetPrefersIncrementalDelivery(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setPrefersIncrementalDelivery:"), value)
}

// The earliest date at which the network load should begin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873413-earliestbegindate?language=objc
func (u_ URLSessionTask) EarliestBeginDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("earliestBeginDate"))
	return rv
}

// The earliest date at which the network load should begin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873413-earliestbegindate?language=objc
func (u_ URLSessionTask) SetEarliestBeginDate(value IDate) {
	objc.Call[objc.Void](u_, objc.Sel("setEarliestBeginDate:"), objc.Ptr(value))
}

// The number of bytes that the task expects to receive in the response body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1410663-countofbytesexpectedtoreceive?language=objc
func (u_ URLSessionTask) CountOfBytesExpectedToReceive() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesExpectedToReceive"))
	return rv
}

// A delegate specific to the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/3746977-delegate?language=objc
func (u_ URLSessionTask) Delegate() URLSessionTaskDelegateWrapper {
	rv := objc.Call[URLSessionTaskDelegateWrapper](u_, objc.Sel("delegate"))
	return rv
}

// A delegate specific to the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/3746977-delegate?language=objc
func (u_ URLSessionTask) SetDelegate(value PURLSessionTaskDelegate) {
	po0 := objc.WrapAsProtocol("NSURLSessionTaskDelegate", value)
	objc.Call[objc.Void](u_, objc.Sel("setDelegate:"), po0)
}

// A delegate specific to the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/3746977-delegate?language=objc
func (u_ URLSessionTask) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The number of bytes that the task has received from the server in the response body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411581-countofbytesreceived?language=objc
func (u_ URLSessionTask) CountOfBytesReceived() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesReceived"))
	return rv
}

// A representation of the overall task progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2908821-progress?language=objc
func (u_ URLSessionTask) Progress() Progress {
	rv := objc.Call[Progress](u_, objc.Sel("progress"))
	return rv
}

// The server’s response to the currently active request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1410586-response?language=objc
func (u_ URLSessionTask) Response() URLResponse {
	rv := objc.Call[URLResponse](u_, objc.Sel("response"))
	return rv
}

// A best-guess upper bound on the number of bytes the client expects to send. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873401-countofbytesclientexpectstosend?language=objc
func (u_ URLSessionTask) CountOfBytesClientExpectsToSend() int64 {
	rv := objc.Call[int64](u_, objc.Sel("countOfBytesClientExpectsToSend"))
	return rv
}

// A best-guess upper bound on the number of bytes the client expects to send. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/2873401-countofbytesclientexpectstosend?language=objc
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToSend(value int64) {
	objc.Call[objc.Void](u_, objc.Sel("setCountOfBytesClientExpectsToSend:"), value)
}

// The URL request object currently being handled by the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411649-currentrequest?language=objc
func (u_ URLSessionTask) CurrentRequest() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("currentRequest"))
	return rv
}

// The original request object passed when the task was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontask/1411572-originalrequest?language=objc
func (u_ URLSessionTask) OriginalRequest() URLRequest {
	rv := objc.Call[URLRequest](u_, objc.Sel("originalRequest"))
	return rv
}