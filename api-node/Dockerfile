FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

RUN npm run build

FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install --omit=dev
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/proto ./proto

ENV GRPC_PORT=50051
ENV LOG_LEVEL=info

EXPOSE 50051
CMD ["npm", "run", "start:prod"]