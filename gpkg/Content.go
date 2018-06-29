package gpkg

import (
	"time"
)

type Content struct {
	ContentTableName         string     `gorm:"column:table_name;unique;not null;primary_key"`
	Identifier               string     `gorm:"column:identifier;unique"`
	Description              string     `gorm:"column:description;default:''"`
	LastChange               *time.Time `gorm:"column:last_change;"`
	MinX                     float64    `gorm:"column:min_x"`
	MinY                     float64    `gorm:"column:min_y"`
	MaxX                     float64    `gorm:"column:max_x"`
	MaxY                     float64    `gorm:"column:max_y"`
	SpatialReferenceSystemId int        `gorm:"column:srs_id"`
}

func (Content) TableName() string {
	return "gpkg_contents"
}
