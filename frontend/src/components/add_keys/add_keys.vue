<template>
  <div ref="addKeysModal">
    <a-modal
      v-model:visible="props.visible"
      title="输入设置密钥"
      :popup-container="addKeysModal"
      @cancel="handleCancel"
      @before-ok="handleBeforeOk"
    >
      <a-form :model="form" ref="fromRef">
        <a-form-item
          field="name"
          label="账户名"
          :rules="[{ required: true, message: '账户名为必填' }]"
          :validate-trigger="['change', 'input']"
        >
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item
          field="key"
          label="密钥"
          :rules="[{ required: true, message: '密钥为必填' }, {
            match: /^[A-Za-z0-9]+$/, 
            message: '密钥值中含有非法字符',
          },]"
          :validate-trigger="['change', 'input']"
        >
          <a-input v-model="form.key" />
        </a-form-item>
        <a-form-item field="keyType" label="密钥类型">
          <a-select v-model="form.keyType">
            <a-option value="time">基于时间</a-option>
            <a-option value="counter">基于计数器</a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, toRaw } from "vue";
import { useGlobalStore , type KeysListItem} from "@/store/global";
import message from "@arco-design/web-vue/es/message";

const fromRef = ref();
const addKeysModal = ref();
const globalStore = useGlobalStore()
const props = defineProps({
    visible: Boolean,
});
const emits = defineEmits(["cloos"]);
const form = reactive(<KeysListItem>{
    name: "",
    key: "",
    keyType: "time",
});

const handleCancel = () => {
    emits("cloos");
};


const handleBeforeOk = async () => {
    fromRef.value.validate()
    if (form.name && form.key) {
        globalStore.addKeys(toRaw(form));
        form.name = "";
        form.key = "";
        message.success("添加成功")
        emits("cloos");
    }else {
        return false;
    }
};
</script>

<style scoped>
:deep(.arco-btn) {
  border-radius: 18px;
}
</style>
