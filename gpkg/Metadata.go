package gpkg

type Metadata struct {
	Id            int    `gorm:"column:id;not null;primary_key"`
	MdScope       string `sql:"type:text" gorm:"column:md_scope;not null;default:'dataset'"`
	MdStandardUri string `sql:"type:text" gorm:"column:md_standard_uri;not null"`
	MimeType      string `sql:"type:text" gorm:"column:mime_type;not null;default:'text/xml'"`
	Metadata      string `sql:"type:text" gorm:"column:metadata;not null;default:''"`
}

func (Metadata) TableName() string {
	return "gpkg_metadata"
}
