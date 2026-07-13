import type { RouteLocationRaw } from 'vue-router'

import { useRouter } from 'vue-router'

export function useAppNavigation() {
  const router = useRouter()

  function goTo(to: RouteLocationRaw) {
    return router.push(to)
  }

  function goToLogin() {
    return goTo({ name: 'login' })
  }

  function goToRegister() {
    return goTo({ name: 'register' })
  }

  return {
    goTo,
    goToLogin,
    goToRegister,
  }
}
