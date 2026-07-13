<script setup lang="ts">
import { ref } from 'vue'

import Button from '../components/Button.vue'
import Input from '../components/Input.vue'
import Typography from '../components/Typography.vue'

defineOptions({
  name: 'RegisterPage',
})

const email = ref('')
const username = ref('')
const password = ref('')
const passwordConfirm = ref('')
const isPasswordVisible = ref(false)
const isPasswordConfirmVisible = ref(false)

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value
}

function togglePasswordConfirmVisibility() {
  isPasswordConfirmVisible.value = !isPasswordConfirmVisible.value
}
</script>

<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-card__header">
        <Typography tag="h1" variant="heading-md" text-align="center">
          Регистрация
        </Typography>
      </div>

      <form class="login-card__form" @submit.prevent>
        <div class="login-card__fields">
          <Input
            v-model="email"
            name="email"
            type="email"
            placeholder="Почта"
            autocomplete="email"
          />

          <Input
            v-model="username"
            name="username"
            placeholder="Имя пользователя"
            autocomplete="username"
          />

          <div class="login-card__password">
            <Input
              v-model="password"
              name="password"
              :type="isPasswordVisible ? 'text' : 'password'"
              placeholder="Пароль"
              autocomplete="new-password"
            >
              <template #append>
                <button
                  type="button"
                  class="login-card__eye"
                  :aria-label="
                    isPasswordVisible ? 'Скрыть пароль' : 'Показать пароль'
                  "
                  @click="togglePasswordVisibility"
                >
                  <svg viewBox="0 0 24 24" fill="none" width="20" height="20">
                    <path
                      d="M1 12s4-7 11-7 11 7 11 7-4 7-11 7-11-7-11-7Z"
                      stroke="currentColor"
                      stroke-width="1.5"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                    <circle
                      cx="12"
                      cy="12"
                      r="3"
                      stroke="currentColor"
                      stroke-width="1.5"
                    />
                  </svg>
                </button>
              </template>
            </Input>
          </div>

          <div class="login-card__password">
            <Input
              v-model="passwordConfirm"
              name="password-confirm"
              :type="isPasswordConfirmVisible ? 'text' : 'password'"
              placeholder="Повторите пароль"
              autocomplete="new-password"
            >
              <template #append>
                <button
                  type="button"
                  class="login-card__eye"
                  :aria-label="
                    isPasswordConfirmVisible
                      ? 'Скрыть пароль'
                      : 'Показать пароль'
                  "
                  @click="togglePasswordConfirmVisibility"
                >
                  <svg viewBox="0 0 24 24" fill="none" width="20" height="20">
                    <path
                      d="M1 12s4-7 11-7 11 7 11 7-4 7-11 7-11-7-11-7Z"
                      stroke="currentColor"
                      stroke-width="1.5"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    />
                    <circle
                      cx="12"
                      cy="12"
                      r="3"
                      stroke="currentColor"
                      stroke-width="1.5"
                    />
                  </svg>
                </button>
              </template>
            </Input>
          </div>
        </div>

        <div class="login-card__actions">
          <Button
            variant="primary"
            content-align="center"
            class="login-card__submit"
          >
            Зарегистрироваться
          </Button>

          <Typography
            tag="p"
            variant="body-sm"
            color="Subdued"
            text-align="center"
          >
            Уже есть аккаунт?
            <Button variant="callback" :to="{ name: 'login' }"> Войти </Button>
          </Typography>
        </div>
      </form>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.login-page {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 2rem;
  background: var(--Background-Default);
}

.login-card {
  display: flex;
  flex-direction: column;
  gap: 2.5rem;
  width: 100%;
  max-width: 37.5rem;
  background: var(--Surface-Default);
  box-shadow: 0 12px 32px 0 rgba(50, 73, 85, 0.1);
  border-radius: 1.875rem;
  overflow: hidden;

  &__header {
    display: flex;
    justify-content: center;
    padding: 1.5rem 2rem 0;
  }

  &__form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    padding: 0 2rem 2rem;
  }

  &__fields {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  &__password {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  &__eye {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    padding: 0;
    color: var(--Icon-Default);
    cursor: pointer;
  }

  &__actions {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
  }

  &__submit {
    width: 100%;
    max-width: 15rem;

    :deep(&.button--primary) {
      border-radius: 1.5rem;
    }
  }
}
</style>
