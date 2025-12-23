<template>
  <div class="user-role-management">
    <a-card :bordered="false">
      <!-- 查询区域 -->
      <div class="table-page-search-wrapper">
        <a-form layout="inline" @finish="handleSearch">
          <a-row :gutter="24">
            <a-col :md="6" :sm="8">
              <a-form-item label="用户名" name="username">
                <a-input placeholder="请输入用户名" v-model:value="queryParams.username" />
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="8">
              <a-form-item label="姓名" name="realname">
                <a-input placeholder="请输入姓名" v-model:value="queryParams.realname" />
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="8">
              <a-form-item label="角色" name="roleName">
                <a-input placeholder="请输入角色名称" v-model:value="queryParams.roleName" />
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="8">
              <a-form-item>
                <a-button type="primary" html-type="submit">查询</a-button>
                <a-button style="margin-left: 8px" @click="handleReset">重置</a-button>
              </a-form-item>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <!-- 操作按钮区域 -->
      <div class="table-operator">
        <a-button @click="handleAssignRole" type="primary" :disabled="!selectedRowKeys.length">
          分配角色
        </a-button>
        <a-button @click="handleBatchRemoveRole" type="primary" danger :disabled="!selectedRowKeys.length">
          批量移除角色
        </a-button>
      </div>

      <!-- 表格区域 -->
      <a-table
        :loading="loading"
        :dataSource="dataSource"
        :columns="columns"
        :pagination="pagination"
        :rowSelection="{
          selectedRowKeys,
          onChange: onSelectChange,
          type: 'checkbox'
        }"
        @change="handleTableChange"
        rowKey="id"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'roles'">
            <a-tag
              v-for="role in record.roles"
              :key="role.id"
              color="blue"
            >
              {{ role.roleName }}
            </a-tag>
          </template>
          <template v-if="column.key === 'action'">
            <a @click="handleEditRoles(record)">编辑角色</a>
            <a-divider type="vertical" />
            <a @click="handleRemoveRole(record)">移除角色</a>
          </template>
        </template>
      </a-table>
    </a-card>

    <!-- 分配角色弹窗 -->
    <a-modal
      v-model:open="assignModalVisible"
      title="分配角色"
      width="600px"
      @ok="handleAssignConfirm"
      @cancel="handleAssignCancel"
    >
      <div v-if="currentUser">
        <p>为用户 <strong>{{ currentUser.realname }} ({{ currentUser.username }})</strong> 分配角色：</p>
        <a-checkbox-group
          v-model:value="selectedRoleIds"
          :options="roleOptions"
        />
      </div>
    </a-modal>
  </div>
</template>

<script>
import { message } from 'ant-design-vue'

// API接口（需要根据实际API调整）
import { getUserRoleList, assignUserRoles, removeUserRole } from '@/api/modules/system/user-role'

export default {
  name: 'UserRoleManagement',
  data() {
    return {
      loading: false,
      dataSource: [],
      selectedRowKeys: [],
      assignModalVisible: false,
      currentUser: null,
      selectedRoleIds: [],
      roleOptions: [],
      queryParams: {
        username: '',
        realname: '',
        roleName: '',
        page: 1,
        pageSize: 10
      },
      pagination: {
        current: 1,
        pageSize: 10,
        total: 0,
        showSizeChanger: true,
        showQuickJumper: true,
        showTotal: (total) => `共 ${total} 条`
      },
      columns: [
        {
          title: '用户名',
          dataIndex: 'username',
          key: 'username'
        },
        {
          title: '姓名',
          dataIndex: 'realname',
          key: 'realname'
        },
        {
          title: '所属角色',
          key: 'roles',
          width: 300
        },
        {
          title: '创建时间',
          dataIndex: 'createTime',
          key: 'createTime',
          sorter: true
        },
        {
          title: '操作',
          key: 'action',
          width: 150,
          align: 'center'
        }
      ]
    }
  },
  mounted() {
    this.loadData()
  },
  methods: {
// 加载数据
async loadData() {
  this.loading = true
  try {
    const res = await getUserRoleList(this.queryParams)
    this.dataSource = res.records
    this.pagination.total = res.total
  } catch (error) {
    message.error('加载数据失败')
  } finally {
    this.loading = false
  }
},

// 表格变化处理
handleTableChange(pag) {
  this.pagination.current = pag.current
  this.pagination.pageSize = pag.pageSize
  this.queryParams.page = pag.current
  this.queryParams.pageSize = pag.pageSize
  this.loadData()
},

// 选择变化
onSelectChange(keys) {
  this.selectedRowKeys = keys
},

// 查询
handleSearch() {
  this.pagination.current = 1
  this.queryParams.page = 1
  this.loadData()
},

// 重置
handleReset() {
  Object.keys(this.queryParams).forEach(key => {
    if (key !== 'page' && key !== 'pageSize') {
      this.queryParams[key] = ''
    }
  })
  this.pagination.current = 1
  this.queryParams.page = 1
  this.loadData()
},

// 分配角色
handleAssignRole() {
  if (this.selectedRowKeys.length !== 1) {
    message.warning('请选择一个用户进行角色分配')
    return
  }

  const user = this.dataSource.find(item => item.id === this.selectedRowKeys[0])
  if (user) {
    this.currentUser = user
    this.selectedRoleIds = user.roles.map(role => role.id)
    this.assignModalVisible = true
    this.loadRoleOptions()
  }
},

// 编辑角色（单个用户）
handleEditRoles(record) {
  this.currentUser = record
  this.selectedRoleIds = record.roles.map(role => role.id)
  this.assignModalVisible = true
  this.loadRoleOptions()
},

// 移除角色
async handleRemoveRole(record) {
  // 这里可以实现移除单个角色的逻辑
  message.info('移除角色功能待实现')
},

// 批量移除角色
handleBatchRemoveRole() {
  message.info('批量移除角色功能待实现')
},

// 加载角色选项
async loadRoleOptions() {
  // 这里应该调用获取所有角色的API
  // 暂时使用模拟数据
  this.roleOptions = [
    { label: '管理员', value: '1' },
    { label: '教师', value: '2' },
    { label: '学生', value: '3' }
  ]
},

// 确认分配
async handleAssignConfirm() {
  if (!this.currentUser) return

  try {
    await assignUserRoles({
      userId: this.currentUser.id,
      roleIds: this.selectedRoleIds
    })
    message.success('角色分配成功')
    this.assignModalVisible = false
    this.loadData()
  } catch (error) {
    message.error('角色分配失败')
  }
},

// 取消分配
handleAssignCancel() {
  this.assignModalVisible = false
  this.currentUser = null
  this.selectedRoleIds = []
}
}
</script>

<style scoped lang="less">
.user-role-management {
  .table-page-search-wrapper {
    margin-bottom: 16px;
  }

  .table-operator {
    margin-bottom: 16px;
  }

  :deep(.ant-table-tbody > tr > td) {
    vertical-align: top;
  }
}
</style>