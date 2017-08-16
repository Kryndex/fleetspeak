// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/messages.proto

package fleetspeak_daemonservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak_monitoring"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A fleetspeak.Message with message type "StdOutput" and data type
// StdOutputData is sent by a daemon service to the server when the daemon
// subprocess produces output on stdout or stderr.
type StdOutputData struct {
	// The pid of the daemon process.
	Pid int64 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	// The index of this message within the set of messages returned for
	// this pid.
	MessageIndex int64  `protobuf:"varint,2,opt,name=message_index,json=messageIndex" json:"message_index,omitempty"`
	Stdout       []byte `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr       []byte `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (m *StdOutputData) Reset()                    { *m = StdOutputData{} }
func (m *StdOutputData) String() string            { return proto.CompactTextString(m) }
func (*StdOutputData) ProtoMessage()               {}
func (*StdOutputData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *StdOutputData) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *StdOutputData) GetMessageIndex() int64 {
	if m != nil {
		return m.MessageIndex
	}
	return 0
}

func (m *StdOutputData) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *StdOutputData) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func init() {
	proto.RegisterType((*StdOutputData)(nil), "fleetspeak.daemonservice.StdOutputData")
}

func init() {
	proto.RegisterFile("fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/messages.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x57, 0x7a, 0x08, 0x2d, 0x48, 0x0e, 0x12, 0x3c, 0x15, 0xbd, 0xf4, 0xd4, 0x1c,
	0x7c, 0x01, 0x0f, 0x5e, 0x3c, 0x89, 0xf5, 0x01, 0x4a, 0x4c, 0xc6, 0x12, 0x6c, 0x32, 0x61, 0x66,
	0xb2, 0xf8, 0xf8, 0xb2, 0x6b, 0x70, 0x59, 0x7b, 0x4b, 0xbe, 0xff, 0xe3, 0xff, 0x47, 0xbd, 0x7d,
	0x9e, 0x01, 0x84, 0x0b, 0xb8, 0x2f, 0xcb, 0xe4, 0xad, 0x3f, 0x47, 0xc8, 0x62, 0x83, 0x83, 0x84,
	0x99, 0x81, 0x86, 0xe8, 0xc1, 0x16, 0x42, 0x41, 0x3b, 0x9b, 0xc7, 0x65, 0x9c, 0x80, 0xd9, 0x9d,
	0x80, 0xf7, 0x93, 0xa7, 0xcd, 0x2c, 0xee, 0x17, 0xe2, 0xdd, 0xd3, 0xff, 0x31, 0x4c, 0x09, 0xf3,
	0x65, 0x7d, 0xc2, 0x1c, 0x05, 0x29, 0xe6, 0x93, 0x25, 0x60, 0xac, 0xe4, 0xe1, 0xb7, 0xfb, 0x7e,
	0x50, 0x9b, 0x77, 0x09, 0xaf, 0x55, 0x4a, 0x95, 0x67, 0x27, 0x4e, 0xdf, 0xa8, 0xbe, 0xc4, 0x60,
	0xba, 0x6d, 0xb7, 0xeb, 0x0f, 0xe3, 0x53, 0x3f, 0xa8, 0x4d, 0x3b, 0xe8, 0x18, 0x73, 0x80, 0x6f,
	0x73, 0x35, 0x65, 0xeb, 0x06, 0x5f, 0x46, 0xa6, 0x6f, 0xd5, 0x8a, 0x25, 0x60, 0x15, 0xd3, 0x6f,
	0xbb, 0xdd, 0xfa, 0xd0, 0x7e, 0x8d, 0x03, 0x91, 0xb9, 0xfe, 0xe3, 0x40, 0xf4, 0xb1, 0x9a, 0xe6,
	0x1f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0xcb, 0xe9, 0xf8, 0x2f, 0x01, 0x00, 0x00,
}