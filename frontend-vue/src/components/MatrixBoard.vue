<script setup lang="ts">
import { ref } from "vue";
import { useMatrix } from "../composables/useMatrix";
import MatrixDisplay from "./MatrixDisplay.vue";

const { processMatrix, results, loading, errorMessage } = useMatrix();

const matrixInput = ref("[\n  [5, 0, 0],\n  [0, 8, 0],\n  [0, 0, 3]\n]");
</script>

<template>
  <section class="dashboard">
    <div class="card input-section">
      <h2>Ingresar Matriz</h2>
      <p class="subtitle">Escribe la matriz en formato JSON array.</p>

      <textarea v-model="matrixInput" rows="8"></textarea>

      <div v-if="errorMessage" class="alert error">{{ errorMessage }}</div>

      <button @click="processMatrix(matrixInput)" :disabled="loading">
        {{ loading ? "Procesando (Go ⚡ Node.js)..." : "Procesar Matriz" }}
      </button>
    </div>

    <div class="card result-section" v-if="results">
      <h2>Resultados</h2>
      <div class="grid-results">
        <div class="result-box">
          <h3>📊 Estadísticas (Node.js)</h3>
          <ul>
            <li>
              <strong>Suma Total:</strong>
              {{ results.statistics_node.total_sum }}
            </li>
            <li>
              <strong>Promedio:</strong>
              {{ results.statistics_node.average.toFixed(2) }}
            </li>
            <li>
              <strong>Valor Máximo:</strong>
              {{ results.statistics_node.max_value }}
            </li>
            <li>
              <strong>Valor Mínimo:</strong>
              {{ results.statistics_node.min_value }}
            </li>
            <li>
              <strong>Es Diagonal:</strong>
              <span
                :class="
                  results.statistics_node.is_diagonal
                    ? 'badge-true'
                    : 'badge-false'
                "
              >
                {{ results.statistics_node.is_diagonal ? "Sí" : "No" }}
              </span>
            </li>
          </ul>
        </div>

        <div class="result-box">
          <h3>📐 Factorización QR (Go)</h3>

          <div class="matrices-wrapper">
            <div class="matrix-section">
              <h4>Matriz Q:</h4>
              <MatrixDisplay :matrix="results.factorization_qr.Q" />
            </div>

            <div class="matrix-section">
              <h4>Matriz R:</h4>
              <MatrixDisplay :matrix="results.factorization_qr.R" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.matrices-wrapper {
  display: flex;
  gap: 2rem;
  flex-direction: column;
}

.matrix-section h4 {
  margin-top: 0;
  margin-bottom: 0.5rem;
  font-size: 1rem;
  color: var(--text-muted);
}

@media (min-width: 768px) {
  .matrices-wrapper {
    flex-direction: row;
    align-items: flex-start;
  }

  .matrix-section {
    flex: 1;
  }
}
</style>
