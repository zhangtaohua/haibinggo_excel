package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ExcelRequest struct {
	Column string `valid:"column" json:"column"`
	Row    string `valid:"row" json:"row"`
	Data   string `valid:"data" json:"data"`
	Table  string `valid:"table" json:"table"`
}

func ExcelSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"column": []string{"required"},
		"row":    []string{"required"},
		"data":   []string{"required"},
		"table":  []string{"required"},
	}
	messages := govalidator.MapData{
		"column": []string{
			"required:列为必填项",
		},
		"row": []string{
			"required:行为必镇项",
		},
		"data": []string{
			"required:数据为必镇项",
		},
		"table": []string{
			"required:表名为必镇项",
		},
	}
	return validate(data, rules, messages)
}

type ExcelQueryRequest struct {
	CreatedBeginAt string `valid:"created_at" form:"created_begin_at,omitempty"`
	CreatedEndAt   string `valid:"created_at" form:"created_end_at,omitempty"`
	Table          string `valid:"table" form:"table"`
}

func ExcelQuerySave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"created_at": []string{},
		"table":      []string{"required"},
	}
	messages := govalidator.MapData{
		"created_at": []string{},
		"table": []string{
			"required:表名为必镇项",
		},
	}
	return validate(data, rules, messages)
}
