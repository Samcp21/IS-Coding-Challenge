import * as dotenv from 'dotenv';
import { z } from 'zod';
dotenv.config();

const envSchema = z.object({
    GRPC_PORT: z.coerce.number(),
    LOG_LEVEL: z.string().default('info')
})

const _env = envSchema.safeParse(process.env);

if (!_env.success) {
    console.error('Invalid environment variables:', _env.error.format());
    throw new Error('Invalid environment variables');
}

export const config = {
    grpcPort: _env.data.GRPC_PORT,
    logLevel: _env.data.LOG_LEVEL
};