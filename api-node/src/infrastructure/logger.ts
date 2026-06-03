import winston from 'winston';
import { config } from '../config/env.js';

const { combine, timestamp, printf, colorize } = winston.format;

const customFormat = printf(({ level, message, timestamp }) => {
    return `[${timestamp}] ${level}: ${message}`;
});

export const logger = winston.createLogger({
    level: config.logLevel,

    format: combine(
        timestamp({ format: 'YYYY-MM-DD HH:mm:ss:ms' }),
        winston.format.splat()
    ),

    transports: [
        new winston.transports.Console({
            format: combine(
                colorize(),
                customFormat
            ),
        }),

    ],
});