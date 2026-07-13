<script setup lang="ts">
import type { CSSProperties } from 'vue'

import { computed } from 'vue'

type Variant =
  | 'callback'
  | 'outlined'
  | 'outlined-secondary'
  | 'burger'
  | 'search'
  | 'product-basket'
  | 'control'
  | 'control-square'
  | 'select'
  | 'primary'
  | 'none'

type ControlVariant = 'primary' | 'secondary'
type TextAlign = 'left' | 'center' | 'right' | 'inherit'
type IconPos = 'left' | 'right'
type ContentAlign = 'flex-start' | 'center' | 'flex-end'

interface IconPosition {
  absolute?: boolean
  top?: CSSProperties['top']
  right?: CSSProperties['right']
  bottom?: CSSProperties['bottom']
  left?: CSSProperties['left']
  transform?: CSSProperties['transform']
}

const props = withDefaults(
  defineProps<{
    variant?: Variant
    state?: boolean
    controlVariant?: ControlVariant
    selectActive?: boolean
    textAlign?: TextAlign
    iconPos?: IconPos
    contentAlign?: ContentAlign
    position?: IconPosition
    gap?: CSSProperties['gap']
    fullWidth?: boolean
  }>(),
  {
    variant: 'primary',
    state: false,
    controlVariant: 'primary',
    selectActive: false,
    textAlign: 'inherit',
    iconPos: 'left',
    contentAlign: 'flex-start',
    position: () => ({
      absolute: false,
    }),
    gap: '0.25rem',
    fullWidth: false,
  },
)

const DEFAULT_ICON_POSITION: Required<IconPosition> = {
  absolute: false,
  top: 'auto',
  right: 'auto',
  bottom: 'auto',
  left: 'auto',
  transform: 'translateY(0)',
}

const normalizedIconPosition = computed<Required<IconPosition>>(() => ({
  ...DEFAULT_ICON_POSITION,
  ...props.position,
}))

const isAbsoluteIcon = computed(() => normalizedIconPosition.value.absolute)

const classes = computed(() => [
  'button',
  `button--${props.variant}`,
  { 'button--burger--active': props.state },
  { [`button--control-${props.controlVariant}`]: props.variant === 'control' },
  { 'button--full-width': props.fullWidth },
])

const contentClasses = computed(() => [
  'button__content',
  { [`button__content--${props.textAlign}`]: props.textAlign !== 'inherit' },
])

const innerClasses = computed(() => [
  'button__inner',
  `button__inner--${props.iconPos}`,
  `button__inner--content-${props.contentAlign}`,
  { 'button__inner--icon-absolute': isAbsoluteIcon.value },
])

const iconClasses = computed(() => [
  'button__icon',
  {
    'button__icon--absolute': isAbsoluteIcon.value,
  },
])

const iconTransform = computed(() => {
  const baseTransform = normalizedIconPosition.value.transform

  if (props.variant === 'select' && props.selectActive) {
    return `${baseTransform} rotate(180deg)`
  }

  return `${baseTransform} rotate(0deg)`
})

const iconStyle = computed(() => {
  if (!isAbsoluteIcon.value) {
    return undefined
  }

  return {
    top: normalizedIconPosition.value.top,
    right: normalizedIconPosition.value.right,
    bottom: normalizedIconPosition.value.bottom,
    left: normalizedIconPosition.value.left,
  } as CSSProperties
})
</script>

<template>
  <button type="button" :class="classes">
    <div :class="innerClasses">
      <span v-if="$slots.icon" :class="iconClasses" :style="iconStyle">
        <slot name="icon" />
      </span>
      <span v-if="$slots.default" :class="contentClasses">
        <slot />
      </span>
    </div>
    <span v-if="variant === 'burger'" class="burger-span"></span>
  </button>
</template>

