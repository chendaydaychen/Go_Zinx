package core

import "fmt"

/*
	AOI区域管理模块
*/

type AOIManager struct {
	// 区域的左边界坐标
	MinX int32
	// 区域的右边界坐标
	MaxX int32
	// 区域的下边界坐标
	MinY int32
	// 区域的上边界坐标
	MaxY int32
	// X方向的Grid数量
	GridCountX int32
	// Y方向的Grid数量
	GridCountY int32
	// 当前区域中有哪些格子map-key=格子ID，value=格子
	Grids map[int32]*Grid
}

// 初始化AOIManager
func NewAOIManager(minX, maxX, minY, maxY, cntsX, cntsY int32) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:       minX,
		MaxX:       maxX,
		MinY:       minY,
		MaxY:       maxY,
		GridCountX: cntsX,
		GridCountY: cntsY,
		Grids:      make(map[int32]*Grid),
	}

	//给AOI初始化区域中的所有格子进行编号和初始化
	for i := int32(0); i < cntsX; i++ {
		for j := int32(0); j < cntsY; j++ {
			// 计算格子的ID
			gid := i*cntsX + j

			// 初始化格子
			aoiMgr.Grids[gid] = NewGrid(gid,
				aoiMgr.MinX+i*aoiMgr.GetGridWidth(),
				aoiMgr.MinX+(i+1)*aoiMgr.GetGridWidth(),
				aoiMgr.MinY+j*aoiMgr.GetGridHeight(),
				aoiMgr.MinY+(j+1)*aoiMgr.GetGridHeight())
		}
	}
	return aoiMgr
}

// 计算得到每个格子在X轴的宽度
func (aoi *AOIManager) GetGridWidth() int32 {
	return (aoi.MaxX - aoi.MinX) / aoi.GridCountX
}

// 计算得到每个格子在Y轴的高度
func (aoi *AOIManager) GetGridHeight() int32 {
	return (aoi.MaxY - aoi.MinY) / aoi.GridCountY
}

// 打印格子信息
func (aoi *AOIManager) Print() {
	// 当前AOIManager的信息
	fmt.Printf("AOIManager: MinX: %d, MaxX: %d, MinY: %d, MaxY: %d, GridCountX: %d, GridCountY: %d\n",
		aoi.MinX, aoi.MaxX, aoi.MinY, aoi.MaxY, aoi.GridCountX, aoi.GridCountY)
	// 每个格子信息
	for _, grid := range aoi.Grids {
		grid.Print()
	}
}

// 根据格子GID获取周边九宫格格子集合
func (aoi *AOIManager) GetSurroundGrids(gid int32) (girds []*Grid) {
	// 判断gid是否在aoi中
	if _, ok := aoi.Grids[gid]; !ok {
		fmt.Println("gid is not in aoi")
		return nil
	}
	// 初始化grids集合
	girds = append(girds, aoi.Grids[gid])
	// 通过gid得到当前格子x轴的编号
	idx := gid % aoi.GridCountX
	// 判断gid左边是否有格子，右边是否有格子,加入集合
	if idx > 0 {
		girds = append(girds, aoi.Grids[gid-1])
	}
	if idx < aoi.GridCountX-1 {
		girds = append(girds, aoi.Grids[gid+1])
	}

	gidsX := make([]int32, 0, len(girds))
	for _, gird := range girds {
		gidsX = append(gidsX, gird.Id)
	}
	// 判断gid上边是否有格子，下边是否有格子
	for _, v := range gidsX {
		idy := v / aoi.GridCountY
		if idy > 0 {
			girds = append(girds, aoi.Grids[v-aoi.GridCountY])
		}
		if idy < aoi.GridCountY-1 {
			girds = append(girds, aoi.Grids[v+aoi.GridCountY])
		}
	}
	return
}

// 通过横纵坐标获取格子id
func (aoi *AOIManager) GetGidByPos(x, y float32) int32 {
	idx := (int32(x) - aoi.MinX) / aoi.GetGridWidth()
	idy := (int32(y) - aoi.MinY) / aoi.GetGridHeight()

	return idy*aoi.GridCountX + idx
}

// 通过横纵坐标得到周边九宫格内全部的PlayerIDs
func (aoi *AOIManager) GetNearbyPlayersByXY(x, y float32) (playerIDs []int32) {
	// 得到当前玩家的格子id
	gid := aoi.GetGidByPos(x, y)
	// 通过格子id得到九宫格内全部的格子
	grids := aoi.GetSurroundGrids(gid)
	// 遍历九宫格内的格子，得到格子内的全部玩家
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetAllPlayers()...)
		fmt.Println("gid:", grid.Id, "playerIDs:", grid.GetAllPlayers())
	}
	return
}

// 添加一个玩家到某个格子
func (aoi *AOIManager) AddPlayer(pID, gID int32) {
	aoi.Grids[gID].AddPlayer(pID)
}

// 移除一个格子中的玩家
func (aoi *AOIManager) RemovePlayer(pID, gID int32) {
	aoi.Grids[gID].RemovePlayer(pID)
}

// 通过GID获取全部的玩家
func (aoi *AOIManager) GetAllPlayersByGridID(gID int32) (pIDs []int32) {
	return aoi.Grids[gID].GetAllPlayers()
}

// 通过坐标添加一个玩家
func (aoi *AOIManager) AddPlayerByPos(pID int32, x, y float32) {
	gid := aoi.GetGidByPos(x, y)
	aoi.AddPlayer(pID, gid)
}

// 通过坐标移除一个玩家
func (aoi *AOIManager) RemovePlayerByPos(pID int32, x, y float32) {
	gid := aoi.GetGidByPos(x, y)
	aoi.RemovePlayer(pID, gid)
}
