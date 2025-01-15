package ziface

/*
   链接管理模块抽象层
*/

type IConnManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(uint32) (IConnection, error)
	Len() int
	ClearAllConn()
}
