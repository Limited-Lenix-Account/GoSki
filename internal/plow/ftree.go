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

func FindCloseMarkerSingle(plow UsePlow, tree rtreego.Rtree) *util.MileMarker {

	// fmt.Printf("Searching Plow: %s\n", plow.ID)
	snowRect := rtreego.Point{plow.Position.Latitude, plow.Position.Longitude}

	results := tree.NearestNeighbor(snowRect)
	close := results.(*util.MileMarker)
	return close

}
