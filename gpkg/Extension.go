package gpkg

type Extension struct {
	Table      string `gorm:"column:table_name"`
	Column     *string `gorm:"column:column_name"`
	Extension  string `gorm:"column:extension_name;not null"`
	Definition string `gorm:"column:definition;not null"`
	Scope      string `gorm:"column:scope;not null"`
}

func (Extension) TableName() string {
	return "gpkg_extensions"
}
