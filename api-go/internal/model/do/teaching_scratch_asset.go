// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingScratchAsset is the golang structure for table teaching_scratch_assets.
type TeachingScratchAsset struct {
	g.Meta     `orm:"table:teaching_scratch_assets, do:true"`
	Id         interface{} // 主键ID
	AssetType  interface{} // 素材类型 1背景 2声音 3造型 4角色
	AssetName  interface{} // 素材名
	AssetData  interface{} // 素材JSON数据
	Md5Ext     interface{} // 素材md5
	Tags       interface{} // 标签
	CreateBy   interface{} // 创建人
	CreateTime interface{} // 创建时间
	UpdateBy   interface{} // 修改人
	UdpateTime interface{} // 修改时间
	DelFlag    interface{} // 删除状态 0正常 1删除
}
