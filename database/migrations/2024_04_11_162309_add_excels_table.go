package migrations

import (
	"database/sql"
	"haibinggo/app/models"
	"haibinggo/pkg/database"
	"haibinggo/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Excel struct {
		models.BaseModel

		A string `gorm:"type:varchar(255);default:null" json:"A"`
		B string `gorm:"type:varchar(255);default:null" json:"B"`
		C string `gorm:"type:varchar(255);default:null" json:"C"`
		D string `gorm:"type:varchar(255);default:null" json:"D"`
		E string `gorm:"type:varchar(255);default:null" json:"E"`
		F string `gorm:"type:varchar(255);default:null" json:"F"`
		G string `gorm:"type:varchar(255);default:0.0" json:"G"`

		H string `gorm:"type:varchar(255);default:0.0" json:"H"`
		I string `gorm:"type:varchar(255);default:0.0" json:"I"`
		J string `gorm:"type:varchar(255);default:0.0" json:"J"`
		K string `gorm:"type:varchar(255);default:0.0" json:"K"`
		L string `gorm:"type:varchar(255);default:1" json:"L"`
		M string `gorm:"type:varchar(255);default:0.0" json:"M"`
		N string `gorm:"type:varchar(255);default:null" json:"N"`

		O string `gorm:"type:varchar(255);default:1.0" json:"O"`
		P string `gorm:"type:varchar(255);default:0.0" json:"P"`
		Q string `gorm:"type:varchar(255);default:0.0" json:"Q"`
		R string `gorm:"type:varchar(255);default:0.0" json:"R"`
		S string `gorm:"type:varchar(255);default:0.0" json:"S"`
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

	var tableNameArr = []string{"order2024", "order2025", "order2026", "order2027", "order2028", "order2029", "order2030"}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		for _, tableName := range tableNameArr {
			if !migrator.HasTable(tableName) {
				database.DB.Table(tableName).AutoMigrate(&Excel{})
			}
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		for _, tableName := range tableNameArr {
			migrator.DropTable(tableName)
		}

	}

	migrate.Add("2024_04_11_162309_add_excels_table", up, down)
}
