package gpkg

type Metadata struct {
	Id            int    `gorm:"column:id;not null;primary_key"`
	MdScope       string `gorm:"column:md_scope;not null;default:dataset"`
	MdStandardUri string `gorm:"column:md_standard_uri;not null"`
	MimeType      string `gorm:"column:mime_type;not null;default:text/xml"`
	Metadata      string `gorm:"column:metadata;not null;default:"`
}

func (Metadata) TableName() string {
	return "gpkg_metadata"
}
