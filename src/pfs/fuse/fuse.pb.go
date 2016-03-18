// Code generated by protoc-gen-go.
// source: pfs/fuse/fuse.proto
// DO NOT EDIT!

/*
Package fuse is a generated protocol buffer package.

It is generated from these files:
	pfs/fuse/fuse.proto

It has these top-level messages:
	CommitMount
	Filesystem
	Node
	Attr
	Dirent
	Root
	DirectoryAttr
	DirectoryLookup
	DirectoryReadDirAll
	DirectoryCreate
	DirectoryMkdir
	FileAttr
	FileRead
	FileOpen
	FileWrite
*/
package fuse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/pfs"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CommitMount struct {
	Commit     *pfs.Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	FromCommit *pfs.Commit `protobuf:"bytes,2,opt,name=from_commit" json:"from_commit,omitempty"`
	Alias      string      `protobuf:"bytes,3,opt,name=alias" json:"alias,omitempty"`
	Shard      *pfs.Shard  `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
}

func (m *CommitMount) Reset()                    { *m = CommitMount{} }
func (m *CommitMount) String() string            { return proto.CompactTextString(m) }
func (*CommitMount) ProtoMessage()               {}
func (*CommitMount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommitMount) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *CommitMount) GetFromCommit() *pfs.Commit {
	if m != nil {
		return m.FromCommit
	}
	return nil
}

func (m *CommitMount) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type Filesystem struct {
	Shard        *pfs.Shard     `protobuf:"bytes,1,opt,name=shard" json:"shard,omitempty"`
	CommitMounts []*CommitMount `protobuf:"bytes,2,rep,name=commit_mounts" json:"commit_mounts,omitempty"`
}

func (m *Filesystem) Reset()                    { *m = Filesystem{} }
func (m *Filesystem) String() string            { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()               {}
func (*Filesystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Filesystem) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Filesystem) GetCommitMounts() []*CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type Node struct {
	File      *pfs.File                   `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	RepoAlias string                      `protobuf:"bytes,2,opt,name=repo_alias" json:"repo_alias,omitempty"`
	Write     bool                        `protobuf:"varint,3,opt,name=write" json:"write,omitempty"`
	Shard     *pfs.Shard                  `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
	Modified  *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=modified" json:"modified,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Node) GetFile() *pfs.File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Node) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Node) GetModified() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Attr struct {
	Mode uint32 `protobuf:"varint,1,opt,name=Mode" json:"Mode,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode" json:"inode,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Dirent) Reset()                    { *m = Dirent{} }
func (m *Dirent) String() string            { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()               {}
func (*Dirent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Root struct {
	Filesystem *Filesystem `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Result     *Node       `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error      string      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Root) GetFilesystem() *Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

func (m *Root) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryAttr struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryAttr) Reset()                    { *m = DirectoryAttr{} }
func (m *DirectoryAttr) String() string            { return proto.CompactTextString(m) }
func (*DirectoryAttr) ProtoMessage()               {}
func (*DirectoryAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DirectoryAttr) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryLookup struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Result    *Node  `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Err       string `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *DirectoryLookup) Reset()                    { *m = DirectoryLookup{} }
func (m *DirectoryLookup) String() string            { return proto.CompactTextString(m) }
func (*DirectoryLookup) ProtoMessage()               {}
func (*DirectoryLookup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DirectoryLookup) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryLookup) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryReadDirAll struct {
	Directory *Node     `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    []*Dirent `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	Error     string    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryReadDirAll) Reset()                    { *m = DirectoryReadDirAll{} }
