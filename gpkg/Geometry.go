package gpkg

type Geometry struct {
	Type        string      `json:"type" bson:"type" yaml:"type" hcl:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates" yaml:"coordinates" hcl:"coordinates"`
}

func NewPoint(lon float64, lat float64) Geometry {
	return Geometry{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}
}

func NewLine(coordinates [][]float64) Geometry {
	return Geometry{
		Type:        "LineString",
		Coordinates: coordinates,
	}
}

func NewPolygon(coordinates [][][2]float64) Geometry {
	return Geometry{
		Type:        "Polygon",
		Coordinates: coordinates,
	}
}
