# 🚀 Talsory - Microservices Architecture

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-244C5A?style=for-the-badge&logo=grpc&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

## 🏗️ Arquitectura del Sistema

El ecosistema se divide en tres componentes principales y aislados:

1. **API Gateway & Orchestrator (Go / Fiber):** Actúa como el único punto de entrada público. Maneja la autenticación (JWT), el enrutamiento y sirve la documentación dinámica (Swagger).
2. **Motor de Cálculo Matemático (Node.js / TypeScript):** Un microservicio interno de alto rendimiento dedicado exclusivamente a la rotación y cálculo estadístico de matrices
3. **Frontend SPA (Vue 3 / Vite / Nginx):**

### ⚡ Comunicación Interna (gRPC)

Se utilizo una comunicacion **gRPC y Protocol Buffers (HTTP/2)**. Esto reduce drásticamente el peso del payload, acelera la serialización y garantiza un contrato estricto de datos entre los servicios.

---

## 🛠️ Prerrequisitos

Para ejecutar este proyecto en cualquier entorno (Local o Cloud), solo necesitas tener instalado:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## ⚙️ Configuración del Entorno (Secrets)

Por seguridad, los archivos de variables de entorno no están en el control de versiones. Antes de levantar la infraestructura, debes crear dos archivos `.env` basándote en esta configuración.

**1. Crea `api-go/.env` en la carpeta raíz de Go:**

```env
PORT=8080
JWT_SECRET=tu_super_clave_secreta_aqui
API_USER=root
API_PASS=root
NODE_GRPC_ADDRESS=api-node:50051
```

**2. Crea `api-node/.env` en la carpeta raíz de Node:**

```env
GRPC_PORT=50051
LOG_LEVEL=info
```

---

## 🚀 Despliegue (Local y Cloud)

La infraestructura utiliza **Multi-stage builds** en Docker, asegurando que las imágenes de producción contengan únicamente los binarios compilados y dependencias estrictas de producción, reduciendo drásticamente el tamaño y la superficie de ataque.

Para compilar y levantar todo el ecosistema, ejecuta el siguiente comando en la raíz del proyecto:

```bash
docker compose up --build -d
```

### 📍 Rutas de Acceso

Una vez que los contenedores estén en ejecución (estado `Started`), los servicios estarán disponibles en:

- **Frontend Web:** `http://localhost`.
- **Documentación API (Swagger):** `http://localhost:8080/swagger/index.html`

### 📊 Monitoreo (Logs)

Para visualizar en tiempo real la orquestación, las peticiones HTTP de Fiber y el procesamiento gRPC de Winston:

```bash
# Ver los logs de toda la infraestructura entrelazados
docker compose logs -f

# Aislar el comportamiento del motor matemático (Node)
docker compose logs -f api-node

# Aislar el comportamiento del API Gateway (Go)
docker compose logs -f api-go
```

---

## ☁️ Estrategia de Nube (IaaS)

El proyecto está preparado para desplegarse de manera transparente en proveedores Cloud como GCP (Compute Engine) o AWS (EC2).

La red virtual (`bridge`) configurada en el `docker-compose.yml` aísla el tráfico interno. Esto significa que la instancia en la nube solo requiere tener expuestos los puertos **80 (HTTP)** y **8080 (Swagger)**.

---

## 🔌 Endpoints de la API (REST)

El API Gateway en Go expone los siguientes endpoints principales. Toda la documentación detallada (schemas, parámetros y códigos de estado) se encuentra interactiva en `http://localhost:8080/swagger/index.html`.

- **`POST /api/auth/login`**: Generación de token JWT para el acceso a las rutas protegidas.
- **`POST /api/matrix/process`**: Endpoint protegido (requiere Bearer Token). Recibe una matriz bidimensional, orquesta su procesamiento y devuelve el resultado rotado junto con las estadísticas.

---

## 🔄 Flujo de Datos (Arquitectura en Acción)

El ecosistema funciona bajo un modelo de delegación de responsabilidades. El viaje de los datos sigue este flujo exacto:

1. **Autenticación:** El cliente (Frontend en Vue) envía credenciales al API Gateway (Go) y recibe un token JWT temporal.
2. **Petición HTTP:** El cliente envía el arreglo matricial en formato JSON al endpoint protegido de Go, adjuntando su token en la cabecera.
3. **Validación y Puente:** Go actúa como escudo protector. Valida el token JWT y, si es correcto, convierte el JSON a formato binario (_Protocol Buffers_).
4. **Llamada gRPC:** Go abre una conexión interna de alta velocidad (gRPC) y le envía los datos binarios al microservicio de Node.js a través de la red privada de Docker.
5. **Cálculo Matemático:** Node.js recibe la petición gRPC, ejecuta la lógica de dominio puro (rotación de la matriz y cálculo de máximo, mínimo, promedio y suma) y envía el resultado de vuelta por el mismo canal.
6. **Respuesta al Cliente:** Go recibe la respuesta binaria de Node.js, la serializa de vuelta a JSON y responde la petición HTTP original del cliente para que Vue renderice los datos.

**Esquema de Comunicación:**
`Cliente (Vue)` ↔️ _[HTTP/REST]_ ↔️ `API Gateway (Go)` ↔️ _[gRPC/Protobuf]_ ↔️ `Math Engine (Node.js)`

## 🌐 Live Demo

El ecosistema completo se encuentra desplegado en una instancia de Google Cloud Platform (GCP) y es accesible públicamente.

- **💻 Frontend (Vue 3):** [http://34.132.98.78](http://34.132.98.78)

- **📖 API Docs (Swagger Go):** [http://34.132.98.78:8080/swagger/index.html](http://34.132.98.78:8080/swagger/index.html)

### 🧪 ¿Cómo probar el flujo completo?

Para evaluar la orquestación entre el Frontend, el API Gateway (Go) y el motor matemático (Node.js vía gRPC), sigue estos pasos:

1. Ingresa al **Frontend** a través del enlace de arriba.
2. Inicia sesión en el portal utilizando las credenciales maestras:
   - **Usuario:** `root`
   - **Contraseña:** `root`
3. Una vez autenticado, ingresa una matriz de prueba en la interfaz (por ejemplo, una matriz de 3x3 o 4x4).
4. Al enviar la petición, el Gateway validará tu token JWT, enviará los datos en binario (gRPC) al microservicio interno, y recibirás la matriz rotada junto con sus estadísticas matemáticas en tiempo real.
