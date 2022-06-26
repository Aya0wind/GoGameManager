package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type L4D2Model struct {
	Db *gorm.DB
}

func (receiver *L4D2Model) QueryUserByUsernameAndPassword(username, password string) (user User, err error) {
	err = receiver.Db.Where("username = ? and password = ?", username, password).Limit(1).Find(&user).Error
	return
}

//地图组相关数据库查询

func (receiver *L4D2Model) QueryAllMapGroup() (groups []MapGroup, err error) {
	d := receiver.Db.Find(&groups)
	if d.Error != nil {
		return nil, d.Error
	}
	return
}

func (receiver *L4D2Model) QueryMapGroupByID(id int64) (group MapGroup, err error) {
	err = receiver.Db.Where("id = ?", id).First(&group).Error
	return
}

func (receiver *L4D2Model) InsertMapGroup(group *MapGroup) (err error) {
	err = receiver.Db.Create(group).Error
	return
}

func (receiver *L4D2Model) DeleteMapGroupAndFilesByID(id int64) (needDeleteFiles []MapFile, err error) {
	//执行数据库事务保证删除原子性
	err = receiver.Db.Transaction(func(tx *gorm.DB) error {
		//先删除地图组
		err = tx.Where("id = ?", id).Delete(&MapGroup{}).Error
		if err != nil {
			return err
		}
		//找到地图组下属的所有地图文件
		err = tx.Where("groupId = ?", id).Find(&needDeleteFiles).Error
		if err != nil {
			return err
		}

		//删除所有地图文件的数据库字段
		err = tx.Where("groupId = ?", id).Delete(&MapFile{}).Error
		return err
	})
	return
}

func (receiver *L4D2Model) DeleteMapGroupByID(id int64) (err error) {
	err = receiver.Db.Where("id = ?", id).Delete(&MapFile{}).Error
	return
}

func (receiver *L4D2Model) UpdateMapGroupPlayTime(id int64) (err error) {
	err = receiver.Db.Model(&MapGroup{}).Where("id = ?", id).Update("lastPlayDateTime", time.Now()).Error
	return
}

//地图文件相关数据库查询

func (receiver *L4D2Model) QueryMapFileByMapGroupID(id int64) (files []MapFile, err error) {
	err = receiver.Db.Find(&files, "groupId = ?", id).Error
	return
}

func (receiver *L4D2Model) InsertMapFile(info *MapFile) (err error) {
	err = receiver.Db.Create(info).Error
	return
}

func (receiver *L4D2Model) InsertMapGroupAndMapFile(mapGroup *MapGroup, fileName string) (err error) {
	err = receiver.Db.Transaction(func(tx *gorm.DB) error {
		newGroupId := mapGroup.Id
		if mapGroup.Id == 0 {
			err = receiver.Db.Create(mapGroup).Error
			if err != nil {
				return err
			}
			var group MapGroup
			err = tx.Last(&group).Error
			if err != nil {
				return err
			}
			newGroupId = group.Id
		} else {
			err = receiver.Db.First(&MapGroup{}, newGroupId).Error
			if err != nil {
				return &StringError{
					ErrorString: "地图组不存在",
				}
			}
		}
		err = tx.Create(&MapFile{
			Id:        0,
			FileName:  fileName,
			GroupID:   newGroupId,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}).Error
		return err
	})
	return
}
