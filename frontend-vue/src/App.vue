<script setup lang="ts">
import { useAuth } from "./composables/useAuth";
import LoginForm from "./components/LoginForm.vue";
import MatrixBoard from "./components/MatrixBoard.vue";

const { token, logout } = useAuth();
</script>

<template>
  <main class="container">
    <header>
      <h1>🧮 Orquestador de Matrices</h1>
      <button v-if="token" @click="logout" class="btn-outline">
        Cerrar Sesión
      </button>
    </header>

    <LoginForm v-if="!token" />
    <MatrixBoard v-else />
  </main>
</template>
<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Fira+Code:wght@400;500&display=swap");

:root {
  --bg-app: #0f172a;
  --bg-card: #1e293b;
  --bg-editor: #020617;
  --primary: #6366f1;
  --primary-hover: #4f46e5;
  --text-main: #f8fafc;
  --text-muted: #94a3b8;
  --border-color: #334155;
  --danger-bg: #7f1d1d;
  --danger-text: #fca5a5;
  --success-bg: #064e3b;
  --success-text: #6ee7b7;
}

body {
  font-family:
    "Inter",
    system-ui,
    -apple-system,
    sans-serif;
  background-color: var(--bg-app);
  color: var(--text-main);
  margin: 0;
  padding: 2rem 1rem;
  min-height: 100vh;
  -webkit-font-smoothing: antialiased;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 3rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid var(--border-color);
}

header h1 {
  margin: 0;
  font-size: 1.75rem;
  font-weight: 700;
  letter-spacing: -0.025em;
  background: linear-gradient(to right, #818cf8, #c084fc);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.card {
  background: var(--bg-card);
  padding: 2.5rem;
  border-radius: 12px;
  border: 1px solid var(--border-color);
  box-shadow:
    0 4px 6px -1px rgba(0, 0, 0, 0.1),
    0 2px 4px -2px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.login-card {
  max-width: 420px;
  margin: 4rem auto;
}

.card h2 {
  margin-top: 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.subtitle {
  color: var(--text-muted);
  font-size: 0.95rem;
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.25rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  font-size: 0.9rem;
  color: var(--text-muted);
}

input {
  width: 100%;
  padding: 0.875rem;
  background-color: var(--bg-app);
  border: 1px solid var(--border-color);
  color: var(--text-main);
  border-radius: 8px;
  font-family: "Inter", sans-serif;
  font-size: 1rem;
  box-sizing: border-box;
  transition: all 0.2s ease;
}

input:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
}

textarea {
  width: 100%;
  padding: 1.25rem;
  background-color: var(--bg-editor);
  border: 1px solid var(--border-color);
  color: #38bdf8;
  border-radius: 8px;
  font-family: "Fira Code", monospace;
  font-size: 0.95rem;
  line-height: 1.5;
  box-sizing: border-box;
  resize: vertical;
  transition: border-color 0.2s ease;
}

textarea:focus {
  outline: none;
  border-color: var(--primary);
}

button {
  background: var(--primary);
  color: white;
  border: none;
  padding: 0.875rem 1.5rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  width: 100%;
  transition:
    background-color 0.2s,
    transform 0.1s;
}

button:hover:not(:disabled) {
  background: var(--primary-hover);
}

button:active:not(:disabled) {
  transform: translateY(1px);
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-outline {
  background: transparent;
  border: 1px solid var(--border-color);
  color: var(--text-main);
  width: auto;
  padding: 0.5rem 1rem;
  font-size: 0.9rem;
}

.btn-outline:hover {
  background: var(--bg-card);
  border-color: var(--text-muted);
}

.alert {
  padding: 1rem;
  border-radius: 8px;
  margin-bottom: 1.5rem;
  font-size: 0.9rem;
}

.error {
  background: var(--danger-bg);
  color: var(--danger-text);
  border: 1px solid rgba(248, 113, 113, 0.2);
}

.grid-results {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}

@media (max-width: 768px) {
  .grid-results {
    grid-template-columns: 1fr;
  }
}

.result-box {
  background: var(--bg-app);
  padding: 1.5rem;
  border-radius: 8px;
  border: 1px solid var(--border-color);
}

.result-box h3 {
  margin-top: 0;
  font-size: 1.1rem;
  color: var(--text-main);
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 0.75rem;
  margin-bottom: 1rem;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

li {
  padding: 0.75rem 0;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

li:last-child {
  border-bottom: none;
}

li strong {
  color: var(--text-muted);
  font-weight: 500;
}

pre {
  background: var(--bg-editor);
  color: #10b981;
  padding: 1.25rem;
  border-radius: 8px;
  overflow-x: auto;
  font-family: "Fira Code", monospace;
  font-size: 0.9rem;
  border: 1px solid var(--border-color);
}

.badge-true {
  background: var(--success-bg);
  color: var(--success-text);
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-size: 0.75rem;
  font-weight: 600;
  border: 1px solid rgba(16, 185, 129, 0.2);
}

.badge-false {
  background: var(--danger-bg);
  color: var(--danger-text);
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-size: 0.75rem;
  font-weight: 600;
  border: 1px solid rgba(239, 68, 68, 0.2);
}
</style>
