<script setup lang="ts">
import { useId } from 'vue'

defineOptions({
  name: 'VChecked',
  inheritAttrs: false,
})

const props = withDefaults(
  defineProps<{
    id?: string
    value?: string | number
    label?: string
    disabled?: boolean
    checked?: boolean
  }>(),
  {
    disabled: false,
    checked: false,
  },
)

const modelValue = defineModel<number[] | undefined>(undefined)

const inputId = props.id || useId()
</script>

<template>
  <label
    class="v-checked"
    :for="inputId"
    :class="{ 'v-checked--disabled': disabled }"
  >
    <input
      :id="inputId"
      v-model="modelValue"
      type="checkbox"
      :value="value"
      class="v-checked__input hidden"
      v-bind="$attrs"
      :checked="checked"
    />

    <div
      class="v-checked__indicator"
      :class="{
        'v-checked__indicator--active':
          value !== undefined && modelValue?.includes(Number(value)),
      }"
    >
      <img src="@assets/img/icons/check.svg" alt="check icon" />
    </div>

    <Typography v-if="label || $slots.default" tag="span" variant="body-sm">
      <slot>{{ label }}</slot>
    </Typography>
  </label>
</template>

<style lang="scss" scoped>
.v-checked {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  &__indicator {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 1.125rem;
    height: 1.125rem;
    border-radius: 0.25rem;
    border: 2px solid var(--Border-Default);
    background: var(--Surface-Default);
    transition: all 0.3s ease;
    &--active {
      background: var(--Action-Primary-Default);
      border-color: var(--Action-Primary-Default);
    }
  }
}

.hidden {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}
</style>
