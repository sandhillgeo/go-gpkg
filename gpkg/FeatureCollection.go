package gpkg

import (
	"encoding/json"
)

type FeatureCollection struct {
	Type                      string                     `json:"type" bson:"type" yaml:"type" hcl:"type"`
	TotalFeatures             int                        `json:"totalFeatures" bson:"totalFeatures" yaml:"totalFeatures" hcl:"totalFeatures"`
	Features                  []Feature                  `json:"features" bson:"features" yaml:"features" hcl:"features"`
	CoordinateReferenceSystem *CoordinateReferenceSystem `json:"crs,omitempty" bson:"crs,omitempty" yaml:"crs,omitempty" hcl:"crs"`
}

func (fc FeatureCollection) ToJson() (string, error) {
	b, err := json.Marshal(fc)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func NewFeatureCollection(features []Feature) FeatureCollection {
	return FeatureCollection{
		Type:                      "FeatureCollection",
		TotalFeatures:             len(features),
		Features:                  features,
		CoordinateReferenceSystem: &CoordinateReferenceSystem{Type: "name", Properties: map[string]interface{}{"name": "EPSG:4326"}},
	}
}
