<template>
  <div class="department-permission-config">
    <a-card :bordered="false">
      <template #title>
        <div class="card-title">
          <span>部门权限配置</span>
          <a-tag v-if="currentDepartment" color="blue">
            {{ currentDepartment.departName }}
          </a-tag>
        </div>
      </template>

      <a-alert
        message="部门权限说明"
        description="配置部门数据权限范围，控制部门成员可以访问的数据范围。支持按部门、机构等维度进行数据隔离。"
        type="info"
        show-icon
        style="margin-bottom: 24px"
      />

      <!-- 权限配置表单 -->
      <a-form :form="form" layout="vertical">
        <a-row :gutter="24">
          <a-col :span="12">
            <a-form-item label="数据权限范围">
              <a-radio-group v-decorator="['dataScope', { initialValue: '1' }]">
                <a-radio value="1">全部数据</a-radio>
                <a-radio value="2">本部门数据</a-radio>
                <a-radio value="3">本部门及下级部门数据</a-radio>
                <a-radio value="4">仅本人数据</a-radio>
                <a-radio value="5">自定义部门</a-radio>
              </a-radio-group>
            </a-form-item>
          </a-col>

          <a-col :span="12">
            <a-form-item label="权限状态">
              <a-switch
                v-decorator="['status', { initialValue: true, valuePropName: 'checked' }]"
                checked-children="启用"
                un-checked-children="禁用"
              />
            </a-form-item>
          </a-col>
        </a-row>

        <!-- 自定义部门选择 -->
        <a-form-item
          v-show="form.getFieldValue('dataScope') === '5'"
          label="选择部门"
        >
          <a-tree-select
            v-decorator="['customDeptIds', {
              rules: [{ required: form.getFieldValue('dataScope') === '5', message: '请选择部门' }]
            }]"
            :tree-data="departmentTree"
            placeholder="请选择部门"
            tree-checkable
            allow-clear
            tree-default-expand-all
            :field-names="{ children: 'children', label: 'title', value: 'key' }"
          />
        </a-form-item>

        <!-- 权限规则列表 -->
        <a-form-item label="权限规则">
          <div class="permission-rules">
            <a-button type="dashed" block @click="addPermissionRule" style="margin-bottom: 16px">
              <PlusOutlined />
              添加权限规则
            </a-button>

            <div v-for="(rule, index) in permissionRules" :key="index" class="rule-item">
              <a-row :gutter="16" align="middle">
                <a-col :span="4">
                  <a-select
                    v-model:value="rule.field"
                    placeholder="选择字段"
                    size="small"
                  >
                    <a-select-option value="create_by">创建人</a-select-option>
                    <a-select-option value="dept_id">部门ID</a-select-option>
                    <a-select-option value="org_code">机构编码</a-select-option>
                  </a-select>
                </a-col>

                <a-col :span="3">
                  <a-select
                    v-model:value="rule.operator"
                    placeholder="条件"
                    size="small"
                  >
                    <a-select-option value="=">=</a-select-option>
                    <a-select-option value="!=">!=</a-select-option>
                    <a-select-option value="in">in</a-select-option>
                    <a-select-option value="like">like</a-select-option>
                  </a-select>
                </a-col>

                <a-col :span="6">
                  <a-input
                    v-model:value="rule.value"
                    placeholder="值"
                    size="small"
                  />
                </a-col>

                <a-col :span="8">
                  <a-input
                    v-model:value="rule.description"
                    placeholder="规则描述"
                    size="small"
                  />
                </a-col>

                <a-col :span="3">
                  <a-button
                    type="link"
                    danger
                    size="small"
                    @click="removePermissionRule(index)"
                  >
                    <DeleteOutlined />
                  </a-button>
                </a-col>
              </a-row>
            </div>
          </div>
        </a-form-item>
      </a-form>

      <!-- 操作按钮 -->
      <div class="action-buttons">
        <a-space>
          <a-button @click="handleReset">重置</a-button>
          <a-button type="primary" @click="handleSave" :loading="saving">
            保存配置
          </a-button>
        </a-space>
      </div>
    </a-card>
  </div>
</template>

<script>
import { message } from 'ant-design-vue'
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { getDepartmentTree, getDepartmentPermissions, saveDepartmentPermissions } from '@/api/modules/system/department'

export default {
  name: 'DepartmentPermissionModal',
  props: {
    departmentId: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      saving: false,
      currentDepartment: null,
      departmentTree: [],
      permissionRules: [],
      form: {
        dataScope: '1',
        status: true,
        customDeptIds: []
      }
    }
  },
  watch: {
    departmentId: {
      handler(newId) {
        if (newId) {
          this.loadDepartmentPermissions()
        }
      },
      immediate: true
    }
  },
  mounted() {
    this.loadDepartmentTree()
  },
  methods: {
async loadDepartmentTree() {
  try {
    const res = await getDepartmentTree()
    this.departmentTree = res
  } catch (error) {
    message.error('加载部门树失败')
  }
},

// 加载部门权限配置
async loadDepartmentPermissions() {
  if (!this.departmentId) return

  try {
    const res = await getDepartmentPermissions(this.departmentId)
    const permission = res.permission

    // 设置表单值
    this.form.dataScope = permission.dataScope
    this.form.status = permission.status
    this.form.customDeptIds = permission.customDeptIds || []

    // 设置权限规则
    this.permissionRules = permission.rules || []

    // 设置当前部门信息
    this.currentDepartment = res.department
  } catch (error) {
    message.error('加载部门权限配置失败')
  }
},

// 添加权限规则
addPermissionRule() {
  this.permissionRules.push({
    field: '',
    operator: '=',
    value: '',
    description: ''
  })
},

// 移除权限规则
removePermissionRule(index) {
  this.permissionRules.splice(index, 1)
},

// 保存配置
async handleSave() {
  if (!this.departmentId) {
    message.error('部门ID不能为空')
    return
  }

  this.saving = true
  try {
    const permissionData = {
      dataScope: this.form.dataScope,
      status: this.form.status,
      customDeptIds: this.form.customDeptIds,
      rules: this.permissionRules.filter(rule =>
        rule.field && rule.operator && rule.value
      )
    }

    await saveDepartmentPermissions({
      departmentId: this.departmentId,
      permission: permissionData
    })

    message.success('部门权限配置保存成功')
    this.$emit('saved')
  } catch (error) {
    message.error('保存失败')
  } finally {
    this.saving = false
  }
},

// 重置
handleReset() {
  this.form.dataScope = '1'
  this.form.status = true
  this.form.customDeptIds = []
  this.permissionRules = []
}
}
</script>

<style scoped lang="less">
.department-permission-config {
  .card-title {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .permission-rules {
    .rule-item {
      padding: 12px;
      border: 1px solid #d9d9d9;
      border-radius: 6px;
      margin-bottom: 8px;

      &:hover {
        border-color: #1890ff;
      }
    }
  }

  .action-buttons {
    text-align: right;
    margin-top: 24px;
    padding-top: 16px;
    border-top: 1px solid #f0f0f0;
  }
}
</style>