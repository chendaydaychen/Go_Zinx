package znet

import (
	"Zinx/zinx/ziface"
	"fmt"
	"sync"
)

/*
	链接管理模块
*/

type ConnManager struct {
	Connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Connections: make(map[uint32]ziface.IConnection),
	}
}
func (cm *ConnManager) Add(conn ziface.IConnection) {
	// 保护共享资源map,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 将conn加入到Connections中
	cm.Connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", len(cm.Connections))
}
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	// 保护共享资源map,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 删除连接
	delete(cm.Connections, conn.GetConnID())
	fmt.Println("connection Remove from ConnManager successfully: conn num = ", len(cm.Connections))
}
func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	// 保护共享资源map,加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	// 从Connections中获取对应ID的Conn
	if conn, ok := cm.Connections[connID]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("connection not found")
	}
}
func (cm *ConnManager) Len() int {
	return len(cm.Connections)
}
func (cm *ConnManager) ClearAllConn() {
	// 保护共享资源map,加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 停止并删除conn conn.Stop()
	for connID, conn := range cm.Connections {
		//停止
		conn.Stop()
		//删除
		delete(cm.Connections, connID)
	}
	fmt.Println("Clear All Connections successfully: conn num = ", len(cm.Connections))
}
