// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/host.proto

/*
Package video_host is a generated protocol buffer package.

It is generated from these files:
	proto/host.proto

It has these top-level messages:
	GetVideoInfoRequest
	GetVideoInfoResponse
	GetVideoRequest
	GetVideoResponse
*/
package video_host

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetVideoInfoRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetVideoInfoRequest) Reset()                    { *m = GetVideoInfoRequest{} }
func (m *GetVideoInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVideoInfoRequest) ProtoMessage()               {}
func (*GetVideoInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetVideoInfoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetVideoInfoResponse struct {
	Id           string                            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title        string                            `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description  string                            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	DateCreated  string                            `protobuf:"bytes,4,opt,name=date_created,json=dateCreated" json:"date_created,omitempty"`
	Views        uint64                            `protobuf:"varint,5,opt,name=views" json:"views,omitempty"`
	Likes        uint64                            `protobuf:"varint,6,opt,name=likes" json:"likes,omitempty"`
	Dislikes     uint64                            `protobuf:"varint,7,opt,name=dislikes" json:"dislikes,omitempty"`
	Resolutions  *GetVideoInfoResponse_Resolutions `protobuf:"bytes,8,opt,name=resolutions" json:"resolutions,omitempty"`
	ThumbnailUrl string                            `protobuf:"bytes,9,opt,name=thumbnail_url,json=thumbnailUrl" json:"thumbnail_url,omitempty"`
}

func (m *GetVideoInfoResponse) Reset()                    { *m = GetVideoInfoResponse{} }
func (m *GetVideoInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVideoInfoResponse) ProtoMessage()               {}
func (*GetVideoInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetVideoInfoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetVideoInfoResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetVideoInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetVideoInfoResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *GetVideoInfoResponse) GetViews() uint64 {
	if m != nil {
		return m.Views
	}
	return 0
}

func (m *GetVideoInfoResponse) GetLikes() uint64 {
	if m != nil {
		return m.Likes
	}
	return 0
}

func (m *GetVideoInfoResponse) GetDislikes() uint64 {
	if m != nil {
		return m.Dislikes
	}
	return 0
}

func (m *GetVideoInfoResponse) GetResolutions() *GetVideoInfoResponse_Resolutions {
	if m != nil {
		return m.Resolutions
	}
	return nil
}

func (m *GetVideoInfoResponse) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

type GetVideoInfoResponse_Resolutions struct {
	Q144P  bool `protobuf:"varint,1,opt,name=q144p" json:"q144p,omitempty"`
	Q240P  bool `protobuf:"varint,2,opt,name=q240p" json:"q240p,omitempty"`
	Q360P  bool `protobuf:"varint,3,opt,name=q360p" json:"q360p,omitempty"`
	Q480P  bool `protobuf:"varint,4,opt,name=q480p" json:"q480p,omitempty"`
	Q720P  bool `protobuf:"varint,5,opt,name=q720p" json:"q720p,omitempty"`
	Q1080P bool `protobuf:"varint,6,opt,name=q1080p" json:"q1080p,omitempty"`
}

func (m *GetVideoInfoResponse_Resolutions) Reset()         { *m = GetVideoInfoResponse_Resolutions{} }
func (m *GetVideoInfoResponse_Resolutions) String() string { return proto.CompactTextString(m) }
func (*GetVideoInfoResponse_Resolutions) ProtoMessage()    {}
func (*GetVideoInfoResponse_Resolutions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *GetVideoInfoResponse_Resolutions) GetQ144P() bool {
	if m != nil {
		return m.Q144P
	}
	return false
}

func (m *GetVideoInfoResponse_Resolutions) GetQ240P() bool {
	if m != nil {
		return m.Q240P
	}
	return false
}

func (m *GetVideoInfoResponse_Resolutions) GetQ360P() bool {
	if m != nil {
		return m.Q360P
	}
	return false
}

func (m *GetVideoInfoResponse_Resolutions) GetQ480P() bool {
	if m != nil {
		return m.Q480P
	}
	return false
}

func (m *GetVideoInfoResponse_Resolutions) GetQ720P() bool {
	if m != nil {
		return m.Q720P
	}
	return false
}

func (m *GetVideoInfoResponse_Resolutions) GetQ1080P() bool {
	if m != nil {
		return m.Q1080P
	}
	return false
}

type GetVideoRequest struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Resolution string `protobuf:"bytes,2,opt,name=resolution" json:"resolution,omitempty"`
}

func (m *GetVideoRequest) Reset()                    { *m = GetVideoRequest{} }
func (m *GetVideoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVideoRequest) ProtoMessage()               {}
func (*GetVideoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetVideoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetVideoRequest) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

type GetVideoResponse struct {
	PresignedUrl string `protobuf:"bytes,1,opt,name=presigned_url,json=presignedUrl" json:"presigned_url,omitempty"`
}

func (m *GetVideoResponse) Reset()                    { *m = GetVideoResponse{} }
func (m *GetVideoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVideoResponse) ProtoMessage()               {}
func (*GetVideoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetVideoResponse) GetPresignedUrl() string {
	if m != nil {
		return m.PresignedUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*GetVideoInfoRequest)(nil), "video_host.GetVideoInfoRequest")
	proto.RegisterType((*GetVideoInfoResponse)(nil), "video_host.GetVideoInfoResponse")
	proto.RegisterType((*GetVideoInfoResponse_Resolutions)(nil), "video_host.GetVideoInfoResponse.Resolutions")
	proto.RegisterType((*GetVideoRequest)(nil), "video_host.GetVideoRequest")
	proto.RegisterType((*GetVideoResponse)(nil), "video_host.GetVideoResponse")
}

func init() { proto.RegisterFile("proto/host.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x9a, 0x86, 0x74, 0x92, 0x42, 0xb5, 0x54, 0x68, 0x15, 0x10, 0x18, 0x57, 0x48,
	0x39, 0x20, 0xe3, 0xba, 0x16, 0xe5, 0x8a, 0x38, 0x14, 0x2e, 0x1c, 0x8c, 0xca, 0x35, 0x72, 0xbb,
	0x03, 0x5d, 0x61, 0xbc, 0x9b, 0xdd, 0x75, 0x79, 0x10, 0x5e, 0x81, 0x47, 0xe2, 0x81, 0xd0, 0x8e,
	0xed, 0xd8, 0xa0, 0x44, 0xbd, 0xe5, 0xff, 0x66, 0xf2, 0x7b, 0xe6, 0x1f, 0x1b, 0x8e, 0xb4, 0x51,
	0x4e, 0xbd, 0xbe, 0x51, 0xd6, 0xc5, 0xf4, 0x93, 0xc1, 0xad, 0x14, 0xa8, 0x56, 0x9e, 0x44, 0x2f,
	0xe1, 0xd1, 0x05, 0xba, 0x2f, 0x1e, 0x7c, 0xac, 0xbe, 0xaa, 0x1c, 0xd7, 0x35, 0x5a, 0xc7, 0x1e,
	0xc0, 0x48, 0x0a, 0x1e, 0x84, 0xc1, 0xf2, 0x20, 0x1f, 0x49, 0x11, 0xfd, 0xd9, 0x83, 0xe3, 0x7f,
	0xfb, 0xac, 0x56, 0x95, 0xc5, 0xff, 0x1b, 0xd9, 0x31, 0xec, 0x3b, 0xe9, 0x4a, 0xe4, 0x23, 0x42,
	0x8d, 0x60, 0x21, 0xcc, 0x04, 0xda, 0x6b, 0x23, 0xb5, 0x93, 0xaa, 0xe2, 0x7b, 0x54, 0x1b, 0x22,
	0xf6, 0x02, 0xe6, 0xa2, 0x70, 0xb8, 0xba, 0x36, 0x58, 0x38, 0x14, 0x7c, 0xdc, 0xb6, 0x14, 0x0e,
	0xdf, 0x37, 0xc8, 0x5b, 0xdf, 0x4a, 0xfc, 0x69, 0xf9, 0x7e, 0x18, 0x2c, 0xc7, 0x79, 0x23, 0x3c,
	0x2d, 0xe5, 0x77, 0xb4, 0x7c, 0xd2, 0x50, 0x12, 0x6c, 0x01, 0x53, 0x21, 0x6d, 0x53, 0xb8, 0x4f,
	0x85, 0x8d, 0x66, 0x9f, 0x60, 0x66, 0xd0, 0xaa, 0xb2, 0xf6, 0x0f, 0xb6, 0x7c, 0x1a, 0x06, 0xcb,
	0x59, 0xfa, 0x2a, 0xee, 0x43, 0x89, 0xb7, 0x6d, 0x1a, 0xe7, 0xfd, 0x7f, 0xf2, 0xa1, 0x01, 0x3b,
	0x81, 0x43, 0x77, 0x53, 0xff, 0xb8, 0xaa, 0x0a, 0x59, 0xae, 0x6a, 0x53, 0xf2, 0x03, 0x9a, 0x7d,
	0xbe, 0x81, 0x97, 0xa6, 0x5c, 0xfc, 0x0a, 0x60, 0x36, 0x70, 0xf0, 0x63, 0xaf, 0x4f, 0xb3, 0x4c,
	0x53, 0x74, 0xd3, 0xbc, 0x11, 0x44, 0xd3, 0x2c, 0xd1, 0x94, 0x9e, 0xa7, 0x5e, 0x10, 0x3d, 0x7b,
	0x93, 0x68, 0xca, 0xcd, 0x53, 0x2f, 0x88, 0x66, 0x6f, 0x13, 0x4d, 0x51, 0x79, 0xea, 0x05, 0xd1,
	0xf3, 0x34, 0xd1, 0x14, 0x92, 0xa7, 0x5e, 0xb0, 0xc7, 0x30, 0x59, 0x9f, 0x26, 0xbe, 0x79, 0x42,
	0xb8, 0x55, 0xd1, 0x3b, 0x78, 0xd8, 0xed, 0xba, 0xe3, 0xf2, 0xec, 0x19, 0x40, 0xbf, 0x6c, 0x7b,
	0xd5, 0x01, 0x89, 0xce, 0xe1, 0xa8, 0xb7, 0x68, 0x5f, 0x8a, 0x13, 0x38, 0xd4, 0x06, 0xad, 0xfc,
	0x56, 0xa1, 0xa0, 0x44, 0x1a, 0xbb, 0xf9, 0x06, 0x5e, 0x9a, 0x32, 0xfd, 0x1d, 0xc0, 0xf8, 0x83,
	0xb2, 0x8e, 0x7d, 0x86, 0xf9, 0x30, 0x70, 0xf6, 0x7c, 0xf7, 0x29, 0x68, 0xc4, 0x45, 0x78, 0xd7,
	0xad, 0xa2, 0x7b, 0xec, 0x02, 0xa6, 0x5d, 0x85, 0x3d, 0xd9, 0xd6, 0xdf, 0x99, 0x3d, 0xdd, 0x5e,
	0xec, 0x8c, 0xae, 0x26, 0xf4, 0xcd, 0x9c, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x15, 0x09,
	0xf8, 0x47, 0x03, 0x00, 0x00,
}