func (m *DirectoryReadDirAll) String() string            { return proto.CompactTextString(m) }
func (*DirectoryReadDirAll) ProtoMessage()               {}
func (*DirectoryReadDirAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DirectoryReadDirAll) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryReadDirAll) GetResult() []*Dirent {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryCreate struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryCreate) Reset()                    { *m = DirectoryCreate{} }
func (m *DirectoryCreate) String() string            { return proto.CompactTextString(m) }
func (*DirectoryCreate) ProtoMessage()               {}
func (*DirectoryCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DirectoryCreate) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryCreate) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryMkdir struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryMkdir) Reset()                    { *m = DirectoryMkdir{} }
func (m *DirectoryMkdir) String() string            { return proto.CompactTextString(m) }
func (*DirectoryMkdir) ProtoMessage()               {}
func (*DirectoryMkdir) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DirectoryMkdir) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryMkdir) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileAttr struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Result *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FileAttr) Reset()                    { *m = FileAttr{} }
func (m *FileAttr) String() string            { return proto.CompactTextString(m) }
func (*FileAttr) ProtoMessage()               {}
func (*FileAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FileAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileRead struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileRead) Reset()                    { *m = FileRead{} }
func (m *FileRead) String() string            { return proto.CompactTextString(m) }
func (*FileRead) ProtoMessage()               {}
func (*FileRead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FileRead) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileOpen struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileOpen) Reset()                    { *m = FileOpen{} }
func (m *FileOpen) String() string            { return proto.CompactTextString(m) }
func (*FileOpen) ProtoMessage()               {}
func (*FileOpen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *FileOpen) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileWrite struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileWrite) Reset()                    { *m = FileWrite{} }
func (m *FileWrite) String() string            { return proto.CompactTextString(m) }
func (*FileWrite) ProtoMessage()               {}
func (*FileWrite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *FileWrite) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitMount)(nil), "fuse.CommitMount")
	proto.RegisterType((*Filesystem)(nil), "fuse.Filesystem")
	proto.RegisterType((*Node)(nil), "fuse.Node")
	proto.RegisterType((*Attr)(nil), "fuse.Attr")
	proto.RegisterType((*Dirent)(nil), "fuse.Dirent")
	proto.RegisterType((*Root)(nil), "fuse.Root")
	proto.RegisterType((*DirectoryAttr)(nil), "fuse.DirectoryAttr")
	proto.RegisterType((*DirectoryLookup)(nil), "fuse.DirectoryLookup")
	proto.RegisterType((*DirectoryReadDirAll)(nil), "fuse.DirectoryReadDirAll")
	proto.RegisterType((*DirectoryCreate)(nil), "fuse.DirectoryCreate")
	proto.RegisterType((*DirectoryMkdir)(nil), "fuse.DirectoryMkdir")
	proto.RegisterType((*FileAttr)(nil), "fuse.FileAttr")
	proto.RegisterType((*FileRead)(nil), "fuse.FileRead")
	proto.RegisterType((*FileOpen)(nil), "fuse.FileOpen")
	proto.RegisterType((*FileWrite)(nil), "fuse.FileWrite")
}

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xdb, 0xb4, 0x6a, 0x6e, 0x16, 0x3e, 0x3c, 0x24, 0x42, 0x01, 0x31, 0x45, 0x20, 0xed,
	0x01, 0xa5, 0x12, 0xe3, 0x0f, 0x4c, 0x43, 0x3c, 0x51, 0x26, 0x06, 0xd2, 0x04, 0x3c, 0x54, 0xd9,
	0x72, 0x53, 0xa2, 0xc5, 0x71, 0x64, 0x3b, 0xa0, 0xfd, 0x06, 0xfe, 0x34, 0xf6, 0x75, 0x93, 0x66,
	0x1b, 0x6c, 0xd3, 0xc4, 0xc3, 0xa6, 0xda, 0xf7, 0xdc, 0x73, 0xce, 0x3d, 0x37, 0x86, 0xed, 0x3a,
	0x57, 0xf3, 0xbc, 0x51, 0x48, 0xff, 0x92, 0x5a, 0x0a, 0x2d, 0x98, 0x67, 0x7f, 0xcf, 0x42, 0x5b,
	0x32, 0x7f, 0xee, 0x72, 0xf6, 0x62, 0x25, 0xc4, 0xaa, 0xc4, 0x39, 0x9d, 0x4e, 0x9a, 0x7c, 0xae,
	0x0b, 0x8e, 0x4a, 0xa7, 0xbc, 0x76, 0x80, 0xf8, 0x27, 0x04, 0x07, 0x82, 0xf3, 0x42, 0x2f, 0x44,
	0x53, 0x69, 0xf6, 0x14, 0x26, 0xa7, 0x74, 0x8c, 0x06, 0x3b, 0x83, 0xdd, 0xe0, 0x4d, 0x90, 0x58,
	0x2e, 0x87, 0x60, 0x3b, 0x10, 0xe4, 0x52, 0xf0, 0xe5, 0x1a, 0x31, 0xbc, 0x8a, 0x08, 0x61, 0x9c,
	0x96, 0x45, 0xaa, 0xa2, 0x91, 0xa9, 0xf9, 0xec, 0x09, 0x8c, 0xd5, 0x8f, 0x54, 0x66, 0x91, 0x47,
	0x50, 0x20, 0xe8, 0x67, 0x7b, 0x13, 0x7f, 0x02, 0x78, 0x5f, 0x94, 0xa8, 0xce, 0x95, 0x46, 0xbe,
	0x01, 0x0e, 0x2e, 0x03, 0xd9, 0x2e, 0x84, 0x4e, 0x6f, 0xc9, 0xad, 0x43, 0x65, 0x64, 0x47, 0x06,
	0xf2, 0x30, 0xa1, 0xd1, 0x7b, 0xde, 0xe3, 0xdf, 0x03, 0xf0, 0x3e, 0x8a, 0x0c, 0xd9, 0x63, 0xf0,
	0x72, 0xc3, 0xbd, 0x26, 0xf3, 0x89, 0xcc, 0x8a, 0x31, 0x06, 0x20, 0xb1, 0x16, 0x4b, 0xe7, 0x71,
	0x48, 0x1e, 0x8d, 0xe5, 0x5f, 0xb2, 0xd0, 0x48, 0x96, 0xa7, 0xd7, 0x58, 0x66, 0xaf, 0x61, 0xca,
	0x45, 0x56, 0xe4, 0x05, 0x66, 0xd1, 0x98, 0xaa, 0xb3, 0xc4, 0xc5, 0x9b, 0xb4, 0xf1, 0x26, 0x5f,
	0xda, 0x78, 0xe3, 0x47, 0xe0, 0xed, 0x6b, 0x2d, 0xd9, 0x16, 0x78, 0x0b, 0x63, 0x8a, 0xcc, 0x84,
	0xf1, 0x2b, 0x98, 0xbc, 0x2b, 0x24, 0x56, 0x14, 0x55, 0x51, 0xb5, 0x05, 0xcf, 0xc2, 0xaa, 0x94,
	0xa3, 0x33, 0x15, 0x1f, 0x83, 0x77, 0x24, 0x84, 0x66, 0x2f, 0x01, 0xf2, 0x2e, 0xa5, 0xf5, 0x3c,
	0x0f, 0xdc, 0xe4, 0xbd, 0xf4, 0x66, 0x30, 0x91, 0xa8, 0x9a, 0xb2, 0x5d, 0x09, 0x38, 0x04, 0x65,
	0x61, 0x64, 0x50, 0x4a, 0x21, 0xdd, 0x46, 0xe2, 0xaf, 0x10, 0x5a, 0xfd, 0x53, 0x2d, 0xe4, 0x39,
	0xd9, 0x7b, 0x0e, 0x7e, 0xd6, 0x5e, 0x74, 0xe9, 0x6f, 0xda, 0xff, 0x41, 0x4d, 0xad, 0x97, 0xa8,
	0x57, 0x70, 0xbf, 0xa3, 0xfe, 0x20, 0xc4, 0x59, 0x53, 0xdf, 0x44, 0x7e, 0x61, 0xe6, 0x9e, 0xd4,
	0xe8, 0x0a, 0x32, 0x80, 0x91, 0x91, 0xa2, 0x9d, 0xf8, 0x71, 0x0a, 0xdb, 0x9d, 0xd0, 0x11, 0xa6,
	0x99, 0x39, 0xec, 0x97, 0xe5, 0x4d, 0x62, 0xcf, 0x7a, 0x93, 0xd8, 0x0f, 0x68, 0xcb, 0xd5, 0x36,
	0xdb, 0xe8, 0xcf, 0xf2, 0xbd, 0x37, 0xcb, 0x81, 0xc4, 0x54, 0xe3, 0x1d, 0x83, 0xfa, 0xdb, 0x0e,
	0xbe, 0xc1, 0xbd, 0x8e, 0x7c, 0x71, 0x66, 0x48, 0xff, 0x23, 0xf7, 0x21, 0x4c, 0xed, 0x87, 0x41,
	0xfb, 0x89, 0x2e, 0x3c, 0x83, 0x3b, 0x6e, 0x75, 0xcf, 0x11, 0xda, 0x9c, 0xaf, 0x21, 0xec, 0x9a,
	0x86, 0xfd, 0xa6, 0xc3, 0x1a, 0xab, 0xdb, 0x37, 0xbd, 0x05, 0xdf, 0x36, 0x1d, 0xdb, 0xc7, 0x78,
	0xeb, 0xae, 0x93, 0x09, 0x3d, 0xbd, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0x77, 0x39,
	0xc4, 0x13, 0x05, 0x00, 0x00,
}