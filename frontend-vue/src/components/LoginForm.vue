<script setup lang="ts">
import { reactive } from "vue";
import { useAuth } from "../composables/useAuth";

const { login, loading, errorMessage } = useAuth();

const credentials = reactive({
  username: "root",
  password: "root",
});
</script>

<template>
  <section class="card login-card">
    <h2>Iniciar Sesión</h2>
    <p class="subtitle">Ingresa tus credenciales para obtener el JWT.</p>

    <form @submit.prevent="login(credentials)">
      <div class="form-group">
        <label>Usuario</label>
        <input type="text" v-model="credentials.username" required />
      </div>
      <div class="form-group">
        <label>Contraseña</label>
        <input type="password" v-model="credentials.password" required />
      </div>

      <div v-if="errorMessage" class="alert error">{{ errorMessage }}</div>

      <button type="submit" :disabled="loading">
        {{ loading ? "Autenticando..." : "Obtener Token" }}
      </button>
    </form>
  </section>
</template>
