<script setup lang="ts">
import { onClickOutside } from '@vueuse/core'
import { ref } from 'vue'

import Button from './Button.vue'
import Typography from './Typography.vue'

interface DataOption {
  label: string
  value: string
}

const props = defineProps<{
  label?: string
  options: DataOption[]
}>()

const isOpen = ref<boolean>(false)
const modelValue = defineModel<string | null>()
const selectRef = ref<HTMLElement | null>(null)

onClickOutside(selectRef, () => {
  isOpen.value = false
})

function selectOption(option: string) {
  modelValue.value = option
  isOpen.value = false
}
</script>

<template>
  <div ref="selectRef" class="select">
    <label>
      <Typography variant="body-sm-medium">
        <slot name="label"></slot>
        {{ props.label }}
      </Typography>
    </label>
    <div class="select__wrapper">
      <Button
        variant="select"
        text-align="left"
        icon-pos="right"
        :position="{
          absolute: true,
          top: '50%',
          right: '0.5rem',
          transform: 'translateY(-50%)',
        }"
        :selectActive="isOpen"
        @click="isOpen = !isOpen"
      >
        {{
          props.options.find(option => option.value === modelValue)?.label ||
          modelValue
        }}
        <template #icon>
          <img
            src="@assets/img/icons/chevron-down_minor.svg"
            alt="chevron down icon"
          />
        </template>
      </Button>
      <ul class="select__list" :class="{ 'select__list--active': isOpen }">
        <li
          v-for="sortItem in props.options"
          :key="sortItem.value"
          class="select__item"
          :class="{ 'select__item--active': sortItem.value === modelValue }"
          @click="selectOption(sortItem.value)"
        >
          <Typography variant="body-sm">
            {{ sortItem.label }}
          </Typography>
        </li>
      </ul>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.select {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex: 0 0 25rem;
}

.select__wrapper {
  position: relative;
  width: 100%;
  max-width: 16.75rem;
}

.select__list {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  position: absolute;
  top: calc(100% + 0.25rem);
  z-index: 2;
  left: 0;
  width: 100%;
  border-radius: 0.5rem;
  background: var(--Surface-Default);
  box-shadow:
    0 0 2px 0 rgba(0, 0, 0, 0.2),
    0 2px 10px 0 rgba(0, 0, 0, 0.1);
  padding: 0.5rem;
  transition: all 0.3s ease;
  transform: scale(0);

  &--active {
    transform: scale(1);
  }
}

.select__item {
  padding: 0.62rem 0.5rem;
  border-radius: 0.25rem;
  cursor: pointer;
  transition: background 0.3s ease;
  position: relative;

  &:hover {
    background: var(--Surface-Hovered);
  }

  &--active {
    background: var(--Surface-Selected-Default);
    &::before {
      content: '';
      position: absolute;
      width: 0.1875rem;
      height: 2.5rem;
      top: 0;
      left: -0.45rem;
      border-radius: 0 0.25rem 0.25rem 0;
      background: var(--Text-Interactive-Default);
    }
    &:hover {
      background: var(--Surface-Selected-Default);
    }
  }
}

@media (max-width: 992px) {
  .select {
    flex: 1 1 auto;
    max-width: none;
    width: 100%;
    flex-direction: column;
    align-items: flex-start;
  }

  .select__wrapper {
    max-width: none;
  }
}
</style>
