<!-- @format -->

<script setup lang="ts">
import { computed } from 'vue'

type Tag = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'p' | 'span'
type ColorVars =
  | 'Default'
  | 'On-Interactive'
  | 'Subdued'
  | 'Highlight'
  | 'Interactive-Default'
  | 'Success'
  | 'Disabled'
  | 'Warning'
  | 'Critical-Default'
type Size = 'xl' | 'lg' | 'md' | 'sm' | 'sales' | 'sm-medium'
type Variant = `heading-${Size}` | `body-${Size}`
type TypographyIconPos = 'left' | 'right'
const props = withDefaults(
  defineProps<{
    tag?: Tag
    variant?: Variant
    iconPosition?: TypographyIconPos
    color?: ColorVars
    textAlign?: 'left' | 'center' | 'right'
    fontSize?: string | undefined
    fontWeight?: string | undefined
  }>(),
  {
    tag: 'p',
    variant: 'body-sm',
    iconPosition: 'left',
    color: 'Default',
    textAlign: 'left',
    fontSize: undefined,
    fontWeight: undefined,
  },
)

const computedColor = computed(() => {
  const colorVars: ColorVars[] = [
    'Default',
    'On-Interactive',
    'Subdued',
    'Highlight',
    'Interactive-Default',
    'Success',
    'Disabled',
    'Warning',
    'Critical-Default',
  ]
  return colorVars.includes(props.color)
    ? `var(--Text-${props.color})`
    : props.color
})

const classes = computed(() => [
  'typography',
  `typography--${props.variant}`,
  `typography-position--${props.iconPosition}`,
  `typography-align--${props.textAlign}`,
])

const styles = computed(() => {
  return {
    fontSize: props.fontSize,
    fontWeight: props.fontWeight,
  }
})
</script>

<template>
  <component :is="tag" :class="classes" :style="styles">
    <slot name="icon" />
    <span>
      <slot />
    </span>
  </component>
</template>

<style lang="scss" scoped>
.typography {
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  gap: 0.38rem;
  color: v-bind(computedColor);
  &:deep(s) {
    text-decoration-color: var(--Text-Critical);
  }
  &--heading-xl {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Display';
    font-size: 2.625rem;
    font-style: normal;
    font-weight: 700;
    line-height: 120%; /* 3.15rem */
  }
  &--heading-lg {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Display';
    font-size: 1.75rem;
    font-style: normal;
    font-weight: 600;
    line-height: 2rem; /* 114.286% */
  }

  &--heading-md {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Display';
    font-size: 1.25rem;
    font-style: normal;
    font-weight: 700;
    line-height: 1.3; /* 1.625rem */
  }

  &--heading-sm {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 1rem;
    font-style: normal;
    font-weight: 600;
    line-height: 1.3; /* 1.3rem */
  }

  &--body-sm {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem; /* 142.857% */
  }

  &--body-sm-medium {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem;
  }

  &--body-md {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 1rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.3; /* 150% */
  }

  &--body-sales {
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.9375rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem;
  }

  & b {
    font-weight: 700;
  }
}
.typography-align {
  &--left {
    align-self: flex-start;
    text-align: left;
  }
  &--center {
    align-self: center;
    text-align: center;
  }
  &--right {
    align-self: flex-end;
    text-align: right;
  }
}

@media (max-width: 992px) {
  .typography {
    &--heading-xl {
      font-size: 1.3125rem;
      line-height: 1.75rem;
    }
    &--heading-lg {
      font-size: 1.3125rem;
      line-height: 1.75rem;
    }
    &--heading-md {
      font-size: 0.9375rem;
      line-height: 1.25rem;
    }
    &--heading-sm {
      font-size: 0.9375rem;
      line-height: 1.25rem;
    }
    &--body-md {
      font-size: 0.9375rem;
      line-height: 1.25rem;
    }
  }
}
</style>
