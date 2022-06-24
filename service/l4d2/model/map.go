package model

import (
	"database/sql"
)

type MapFile struct {
	Id        int64  `gorm:"column:id;primaryKey" json:"id"`
	FileName  string `gorm:"column:fileName" json:"fileName"`
	GroupID   int64  `gorm:"column:groupId" json:"groupId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type MapGroup struct {
	Id           int64          `gorm:"column:id;primaryKey" json:"id"`
	Name         sql.NullString `gorm:"column:name" json:"name"`
	StartName    sql.NullString `gorm:"column:startName" json:"startName"`
	LastPlayTime int64          `gorm:"column:lastPlayTime" json:"lastPlayTime"`
	CreatedTime  int64          `json:"createdAt"`
	UpdatedTime  int64          `json:"updatedAt"`
}

func (MapFile) TableName() string {
	return "mapFile"
}

func (MapGroup) TableName() string {
	return "mapGroup"
}
