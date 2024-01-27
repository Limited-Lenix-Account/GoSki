package plow

import (
	"fmt"

	"github.com/dhconnelly/rtreego"
	"traffic.go/util"
)

func DeterminePlowPos() *[]UsePlow {

	plows := ParsePlows()
	tree := MakeTree()
	for i := range *plows {
		if (*plows)[i].Active {
			(*plows)[i].ClosestMile = FindCloseMarkerSingle((*plows)[i], *tree)
		}

	}
	for _, v := range *plows {
		fmt.Println(v.ID, v.ClosestMile)
	}
	return plows

}

func MakeTree() *rtreego.Rtree {

	markers := util.ReadMileMarker()
	tree := rtreego.NewTree(2, 25, 50)

	for i := range markers {
		tree.Insert(&markers[i])
	}
	return tree
}

func FindCloseMarkerSingle(plow UsePlow, tree rtreego.Rtree) *util.MileMarker {

	snowRect := rtreego.Point{plow.Position.Latitude, plow.Position.Longitude}

	results := tree.NearestNeighbors(1, snowRect)
	close := results[0].(*util.MileMarker)
	return close

}
