package gpkg

type SQLiteSequence struct {
	Name  string `gorm:"column:name;unique;not null;primary_key"`
	Value int64  `gorm:"column:seq;not null"`
}

func (SQLiteSequence) TableName() string {
	return "sqlite_sequence"
}
