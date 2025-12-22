// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package entity

// SysRolePermission is the golang structure for table sys_role_permission.
type SysRolePermission struct {
	Id           string `json:"id"           description:"主键id"`
	RoleId       string `json:"roleId"       description:"角色id"`
	PermissionId string `json:"permissionId" description:"权限id"`
	DataRuleIds  string `json:"dataRuleIds"  description:"数据规则ID"`
}
