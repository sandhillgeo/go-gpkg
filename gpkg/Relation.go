package gpkg

type Relation struct {
	Id                   int    `gorm:"column:id;uniqu;not null;primary_key"`
	BaseTableName        string `gorm:"column:base_table_name;not null"`
	BasePrimaryColumn    string `gorm:"column:base_primary_column;not null;default:'id'"`
	RelatedTableName     string `gorm:"column:related_table_name;not null"`
	RelatedPrimaryColumn string `gorm:"column:related_primary_column;not null;default:'id'"`
	RelationName         string `gorm:"column:relation_name;not null"`
	MappingTableName     string `gorm:"column:mapping_table_name;not null"`
}

func (Relation) TableName() string {
	return "gpkgext_relations"
}
