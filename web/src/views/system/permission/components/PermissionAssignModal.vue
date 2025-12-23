<template>
  <a-drawer
    :title="title"
    :width="650"
    :open="visible"
    :closable="true"
    @close="handleClose"
    placement="right"
    :destroyOnClose="true"
  >
    <template #extra>
      <a-space>
        <a-button @click="handleClose">取消</a-button>
        <a-button type="primary" @click="handleSubmit(false)" :loading="loading">
          仅保存
        </a-button>
        <a-button type="primary" @click="handleSubmit(true)" :loading="loading">
          保存并关闭
        </a-button>
      </a-space>
    </template>

    <div class="permission-assign-content">
      <a-alert
        message="权限配置说明"
        description="勾选需要分配给角色的权限，支持父子权限联动。点击权限节点可查看数据规则配置。"
        type="info"
        show-icon
        style="margin-bottom: 16px"
      />

      <div class="permission-tree-section">
        <div class="tree-header">
          <h4>权限列表</h4>
          <a-dropdown :trigger="['click']">
            <a-button>树操作 <DownOutlined /></a-button>
            <template #overlay>
              <a-menu @click="handleTreeOperation">
                <a-menu-item key="parent-child">父子关联</a-menu-item>
                <a-menu-item key="cancel-association">取消关联</a-menu-item>
                <a-menu-item key="check-all">全部勾选</a-menu-item>
                <a-menu-item key="cancel-all">取消全选</a-menu-item>
                <a-menu-item key="expand-all">展开所有</a-menu-item>
                <a-menu-item key="collapse-all">合并所有</a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>

        <div class="permission-tree">
          <a-tree
            checkable
            :tree-data="permissionTree"
            :checked-keys="checkedKeys"
            :expanded-keys="expandedKeys"
            :check-strictly="checkStrictly"
            @check="handleCheck"
            @expand="handleExpand"
            @select="handleSelect"
            :selected-keys="selectedKeys"
            :field-names="{ children: 'children', title: 'name', key: 'id' }"
          >
            <template #title="{ name, ruleFlag }">
              <span>
                {{ name }}
                <a-icon
                  v-if="ruleFlag"
                  type="align-left"
                  style="margin-left: 5px; color: red"
                  title="存在数据规则"
                />
              </span>
            </template>
          </a-tree>
        </div>
      </div>
    </div>

    <!-- 数据规则配置弹窗 -->
    <PermissionDataRuleModal ref="dataRuleModal" />
  </a-drawer>
</template>

<script>
import { message } from 'ant-design-vue'
import { DownOutlined } from '@ant-design/icons-vue'
import PermissionDataRuleModal from './PermissionDataRuleModal.vue'
import { getPermissionTree, getRolePermissions, assignRolePermissions } from '@/api/modules/system/permission'

export default {
  name: 'PermissionAssignModal',
  components: {
    PermissionDataRuleModal
  },
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    roleId: {
      type: String,
      default: ''
    },
    roleName: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      loading: false,
      permissionTree: [],
      checkedKeys: [],
      defaultCheckedKeys: [],
      expandedKeys: [],
      selectedKeys: [],
      checkStrictly: true,
      allPermissionIds: []
    }
  },
  computed: {
    title() {
      return this.roleName ? `为角色"${this.roleName}"分配权限` : '权限分配'
    }
  },
  watch: {
    visible(newVisible) {
      if (newVisible) {
        this.loadPermissionTree()
        this.loadRolePermissions()
      }
    },
    roleId() {
      if (this.visible) {
        this.loadRolePermissions()
      }
    }
  },
  mounted() {
    // 初始化逻辑如果需要
  },
  methods: {
async loadPermissionTree() {
  try {
    const res = await getPermissionTree()
    this.permissionTree = res.treeList
    this.allPermissionIds = res.ids
    this.expandedKeys = res.ids // 默认展开所有
  } catch (error) {
    message.error('加载权限树失败')
  }
},

// 加载角色权限
async loadRolePermissions() {
  if (!this.roleId) return

  try {
    const res = await getRolePermissions(this.roleId)
    this.checkedKeys = res.permissionIds || []
    this.defaultCheckedKeys = [...this.checkedKeys]
  } catch (error) {
    message.error('加载角色权限失败')
  }
},

// 权限勾选处理
handleCheck(checked) {
  if (this.checkStrictly) {
    this.checkedKeys = checked
  } else {
    this.checkedKeys = checked.checked || []
  }
},

// 展开/收起处理
handleExpand(expanded) {
  this.expandedKeys = expanded
},

// 选择节点处理
handleSelect(selected) {
  this.selectedKeys = selected
  if (selected.length > 0) {
    // 打开数据规则配置
    this.$refs.dataRuleModal?.show(selected[0], this.roleId)
  }
},

// 树操作处理
handleTreeOperation({ key }) {
  switch (key) {
    case 'parent-child':
      this.checkStrictly = false
      break
    case 'cancel-association':
      this.checkStrictly = true
      break
    case 'check-all':
      this.checkedKeys = [...this.allPermissionIds]
      break
    case 'cancel-all':
      this.checkedKeys = []
      break
    case 'expand-all':
      this.expandedKeys = [...this.allPermissionIds]
      break
    case 'collapse-all':
      this.expandedKeys = []
      break
  }
},

// 提交权限分配
async handleSubmit(closeAfterSave = false) {
  if (!this.roleId) {
    message.error('角色ID不能为空')
    return
  }

  this.loading = true
  try {
    const params = {
      roleId: this.roleId,
      permissionIds: this.checkedKeys.join(','),
      lastPermissionIds: this.defaultCheckedKeys.join(',')
    }

    await assignRolePermissions(params)
    message.success('权限分配成功')

    if (closeAfterSave) {
      this.handleClose()
    } else {
      // 更新默认选中状态
      this.defaultCheckedKeys = [...this.checkedKeys]
    }

    this.$emit('ok')
  } catch (error) {
    message.error('权限分配失败')
  } finally {
    this.loading = false
  }
},

// 关闭抽屉
handleClose() {
  this.$emit('update:visible', false)
  this.resetData()
},

// 重置数据
resetData() {
  this.checkedKeys = []
  this.defaultCheckedKeys = []
  this.expandedKeys = []
  this.selectedKeys = []
  this.checkStrictly = true
}
}
</script>

<style scoped lang="less">
.permission-assign-content {
  .permission-tree-section {
    .tree-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 16px;

      h4 {
        margin: 0;
        color: rgba(0, 0, 0, 0.85);
      }
    }

    .permission-tree {
      max-height: calc(100vh - 300px);
      overflow-y: auto;
      border: 1px solid #d9d9d9;
      border-radius: 6px;
      padding: 16px;
    }
  }
}

:deep(.ant-drawer-body) {
  padding: 24px;
}
</style>