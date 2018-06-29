package gpkg

type Extent struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func (e Extent) GetCenter() []float64 {
	return []float64{
		(e.MinX + e.MaxX) / 2.0,
		(e.MinY + e.MaxY) / 2.0,
	}
}

func (e Extent) GetCenterX() float64 {
	return (e.MinX + e.MaxX) / 2.0
}

func (e Extent) GetCenterY() float64 {
	return (e.MinY + e.MaxY) / 2.0
}
