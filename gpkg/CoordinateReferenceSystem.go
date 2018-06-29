package gpkg

type CoordinateReferenceSystem struct {
	Type                      string                     `json:"type" bson:"type" yaml:"type" hcl:"type"`
	Properties                map[string]interface{}     `json:"properties" bson:"properties" yaml:"properties" hcl:"properties"`
	CoordinateReferenceSystem *CoordinateReferenceSystem `json:"crs,omitempty" bson:"crs,omitempty" yaml:"crs,omitempty" hcl:"crs"`
}
