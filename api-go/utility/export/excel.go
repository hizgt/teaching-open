package export

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExcelExporter Excel导出器
type ExcelExporter struct {
	file      *excelize.File
	sheetName string
	rowIndex  int
}

// NewExcelExporter 创建Excel导出器
func NewExcelExporter(sheetName string) *ExcelExporter {
	f := excelize.NewFile()
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	return &ExcelExporter{
		file:      f,
		sheetName: sheetName,
		rowIndex:  1,
	}
}

// SetHeaders 设置表头
func (e *ExcelExporter) SetHeaders(headers []string) error {
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", columnName(i), e.rowIndex)
		if err := e.file.SetCellValue(e.sheetName, cell, header); err != nil {
			return err
		}
	}

	// 设置表头样式
	style, err := e.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E0E0E0"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return err
	}

	endCol := columnName(len(headers) - 1)
	if err := e.file.SetCellStyle(e.sheetName, fmt.Sprintf("A%d", e.rowIndex), fmt.Sprintf("%s%d", endCol, e.rowIndex), style); err != nil {
		return err
	}

	e.rowIndex++
	return nil
}

// AddRow 添加数据行
func (e *ExcelExporter) AddRow(values []interface{}) error {
	for i, value := range values {
		cell := fmt.Sprintf("%s%d", columnName(i), e.rowIndex)
		if err := e.file.SetCellValue(e.sheetName, cell, value); err != nil {
			return err
		}
	}
	e.rowIndex++
	return nil
}

// AddRows 批量添加数据行
func (e *ExcelExporter) AddRows(rows [][]interface{}) error {
	for _, row := range rows {
		if err := e.AddRow(row); err != nil {
			return err
		}
	}
	return nil
}

// SetColumnWidth 设置列宽
func (e *ExcelExporter) SetColumnWidth(colIndex int, width float64) error {
	col := columnName(colIndex)
	return e.file.SetColWidth(e.sheetName, col, col, width)
}

// AutoFilter 设置自动筛选
func (e *ExcelExporter) AutoFilter(startCol, endCol int) error {
	start := fmt.Sprintf("%s1", columnName(startCol))
	end := fmt.Sprintf("%s%d", columnName(endCol), e.rowIndex-1)
	return e.file.AutoFilter(e.sheetName, start, end, []excelize.AutoFilterOptions{})
}

// SaveToFile 保存到文件
func (e *ExcelExporter) SaveToFile(filename string) error {
	return e.file.SaveAs(filename)
}

// GetBuffer 获取文件字节流
func (e *ExcelExporter) GetBuffer() ([]byte, error) {
	return e.file.WriteToBuffer()
}

// columnName 将列索引转换为Excel列名 (0->A, 1->B, ..., 25->Z, 26->AA)
func columnName(index int) string {
	name := ""
	for index >= 0 {
		name = string(rune('A'+index%26)) + name
		index = index/26 - 1
	}
	return name
}

// ExportUserList 导出用户列表示例
func ExportUserList(users []map[string]interface{}) ([]byte, error) {
	exporter := NewExcelExporter("用户列表")

	// 设置表头
	headers := []string{"用户名", "姓名", "性别", "手机号", "邮箱", "部门", "状态", "创建时间"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	// 设置列宽
	widths := []float64{15, 15, 10, 15, 20, 20, 10, 20}
	for i, width := range widths {
		exporter.SetColumnWidth(i, width)
	}

	// 添加数据
	for _, user := range users {
		row := []interface{}{
			user["username"],
			user["realname"],
			getSexText(user["sex"]),
			user["phone"],
			user["email"],
			user["departName"],
			getStatusText(user["status"]),
			formatTime(user["createTime"]),
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	// 设置自动筛选
	exporter.AutoFilter(0, len(headers)-1)

	return exporter.GetBuffer()
}

// ExportWorkList 导出作品列表
func ExportWorkList(works []map[string]interface{}) ([]byte, error) {
	exporter := NewExcelExporter("作品列表")

	headers := []string{"作品标题", "学生姓名", "课程名称", "作品类型", "状态", "分数", "提交时间", "批改时间"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	widths := []float64{30, 15, 20, 15, 10, 10, 20, 20}
	for i, width := range widths {
		exporter.SetColumnWidth(i, width)
	}

	for _, work := range works {
		row := []interface{}{
			work["title"],
			work["studentName"],
			work["courseName"],
			work["type"],
			getWorkStatusText(work["status"]),
			work["score"],
			formatTime(work["submitTime"]),
			formatTime(work["correctTime"]),
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	exporter.AutoFilter(0, len(headers)-1)

	return exporter.GetBuffer()
}

// ExportLogList 导出日志列表
func ExportLogList(logs []map[string]interface{}) ([]byte, error) {
	exporter := NewExcelExporter("操作日志")

	headers := []string{"操作用户", "操作类型", "操作内容", "IP地址", "操作时间", "耗时(ms)"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	widths := []float64{15, 15, 30, 15, 20, 10}
	for i, width := range widths {
		exporter.SetColumnWidth(i, width)
	}

	for _, log := range logs {
		row := []interface{}{
			log["username"],
			log["logType"],
			log["logContent"],
			log["ip"],
			formatTime(log["createTime"]),
			log["costTime"],
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	exporter.AutoFilter(0, len(headers)-1)

	return exporter.GetBuffer()
}

// 辅助函数
func getSexText(sex interface{}) string {
	switch sex {
	case 1, "1":
		return "男"
	case 2, "2":
		return "女"
	default:
		return "未知"
	}
}

func getStatusText(status interface{}) string {
	switch status {
	case 1, "1":
		return "正常"
	case 2, "2":
		return "冻结"
	default:
		return "未知"
	}
}

func getWorkStatusText(status interface{}) string {
	switch status {
	case 0, "0":
		return "草稿"
	case 1, "1":
		return "已提交"
	case 2, "2":
		return "已批改"
	default:
		return "未知"
	}
}

func formatTime(t interface{}) string {
	if t == nil {
		return ""
	}

	switch v := t.(type) {
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case string:
		return v
	default:
		return fmt.Sprintf("%v", t)
	}
}
