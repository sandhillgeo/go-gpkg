package gpkg

import (
	"io"
)

type TileMatrixSetIterator struct {
	tileMatrixSets []TileMatrixSet
	index          int
}

func (it *TileMatrixSetIterator) Next() (*TileMatrixSet, error) {
	if it.index > len(it.tileMatrixSets)-1 {
		return &TileMatrixSet{}, io.EOF
	}
	tms := it.tileMatrixSets[it.index]
	it.index += 1
	return &tms, nil
}
