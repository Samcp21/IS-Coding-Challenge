import { StatisticsCalculator } from './statistics.js';

describe('StatisticsCalculator', () => {
    it('debe calcular las estadísticas correctas para una matriz estándar', () => {
        const matrix = [
            [1, 2],
            [3, 4]
        ];

        const stats = StatisticsCalculator.calculate(matrix);

        expect(stats.max_value).toBe(4);
        expect(stats.min_value).toBe(1);
        expect(stats.total_sum).toBe(10);
        expect(stats.average).toBe(2.5);
        expect(stats.is_diagonal).toBe(false);
    });

    it('debe identificar correctamente una matriz diagonal verdadera', () => {
        const matrix = [
            [5, 0, 0],
            [0, 8, 0],
            [0, 0, 3]
        ];
        const stats = StatisticsCalculator.calculate(matrix);
        expect(stats.is_diagonal).toBe(true);
    });

    it('debe indicar que NO es diagonal si hay números fuera de la diagonal principal', () => {
        const matrix = [
            [0, 0, 5],
            [0, 8, 0],
            [3, 0, 0]
        ];
        const stats = StatisticsCalculator.calculate(matrix);
        expect(stats.is_diagonal).toBe(false);
        expect(stats.min_value).toBe(0);
    });
});