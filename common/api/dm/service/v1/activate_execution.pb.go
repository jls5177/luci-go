// Code generated by protoc-gen-go.
// source: activate_execution.proto
// DO NOT EDIT!

/*
Package dm is a generated protocol buffer package.

It is generated from these files:
	activate_execution.proto
	add_deps.proto
	claim_execution.proto
	ensure_attempt.proto
	ensure_quests.proto
	finish_attempt.proto
	graph_data.proto
	graph_query.proto
	service.proto
	types.proto
	walk_graph.proto

It has these top-level messages:
	ActivateExecutionReq
	AddDepsReq
	AddDepsRsp
	ClaimExecutionRsp
	EnsureAttemptReq
	EnsureQuestsReq
	EnsureQuestsRsp
	FinishAttemptReq
	Quest
	Attempt
	Execution
	GraphData
	GraphQuery
	MultiPropertyValue
	PropertyValue
	AttemptList
	WalkGraphReq
*/
package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// ActivateExecutionReq allows a currently-running Execution to activate itself.
// Doing this allows DM to know that the Execution has started, and also enables
// the Execution to access other APIs like WalkGraph, AddDeps, and
// FinishAttempt.
//
// ActivateExecution must be called with the ExecutionID and Activation token
// that DM provided when the Execution was started with the distributor.
//
// If the Execution has not been activated, the Execution will be marked as
// 'activating' and this will return an OK code. At this point, your client
// may use the ExecutionToken with any RPCs that have an ExecutionAuth field.
//
// This RPC may return:
//   * OK - The Execution is now activated.
//   * InvalidArgmument - The request was malformed. Retrying will not help.
//   * PermissionDenied - The provided activation token was incorrect.
//     Retrying will not help.
//   * AlreadyExists - The activation token was correct, but some other entity
//     already activated this Execution. The client should cease operations.
//     Retrying will not help.
//
// All other errors should be retried with the exact same ActivateExecutionReq
// data.
type ActivateExecutionReq struct {
	// Auth is the Execution_Auth containing the Activation Token, as provided
	// to the distributor when the Execution was created.
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// ExecutionToken should be randomly generated by the machine running the
	// execution, or by the distributor service such that if two racing Executions
	// both attempt to Activate with the same ExecutionID and ActivationToken, the
	// ExecutionToken will (probably) be different for them so that only one will
	// win.
	ExecutionToken []byte `protobuf:"bytes,2,opt,name=execution_token,json=executionToken,proto3" json:"execution_token,omitempty"`
}

func (m *ActivateExecutionReq) Reset()                    { *m = ActivateExecutionReq{} }
func (m *ActivateExecutionReq) String() string            { return proto.CompactTextString(m) }
func (*ActivateExecutionReq) ProtoMessage()               {}
func (*ActivateExecutionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivateExecutionReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivateExecutionReq)(nil), "dm.ActivateExecutionReq")
}

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4c, 0x2e, 0xc9,
	0x2c, 0x4b, 0x2c, 0x49, 0x8d, 0x4f, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0xc9, 0x95, 0x12, 0x48, 0x2f, 0x4a, 0x2c, 0xc8, 0x88,
	0x4f, 0x49, 0x2c, 0x49, 0x84, 0x88, 0x2a, 0xa5, 0x73, 0x89, 0x38, 0x42, 0x75, 0xb8, 0xc2, 0x34,
	0x04, 0xa5, 0x16, 0x0a, 0xa9, 0x71, 0xb1, 0x24, 0x96, 0x96, 0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0x09, 0xe9, 0xa5, 0xe4, 0xea, 0xc1, 0xe5, 0xf5, 0x1c, 0x81, 0x32, 0x41, 0x60, 0x79,
	0x21, 0x75, 0x2e, 0x7e, 0xb8, 0x45, 0xf1, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x12, 0x4c, 0x40, 0x2d,
	0x3c, 0x41, 0x7c, 0x70, 0xe1, 0x10, 0x90, 0x68, 0x12, 0x1b, 0xd8, 0x3e, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xfe, 0x42, 0x6f, 0xa1, 0x00, 0x00, 0x00,
}
