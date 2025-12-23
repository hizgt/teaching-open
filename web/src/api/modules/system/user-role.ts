import { request } from '@/api/request'

// 用户角色相关接口

export interface UserRoleQuery {
  username?: string
  realname?: string
  roleName?: string
  page: number
  pageSize: number
}

export interface UserRole {
  id: string
  username: string
  realname: string
  roles: Role[]
  createTime: string
}

export interface Role {
  id: string
  roleName: string
  roleCode: string
}

export interface AssignUserRolesParams {
  userId: string
  roleIds: string[]
}

export interface UserRoleListResult {
  records: UserRole[]
  total: number
  pageNo: number
  pageSize: number
}

// 获取用户角色列表
export function getUserRoleList(params: UserRoleQuery) {
  return request.get<UserRoleListResult>('/sys/user/roleList', params)
}

// 分配用户角色
export function assignUserRoles(data: AssignUserRolesParams) {
  return request.post('/sys/user/assignRoles', data)
}

// 移除用户角色
export function removeUserRole(userId: string, roleId: string) {
  return request.delete(`/sys/user/removeRole`, { userId, roleId })
}

// 获取所有角色列表
export function getAllRoles() {
  return request.get<Role[]>('/sys/role/listAll')
}