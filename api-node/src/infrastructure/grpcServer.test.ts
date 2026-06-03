import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import path from 'path';
import { GrpcServerAdapter } from './grpcServer.js';
import { MatrixService } from '../application/matrixService.js';

const PROTO_PATH = path.join(__dirname, '../../proto/matrix.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true, longs: String, enums: String, defaults: true, oneofs: true
});
const matrixProto = (grpc.loadPackageDefinition(packageDefinition) as any).matrix;

describe('GrpcServerAdapter (Prueba de Integración)', () => {
    let serverAdapter: GrpcServerAdapter;
    let testClient: any;
    const TEST_PORT = 50052;

    beforeAll((done: jest.DoneCallback) => {
        const service = new MatrixService();
        serverAdapter = new GrpcServerAdapter(service, TEST_PORT);
        serverAdapter.start();

        testClient = new matrixProto.MatrixService(
            `localhost:${TEST_PORT}`,
            grpc.credentials.createInsecure()
        );

        setTimeout(done, 100);
    });

    afterAll(() => {
        serverAdapter.stop();
    });

    it('debe recibir una petición gRPC, procesarla y devolver estadísticas', (done: jest.DoneCallback) => {
        const gRpcPayload = {
            rows: [
                { values: [1, 2] },
                { values: [3, 4] }
            ]
        };

        testClient.ProcessRotatedMatrix(gRpcPayload, (error: any, response: any) => {
            expect(error).toBeNull();

            expect(response).toBeDefined();
            expect(response.total_sum).toBe(10);
            expect(response.max_value).toBe(4);
            expect(response.is_diagonal).toBe(false);

            done();
        });
    });

    it('debe devolver un error gRPC (INVALID_ARGUMENT) si la matriz está vacía', (done: jest.DoneCallback) => {
        const emptyPayload = { rows: [] };

        testClient.ProcessRotatedMatrix(emptyPayload, (error: any, response: any) => {
            expect(error).toBeDefined();
            expect(error.code).toBe(grpc.status.INVALID_ARGUMENT);
            expect(error.details).toContain("empty or invalid");

            done();
        });
    });
});