import { MatrixService } from './matrixService.js';

describe('MatrixService', () => {
    let service: MatrixService;

    beforeEach(() => {
        service = new MatrixService();
    });

    it('debe lanzar un error si la matriz está vacía', () => {
        expect(() => {
            service.processMatrix([]);
        }).toThrow("empty or invalid");
    });

    it('debe procesar exitosamente una matriz válida', () => {
        const result = service.processMatrix([[10]]);
        expect(result.max_value).toBe(10);
        expect(result.total_sum).toBe(10);
    });
});