package gpkg

type TileMatrixSet struct {
	Name                     string   `gorm:"column:table_name;not null;primary_key"`
	SpatialReferenceSystemId *int     `gorm:"column:srs_id;not null"`
	MinX                     *float64 `gorm:"column:min_x;not null"`
	MinY                     *float64 `gorm:"column:min_y;not null"`
	MaxX                     *float64 `gorm:"column:max_x;not null"`
	MaxY                     *float64 `gorm:"column:max_y;not null"`
}

func (TileMatrixSet) TableName() string {
	return "gpkg_tile_matrix_set"
}

func (tms TileMatrixSet) GetSpatialReferenceSystemId() int {
	if tms.SpatialReferenceSystemId == nil {
		return 0
	}
	return *tms.SpatialReferenceSystemId
}
