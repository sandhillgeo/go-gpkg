package gpkg

type VectorLayerList struct {
	vectorLayers []VectorLayer
}

func (l *VectorLayerList) Size() int {
	return len(l.vectorLayers)
}

func (l *VectorLayerList) Item(i int) *VectorLayer {
	return &l.vectorLayers[i]
}
