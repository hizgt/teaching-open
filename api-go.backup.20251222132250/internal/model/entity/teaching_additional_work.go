package entity

import "github.com/gogf/gf/v2/os/gtime"

// TeachingAdditionalWork 附加作业实体
type TeachingAdditionalWork struct {
	Id              string      `json:"id" orm:"id,primary"`
	CreateBy        string      `json:"createBy" orm:"create_by"`
	CreateTime      *gtime.Time `json:"createTime" orm:"create_time"`
	UpdateBy        string      `json:"updateBy" orm:"update_by"`
	UpdateTime      *gtime.Time `json:"updateTime" orm:"update_time"`
	SysOrgCode      string      `json:"sysOrgCode" orm:"sys_org_code"`
	CodeType        string      `json:"codeType" orm:"code_type"`
	WorkName        string      `json:"workName" orm:"work_name"`
	WorkDesc        string      `json:"workDesc" orm:"work_desc"`
	WorkCover       string      `json:"workCover" orm:"work_cover"`
	WorkUrl         string      `json:"workUrl" orm:"work_url"`
	WorkDept        string      `json:"workDept" orm:"work_dept"`
	Status          int         `json:"status" orm:"status"`
	WorkDocumentUrl string      `json:"workDocumentUrl" orm:"work_document_url"`
}