<style lang="scss" scoped>
.button {
  background: none;
  transition: all 0.3s ease;
  position: relative;
  touch-action: manipulation;

  &--full-width {
    width: 100%;
  }

  &--primary {
    padding: 0.75rem 1.5rem;
    border-radius: 0.25rem;
    background: var(--Icon-Highlight);
    box-shadow:
      0 1px 0 0 rgba(0, 0, 0, 0.08),
      0 -1px 0 0 rgba(0, 0, 0, 0.2) inset;
    color: var(--Text-On-Primary);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem; /* 142.857% */

    &:hover {
      border-radius: 0.25rem;
      background: var(--Text-Highlight);
      box-shadow:
        0 1px 0 0 rgba(0, 0, 0, 0.08),
        0 -1px 0 0 rgba(0, 0, 0, 0.2) inset;
    }
  }

  &__inner {
    display: flex;
    align-items: center;
    gap: v-bind(gap);

    &--left {
      flex-direction: row;
    }

    &--right {
      flex-direction: row-reverse;
    }

    &--icon-absolute {
      flex-direction: row;
    }

    &--content-flex-start {
      justify-content: flex-start;
    }

    &--content-center {
      justify-content: center;
    }

    &--content-flex-end {
      justify-content: flex-end;
    }
  }

  &__content {
    display: block;

    &--left {
      text-align: left;
    }

    &--center {
      text-align: center;
    }

    &--right {
      text-align: right;
    }
  }

  &--callback {
    color: var(--Interactive-Default, #2c6ecb);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem; /* 142.857% */
  }
  &--outlined {
    padding: 0.5rem 1rem;
    border-radius: 0.25rem;
    border: 1px solid var(--Icon-Highlight);
    color: var(--Text-Default);
    text-align: center;
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem; /* 142.857% */
    position: relative;
    top: 0;
    left: 0;

    &:active {
      top: 1px;
    }

    &:hover {
      border-radius: 0.25rem;
      border: 1px solid var(--Icon-Highlight);
      background: var(--Icon-Highlight);
      color: #fff;
    }
  }

  &--outlined-secondary {
    padding: 0.5rem 0.75rem;
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Action-Secondary-Default);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
    color: var(--Text-Default);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.9375rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem; /* 133.333% */
    position: relative;
    top: 0;
    left: 0;

    &:active {
      top: 1px;
    }
  }
  &--burger {
    width: 1.25rem;
    height: 1.25rem;
    position: relative;
    z-index: 10;

    & span {
      position: absolute;
      top: 50%;
      left: 0;
      transform: translateY(-50%);
      width: 100%;
      height: 0.125rem;
      background: var(--Icon-Default);
      border-radius: 1rem;
      transition: all 0.3s ease;
    }
    &::after,
    &::before {
      transition: all 0.3s ease;
      content: '';
      position: absolute;
      left: 0;
      width: 100%;
      height: 0.125rem;
      background: var(--Icon-Default);
      border-radius: 1rem;
    }
    &::after {
      top: 0.125rem;
    }
    &::before {
      bottom: 0.125rem;
    }
    &--active {
      & span {
        opacity: 0;
      }
      &::after {
        transform: rotate(45deg);
        top: 50%;
      }
      &::before {
        transform: rotate(-45deg);
        top: 50%;
      }
    }
  }
  &--product-basket {
    padding: 0.31rem 0.5rem 0.31rem 0.75rem;
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Action-Secondary-Default);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
    color: #5c5f62;
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.9375rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem;
    display: flex;
    align-items: center;
    gap: 0.25rem;
    position: relative;
    top: 0;
    left: 0;

    &:active {
      top: 1px;
    }
  }
  &--control {
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;

    &:hover {
      background: var(--Icon-Highlight);
      & :deep(path) {
        fill: #fff;
      }
    }
    &-primary {
      background: #fff;
    }
    &-secondary {
      background: var(--Background-Default);
    }

    &:active {
      position: relative;
      top: 0;
      left: 0;

      &:active {
        top: 1px;
      }
    }
  }
  &--control-square {
    width: 2.25rem;
    height: 2.25rem;
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Action-Secondary-Default);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
  }
  &--select {
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Surface-Default);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
    padding: 0.5rem 0.5rem 0.5rem 0.75rem;
    color: var(--Text-Default);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem;
    width: 100%;
    position: relative;
  }

  &__icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    transition: transform 0.3s ease;
    will-change: transform;
    flex: 0 0 auto;

    & :deep(img),
    & :deep(svg) {
      width: 1.25rem;
      height: 1.25rem;
      display: block;
    }

    &--absolute {
      position: absolute;
      transform: v-bind(iconTransform);
    }
  }
}

@media (max-width: 992px) {
  .button {
    &--primary {
      padding: 0.5rem 1rem;
    }
  }
}
</style>
