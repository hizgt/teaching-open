package export

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

// CSVExporter CSV导出器
type CSVExporter struct {
	buffer *bytes.Buffer
	writer *csv.Writer
}

// NewCSVExporter 创建CSV导出器
func NewCSVExporter() *CSVExporter {
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	// 写入UTF-8 BOM，让Excel正确识别中文
	buffer.Write([]byte{0xEF, 0xBB, 0xBF})

	return &CSVExporter{
		buffer: buffer,
		writer: writer,
	}
}

// SetHeaders 设置表头
func (e *CSVExporter) SetHeaders(headers []string) error {
	return e.writer.Write(headers)
}

// AddRow 添加数据行
func (e *CSVExporter) AddRow(values []string) error {
	return e.writer.Write(values)
}

// AddRows 批量添加数据行
func (e *CSVExporter) AddRows(rows [][]string) error {
	for _, row := range rows {
		if err := e.AddRow(row); err != nil {
			return err
		}
	}
	return nil
}

// GetBuffer 获取CSV字节流
func (e *CSVExporter) GetBuffer() []byte {
	e.writer.Flush()
	return e.buffer.Bytes()
}

// ExportUserListCSV 导出用户列表为CSV
func ExportUserListCSV(users []map[string]interface{}) ([]byte, error) {
	exporter := NewCSVExporter()

	// 设置表头
	headers := []string{"用户名", "姓名", "性别", "手机号", "邮箱", "部门", "状态", "创建时间"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	// 添加数据
	for _, user := range users {
		row := []string{
			toString(user["username"]),
			toString(user["realname"]),
			getSexText(user["sex"]),
			toString(user["phone"]),
			toString(user["email"]),
			toString(user["departName"]),
			getStatusText(user["status"]),
			formatTime(user["createTime"]),
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	return exporter.GetBuffer(), nil
}

// ExportWorkListCSV 导出作品列表为CSV
func ExportWorkListCSV(works []map[string]interface{}) ([]byte, error) {
	exporter := NewCSVExporter()

	headers := []string{"作品标题", "学生姓名", "课程名称", "作品类型", "状态", "分数", "提交时间", "批改时间"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	for _, work := range works {
		row := []string{
			toString(work["title"]),
			toString(work["studentName"]),
			toString(work["courseName"]),
			toString(work["type"]),
			getWorkStatusText(work["status"]),
			toString(work["score"]),
			formatTime(work["submitTime"]),
			formatTime(work["correctTime"]),
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	return exporter.GetBuffer(), nil
}

// ExportLogListCSV 导出日志列表为CSV
func ExportLogListCSV(logs []map[string]interface{}) ([]byte, error) {
	exporter := NewCSVExporter()

	headers := []string{"操作用户", "操作类型", "操作内容", "IP地址", "操作时间", "耗时(ms)"}
	if err := exporter.SetHeaders(headers); err != nil {
		return nil, err
	}

	for _, log := range logs {
		row := []string{
			toString(log["username"]),
			toString(log["logType"]),
			toString(log["logContent"]),
			toString(log["ip"]),
			formatTime(log["createTime"]),
			toString(log["costTime"]),
		}
		if err := exporter.AddRow(row); err != nil {
			return nil, err
		}
	}

	return exporter.GetBuffer(), nil
}

// toString 将interface{}转换为string
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
