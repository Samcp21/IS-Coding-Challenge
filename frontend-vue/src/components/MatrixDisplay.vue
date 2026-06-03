<script setup lang="ts">
defineProps<{
  matrix: number[][];
}>();
</script>

<template>
  <div class="matrix-container">
    <div class="matrix-bracket left"></div>
    <div class="matrix-content">
      <div
        v-for="(row, rowIndex) in matrix"
        :key="'row-' + rowIndex"
        class="matrix-row"
      >
        <span
          v-for="(cell, colIndex) in row"
          :key="'cell-' + rowIndex + '-' + colIndex"
          class="matrix-cell"
        >
          {{ Number(cell.toFixed(4)).toString() }}
        </span>
      </div>
    </div>
    <div class="matrix-bracket right"></div>
  </div>
</template>

<style scoped>
.matrix-container {
  display: inline-flex;
  align-items: stretch;
  margin: 1rem 0;
  padding: 0.5rem;
  background-color: var(--bg-editor);
  border-radius: 8px;
  border: 1px solid var(--border-color);
  font-family: "Fira Code", monospace;
  color: #38bdf8;
  overflow-x: auto;
  max-width: 100%;
}

.matrix-bracket {
  width: 12px;
  border: 2px solid var(--text-muted);
}

.matrix-bracket.left {
  border-right: none;
  border-top-left-radius: 4px;
  border-bottom-left-radius: 4px;
  margin-right: 0.5rem;
}

.matrix-bracket.right {
  border-left: none;
  border-top-right-radius: 4px;
  border-bottom-right-radius: 4px;
  margin-left: 0.5rem;
}

.matrix-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.matrix-row {
  display: flex;
  justify-content: center;
}

.matrix-cell {
  width: 3.5rem;
  text-align: center;
  padding: 0.4rem;
  font-size: 0.95rem;
}
</style>
