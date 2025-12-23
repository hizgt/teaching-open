import { request } from '@/api/request'

// 部门相关接口

export interface DepartmentNode {
  id: string
  departName: string
  parentId?: string
  key: string
  title: string
  children?: DepartmentNode[]
}

export interface DepartmentPermission {
  dataScope: string
  status: boolean
  customDeptIds?: string[]
  rules: PermissionRule[]
}

export interface PermissionRule {
  field: string
  operator: string
  value: string
  description: string
}

export interface DepartmentPermissionResult {
  department: DepartmentNode
  permission: DepartmentPermission
}

export interface SaveDepartmentPermissionParams {
  departmentId: string
  permission: DepartmentPermission
}

// 获取部门树
export function getDepartmentTree() {
  return request.get<DepartmentNode[]>('/sys/depart/treeList')
}

// 获取部门权限配置
export function getDepartmentPermissions(departmentId: string) {
  return request.get<DepartmentPermissionResult>(`/sys/depart/permission/${departmentId}`)
}

// 保存部门权限配置
export function saveDepartmentPermissions(data: SaveDepartmentPermissionParams) {
  return request.post('/sys/depart/permission/save', data)
}

// 获取部门列表（分页）
export function getDepartmentList(params: {
  departName?: string
  orgCode?: string
  page: number
  pageSize: number
}) {
  return request.get('/sys/depart/list', params)
}

// 添加部门
export function addDepartment(data: Partial<DepartmentNode>) {
  return request.post('/sys/depart/add', data)
}

// 编辑部门
export function editDepartment(data: Partial<DepartmentNode>) {
  return request.put('/sys/depart/edit', data)
}

// 删除部门
export function deleteDepartment(id: string) {
  return request.delete(`/sys/depart/delete`, { id })
}