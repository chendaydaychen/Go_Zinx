package znet

import (
	"Zinx/zinx/utils"
	"Zinx/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

//封包，拆包具体处理模块

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	//id uint32(4字节) +  dataLen uint32(4字节)
	return 8
}

// 封包方法，将消息头和消息体拼接成最终的二进制数据包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//将dataLen写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将dataId写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//将data写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法，将二进制数据包拆成消息头和消息体
func (dp *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压读head，得到dataLen和id
	msg := &Message{}

	//从dataBuff读取dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读取Msgid
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data received")
	}

	return msg, nil

}
