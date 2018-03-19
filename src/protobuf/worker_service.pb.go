// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker_service.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShutDownRequest struct {
}

func (m *ShutDownRequest) Reset()                    { *m = ShutDownRequest{} }
func (m *ShutDownRequest) String() string            { return proto.CompactTextString(m) }
func (*ShutDownRequest) ProtoMessage()               {}
func (*ShutDownRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ShutDownResponse struct {
	IterationNum int32 `protobuf:"varint,1,opt,name=iterationNum" json:"iterationNum,omitempty"`
}

func (m *ShutDownResponse) Reset()                    { *m = ShutDownResponse{} }
func (m *ShutDownResponse) String() string            { return proto.CompactTextString(m) }
func (*ShutDownResponse) ProtoMessage()               {}
func (*ShutDownResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ShutDownResponse) GetIterationNum() int32 {
	if m != nil {
		return m.IterationNum
	}
	return 0
}

type PEvalRequest struct {
}

func (m *PEvalRequest) Reset()                    { *m = PEvalRequest{} }
func (m *PEvalRequest) String() string            { return proto.CompactTextString(m) }
func (*PEvalRequest) ProtoMessage()               {}
func (*PEvalRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type WorkerCommunicationSize struct {
	WorkerID          int32 `protobuf:"varint,1,opt,name=workerID" json:"workerID,omitempty"`
	CommunicationSize int32 `protobuf:"varint,2,opt,name=communicationSize" json:"communicationSize,omitempty"`
}

func (m *WorkerCommunicationSize) Reset()                    { *m = WorkerCommunicationSize{} }
func (m *WorkerCommunicationSize) String() string            { return proto.CompactTextString(m) }
func (*WorkerCommunicationSize) ProtoMessage()               {}
func (*WorkerCommunicationSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *WorkerCommunicationSize) GetWorkerID() int32 {
	if m != nil {
		return m.WorkerID
	}
	return 0
}

func (m *WorkerCommunicationSize) GetCommunicationSize() int32 {
	if m != nil {
		return m.CommunicationSize
	}
	return 0
}

type PEvalResponseBody struct {
	IterationNum int64 `protobuf:"varint,1,opt,name=iterationNum" json:"iterationNum,omitempty"`
	// duration time of partial SSSP loop
	IterationSeconds float64 `protobuf:"fixed64,2,opt,name=iterationSeconds" json:"iterationSeconds,omitempty"`
	// duration time of combine message
	CombineSeconds float64 `protobuf:"fixed64,3,opt,name=combineSeconds" json:"combineSeconds,omitempty"`
	// number of updated boarders node pair
	UpdatePairNum int32 `protobuf:"varint,4,opt,name=updatePairNum" json:"updatePairNum,omitempty"`
	// number of destinations which message send to
	DstPartitionNum int32 `protobuf:"varint,5,opt,name=dstPartitionNum" json:"dstPartitionNum,omitempty"`
	// duration of a worker send to message to all other workers
	AllPeerSend float64 `protobuf:"fixed64,6,opt,name=allPeerSend" json:"allPeerSend,omitempty"`
	// size of worker to worker communication pairs
	PairNum []*WorkerCommunicationSize `protobuf:"bytes,7,rep,name=pairNum" json:"pairNum,omitempty"`
}

func (m *PEvalResponseBody) Reset()                    { *m = PEvalResponseBody{} }
func (m *PEvalResponseBody) String() string            { return proto.CompactTextString(m) }
func (*PEvalResponseBody) ProtoMessage()               {}
func (*PEvalResponseBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PEvalResponseBody) GetIterationNum() int64 {
	if m != nil {
		return m.IterationNum
	}
	return 0
}

func (m *PEvalResponseBody) GetIterationSeconds() float64 {
	if m != nil {
		return m.IterationSeconds
	}
	return 0
}

func (m *PEvalResponseBody) GetCombineSeconds() float64 {
	if m != nil {
		return m.CombineSeconds
	}
	return 0
}

func (m *PEvalResponseBody) GetUpdatePairNum() int32 {
	if m != nil {
		return m.UpdatePairNum
	}
	return 0
}

func (m *PEvalResponseBody) GetDstPartitionNum() int32 {
	if m != nil {
		return m.DstPartitionNum
	}
	return 0
}

func (m *PEvalResponseBody) GetAllPeerSend() float64 {
	if m != nil {
		return m.AllPeerSend
	}
	return 0
}

func (m *PEvalResponseBody) GetPairNum() []*WorkerCommunicationSize {
	if m != nil {
		return m.PairNum
	}
	return nil
}

type PEvalResponse struct {
	Ok   bool               `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Body *PEvalResponseBody `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *PEvalResponse) Reset()                    { *m = PEvalResponse{} }
func (m *PEvalResponse) String() string            { return proto.CompactTextString(m) }
func (*PEvalResponse) ProtoMessage()               {}
func (*PEvalResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *PEvalResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *PEvalResponse) GetBody() *PEvalResponseBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type IncEvalRequest struct {
}

func (m *IncEvalRequest) Reset()                    { *m = IncEvalRequest{} }
func (m *IncEvalRequest) String() string            { return proto.CompactTextString(m) }
func (*IncEvalRequest) ProtoMessage()               {}
func (*IncEvalRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type IncEvalResponseBody struct {
	// duration time of aggregator
	AggregatorSeconds     float64 `protobuf:"fixed64,1,opt,name=aggregatorSeconds" json:"aggregatorSeconds,omitempty"`
	AggregatorOriSize     int32   `protobuf:"varint,2,opt,name=aggregatorOriSize" json:"aggregatorOriSize,omitempty"`
	AggregatorReducedSize int32   `protobuf:"varint,3,opt,name=aggregatorReducedSize" json:"aggregatorReducedSize,omitempty"`
	IterationNum          int64   `protobuf:"varint,4,opt,name=iterationNum" json:"iterationNum,omitempty"`
	// duration time of partial SSSP loop
	IterationSeconds float64 `protobuf:"fixed64,5,opt,name=iterationSeconds" json:"iterationSeconds,omitempty"`
	// duration time of combine message
	CombineSeconds float64 `protobuf:"fixed64,6,opt,name=combineSeconds" json:"combineSeconds,omitempty"`
	// number of updated boarders node pair
	UpdatePairNum int32 `protobuf:"varint,7,opt,name=updatePairNum" json:"updatePairNum,omitempty"`
	// number of destinations which message send to
	DstPartitionNum int32 `protobuf:"varint,8,opt,name=dstPartitionNum" json:"dstPartitionNum,omitempty"`
	// duration of a worker send to message to all other workers
	AllPeerSend float64 `protobuf:"fixed64,9,opt,name=allPeerSend" json:"allPeerSend,omitempty"`
	// size of worker to worker communication pairs
	PairNum []*WorkerCommunicationSize `protobuf:"bytes,10,rep,name=pairNum" json:"pairNum,omitempty"`
}

func (m *IncEvalResponseBody) Reset()                    { *m = IncEvalResponseBody{} }
func (m *IncEvalResponseBody) String() string            { return proto.CompactTextString(m) }
func (*IncEvalResponseBody) ProtoMessage()               {}
func (*IncEvalResponseBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *IncEvalResponseBody) GetAggregatorSeconds() float64 {
	if m != nil {
		return m.AggregatorSeconds
	}
	return 0
}

func (m *IncEvalResponseBody) GetAggregatorOriSize() int32 {
	if m != nil {
		return m.AggregatorOriSize
	}
	return 0
}

func (m *IncEvalResponseBody) GetAggregatorReducedSize() int32 {
	if m != nil {
		return m.AggregatorReducedSize
	}
	return 0
}

func (m *IncEvalResponseBody) GetIterationNum() int64 {
	if m != nil {
		return m.IterationNum
	}
	return 0
}

func (m *IncEvalResponseBody) GetIterationSeconds() float64 {
	if m != nil {
		return m.IterationSeconds
	}
	return 0
}

func (m *IncEvalResponseBody) GetCombineSeconds() float64 {
	if m != nil {
		return m.CombineSeconds
	}
	return 0
}

func (m *IncEvalResponseBody) GetUpdatePairNum() int32 {
	if m != nil {
		return m.UpdatePairNum
	}
	return 0
}

func (m *IncEvalResponseBody) GetDstPartitionNum() int32 {
	if m != nil {
		return m.DstPartitionNum
	}
	return 0
}

func (m *IncEvalResponseBody) GetAllPeerSend() float64 {
	if m != nil {
		return m.AllPeerSend
	}
	return 0
}

func (m *IncEvalResponseBody) GetPairNum() []*WorkerCommunicationSize {
	if m != nil {
		return m.PairNum
	}
	return nil
}

type IncEvalResponse struct {
	Update bool                 `protobuf:"varint,1,opt,name=update" json:"update,omitempty"`
	Body   *IncEvalResponseBody `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *IncEvalResponse) Reset()                    { *m = IncEvalResponse{} }
func (m *IncEvalResponse) String() string            { return proto.CompactTextString(m) }
func (*IncEvalResponse) ProtoMessage()               {}
func (*IncEvalResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *IncEvalResponse) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

func (m *IncEvalResponse) GetBody() *IncEvalResponseBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type AssembleRequest struct {
}

func (m *AssembleRequest) Reset()                    { *m = AssembleRequest{} }
func (m *AssembleRequest) String() string            { return proto.CompactTextString(m) }
func (*AssembleRequest) ProtoMessage()               {}
func (*AssembleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type AssembleResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *AssembleResponse) Reset()                    { *m = AssembleResponse{} }
func (m *AssembleResponse) String() string            { return proto.CompactTextString(m) }
func (*AssembleResponse) ProtoMessage()               {}
func (*AssembleResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *AssembleResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type SSSPMessageRequest struct {
	Pair []*SSSPMessageStruct `protobuf:"bytes,1,rep,name=pair" json:"pair,omitempty"`
}

func (m *SSSPMessageRequest) Reset()                    { *m = SSSPMessageRequest{} }
func (m *SSSPMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SSSPMessageRequest) ProtoMessage()               {}
func (*SSSPMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *SSSPMessageRequest) GetPair() []*SSSPMessageStruct {
	if m != nil {
		return m.Pair
	}
	return nil
}

type SSSPMessageStruct struct {
	NodeID   int64 `protobuf:"varint,1,opt,name=nodeID" json:"nodeID,omitempty"`
	Distance int64 `protobuf:"varint,2,opt,name=distance" json:"distance,omitempty"`
}

func (m *SSSPMessageStruct) Reset()                    { *m = SSSPMessageStruct{} }
func (m *SSSPMessageStruct) String() string            { return proto.CompactTextString(m) }
func (*SSSPMessageStruct) ProtoMessage()               {}
func (*SSSPMessageStruct) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *SSSPMessageStruct) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *SSSPMessageStruct) GetDistance() int64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

type SSSPMessageResponse struct {
}

func (m *SSSPMessageResponse) Reset()                    { *m = SSSPMessageResponse{} }
func (m *SSSPMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*SSSPMessageResponse) ProtoMessage()               {}
func (*SSSPMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

type SimMessageRequest struct {
	Pair []*SimMessageStruct `protobuf:"bytes,1,rep,name=pair" json:"pair,omitempty"`
}

func (m *SimMessageRequest) Reset()                    { *m = SimMessageRequest{} }
func (m *SimMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SimMessageRequest) ProtoMessage()               {}
func (*SimMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *SimMessageRequest) GetPair() []*SimMessageStruct {
	if m != nil {
		return m.Pair
	}
	return nil
}

type SimMessageResponse struct {
}

func (m *SimMessageResponse) Reset()                    { *m = SimMessageResponse{} }
func (m *SimMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*SimMessageResponse) ProtoMessage()               {}
func (*SimMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

type SimMessageStruct struct {
	PatternId int64 `protobuf:"varint,1,opt,name=patternId" json:"patternId,omitempty"`
	DataId    int64 `protobuf:"varint,2,opt,name=dataId" json:"dataId,omitempty"`
}

func (m *SimMessageStruct) Reset()                    { *m = SimMessageStruct{} }
func (m *SimMessageStruct) String() string            { return proto.CompactTextString(m) }
func (*SimMessageStruct) ProtoMessage()               {}
func (*SimMessageStruct) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *SimMessageStruct) GetPatternId() int64 {
	if m != nil {
		return m.PatternId
	}
	return 0
}

func (m *SimMessageStruct) GetDataId() int64 {
	if m != nil {
		return m.DataId
	}
	return 0
}

type PRMessageRequest struct {
	Pair []*PRMessageStruct `protobuf:"bytes,1,rep,name=pair" json:"pair,omitempty"`
}

func (m *PRMessageRequest) Reset()                    { *m = PRMessageRequest{} }
func (m *PRMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*PRMessageRequest) ProtoMessage()               {}
func (*PRMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *PRMessageRequest) GetPair() []*PRMessageStruct {
	if m != nil {
		return m.Pair
	}
	return nil
}

type PRMessageResponse struct {
}

func (m *PRMessageResponse) Reset()                    { *m = PRMessageResponse{} }
func (m *PRMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*PRMessageResponse) ProtoMessage()               {}
func (*PRMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

type PRMessageStruct struct {
	NodeID int64   `protobuf:"varint,1,opt,name=nodeID" json:"nodeID,omitempty"`
	PrVal  float64 `protobuf:"fixed64,2,opt,name=prVal" json:"prVal,omitempty"`
}

func (m *PRMessageStruct) Reset()                    { *m = PRMessageStruct{} }
func (m *PRMessageStruct) String() string            { return proto.CompactTextString(m) }
func (*PRMessageStruct) ProtoMessage()               {}
func (*PRMessageStruct) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *PRMessageStruct) GetNodeID() int64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *PRMessageStruct) GetPrVal() float64 {
	if m != nil {
		return m.PrVal
	}
	return 0
}

func init() {
	proto.RegisterType((*ShutDownRequest)(nil), "protobuf.shutDownRequest")
	proto.RegisterType((*ShutDownResponse)(nil), "protobuf.shutDownResponse")
	proto.RegisterType((*PEvalRequest)(nil), "protobuf.pEvalRequest")
	proto.RegisterType((*WorkerCommunicationSize)(nil), "protobuf.workerCommunicationSize")
	proto.RegisterType((*PEvalResponseBody)(nil), "protobuf.pEvalResponseBody")
	proto.RegisterType((*PEvalResponse)(nil), "protobuf.pEvalResponse")
	proto.RegisterType((*IncEvalRequest)(nil), "protobuf.incEvalRequest")
	proto.RegisterType((*IncEvalResponseBody)(nil), "protobuf.incEvalResponseBody")
	proto.RegisterType((*IncEvalResponse)(nil), "protobuf.incEvalResponse")
	proto.RegisterType((*AssembleRequest)(nil), "protobuf.assembleRequest")
	proto.RegisterType((*AssembleResponse)(nil), "protobuf.assembleResponse")
	proto.RegisterType((*SSSPMessageRequest)(nil), "protobuf.sSSPMessageRequest")
	proto.RegisterType((*SSSPMessageStruct)(nil), "protobuf.sSSPMessageStruct")
	proto.RegisterType((*SSSPMessageResponse)(nil), "protobuf.sSSPMessageResponse")
	proto.RegisterType((*SimMessageRequest)(nil), "protobuf.simMessageRequest")
	proto.RegisterType((*SimMessageResponse)(nil), "protobuf.simMessageResponse")
	proto.RegisterType((*SimMessageStruct)(nil), "protobuf.simMessageStruct")
	proto.RegisterType((*PRMessageRequest)(nil), "protobuf.pRMessageRequest")
	proto.RegisterType((*PRMessageResponse)(nil), "protobuf.pRMessageResponse")
	proto.RegisterType((*PRMessageStruct)(nil), "protobuf.pRMessageStruct")
}

func init() { proto.RegisterFile("worker_service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x86, 0x45, 0xd1, 0xba, 0xf8, 0xd8, 0x96, 0xc4, 0xf1, 0x4d, 0x96, 0x6d, 0xc1, 0x1d, 0x14,
	0x85, 0x51, 0xd8, 0x32, 0xea, 0x16, 0x5d, 0xb4, 0x8b, 0xd6, 0xb7, 0x24, 0x5a, 0x24, 0x21, 0x48,
	0x20, 0xd9, 0x04, 0x08, 0x28, 0x72, 0x22, 0x13, 0x16, 0x39, 0x0c, 0x39, 0xb4, 0xe1, 0x3c, 0x5a,
	0x9e, 0xc0, 0xab, 0x20, 0x8f, 0x90, 0x78, 0x95, 0xc7, 0x08, 0x38, 0x1c, 0x4a, 0xbc, 0x29, 0x50,
	0x56, 0xd2, 0xb9, 0xce, 0x7f, 0x0e, 0x3f, 0x0e, 0x61, 0xe3, 0x8e, 0xfa, 0x37, 0xc4, 0x7f, 0x1b,
	0x10, 0xff, 0xd6, 0x36, 0xc9, 0xc0, 0xf3, 0x29, 0xa3, 0xa8, 0xc9, 0x7f, 0x46, 0xe1, 0xbb, 0xde,
	0xf1, 0xd8, 0x66, 0xd7, 0xe1, 0x68, 0x60, 0x52, 0xe7, 0x64, 0x4c, 0xc7, 0xf4, 0x24, 0x89, 0x70,
	0x8b, 0x1b, 0xfc, 0x5f, 0x5c, 0x88, 0x15, 0x68, 0x07, 0xd7, 0x21, 0xbb, 0xa4, 0x77, 0xae, 0x46,
	0xde, 0x87, 0x24, 0x60, 0xf8, 0x6f, 0xe8, 0xcc, 0x5c, 0x81, 0x47, 0xdd, 0x80, 0x20, 0x0c, 0xab,
	0x36, 0x23, 0xbe, 0xc1, 0x6c, 0xea, 0xbe, 0x08, 0x9d, 0xae, 0x74, 0x20, 0x1d, 0xd6, 0xb4, 0x8c,
	0x0f, 0xb7, 0x60, 0xd5, 0xbb, 0xba, 0x35, 0x26, 0x49, 0x1f, 0x13, 0xb6, 0x63, 0xad, 0x17, 0xd4,
	0x71, 0x42, 0xd7, 0x36, 0x79, 0xa6, 0x6e, 0x7f, 0x20, 0xa8, 0x07, 0xcd, 0x38, 0x34, 0xbc, 0x14,
	0xad, 0xa6, 0x36, 0x3a, 0x02, 0xc5, 0xcc, 0x17, 0x74, 0xab, 0x3c, 0xa9, 0x18, 0xc0, 0x1f, 0xab,
	0xa0, 0x88, 0x53, 0x63, 0xa9, 0xe7, 0xd4, 0xba, 0x2f, 0x95, 0x2b, 0x67, 0xe5, 0xa2, 0xdf, 0xa1,
	0x33, 0xb5, 0x75, 0x62, 0x52, 0xd7, 0x0a, 0xf8, 0x31, 0x92, 0x56, 0xf0, 0xa3, 0xdf, 0xa0, 0x65,
	0x52, 0x67, 0x64, 0xbb, 0x24, 0xc9, 0x94, 0x79, 0x66, 0xce, 0x8b, 0x7e, 0x85, 0xb5, 0xd0, 0xb3,
	0x0c, 0x46, 0x54, 0xc3, 0xf6, 0xa3, 0x83, 0x97, 0xb8, 0xee, 0xac, 0x13, 0x1d, 0x42, 0xdb, 0x0a,
	0x98, 0x6a, 0xf8, 0xcc, 0x4e, 0x04, 0xd6, 0x78, 0x5e, 0xde, 0x8d, 0x0e, 0x60, 0xc5, 0x98, 0x4c,
	0x54, 0x42, 0x7c, 0x9d, 0xb8, 0x56, 0xb7, 0xce, 0x0f, 0x4d, 0xbb, 0xd0, 0xbf, 0xd0, 0xf0, 0xc4,
	0x59, 0x8d, 0x03, 0xf9, 0x70, 0xe5, 0xf4, 0x97, 0x41, 0xf2, 0xc0, 0x07, 0x73, 0xb6, 0xaf, 0x25,
	0x15, 0x58, 0x85, 0xb5, 0xcc, 0xee, 0x50, 0x0b, 0xaa, 0xf4, 0x86, 0x6f, 0xab, 0xa9, 0x55, 0xe9,
	0x0d, 0x3a, 0x81, 0xa5, 0x11, 0xb5, 0xee, 0xf9, 0x5e, 0x56, 0x4e, 0x77, 0x67, 0xad, 0x0b, 0x2b,
	0xd7, 0x78, 0x22, 0xee, 0x40, 0xcb, 0x76, 0xcd, 0x34, 0x05, 0x9f, 0x64, 0x58, 0x9f, 0xba, 0x52,
	0x8f, 0xe8, 0x08, 0x14, 0x63, 0x3c, 0xf6, 0xc9, 0xd8, 0x60, 0xd4, 0x4f, 0xb6, 0x2a, 0xf1, 0x01,
	0x8b, 0x81, 0x6c, 0xf6, 0x4b, 0xdf, 0x4e, 0x43, 0x51, 0x08, 0xa0, 0xbf, 0x60, 0x73, 0xe6, 0xd4,
	0x88, 0x15, 0x9a, 0xc4, 0xe2, 0x15, 0x32, 0xaf, 0x28, 0x0f, 0x16, 0xa0, 0x59, 0x5a, 0x10, 0x9a,
	0xda, 0xc2, 0xd0, 0xd4, 0x17, 0x83, 0xa6, 0xb1, 0x20, 0x34, 0xcd, 0x85, 0xa0, 0x59, 0xfe, 0x21,
	0x34, 0xf0, 0xd3, 0xd0, 0xbc, 0x81, 0x76, 0xee, 0x79, 0xa2, 0x2d, 0xa8, 0xc7, 0x62, 0x05, 0x3a,
	0xc2, 0x42, 0x7f, 0x64, 0xf0, 0xd9, 0x9f, 0x1d, 0x52, 0x02, 0x84, 0x00, 0x48, 0x81, 0xb6, 0x11,
	0x04, 0xc4, 0x19, 0x4d, 0x48, 0x42, 0x10, 0x86, 0xce, 0xcc, 0x55, 0x0e, 0x2a, 0xbe, 0x02, 0x14,
	0xe8, 0xba, 0xfa, 0x9c, 0x04, 0x81, 0x31, 0x4e, 0x2a, 0x23, 0x7c, 0x23, 0xd5, 0x5d, 0x89, 0x0f,
	0x99, 0xc2, 0x37, 0x95, 0xab, 0x33, 0x3f, 0x34, 0x99, 0xc6, 0x13, 0xf1, 0x53, 0x50, 0x0a, 0xa1,
	0x68, 0x3a, 0x97, 0x5a, 0x44, 0x5c, 0x55, 0xb2, 0x26, 0xac, 0xe8, 0x12, 0xb3, 0xec, 0x80, 0x19,
	0xae, 0x19, 0xa3, 0x28, 0x6b, 0x53, 0x1b, 0x6f, 0xc2, 0x7a, 0x46, 0x4f, 0x2c, 0x1b, 0x5f, 0x80,
	0x12, 0xd8, 0x4e, 0x4e, 0xe5, 0x20, 0xa3, 0xb2, 0x97, 0x52, 0x39, 0x4d, 0xcd, 0x88, 0xdc, 0x00,
	0x94, 0x6e, 0x22, 0x5a, 0x3f, 0x83, 0x4e, 0x3e, 0x1f, 0xed, 0xc1, 0xb2, 0x67, 0x30, 0x46, 0x7c,
	0x77, 0x68, 0x09, 0xf1, 0x33, 0x47, 0x34, 0x97, 0x65, 0x30, 0x63, 0x68, 0x09, 0xf5, 0xc2, 0xc2,
	0x67, 0xd0, 0xf1, 0xb4, 0x9c, 0xc6, 0xe3, 0x8c, 0xc6, 0x9d, 0xd4, 0x45, 0xa0, 0x95, 0x49, 0x5c,
	0x07, 0x25, 0xd5, 0x42, 0x28, 0xfc, 0x0f, 0xda, 0xb9, 0xec, 0xb9, 0xab, 0xdd, 0x80, 0x9a, 0xe7,
	0xbf, 0x32, 0x26, 0xe2, 0x42, 0x8e, 0x8d, 0xd3, 0x6f, 0x32, 0xd4, 0x5f, 0x73, 0x3c, 0xd1, 0x05,
	0x34, 0x75, 0xf1, 0x8d, 0x42, 0x29, 0x35, 0xb9, 0x4f, 0x59, 0xaf, 0x57, 0x16, 0x12, 0x72, 0x2a,
	0xe8, 0x1f, 0xa8, 0xa9, 0x11, 0x86, 0x68, 0xab, 0x70, 0xb1, 0xc5, 0xe5, 0xdb, 0x73, 0x2e, 0x3c,
	0x5c, 0x41, 0xff, 0x43, 0x63, 0x18, 0x43, 0x8c, 0xba, 0x25, 0x5c, 0xc7, 0xf5, 0x3b, 0x73, 0x89,
	0xc7, 0x95, 0x68, 0x84, 0x33, 0x81, 0x75, 0x7a, 0x84, 0x1c, 0xfd, 0xe9, 0x11, 0xf2, 0x6f, 0x01,
	0xae, 0xa0, 0x21, 0x34, 0x75, 0x5d, 0x57, 0xf9, 0x5b, 0xbd, 0x57, 0xca, 0x77, 0xd2, 0x67, 0x7f,
	0x4e, 0x74, 0xda, 0xea, 0x09, 0x34, 0x02, 0xdb, 0xe1, 0x9d, 0x76, 0xcb, 0x18, 0x4c, 0x1a, 0xed,
	0x95, 0x07, 0x53, 0x73, 0xd5, 0x55, 0x8d, 0xb7, 0xe9, 0x95, 0x60, 0x92, 0x74, 0xd9, 0x2d, 0x8d,
	0x25, 0x4d, 0xce, 0x3b, 0x0f, 0x5f, 0xfb, 0x95, 0x87, 0xc7, 0xbe, 0xf4, 0xf9, 0xb1, 0x2f, 0x7d,
	0x79, 0xec, 0x4b, 0xa3, 0x3a, 0xcf, 0xff, 0xf3, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x33,
	0xd6, 0xe6, 0x00, 0x09, 0x00, 0x00,
}
