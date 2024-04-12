package v1

import (
	"fmt"
	"haibinggo/app/models/excel"
	"haibinggo/app/models/user"
	"haibinggo/app/requests"
	"haibinggo/pkg/app"
	"haibinggo/pkg/auth"
	"haibinggo/pkg/config"
	"haibinggo/pkg/helpers"
	"haibinggo/pkg/response"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

type ExcelsController struct {
	BaseAPIController
}

func (ctrl *ExcelsController) Index(c *gin.Context) {
	userId := auth.CurrentUID(c)

	request := requests.ExcelQueryRequest{}
	if ok := requests.Validate(c, &request, requests.ExcelQuerySave); !ok {
		return
	}

	var excelAll []excel.Excel
	tableName := request.Table
	whereFields := []interface{}{}

	if userId == "1" || userId == "2" || userId == "3" {
		if !helpers.Empty(request.CreatedBeginAt) && !helpers.Empty(request.CreatedEndAt) {
			whereFields = append(whereFields, `created_at BETWEEN ? AND ?`)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)

			excelAll = excel.AllByFields(tableName, whereFields)
		} else {
			excelAll = excel.All(tableName)
		}
	} else {
		if !helpers.Empty(request.CreatedBeginAt) && !helpers.Empty(request.CreatedEndAt) {
			whereFields = append(whereFields, `( y = ? AND created_at BETWEEN ? AND ? ) OR ( y IN ? AND ( z IS NULL OR z = ? )  AND created_at BETWEEN ? AND ? ) `)
			whereFields = append(whereFields, userId)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)
			whereFields = append(whereFields, []interface{}{"1", "2", "3"})
			whereFields = append(whereFields, userId)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)

			excelAll = excel.AllByFields(tableName, whereFields)
		} else {
			excelAll = excel.All(tableName)
		}
	}
	response.Data(c, excelAll)
}

