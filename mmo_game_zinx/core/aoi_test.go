package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoi := NewAOIManager(-100, 100, -100, 100, 10, 10)
	aoi.Print()

}

func TestAOIManagerSuroundGridsByGid(t *testing.T) {
	aoi := NewAOIManager(-100, 100, -100, 100, 10, 10)

	for gid := range aoi.Grids {
		grids := aoi.GetSurroundGrids(gid)
		fmt.Println("gid:", gid, "grids len = ", len(grids))
		for _, grid := range grids {
			grid.Print()
		}
	}
}
