package gpkg

import (
	"strconv"
	"strings"
)

type SpatialReferenceSystem struct {
	Name                           string `gorm:"column:srs_name;unique;not null;primary_key"`
	SpatialReferenceSystemId       *int   `gorm:"column:srs_id;unique;not null;primary_key"`
	Organization                   string `gorm:"column:organization;not null" json:"org"`
	OrganizationCoordinateSystemId *int   `gorm:"column:organization_coordsys_id;not null" json:"org_id"`
	Definition                     string `gorm:"column:definition;not null" json:"def"`
	Description                    string `gorm:"column:description" json:"description"`
}

func (srs *SpatialReferenceSystem) Code() string {
	if len(srs.Organization) > 0 && srs.OrganizationCoordinateSystemId != nil {
		return strings.ToUpper(srs.Organization + ":" + strconv.Itoa(*srs.OrganizationCoordinateSystemId))
	}
	return ""
}

func (SpatialReferenceSystem) TableName() string {
	return "gpkg_spatial_ref_sys"
}
