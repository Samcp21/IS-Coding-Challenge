export interface MatrixStatistics {
    max_value: number;
    min_value: number;
    average: number;
    total_sum: number;
    is_diagonal: boolean;
}

export class StatisticsCalculator {
    public static calculate(matrix: number[][]): MatrixStatistics {
        let max = -Infinity;
        let min = Infinity;
        let sum = 0;
        let count = 0;
        let isDiagonal = true;

        const rows = matrix.length;
        const cols = matrix[0].length;

        for (let i = 0; i < rows; i++) {
            for (let j = 0; j < cols; j++) {
                const val = matrix[i][j];

                if (val > max) max = val;
                if (val < min) min = val;
                sum += val;
                count++;

                if (i !== j && val !== 0) {
                    isDiagonal = false;
                }
            }
        }

        if (rows !== cols) {
            isDiagonal = false;
        }

        return {
            max_value: max === -Infinity ? 0 : max,
            min_value: min === Infinity ? 0 : min,
            average: count === 0 ? 0 : sum / count,
            total_sum: sum,
            is_diagonal: isDiagonal
        };
    }
}