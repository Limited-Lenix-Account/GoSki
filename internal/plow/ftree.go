package plow

import (
	"github.com/dhconnelly/rtreego"
	"github.com/schollz/progressbar/v3"
	"traffic.go/util"
)

func DeterminePlowPos(tree *rtreego.Rtree) *[]UsePlow {

	plows := ParsePlows()

	plowLen := len(*plows)
	QueryBar := progressbar.Default(int64(plowLen))

	for i := range *plows {
		QueryBar.Add(1)
		if (*plows)[i].Active {
			(*plows)[i].ClosestMile = FindCloseMarkerSingle((*plows)[i], *tree)
		}

	}

	return plows

}

// TODO: this function should be changed to use just a pair of coordinates rather than the Plow Object
// i'll fix this later probably but for now i'm gonna make a new function -------------
func FindCloseMarkerSingle(plow UsePlow, tree rtreego.Rtree) *util.MileMarker {
	snowRect := rtreego.Point{plow.Position.Latitude, plow.Position.Longitude}
	results := tree.NearestNeighbor(snowRect)
	close := results.(*util.MileMarker)
	return close

}

func FindClosestMileFromPoint(point rtreego.Point, tree *rtreego.Rtree) *util.MileMarker {
	results := tree.NearestNeighbor(point)
	close := results.(*util.MileMarker)
	return close
}
