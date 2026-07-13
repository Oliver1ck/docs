<script setup lang="ts">
import { ref } from 'vue'

import Button from '../components/Button.vue'
import Checked from '../components/Checked.vue'
import Input from '../components/Input.vue'
import Typography from '../components/Typography.vue'

defineOptions({
  name: 'LoginPage',
})

const login = ref('')
const password = ref('')
const rememberMe = ref(false)
const isPasswordVisible = ref(false)

function togglePasswordVisibility() {
  isPasswordVisible.value = !isPasswordVisible.value
}
</script>

<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-card__header">
        <Typography tag="h1" variant="heading-md" text-align="center">
          Авторизация
        </Typography>
      </div>

      <form class="login-card__form" @submit.prevent>
        <div class="login-card__fields">
          <Input
            v-model="login"
            name="login"
            placeholder="Логин или почта"
            autocomplete="username"
          />

          <div class="login-card__password">
            <Input
              v-model="password"
              name="password"
              :type="isPasswordVisible ? 'text' : 'password'"
              placeholder="Пароль"
              autocomplete="current-password"
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
                  <svg
                    v-if="!isPasswordVisible"
                    width="16"
                    height="16"
                    viewBox="0 0 16 16"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      fill-rule="evenodd"
                      clip-rule="evenodd"
                      d="M15.946 8.33295L15.9488 8.32486C15.9689 8.26559 15.9835 8.20435 15.9923 8.14194C16.0144 7.98413 15.9985 7.82333 15.946 7.67291C15.3777 6.01746 14.3066 4.58087 12.8822 3.56378C11.4578 2.54672 9.75122 2 8.00097 2C6.25073 2 4.54418 2.54672 3.11975 3.56378C1.69531 4.58087 0.624202 6.01746 0.0559646 7.67291C0.0251419 7.76114 0.00689238 7.85327 0.00158232 7.9462L0.000849898 7.96133C0.000117476 7.97974 -0.0001877 7.99817 0.000117476 8.01663C0.000300582 8.03098 0.000788863 8.04532 0.00158232 8.05966C0.00689238 8.15298 0.025264 8.24506 0.0559646 8.33295C0.631343 9.98459 1.70507 11.417 3.12902 12.4326C4.55297 13.4481 6.25695 13.9968 8.00598 14.0029C9.75323 13.9948 11.455 13.4451 12.8769 12.4297C14.2989 11.4142 15.3711 9.98297 15.946 8.33295ZM4.27984 10.8071C3.28991 10.0996 2.52423 9.12619 2.06921 8.00293C2.52423 6.87967 3.28991 5.90622 4.27984 5.19879C5.36553 4.42294 6.66656 4.00583 8.00097 4.00583C9.33538 4.00583 10.6364 4.42294 11.7221 5.19879C12.712 5.90622 13.4777 6.87967 13.9327 8.00293C13.4777 9.12619 12.712 10.0996 11.7221 10.8071C10.6364 11.5829 9.33538 12 8.00097 12C6.66656 12 5.36553 11.5829 4.27984 10.8071ZM6.88885 9.66586C7.21777 9.88562 7.60443 10.0029 8 10.0029C8.26263 10.0029 8.5227 9.9512 8.76538 9.85068C9.00799 9.75018 9.22851 9.60284 9.41424 9.41714C9.59991 9.23141 9.74725 9.01096 9.84777 8.76831C9.94824 8.52563 10 8.26556 10 8.00293C10 7.60736 9.88269 7.22067 9.66296 6.89178C9.44317 6.5629 9.13079 6.30655 8.76538 6.15518C8.3999 6.00378 7.9978 5.9642 7.6098 6.04135C7.22186 6.11853 6.86547 6.30902 6.58581 6.58871C6.30609 6.86841 6.1156 7.22479 6.03845 7.61276C5.96124 8.0007 6.00085 8.40283 6.15222 8.76831C6.30365 9.13376 6.55993 9.44611 6.88885 9.66586Z"
                      fill="#9FA1AD"
                    />
                  </svg>
                  <svg
                    v-else
                    width="16"
                    height="16"
                    viewBox="0 0 16 16"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      fill-rule="evenodd"
                      clip-rule="evenodd"
                      d="M13.4268 3.98744L14.7071 2.70711C15.0976 2.31658 15.0976 1.68342 14.7071 1.29289C14.3166 0.902369 13.6834 0.902369 13.2929 1.29289L1.29289 13.2929C0.902369 13.6834 0.902369 14.3166 1.29289 14.7071C1.68342 15.0976 2.31658 15.0976 2.70711 14.7071L4.28779 13.1264C5.43872 13.696 6.70998 13.9984 8.00598 14.0029C9.75323 13.9948 11.455 13.4451 12.8769 12.4297C14.2989 11.4142 15.3711 9.98297 15.946 8.33295L15.9488 8.32486C15.9689 8.26559 15.9835 8.20435 15.9923 8.14194C16.0144 7.98413 15.9985 7.82333 15.946 7.67291C15.4527 6.2357 14.5803 4.96345 13.4268 3.98744ZM12.0029 5.41135L9.93123 7.48299C9.97663 7.65163 10 7.8264 10 8.00293C10 8.26556 9.94824 8.52563 9.84777 8.76831C9.74725 9.01096 9.59991 9.23141 9.41424 9.41714C9.22851 9.60284 9.00799 9.75018 8.76538 9.85068C8.5227 9.9512 8.26263 10.0029 8 10.0029C7.82346 10.0029 7.6487 9.97957 7.48006 9.93416L5.80333 11.6109C6.50332 11.8668 7.24681 12 8.00097 12C9.33538 12 10.6364 11.5829 11.7221 10.8071C12.712 10.0996 13.4777 9.12619 13.9327 8.00293C13.5217 6.98828 12.8571 6.09586 12.0029 5.41135Z"
                      fill="#9FA1AD"
                    />
                    <path
                      d="M8.00097 2C8.54163 2 9.07812 2.05217 9.6034 2.15426L7.74678 4.01088C6.50192 4.06034 5.29648 4.47229 4.27984 5.19879C3.28991 5.90622 2.52423 6.87967 2.06921 8.00293C2.22645 8.3911 2.4208 8.76139 2.64868 9.10898L1.21561 10.542C0.724759 9.87106 0.332715 9.12737 0.0559646 8.33295C0.025264 8.24506 0.00689238 8.15298 0.00158232 8.05966C0.000788863 8.04532 0.000300582 8.03098 0.000117476 8.01663C-0.0001877 7.99817 0.000117476 7.97974 0.000849898 7.96133L0.00158232 7.9462C0.00689238 7.85327 0.0251419 7.76114 0.0559646 7.67291C0.624202 6.01746 1.69531 4.58087 3.11975 3.56378C4.54418 2.54672 6.25073 2 8.00097 2Z"
                      fill="#9FA1AD"
                    />
                  </svg>
                </button>
              </template>
            </Input>

            <button type="button" class="login-card__forgot">
              <Typography tag="span" variant="body-sm" color="Subdued">
                Восстановить пароль
              </Typography>
            </button>
          </div>
        </div>

        <Checked v-model="rememberMe"> 
          Запомнить меня
        </Checked>

        <div class="login-card__actions">
          <Button
            variant="primary"
            content-align="center"
            class="login-card__submit"
          >
            Войти
          </Button>

          <Typography
            tag="p"
            variant="body-sm"
            color="Subdued"
            text-align="center"
          >
            Нет аккаунта?
            <Button variant="callback" :to="{ name: 'register' }">
              Зарегистрироваться
            </Button>
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

  &__forgot {
    align-self: flex-start;
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
  }

  &__checkbox {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    cursor: pointer;
  }

  &__checkbox-input {
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

  &__checkbox-box {
    display: inline-flex;
    width: 1.125rem;
    height: 1.125rem;
    flex-shrink: 0;
    border-radius: 0.25rem;
    border: 1px solid var(--Border-Default);
    background: var(--Surface-Default);
    transition: all 0.2s ease;
  }

  &__checkbox-input:checked + &__checkbox-box {
    background: var(--Icon-Highlight);
    border-color: var(--Icon-Highlight);
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
