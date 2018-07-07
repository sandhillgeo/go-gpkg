package gpkg

import (
//"database/sql"
)

type TileMatrix struct {
	Name string `sql:"type:text" gorm:"column:table_name;not null"`
	//ZoomLevel sql.NullInt64 `gorm:"column:zoom_level;not null"`
	ZoomLevel    NullInt8 `gorm:"column:zoom_level;not null"`
	MatrixWidth  uint64   `gorm:"column:matrix_width;not null"`
	MatrixHeight uint64   `gorm:"column:matrix_height;not null"`
	TileWidth    uint32   `gorm:"column:tile_width;not null"`
	TileHeight   uint32   `gorm:"column:tile_height;not null"`
	PixelXSize   float64  `gorm:"column:pixel_x_size;not null"`
	PixelYSize   float64  `gorm:"column:pixel_y_size;not null"`
}

func (TileMatrix) TableName() string {
	return "gpkg_tile_matrix"
}
