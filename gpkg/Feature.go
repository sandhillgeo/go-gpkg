package gpkg

type Feature struct {
	Id           interface{}            `json:"id" bson:"id" yaml:"id" hcl:"id"`
	Type         string                 `json:"type" bson:"type" yaml:"type" hcl:"type"`
	Properties   map[string]interface{} `json:"properties" bson:"properties" yaml:"properties" hcl:"properties"`
	GeometryName string                 `json:"geometry_name" bson:"geometry_name" yaml:"geometry_name" hcl:"geometry_name"`
	Geometry     Geometry               `json:"geometry" bson:"geometry" yaml:"geometry" hcl:"geometry"`
}

func NewFeature(id interface{}, properties map[string]interface{}, geom Geometry) Feature {

	properties["id"] = id

	f := Feature{
		Id:           id,
		Type:         "Feature",
		Properties:   properties,
		GeometryName: "the_geom",
		Geometry:     geom,
	}

	return f
}
