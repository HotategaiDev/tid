// Code generated by protoc-gen-go.
// source: tracking.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	tracking.proto

It has these top-level messages:
	SysStatus
	TrackingTimesheet
	TrackingEntry
	TrackingEntryRef
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// TrackingStatus keeps hold of information about the current "state" of tracking, i.e. whether or
// not a timer is running, and which timer we last modified.
type SysStatus struct {
	// Whether or not a (any) timer is running.
	IsRunning bool `protobuf:"varint,1,opt,name=is_running,json=isRunning" json:"is_running,omitempty"`
	// The date of the timesheet currently being tracked.
	Timesheet string `protobuf:"bytes,2,opt,name=timesheet" json:"timesheet,omitempty"`
	// The hash of the entry currently being tracked.
	Entry string `protobuf:"bytes,3,opt,name=entry" json:"entry,omitempty"`
	// The name of the workspace currently being tracked.
	Workspace string `protobuf:"bytes,4,opt,name=workspace" json:"workspace,omitempty"`
}

func (m *SysStatus) Reset()                    { *m = SysStatus{} }
func (m *SysStatus) String() string            { return proto1.CompactTextString(m) }
func (*SysStatus) ProtoMessage()               {}
func (*SysStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SysStatus) GetIsRunning() bool {
	if m != nil {
		return m.IsRunning
	}
	return false
}

func (m *SysStatus) GetTimesheet() string {
	if m != nil {
		return m.Timesheet
	}
	return ""
}

func (m *SysStatus) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func (m *SysStatus) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

// TrackingTimesheet represents a timesheet. It contains all of the entries for a time period.
type TrackingTimesheet struct {
	// The key of this timesheet.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// A timesheet consists of many entries. A timesheet can be totally empty, it's just there to
	// hold the reference to a bunch of entries under a date, i.e. for easy access.
	Entries []string `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *TrackingTimesheet) Reset()                    { *m = TrackingTimesheet{} }
func (m *TrackingTimesheet) String() string            { return proto1.CompactTextString(m) }
func (*TrackingTimesheet) ProtoMessage()               {}
func (*TrackingTimesheet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TrackingTimesheet) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TrackingTimesheet) GetEntries() []string {
	if m != nil {
		return m.Entries
	}
	return nil
}

// TrackingEntry represents an entry in the timesheet. This will have enough information to commit
// to the duration, and also enough information to identify the entry.
type TrackingEntry struct {
	// The key of this entry.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The key of the timesheet this entry belongs to.
	Timesheet string `protobuf:"bytes,2,opt,name=timesheet" json:"timesheet,omitempty"`
	// The note associated with this entry.
	Note string `protobuf:"bytes,3,opt,name=note" json:"note,omitempty"`
	// The unix timestamp of when this entry was created.
	Created uint64 `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	// The unix timestamp of when this entry was last updated.
	Updated uint64 `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	// The number of seconds this has been tracked for (once committed).
	Duration uint64 `protobuf:"varint,6,opt,name=duration" json:"duration,omitempty"`
}

func (m *TrackingEntry) Reset()                    { *m = TrackingEntry{} }
func (m *TrackingEntry) String() string            { return proto1.CompactTextString(m) }
func (*TrackingEntry) ProtoMessage()               {}
func (*TrackingEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TrackingEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TrackingEntry) GetTimesheet() string {
	if m != nil {
		return m.Timesheet
	}
	return ""
}

func (m *TrackingEntry) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *TrackingEntry) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TrackingEntry) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TrackingEntry) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

// TrackingEntryRef represents a reference from an entry's short key to it's full key.
type TrackingEntryRef struct {
	// The key of this entry reference.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The key of the entry this reference belongs to.
	Entry string `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *TrackingEntryRef) Reset()                    { *m = TrackingEntryRef{} }
