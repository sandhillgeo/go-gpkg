package gpkg

import (
	"time"
)

type Content struct {
	ContentTableName         string     `sql:"type:text" gorm:"column:table_name;unique;not null;primary_key"`
	DataType                 string     `sql:"type:text" gorm:"column:data_type;not null"`
	Identifier               string     `sql:"type:text" gorm:"column:identifier;unique"`
	Description              string     `sql:"type:text" gorm:"column:description;default:''"`
	LastChange               *time.Time `gorm:"column:last_change;not null"`
	MinX                     float64    `gorm:"column:min_x"`
	MinY                     float64    `gorm:"column:min_y"`
	MaxX                     float64    `gorm:"column:max_x"`
	MaxY                     float64    `gorm:"column:max_y"`
	SpatialReferenceSystemId int        `sql:"type:integer REFERENCES gpkg_spatial_ref_sys(srs_id)" gorm:"column:srs_id"`
}

func (Content) TableName() string {
	return "gpkg_contents"
}
