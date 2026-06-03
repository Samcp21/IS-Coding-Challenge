import { MatrixService } from './application/matrixService.js';
import { config } from './config/env.js';
import { GrpcServerAdapter } from './infrastructure/grpcServer.js';

function bootstrap() {
    const matrixService = new MatrixService();
    const grpcServer = new GrpcServerAdapter(matrixService, config.grpcPort);
    grpcServer.start();
}

bootstrap();