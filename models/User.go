/*
 * @title: 这里写标题
 * @Date: 2022-02-15 17:37:00
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/dao/User.go
 */
package models

type User struct {
	Model
	Name    string `json:"name" gorm:"name"`
	IdCard  string `json:"id_card" gorm:"id_card"`
	Phone   string `json:"phone" gorm:"phone"`
	Address string `json:"address" gorm:"address"`
	Rank    int8   `json:"rank" gorm:"rank"`
}
