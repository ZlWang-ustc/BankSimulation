// Code generated by protoc-gen-go.
// source: AnimationTrigger.proto
// DO NOT EDIT!

/*
Package banksimulation is a generated protocol buffer package.

It is generated from these files:
	AnimationTrigger.proto

It has these top-level messages:
	BankSimulatorState
	ServerAck
*/
package banksimulation

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type BankSimulatorState_Action int32

const (
	BankSimulatorState_EVENTS_LOADED                     BankSimulatorState_Action = 0
	BankSimulatorState_CUSTOMER_SENT_DIRECTLY_TO_TELLER  BankSimulatorState_Action = 1
	BankSimulatorState_CUSTOMER_SENT_TO_TELLER_LINE      BankSimulatorState_Action = 2
	BankSimulatorState_CUSTOMER_EXIT                     BankSimulatorState_Action = 3
	BankSimulatorState_CUSTOMER_SENT_TO_TELLER_FROM_LINE BankSimulatorState_Action = 4
	BankSimulatorState_SIMULATION_END                    BankSimulatorState_Action = 5
)

var BankSimulatorState_Action_name = map[int32]string{
	0: "EVENTS_LOADED",
	1: "CUSTOMER_SENT_DIRECTLY_TO_TELLER",
	2: "CUSTOMER_SENT_TO_TELLER_LINE",
	3: "CUSTOMER_EXIT",
	4: "CUSTOMER_SENT_TO_TELLER_FROM_LINE",
	5: "SIMULATION_END",
}
var BankSimulatorState_Action_value = map[string]int32{
	"EVENTS_LOADED":                     0,
	"CUSTOMER_SENT_DIRECTLY_TO_TELLER":  1,
	"CUSTOMER_SENT_TO_TELLER_LINE":      2,
	"CUSTOMER_EXIT":                     3,
	"CUSTOMER_SENT_TO_TELLER_FROM_LINE": 4,
	"SIMULATION_END":                    5,
}

func (x BankSimulatorState_Action) String() string {
	return proto.EnumName(BankSimulatorState_Action_name, int32(x))
}

type BankSimulatorState struct {
	Pid           int32                           `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
	CurrentAction BankSimulatorState_Action       `protobuf:"varint,2,opt,name=currentAction,enum=banksimulation.BankSimulatorState_Action" json:"currentAction,omitempty"`
	TellerNum     int32                           `protobuf:"varint,3,opt,name=tellerNum" json:"tellerNum,omitempty"`
	TellerCount   int32                           `protobuf:"varint,4,opt,name=tellerCount" json:"tellerCount,omitempty"`
	CurrentTime   int32                           `protobuf:"varint,5,opt,name=currentTime" json:"currentTime,omitempty"`
	ArrivalEvent  []*BankSimulatorState_BankEvent `protobuf:"bytes,6,rep,name=arrivalEvent" json:"arrivalEvent,omitempty"`
}

func (m *BankSimulatorState) Reset()         { *m = BankSimulatorState{} }
func (m *BankSimulatorState) String() string { return proto.CompactTextString(m) }
func (*BankSimulatorState) ProtoMessage()    {}

func (m *BankSimulatorState) GetArrivalEvent() []*BankSimulatorState_BankEvent {
	if m != nil {
		return m.ArrivalEvent
	}
	return nil
}

type BankSimulatorState_BankEvent struct {
	ArrivalTime int32 `protobuf:"varint,1,opt,name=arrivalTime" json:"arrivalTime,omitempty"`
	Duration    int32 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *BankSimulatorState_BankEvent) Reset()         { *m = BankSimulatorState_BankEvent{} }
func (m *BankSimulatorState_BankEvent) String() string { return proto.CompactTextString(m) }
func (*BankSimulatorState_BankEvent) ProtoMessage()    {}

// The response message
type ServerAck struct {
	Ack bool `protobuf:"varint,1,opt,name=ack" json:"ack,omitempty"`
}

func (m *ServerAck) Reset()         { *m = ServerAck{} }
func (m *ServerAck) String() string { return proto.CompactTextString(m) }
func (*ServerAck) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("banksimulation.BankSimulatorState_Action", BankSimulatorState_Action_name, BankSimulatorState_Action_value)
}

// Client API for BankSimulatorStateProxy service

type BankSimulatorStateProxyClient interface {
	SendState(ctx context.Context, in *BankSimulatorState, opts ...grpc.CallOption) (*ServerAck, error)
}

type bankSimulatorStateProxyClient struct {
	cc *grpc.ClientConn
}

func NewBankSimulatorStateProxyClient(cc *grpc.ClientConn) BankSimulatorStateProxyClient {
	return &bankSimulatorStateProxyClient{cc}
}

func (c *bankSimulatorStateProxyClient) SendState(ctx context.Context, in *BankSimulatorState, opts ...grpc.CallOption) (*ServerAck, error) {
	out := new(ServerAck)
	err := grpc.Invoke(ctx, "/banksimulation.BankSimulatorStateProxy/SendState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BankSimulatorStateProxy service

type BankSimulatorStateProxyServer interface {
	SendState(context.Context, []byte, *BankSimulatorState) (*ServerAck, error)
}

func RegisterBankSimulatorStateProxyServer(s *grpc.Server, srv BankSimulatorStateProxyServer) {
	s.RegisterService(&_BankSimulatorStateProxy_serviceDesc, srv)
}

func _BankSimulatorStateProxy_SendState_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(BankSimulatorState)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(BankSimulatorStateProxyServer).SendState(ctx, buf, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BankSimulatorStateProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "banksimulation.BankSimulatorStateProxy",
	HandlerType: (*BankSimulatorStateProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendState",
			Handler:    _BankSimulatorStateProxy_SendState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
