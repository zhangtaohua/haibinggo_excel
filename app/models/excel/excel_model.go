//Package excel 模型
package excel

import (
	"haibinggo/app/models"
	"haibinggo/pkg/database"
)

type Excel struct {
	models.BaseModel

	// Put fields in here
	A string `gorm:"type:varchar(255);default:null" json:"A"`
	B string `gorm:"type:varchar(255);default:null" json:"B"`
	C string `gorm:"type:varchar(255);default:null" json:"C"`
	D string `gorm:"type:varchar(255);default:null" json:"D"`
	E string `gorm:"type:varchar(255);default:null" json:"E"`
	F string `gorm:"type:varchar(255);default:null" json:"F"`
	G string `gorm:"type:varchar(255);default:null" json:"G"`

	H string `gorm:"type:varchar(255);default:null" json:"H"`
	I string `gorm:"type:varchar(255);default:null" json:"I"`
	J string `gorm:"type:varchar(255);default:null" json:"J"`
	K string `gorm:"type:varchar(255);default:null" json:"K"`
	L string `gorm:"type:varchar(255);default:null" json:"L"`
	M string `gorm:"type:varchar(255);default:null" json:"M"`
	N string `gorm:"type:varchar(255);default:null" json:"N"`

	O string `gorm:"type:varchar(255);default:null" json:"O"`
	P string `gorm:"type:varchar(255);default:null" json:"P"`
	Q string `gorm:"type:varchar(255);default:null" json:"Q"`
	R string `gorm:"type:varchar(255);default:null" json:"R"`
	S string `gorm:"type:varchar(255);default:null" json:"S"`
	T string `gorm:"type:varchar(255);default:null" json:"T"`

	U string `gorm:"type:varchar(255);default:null" json:"U"`
	V string `gorm:"type:varchar(255);default:null" json:"V"`
	W string `gorm:"type:varchar(255);default:null" json:"W"`
	X string `gorm:"type:varchar(255);default:null" json:"X"`
	Y string `gorm:"type:varchar(255);default:null" json:"Y"`
	Z string `gorm:"type:varchar(255);default:null" json:"Z"`

	AA string `gorm:"type:varchar(255);default:null" json:"AA"`
	AB string `gorm:"type:varchar(255);default:null" json:"AB"`
	AC string `gorm:"type:varchar(255);default:null" json:"AC"`
	AD string `gorm:"type:varchar(255);default:null" json:"AD"`
	AE string `gorm:"type:varchar(255);default:null" json:"AE"`
	AF string `gorm:"type:varchar(255);default:null" json:"AF"`
	AG string `gorm:"type:varchar(255);default:null" json:"AG"`

	TblName string `gorm:"type:varchar(255);default:null" json:"table_name"`

	models.CommonTimestampsField
}

type UpdateRes struct {
	Error        error
	RowsAffected int64
}

func (excel *Excel) Create(tableName string) {
	database.DB.Table(tableName).Create(&excel)
}

func (excel *Excel) Save(tableName string) (rowsAffected int64) {
	result := database.DB.Table(tableName).Save(&excel)
	return result.RowsAffected
}

func (excel *Excel) Update(tableName string, id string, column string, data string) (res *UpdateRes) {
	result := database.DB.Table(tableName).Where("id = ?", id).Update(column, data)
	return &UpdateRes{result.Error, result.RowsAffected}
}

func (excel *Excel) Delete(tableName string) (rowsAffected int64) {
	result := database.DB.Table(tableName).Delete(&excel)
	return result.RowsAffected
}
