package core

import (
	"fmt"
	"sync"
)

/*
	一个AOI地图的格子模块
*/

type Grid struct {
	// 格子ID
	Id int32
	// 格子的左边界坐标
	MinX int32
	// 格子的右边界坐标
	MaxX int32
	// 格子的下边界坐标
	MinY int32
	// 格子的上边界坐标
	MaxY int32
	// 格子内的玩家或物品的集合
	PlayerIDs map[int32]bool
	// 保护当前集合的锁
	pIDLock sync.RWMutex
}

// 初始化当前格子方法
func NewGrid(id int32, minX, maxX, minY, maxY int32) *Grid {
	return &Grid{
		Id:        id,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIDs: make(map[int32]bool),
	}
}

// 给格子添加一个玩家
func (g *Grid) AddPlayer(playerID int32) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.PlayerIDs[playerID] = true
}

// 从格子删除一个玩家
func (g *Grid) RemovePlayer(playerID int32) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.PlayerIDs, playerID)
}

// 得到当前格子的所有玩家
func (g *Grid) GetAllPlayers() []int32 {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	players := make([]int32, 0)
	for playerID := range g.PlayerIDs {
		players = append(players, playerID)
	}
	return players
}

// 调试打印格子的基本信息
func (g *Grid) Print() {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	fmt.Printf("Grid ID: %d, MinX: %d, MaxX: %d, MinY: %d, MaxY: %d, PlayerIDs: %v\n",
		g.Id, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIDs)
}