func (ctrl *ExcelsController) Store(c *gin.Context) {
	request := requests.ExcelRequest{}
	if ok := requests.Validate(c, &request, requests.ExcelSave); !ok {
		return
	}

	tableName := request.Table
	userId := auth.CurrentUID(c)

	excelModel := excel.Excel{
		TblName: tableName,
		Y:       userId,
	}

	v := reflect.ValueOf(&excelModel).Elem()

	f := v.FieldByName(request.Column)
	if f.IsValid() && f.CanSet() {
		f.SetString(request.Data)
	}

	excelModel.Create(tableName)
	if excelModel.ID > 0 {
		response.Created(c, excelModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ExcelsController) Update(c *gin.Context) {
	tableName := c.Param("table")
	userId := auth.CurrentUID(c)

	excelModel := excel.Get(tableName, c.Param("row"))
	if excelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	createId := excelModel.Y
	yunyingId := excelModel.Z

	if userId != "1" && userId != "2" && userId != "3" {
		if createId == "1" || createId == "2" || createId == "3" {
			if !helpers.Empty(yunyingId) && yunyingId != userId {
				response.Abort403(c)
				return
			}
		} else {
			if createId != userId {
				response.Abort403(c)
				return
			}
		}
	}

	request := requests.ExcelRequest{}
	bindOk := requests.Validate(c, &request, requests.ExcelSave)
	if !bindOk {
		return
	}

	rowsAffected := excelModel.Update(tableName, request.Row, request.Column, request.Data)
	if rowsAffected.Error != nil {
		response.Abort500(c, "更新失败，请稍后尝试~")
	} else {
		response.Success(c)

	}
}

func (ctrl *ExcelsController) Delete(c *gin.Context) {
	tableName := c.Param("table")
	userId := auth.CurrentUID(c)

	excelModel := excel.Get(tableName, c.Param("row"))
	if excelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	createId := excelModel.Y
	if userId != "1" && userId != "2" && userId != "3" {
		if createId != userId {
			response.Abort403(c)
			return
		}
	}

	rowsAffected := excelModel.Delete(tableName)
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *ExcelsController) Download(c *gin.Context) {
	userId := auth.CurrentUID(c)

	request := requests.ExcelQueryRequest{}
	if ok := requests.Validate(c, &request, requests.ExcelQuerySave); !ok {
		return
	}

	var excelAll []excel.Excel
	tableName := request.Table
	whereFields := []interface{}{}

	if userId == "1" || userId == "2" || userId == "3" {
		if !helpers.Empty(request.CreatedBeginAt) && !helpers.Empty(request.CreatedEndAt) {
			whereFields = append(whereFields, `created_at BETWEEN ? AND ?`)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)

			excelAll = excel.AllByFields(tableName, whereFields)
		} else {
			excelAll = excel.All(tableName)
		}
	} else {
		if !helpers.Empty(request.CreatedBeginAt) && !helpers.Empty(request.CreatedEndAt) {
			whereFields = append(whereFields, `( y = ? AND created_at BETWEEN ? AND ? ) OR ( y IN ? AND ( z IS NULL OR z = ? )  AND created_at BETWEEN ? AND ? ) `)
			whereFields = append(whereFields, userId)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)
			whereFields = append(whereFields, []interface{}{"1", "2", "3"})
			whereFields = append(whereFields, userId)
			whereFields = append(whereFields, request.CreatedBeginAt)
			whereFields = append(whereFields, request.CreatedEndAt)

			excelAll = excel.AllByFields(tableName, whereFields)
		} else {
			excelAll = excel.All(tableName)
		}
	}

	f := excelize.NewFile()

	index, err := f.NewSheet("Sheet1")
	if err != nil {
		response.Abort500(c, "创建Sheet失败，请稍后尝试~")
		return
	}
	f.SetActiveSheet(index)

	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "表名")
	f.SetCellValue("Sheet1", "C1", "ID")

	f.SetCellValue("Sheet1", "D1", "商品名称")
	f.SetCellValue("Sheet1", "E1", "商品串号")
	f.SetCellValue("Sheet1", "F1", "商品类型")
	f.SetCellValue("Sheet1", "G1", "供应商")

	f.SetCellValue("Sheet1", "H1", "采购员")
	f.SetCellValue("Sheet1", "I1", "入库时间")
	f.SetCellValue("Sheet1", "J1", "入库单价(¥)")

	f.SetCellValue("Sheet1", "K1", "刷机费(¥)")
	f.SetCellValue("Sheet1", "L1", "耗材费(¥)")
	f.SetCellValue("Sheet1", "M1", "附加费(¥)")
	f.SetCellValue("Sheet1", "N1", "物流费(¥)")
	f.SetCellValue("Sheet1", "O1", "销售数量")
	f.SetCellValue("Sheet1", "P1", "总成本(¥)")
	f.SetCellValue("Sheet1", "Q1", "销售币种")

	f.SetCellValue("Sheet1", "R1", "汇率(¥)")
	f.SetCellValue("Sheet1", "S1", "销售单价")
	f.SetCellValue("Sheet1", "T1", "销售单价(¥)")
	f.SetCellValue("Sheet1", "U1", "销售总价(¥)")
	f.SetCellValue("Sheet1", "V1", "利润(¥)")
	f.SetCellValue("Sheet1", "W1", "销售平台")

	f.SetCellValue("Sheet1", "X1", "店铺名称")
	f.SetCellValue("Sheet1", "Y1", "销售平台订单号")
	f.SetCellValue("Sheet1", "Z1", "物流平台")
	f.SetCellValue("Sheet1", "AA1", "物流单号")

	f.SetCellValue("Sheet1", "AB1", "入库员")

	f.SetCellValue("Sheet1", "AC1", "运营员")

	f.SetCellValue("Sheet1", "AD1", "备注")
	f.SetCellValue("Sheet1", "AE1", "创建时间")
	f.SetCellValue("Sheet1", "AF1", "更新时间")
	f.SetCellValue("Sheet1", "AG1", "删除时间")

	users := user.All()
	for rowIndex, rowData := range excelAll {
		rowId := cast.ToString(rowIndex + 2)
		f.SetCellValue("Sheet1", "A"+rowId, rowIndex)
		f.SetCellValue("Sheet1", "B"+rowId, rowData.TblName)
		f.SetCellValue("Sheet1", "C"+rowId, rowData.ID)

		f.SetCellValue("Sheet1", "D"+rowId, rowData.A)
		f.SetCellValue("Sheet1", "E"+rowId, rowData.B)
		f.SetCellValue("Sheet1", "F"+rowId, rowData.C)
		f.SetCellValue("Sheet1", "G"+rowId, rowData.D)

		username := ""
		for _, user := range users {
			if cast.ToString(user.ID) == rowData.E {
				username = user.NickName
			}
		}
		f.SetCellValue("Sheet1", "H"+rowId, username)
		f.SetCellValue("Sheet1", "I"+rowId, rowData.F)
		f.SetCellValue("Sheet1", "J"+rowId, rowData.G)

		f.SetCellValue("Sheet1", "K"+rowId, rowData.H)
		f.SetCellValue("Sheet1", "L"+rowId, rowData.I)
		f.SetCellValue("Sheet1", "M"+rowId, rowData.J)
		f.SetCellValue("Sheet1", "N"+rowId, rowData.K)
		f.SetCellValue("Sheet1", "O"+rowId, rowData.L)
		f.SetCellValue("Sheet1", "P"+rowId, rowData.M)
		f.SetCellValue("Sheet1", "Q"+rowId, rowData.N)

		f.SetCellValue("Sheet1", "R"+rowId, rowData.O)
		f.SetCellValue("Sheet1", "S"+rowId, rowData.P)
		f.SetCellValue("Sheet1", "T"+rowId, rowData.Q)
		f.SetCellValue("Sheet1", "U"+rowId, rowData.R)
		f.SetCellValue("Sheet1", "V"+rowId, rowData.S)
		f.SetCellValue("Sheet1", "W"+rowId, rowData.T)

		f.SetCellValue("Sheet1", "X"+rowId, rowData.U)
		f.SetCellValue("Sheet1", "Y"+rowId, rowData.V)
		f.SetCellValue("Sheet1", "Z"+rowId, rowData.W)
		f.SetCellValue("Sheet1", "AA"+rowId, rowData.X)

		for _, user := range users {
			if cast.ToString(user.ID) == rowData.Y {
				username = user.NickName
			}
		}

		f.SetCellValue("Sheet1", "AB"+rowId, username)

		for _, user := range users {
			if cast.ToString(user.ID) == rowData.Z {
				username = user.NickName
			}
		}
		f.SetCellValue("Sheet1", "AC"+rowId, username)

		f.SetCellValue("Sheet1", "AD"+rowId, rowData.AA)

		f.SetCellValue("Sheet1", "AE"+rowId, rowData.CreatedAt)
		f.SetCellValue("Sheet1", "AF"+rowId, rowData.UpdatedAt)
		f.SetCellValue("Sheet1", "AG"+rowId, rowData.DeletedAt)
	}

	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/excel/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)

	// Save spreadsheet by the given path.
	fileName := helpers.RandomString(16) + ".xlsx"

	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	xlsxPath := publicPath + dirName + fileName

	fmt.Printf("保存的地址为： %s", xlsxPath)
	if err := f.SaveAs(xlsxPath); err != nil {
		response.Abort500(c, "保存失败，请稍后尝试~")
		return
	}

	if err := f.Close(); err != nil {
		response.Abort500(c, "关闭excel失败，请稍后尝试~")
		return
	}

	resPath := config.GetString("app.url") + dirName + fileName

	response.Data(c, gin.H{
		"success": true,
		"message": "操作成功！",
		"url":     resPath,
	})
}
