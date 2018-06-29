package gpkg

type GeometryColumn struct {
	GeometryColumnTableName  string `gorm:"column:table_name;unique;not null;primary_key"`
	ColumnName               string `gorm:"column:column_name;not null"`
	GeometryType             string `gorm:"column:geometry_type_name;not null"`
	SpatialReferenceSystemId int    `gorm:"column:srs_id"`
	Z                        int    `gorm:"column:z;not null"`
	M                        int    `gorm:"column:m;not null"`
}

func (GeometryColumn) TableName() string {
	return "gpkg_geometry_columns"
}
