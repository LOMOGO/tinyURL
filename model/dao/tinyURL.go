package dao

import (
	"tinyURL/global"
	"tinyURL/model"
	"tinyURL/utils"
)

//生成短链接
func GenderURLCode(u *model.URLInfo) (string, error) {
	if err := insertURL(u); err != nil {
		return "", err
	}
	u.URLCode = utils.Encode(int64(u.ID))
	if err := updateByID(u); err != nil {
		return "", err
	}
	return u.URLCode, nil
}

func SelectByCode(u *model.URLInfo) error {
	tx := global.DB.Where("url_code = ?", u.URLCode).First(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func updateByID(u *model.URLInfo) error {
	tx := global.DB.Model(&model.URLInfo{}).Where("id = ?", u.ID).Update("url_code", u.URLCode)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func insertURL(u *model.URLInfo) error {
	tx := global.DB.Create(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}