func (m *TrackingEntryRef) String() string            { return proto1.CompactTextString(m) }
func (*TrackingEntryRef) ProtoMessage()               {}
func (*TrackingEntryRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TrackingEntryRef) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TrackingEntryRef) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func init() {
	proto1.RegisterType((*SysStatus)(nil), "proto.SysStatus")
	proto1.RegisterType((*TrackingTimesheet)(nil), "proto.TrackingTimesheet")
	proto1.RegisterType((*TrackingEntry)(nil), "proto.TrackingEntry")
	proto1.RegisterType((*TrackingEntryRef)(nil), "proto.TrackingEntryRef")
}

func init() { proto1.RegisterFile("tracking.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xe9, 0xbf, 0x75, 0x33, 0xa0, 0xac, 0xc1, 0x43, 0x10, 0x85, 0xd2, 0x53, 0x4f, 0x5e,
	0xbc, 0x79, 0xf1, 0xe4, 0x0b, 0x64, 0xf7, 0x2e, 0xb1, 0x1d, 0xd7, 0x50, 0x4c, 0x4a, 0x32, 0x45,
	0x0a, 0xbe, 0x8d, 0x2f, 0x2a, 0x49, 0xff, 0xa8, 0x28, 0x9e, 0x32, 0xdf, 0x7c, 0xcc, 0x7c, 0x3f,
	0x26, 0x70, 0x46, 0x4e, 0x35, 0x9d, 0x36, 0xc7, 0x9b, 0xde, 0x59, 0xb2, 0xbc, 0x88, 0x4f, 0xf5,
	0x0e, 0x6c, 0x3f, 0xfa, 0x3d, 0x29, 0x1a, 0x3c, 0xbf, 0x06, 0xd0, 0xfe, 0xd1, 0x0d, 0xc6, 0x68,
	0x73, 0x14, 0x49, 0x99, 0xd4, 0x5b, 0xc9, 0xb4, 0x97, 0x53, 0x83, 0x5f, 0x01, 0x23, 0xfd, 0x8a,
	0xfe, 0x05, 0x91, 0x44, 0x5a, 0x26, 0x35, 0x93, 0x5f, 0x0d, 0x7e, 0x01, 0x05, 0x1a, 0x72, 0xa3,
	0xc8, 0xa2, 0x33, 0x89, 0x30, 0xf3, 0x66, 0x5d, 0xe7, 0x7b, 0xd5, 0xa0, 0xc8, 0xa7, 0x99, 0xb5,
	0x51, 0xdd, 0xc3, 0xf9, 0x61, 0xc6, 0x3a, 0xac, 0x8b, 0x76, 0x90, 0x75, 0x38, 0xc6, 0x78, 0x26,
	0x43, 0xc9, 0x05, 0x9c, 0x84, 0x6d, 0x1a, 0xbd, 0x48, 0xcb, 0xac, 0x66, 0x72, 0x91, 0xd5, 0x47,
	0x02, 0xa7, 0xcb, 0x86, 0x87, 0x18, 0xf8, 0x7b, 0xfa, 0x7f, 0x6c, 0x0e, 0xb9, 0xb1, 0x84, 0x33,
	0x75, 0xac, 0x43, 0x5e, 0xe3, 0x50, 0x11, 0xb6, 0x11, 0x39, 0x97, 0x8b, 0x0c, 0xce, 0xd0, 0xb7,
	0xd1, 0x29, 0x26, 0x67, 0x96, 0xfc, 0x12, 0xb6, 0xed, 0xe0, 0x14, 0x69, 0x6b, 0xc4, 0x26, 0x5a,
	0xab, 0xae, 0xee, 0x60, 0xf7, 0x03, 0x52, 0xe2, 0xf3, 0x1f, 0x9c, 0xeb, 0x01, 0xd3, 0x6f, 0x07,
	0x7c, 0xda, 0xc4, 0x7f, 0xba, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x91, 0xc8, 0x45, 0xb8, 0xc0,
	0x01, 0x00, 0x00,
}
