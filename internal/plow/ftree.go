package plow

import (
	"fmt"

	"github.com/dhconnelly/rtreego"
	"traffic.go/util"
)

func DeterminePlowPos() {
	plows := ParsePlows()
	tree := MakeTree()

	for _, v := range *plows {
		if v.Active {
			v.ClosestMile = FindCloseMarkerSingle(v, *tree)

			fmt.Println(v.ID, v.ClosestMile.RoadType, v.ClosestMile.Marker)

		}
	}

}

func MakeTree() *rtreego.Rtree {

	markers := util.ReadMileMarker()
	tree := rtreego.NewTree(2, 25, 50)

	for i := range markers {
		tree.Insert(&markers[i])
	}
	return tree
}

func FindCloseMarker(plows []UsePlow, tree rtreego.Rtree) {

	for _, plow := range plows {
		snowRect := rtreego.Point{plow.Position.Latitude, plow.Position.Longitude}

		results := tree.NearestNeighbors(1, snowRect)
		close := results[0].(*util.MileMarker)

		fmt.Println(close)

	}

}

func FindCloseMarkerSingle(plow UsePlow, tree rtreego.Rtree) *util.MileMarker {

	snowRect := rtreego.Point{plow.Position.Latitude, plow.Position.Longitude}

	results := tree.NearestNeighbors(1, snowRect)
	close := results[0].(*util.MileMarker)
	return close
	// fmt.Println(close)

}
