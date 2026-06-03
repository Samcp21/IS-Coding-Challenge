import { MatrixStatistics, StatisticsCalculator } from '../domain/statistics.js';

export class MatrixService {
    public processMatrix(matrix: number[][]): MatrixStatistics {
        if (!matrix || matrix.length === 0 || matrix[0].length === 0) {
            throw new Error("The input matrix is empty or invalid.");
        }

        return StatisticsCalculator.calculate(matrix);
    }
}