// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingScratchAsset is the golang structure for table teaching_scratch_assets.
type TeachingScratchAsset struct {
	Id         string      `json:"id"         description:"主键ID"`
	AssetType  int         `json:"assetType"  description:"素材类型 1背景 2声音 3造型 4角色"`
	AssetName  string      `json:"assetName"  description:"素材名"`
	AssetData  string      `json:"assetData"  description:"素材JSON数据"`
	Md5Ext     string      `json:"md5Ext"     description:"素材md5"`
	Tags       string      `json:"tags"       description:"标签"`
	CreateBy   string      `json:"createBy"   description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"修改人"`
	UdpateTime *gtime.Time `json:"udpateTime" description:"修改时间"`
	DelFlag    int         `json:"delFlag"    description:"删除状态 0正常 1删除"`
}
