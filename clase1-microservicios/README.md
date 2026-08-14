# Starter

Proyecto monolítico funcional que reciben los alumnos.

## Ejecutar

```bash
go run ./cmd/api
```

Escucha en `localhost:8080`.

## Ejecutar con Docker

Construir imagen:

```bash
docker build -t clase1-starter .
```

Levantar contenedor:

```bash
docker run --rm -p 8080:8080 clase1-starter
```

## Endpoints

```text
GET  /health
GET  /products
GET  /products/{id}
GET  /orders
POST /orders
```

## Ejemplos

Consultar productos:

```bash
curl http://localhost:8080/products
```

Crear orden:

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"productId":"1","quantity":2}'
```

Consultar órdenes:

```bash
curl http://localhost:8080/orders
```

## Exercise

Transformar esta aplicación monolítica en dos microservicios independientes:

```text
              CLIENT
                 |
                 v
        ORDER SERVICE :8082
                 |
                 | HTTP
                 v
       PRODUCT SERVICE :8081
```

Objetivos:

- Separar Products y Orders en procesos distintos.
- Reemplazar la dependencia directa a `ProductRepository` por comunicación HTTP.
- Mantener el comportamiento funcional y el manejo simple de errores.
- Nuevo comportamiento de `Order Service`: cuando una orden se crea exitosamente, debe llamar a `Product Service` con `PUT /products/{id}/stock` y enviar en el body la nueva cantidad para descontar el stock del producto.
