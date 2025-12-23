import { request } from '@/api/request'

// 权限相关接口

export interface PermissionNode {
  id: string
  name: string
  parentId?: string
  ruleFlag?: boolean
  children?: PermissionNode[]
}

export interface PermissionTreeResult {
  treeList: PermissionNode[]
  ids: string[]
}

export interface RolePermissionResult {
  permissionIds: string[]
}

export interface AssignRolePermissionParams {
  roleId: string
  permissionIds: string
  lastPermissionIds: string
}

// 获取权限树
export function getPermissionTree() {
  return request.get<PermissionTreeResult>('/sys/permission/treeList')
}

// 获取角色权限
export function getRolePermissions(roleId: string) {
  return request.get<RolePermissionResult>(`/sys/role/permission/${roleId}`)
}

// 分配角色权限
export function assignRolePermissions(data: AssignRolePermissionParams) {
  return request.post('/sys/role/permission/assign', data)
}

// 获取权限列表（分页）
export function getPermissionList(params: {
  name?: string
  type?: string
  page: number
  pageSize: number
}) {
  return request.get('/sys/permission/list', params)
}

// 添加权限
export function addPermission(data: Partial<PermissionNode>) {
  return request.post('/sys/permission/add', data)
}

// 编辑权限
export function editPermission(data: Partial<PermissionNode>) {
  return request.put('/sys/permission/edit', data)
}

// 删除权限
export function deletePermission(id: string) {
  return request.delete(`/sys/permission/delete`, { id })
}