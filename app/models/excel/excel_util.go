package excel

import (
	"haibinggo/pkg/database"
)

func Get(tableName, idstr string) (excel Excel) {
	database.DB.Table(tableName).Where("id", idstr).First(&excel)
	return
}

func GetBy(tableName, field, value string) (excel Excel) {
	database.DB.Table(tableName).Where("? = ?", field, value).First(&excel)
	return
}

func All(tableName string) (excels []Excel) {
	database.DB.Table(tableName).Find(&excels)
	return
}

func AllByFields(tableName string, whereFields []interface{}) (excels []Excel) {
	whereFieldsValue := whereFields[1:]
	database.DB.Table(tableName).Where(whereFields[0], whereFieldsValue...).Find(&excels)
	return
}

func IsExist(tableName, field, value string) bool {
	var count int64
	database.DB.Table(tableName).Model(Excel{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}
