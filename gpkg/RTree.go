package gpkg

type RTree struct {
	Table string  `gorm:"-"`
	Id    int     `gorm:"column:id;unique;not null;primary_key"`
	MinX  float64 `gorm:"column:minx;not null"`
	MaxX  float64 `gorm:"column:maxx;not null"`
	MinY  float64 `gorm:"column:miny;not null"`
	MaxY  float64 `gorm:"column:maxy;not null"`
}

func (rt RTree) ResourceName() string {
	return rt.Table
}

func (rt RTree) TableName() string {
	return rt.Table
}
