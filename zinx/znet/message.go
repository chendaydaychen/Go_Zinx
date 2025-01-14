package znet

type Message struct {
	Id      uint32 // 消息的id
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// 创建一个消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 获取方法
func (m *Message) GetMsgId() uint32 {
	return m.Id
}
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}
func (m *Message) GetData() []byte {
	return m.Data
}

// 设置方法
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
