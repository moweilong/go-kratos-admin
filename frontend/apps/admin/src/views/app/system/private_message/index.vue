<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { PrivateMessage } from '#/generated/api/internal_message/service/v1/private_message.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { usePrivateMessageStore } from '#/stores';

import PrivateMessageDrawer from './private-message-drawer.vue';

const privateMessageStore = usePrivateMessageStore();

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'subject',
      label: $t('page.privateMessage.subject'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<PrivateMessage> = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  height: 'auto',
  exportConfig: {},
  pagerConfig: {
    enabled: false,
  },
  rowConfig: {
    isHover: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        return await privateMessageStore.listPrivateMessage(
          true,
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    { title: $t('page.privateMessage.subject'), field: 'subject' },
    {
      title: $t('ui.table.createTime'),
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: $t('ui.table.remark'), field: 'remark' },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 90,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [Drawer, drawerApi] = useVbenDrawer({
  // 连接抽离的组件
  connectedComponent: PrivateMessageDrawer,
});

/* 打开模态窗口 */
function openModal(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });

  drawerApi.open();
}

/* 创建 */
function handleCreate() {
  console.log('创建');
  // openModal(true);
}

/* 编辑 */
function handleEdit(row: any) {
  console.log('编辑', row);
  // openModal(false, row);
}

/* 删除 */
async function handleDelete(row: any) {
  console.log('删除', row);

  try {
    await privateMessageStore.deletePrivateMessage(row.id);

    notification.success({
      message: $t('ui.notification.delete_success'),
    });

    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('ui.notification.delete_failed'),
    });
  }
}

/* 修改状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await privateMessageStore.updatePrivateMessage(row.id, {
      status: row.status,
    });

    notification.success({
      message: $t('ui.notification.update_status_success'),
    });
  } catch {
    notification.error({
      message: $t('ui.notification.update_status_failed'),
    });
  } finally {
    row.pending = false;
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.system.privateMessage')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.privateMessage.button.create') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-switch
          :checked="row.status === 'ON'"
          :loading="row.pending"
          :checked-children="$t('ui.switch.active')"
          :un-checked-children="$t('ui.switch.inactive')"
          @change="
            (checked: any) => handleStatusChanged(row, checked as boolean)
          "
        />
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click="() => handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('page.privateMessage.moduleName'),
            })
          "
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
