<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'

type Variant =
  | 'primary'
  | 'secondary'
  | 'tertiary'
  | 'outline'
  | 'semibold'
  | 'basket'
  | 'searchItem'
  | 'basket-outline'
  | 'interactive'
  | 'primary-button'
  | 'breadCrumbs'
  | 'title'

// sm - 0.875rem, md - 1rem, lg - 1.125rem
type Size = 'sm' | 'md' | 'lg'
type LinkIconPos = 'left' | 'right'

defineOptions({
  name: 'VLink',
})

const props = withDefaults(
  defineProps<{
    to: string
    variant?: Variant
    size?: Size
    iconPosition?: LinkIconPos
    contentPosition?: 'left' | 'center' | 'right'
    position?: 'left' | 'center' | 'right' | 'fullWidth'
    padding?: string | undefined
  }>(),
  {
    variant: 'primary',
    size: 'md',
    iconPosition: 'left',
    contentPosition: 'left',
    position: 'left',
    padding: undefined,
  },
)

const route = useRouter().currentRoute.value

function filterPathSegments(path: string): string[] {
  return path
    .split('/')
    .filter(item => item.length > 0 && Number.isNaN(Number(item)))
}

const primaryActive = computed(() => {
  if (props.to === '/' && route.path === '/') {
    return true
  }

  const routeSegments = filterPathSegments(route.path)
  const propsSegments = filterPathSegments(props.to)
  const firstPropsSegment = propsSegments[0]

  if (!firstPropsSegment) {
    return false
  }

  return routeSegments.includes(firstPropsSegment)
})

const classes = computed(() => [
  'link',
  `link--${props.variant}`,
  `link--${props.size}`,
  `link--icon-${props.iconPosition}`,
  `link--content-${props.contentPosition}`,
  `link--position-${props.position}`,
  { 'link--primary--active': primaryActive.value },
])
</script>

<template>
  <NuxtLink :to="props.to" :class="classes" :style="{ padding: props.padding }">
    <slot name="icon" />
    <span>
      <slot />
    </span>
  </NuxtLink>
</template>

<style lang="scss" scoped>
.link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.3s ease;
  &--primary {
    color: var(--Text-On-Interactive);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 500;

    &:hover {
      text-decoration: underline;
      text-underline-offset: 0.25rem;
      text-decoration-color: var(--Text-On-Interactive);
    }
    &--active {
      color: var(--Text-Active);
      &:hover {
        text-decoration: none;
      }
    }
  }

  &--secondary {
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 500;
    &:hover {
      text-decoration: underline;
      text-underline-offset: 0.25rem;
      text-decoration-color: inherit;
    }
  }
  &--outline {
    border-radius: 0.25rem;
    border: 1px solid var(--Border-Default);
    background: #fff;
    padding: 0.5rem 1rem;
    color: var(--Text-Default);
    text-align: center;
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 500;

    &:hover {
      background: var(--Action-Secondary-Hovered);
    }
  }
  &--semibold {
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 600;
  }
  &--basket {
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Action-Secondary-Hovered);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
    padding: 0.5rem 0.88rem 0.5rem 0.75rem;
    color: var(--Text-Default);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.9375rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem; /* 133.333% */
  }
  &--basket-outline {
    border-radius: 0.25rem;
    border: 1px solid var(--Button-border-gradient);
    background: var(--Action-Secondary-Default);
    box-shadow: 0 1px 0 0 rgba(0, 0, 0, 0.05);
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 500;
    line-height: 1.25rem; /* 142.857% */
    padding: 0.75rem 1.5rem;
  }
  &--searchItem {
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem;
    padding: 0.5rem 1rem;

    &:deep(img) {
      width: 2.5rem;
      height: 2.5rem;
      object-fit: cover;
      border-radius: 0.1875rem;
      border: 1px solid var(--Border-Neutral-Subdued);
    }
    &:hover {
      background: var(--Action-Secondary-Pressed);
    }
  }
  &--breadCrumbs {
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 400;
    line-height: 1.25rem;
  }
  &--interactive {
    color: var(--Text-Interactive-Default);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-size: 0.875rem;
    font-style: normal;
    font-weight: 400;
    line-height: 1.5rem; /* 171.429% */
  }
  &--primary-button {
    padding: 0.75rem 1.5rem;
    border-radius: 0.25rem;
    background: var(--Icon-Highlight);
    box-shadow:
      0 1px 0 0 rgba(0, 0, 0, 0.08),
      0 -1px 0 0 rgba(0, 0, 0, 0.2) inset;
    color: var(--Text-On-Primary);
    text-align: center;
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

  &--title {
    color: var(--Text-Default, #202223);
    font-feature-settings:
      'liga' off,
      'clig' off;
    font-family: 'SF Pro Text';
    font-style: normal;
    font-weight: 600;
  }
  &--sm {
    font-size: 0.875rem;
    line-height: 1.25rem; /* 142.857% */
  }
  &--md {
    font-size: 1rem;
    line-height: 130%; /* 1.3rem */
  }
}

.link--content {
  &-left {
    justify-content: flex-start;
  }
  &-center {
    justify-content: center;
  }
  &-right {
    justify-content: flex-end;
  }
}

.link--position {
  &-left {
    align-self: flex-start;
  }
  &-center {
    align-self: center;
  }
  &-right {
    align-self: flex-end;
  }

  &-fullWidth {
    align-self: stretch;
    width: 100%;
  }
}

.link--icon {
  &-left {
    flex-direction: row;
  }
  &-right {
    flex-direction: row-reverse;
  }
}

@media (max-width: 992px) {
  .link {
    &--basket {
      padding: 0;
      background: none;
      border: none;
      box-shadow: none;
    }
    &--primary {
      color: var(--Text-Default);
      font-feature-settings:
        'liga' off,
        'clig' off;
      font-family: 'SF Pro Display';
      font-size: 1.25rem;
      font-style: normal;
      font-weight: 600;
      line-height: 1.75rem; /* 140% */
    }
    &--primary-button {
      padding: 0.5rem 1rem;
    }

    &--md {
      font-size: 0.9375rem;
      line-height: 1.25rem;
    }
  }
}
</style>
