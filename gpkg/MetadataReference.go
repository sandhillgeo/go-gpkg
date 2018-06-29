package gpkg

import (
	"time"
)

type MetadataReference struct {
	ReferenceScope string     `gorm:"column:reference_scope;not null"`
	Name           string     `gorm:"column:table_name"`
	ColumnName     string     `gorm:"column:column_name"`
	RowIdValue     *int       `gorm:"column:row_id_value"`
	Timestamp      *time.Time `gorm:"column:timestamp;not null;"`
	MdFileId       int        `gorm:"column:md_file_id;not null"`
	MdParentId     *int       `gorm:"column:md_parent_id"`
}

func (MetadataReference) TableName() string {
	return "gpkg_metadata_reference"
}
