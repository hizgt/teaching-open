<template>
  <a-modal
    :title="`数据规则配置 - ${permissionName}`"
    :open="visible"
    :width="800"
    @ok="handleSubmit"
    @cancel="handleCancel"
    :confirmLoading="loading"
  >
    <a-form :form="form" layout="vertical">
      <a-form-item label="规则名称">
        <a-input
          v-decorator="['ruleName', { rules: [{ required: true, message: '请输入规则名称' }] }]"
          placeholder="请输入规则名称"
        />
      </a-form-item>

      <a-form-item label="规则条件">
        <a-select
          v-decorator="['ruleColumn']"
          placeholder="选择字段"
          style="width: 200px; margin-right: 8px"
        >
          <a-select-option value="create_by">创建人</a-select-option>
          <a-select-option value="dept_id">部门ID</a-select-option>
          <a-select-option value="org_code">机构编码</a-select-option>
        </a-select>

        <a-select
          v-decorator="['ruleConditions']"
          placeholder="选择条件"
          style="width: 150px; margin-right: 8px"
        >
          <a-select-option value="=">=</a-select-option>
          <a-select-option value="!=">!=</a-select-option>
          <a-select-option value="in">in</a-select-option>
          <a-select-option value="not in">not in</a-select-option>
        </a-select>

        <a-input
          v-decorator="['ruleValue']"
          placeholder="输入值"
          style="width: 200px"
        />
      </a-form-item>

      <a-form-item label="规则描述">
        <a-textarea
          v-decorator="['ruleDescription']"
          placeholder="请输入规则描述"
          :rows="3"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script>
import { message } from 'ant-design-vue'

export default {
  name: 'PermissionDataRuleModal',
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      loading: false,
      permissionId: '',
      roleId: '',
      permissionName: '',
      form: {
        ruleName: '',
        ruleColumn: '',
        ruleConditions: '',
        ruleValue: '',
        ruleDescription: ''
      }
    }
  },
  methods: {
// 显示弹窗
show(permId, rId, permName = '') {
  this.permissionId = permId
  this.roleId = rId
  this.permissionName = permName
  this.$emit('update:visible', true)
  // 这里可以加载现有的数据规则
  this.loadDataRule()
},

// 加载数据规则
async loadDataRule() {
  // 模拟加载数据规则
  // const res = await getDataRule(this.permissionId, this.roleId)
  // this.form = res
},

// 提交
async handleSubmit() {
  this.loading = true
  try {
    // 这里应该调用保存数据规则的API
    // await saveDataRule({ ...this.form, permissionId: this.permissionId, roleId: this.roleId })
    message.success('数据规则保存成功')
    this.$emit('update:visible', false)
    this.$emit('ok')
  } catch (error) {
    message.error('保存失败')
  } finally {
    this.loading = false
  }
},

// 取消
handleCancel() {
  this.$emit('update:visible', false)
}
}
</script>

<style scoped>
:deep(.ant-form-item) {
  margin-bottom: 16px;
}
</style>