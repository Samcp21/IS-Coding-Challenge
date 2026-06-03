import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import path from 'path';
import { fileURLToPath } from 'url';
import { MatrixService } from '../application/matrixService.js';
import { logger } from './logger.js';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export class GrpcServerAdapter {
    private server: grpc.Server;

    constructor(
        private readonly matrixService: MatrixService,
        private readonly port: number
    ) {
        this.server = new grpc.Server();
        this.setupRoutes();
    }

    private setupRoutes(): void {
        const PROTO_PATH = path.join(__dirname, '../../proto/matrix.proto');
        const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
            keepCase: true, longs: String, enums: String, defaults: true, oneofs: true
        });

        const matrixProto = (grpc.loadPackageDefinition(packageDefinition) as any).matrix;

        this.server.addService(matrixProto.MatrixService.service, {
            ProcessRotatedMatrix: this.handleProcessMatrix.bind(this)
        });
    }

    private handleProcessMatrix(call: any, callback: any): void {
        try {
            logger.info("\n[gRPC] Request received from Go.");

            const rawRows = call.request.rows || [];
            const jsMatrix: number[][] = rawRows.map((row: any) => row.values || []);

            const result = this.matrixService.processMatrix(jsMatrix);

            logger.info("[gRPC] Statistics successfully calculated.");
            callback(null, result);

        } catch (error: any) {
            logger.error("[gRPC] Error processing request:", error.message);
            callback({
                code: grpc.status.INVALID_ARGUMENT,
                details: error.message || "Error processing matrix"
            });
        }
    }

    public start(): void {

        const bindAddress = `0.0.0.0:${this.port}`;

        this.server.bindAsync(bindAddress, grpc.ServerCredentials.createInsecure(), (error, port) => {
            if (error) {
                logger.error("Failed to start gRPC:", error);
                return;
            }
            logger.info(`🚀 Node.js (gRPC) Microservice listening on port ${port}`);
        });
    }

    public stop(): void {
        this.server.forceShutdown();
    }
}