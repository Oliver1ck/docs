<script setup lang="ts">
import type { InputHTMLAttributes } from 'vue'

import { useId } from 'vue'

defineOptions({
  name: 'VInput',
  inheritAttrs: false,
})

const props = withDefaults(
  defineProps<{
    id?: string
    type?: InputHTMLAttributes['type']
    label?: string
    error?: string | boolean
    disabled?: boolean
  }>(),
  {
    type: 'text',
    disabled: false,
  },
)
const modelValue = defineModel<string | number>()

const inputId = props.id || useId()
</script>

<template>
  <div
    class="v-input"
    :class="{
      'v-input--disabled': disabled,
      'v-input--error': !!error,
    }"
  >
    <label v-if="label || $slots.label" :for="inputId" class="v-input__label">
      <slot name="label">{{ label }}</slot>
    </label>

    <label class="v-input__wrapper">
      <div v-if="$slots.prepend" class="v-input__icon v-input__icon--prepend">
        <slot name="prepend" />
      </div>
      <input
        :id="inputId"
        v-model="modelValue"
        :type="type"
        :disabled="disabled"
        class="v-input__control"
        v-bind="$attrs"
      />
      <div v-if="$slots.append" class="v-input__icon v-input__icon--append">
        <slot name="append" />
      </div>
    </label>
    <div v-if="error" class="v-input__error-msg">
      {{ typeof error === 'string' ? error : '' }}
    </div>
  </div>
</template>

<style lang="scss" scoped>
.v-input {
  display: flex;
  flex-direction: column;
  gap: 0.375rem;
  width: 100%;

  &__label {
    font-size: 0.875rem;
    font-weight: 500;
    color: var(--Text-Default);
  }

  &__wrapper {
    display: flex;
    align-items: center;
    border-radius: 0.25rem;
    border: 1px solid var(--Text-input-border);
    background: var(--Surface-Default);
    padding: 0.5rem 0.75rem;
  }

  &__control {
    flex: 1;
    width: 100%;
    border: none;
    background: transparent;
    outline: none;
    color: var(--Text-Subdued, #374151);
    font-family: 'SF Pro Text', sans-serif;
    font-size: 0.9375rem;
    padding: 0;

    &::placeholder {
      color: var(--Text-Placeholder, #9ca3af);
    }
  }

  &__icon {
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--Icon-Color, #6b7280);
    flex-shrink: 0;

    &--prepend {
      margin-right: 0.5rem;
    }

    &--append {
      margin-left: 0.5rem;
    }
  }

  &--disabled {
    .v-input__wrapper {
      background-color: var(--Surface-Disabled, #f3f4f6);
      cursor: not-allowed;
    }

    .v-input__control {
      cursor: not-allowed;
      color: var(--Text-Disabled, #9ca3af);
    }
  }

  &--error {
    .v-input__wrapper {
      border-color: var(--Error-Color, #ef4444);
    }
  }

  &__error-msg {
    color: var(--Error-Color, #ef4444);
    font-size: 0.75rem;
  }
}

@media (max-width: 992px) {
  .v-input__control {
    font-size: 1rem;
  }
}
</style>
