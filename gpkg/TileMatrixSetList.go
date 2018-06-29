package gpkg

type TileMatrixSetList struct {
	tileMatrixSets []TileMatrixSet
}

func (l *TileMatrixSetList) Size() int {
	return len(l.tileMatrixSets)
}

func (l *TileMatrixSetList) Item(i int) *TileMatrixSet {
	return &l.tileMatrixSets[i]
}
