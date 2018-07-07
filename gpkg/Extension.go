package gpkg

type Extension struct {
	Table      string  `sql:"type:text" gorm:"column:table_name"`
	Column     *string `sql:"type:text" gorm:"column:column_name"`
	Extension  string  `sql:"type:text" gorm:"column:extension_name;not null"`
	Definition string  `sql:"type:text" gorm:"column:definition;not null"`
	Scope      string  `sql:"type:text" gorm:"column:scope;not null"`
}

func (Extension) TableName() string {
	return "gpkg_extensions"
}
