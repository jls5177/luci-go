// Code generated by protoc-gen-go.
// source: acquisition_task.proto
// DO NOT EDIT!

package ts_mon_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task_TypeId int32

const (
	Task_MESSAGE_TYPE_ID Task_TypeId = 34049749
)

var Task_TypeId_name = map[int32]string{
	34049749: "MESSAGE_TYPE_ID",
}
var Task_TypeId_value = map[string]int32{
	"MESSAGE_TYPE_ID": 34049749,
}

func (x Task_TypeId) Enum() *Task_TypeId {
	p := new(Task_TypeId)
	*p = x
	return p
}
func (x Task_TypeId) String() string {
	return proto.EnumName(Task_TypeId_name, int32(x))
}
func (x *Task_TypeId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Task_TypeId_value, data, "Task_TypeId")
	if err != nil {
		return err
	}
	*x = Task_TypeId(value)
	return nil
}
func (Task_TypeId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type Task struct {
	ServiceName      *string `protobuf:"bytes,20,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	JobName          *string `protobuf:"bytes,30,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	DataCenter       *string `protobuf:"bytes,40,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	HostName         *string `protobuf:"bytes,50,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	TaskNum          *int32  `protobuf:"varint,60,opt,name=task_num,json=taskNum" json:"task_num,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Task) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Task) GetJobName() string {
	if m != nil && m.JobName != nil {
		return *m.JobName
	}
	return ""
}

func (m *Task) GetDataCenter() string {
	if m != nil && m.DataCenter != nil {
		return *m.DataCenter
	}
	return ""
}

func (m *Task) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

func (m *Task) GetTaskNum() int32 {
	if m != nil && m.TaskNum != nil {
		return *m.TaskNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "ts_mon.proto.Task")
	proto.RegisterEnum("ts_mon.proto.Task_TypeId", Task_TypeId_name, Task_TypeId_value)
}

var fileDescriptor1 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4c, 0x2e, 0x2c,
	0xcd, 0x2c, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x29, 0x8e, 0xcf, 0xcd, 0xcf, 0x83, 0xf0, 0x94, 0xf6, 0x33, 0x72,
	0xb1, 0x84, 0x00, 0x25, 0x85, 0x14, 0xb9, 0x78, 0x8a, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xe3,
	0xf3, 0x12, 0x73, 0x53, 0x25, 0x44, 0x14, 0x18, 0x35, 0x38, 0x83, 0xb8, 0xa1, 0x62, 0x7e, 0x40,
	0x21, 0x21, 0x49, 0x2e, 0x8e, 0xac, 0xfc, 0x24, 0x88, 0xb4, 0x1c, 0x58, 0x9a, 0x1d, 0xc8, 0x07,
	0x4b, 0xc9, 0x73, 0x71, 0xa7, 0x24, 0x96, 0x24, 0xc6, 0x27, 0xa7, 0xe6, 0x95, 0xa4, 0x16, 0x49,
	0x68, 0x80, 0x65, 0xb9, 0x40, 0x42, 0xce, 0x60, 0x11, 0x21, 0x69, 0x2e, 0xce, 0x8c, 0xfc, 0xe2,
	0x12, 0x88, 0x66, 0x23, 0xb0, 0x34, 0x07, 0x48, 0x00, 0x66, 0x30, 0xc8, 0x81, 0xf1, 0x79, 0xa5,
	0xb9, 0x12, 0x36, 0x40, 0x39, 0xd6, 0x20, 0x76, 0x10, 0xdf, 0xaf, 0x34, 0x57, 0x49, 0x81, 0x8b,
	0x2d, 0xa4, 0xb2, 0x20, 0xd5, 0x33, 0x45, 0x48, 0x8c, 0x8b, 0xdf, 0xd7, 0x35, 0x38, 0xd8, 0xd1,
	0xdd, 0x35, 0x3e, 0x24, 0x32, 0xc0, 0x35, 0xde, 0xd3, 0x45, 0xe0, 0xea, 0xdc, 0x79, 0x02, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1a, 0x46, 0x6a, 0xe8, 0x00, 0x00, 0x00,
}
