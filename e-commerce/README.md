# Sistema de Gestión E-Commerce - API REST

## Información del proyecto

**Nombre:** Sistema de Gestión E-Commerce  
**Autor:** Stalyn Guacales  
**Lenguaje:** Go  
**Arquitectura:** API REST  
**Formato de intercambio:** JSON  
**Fecha:** 2026  

---

## 1. Descripción del proyecto

El Sistema de Gestión E-Commerce es una aplicación web desarrollada en
lenguaje Go que permite administrar las principales operaciones de una
plataforma de comercio electrónico.

El sistema proporciona servicios web mediante una API REST para gestionar
usuarios, productos, carritos de compras, pedidos y pagos.

La aplicación fue desarrollada aplicando conceptos de Programación
Orientada a Objetos, estructuras de datos, interfaces, manejo de errores,
serialización JSON y programación concurrente.

---

## 2. Objetivo general

Desarrollar un sistema de gestión E-Commerce mediante una API REST
implementada en Go que permita administrar usuarios, productos, carritos,
pedidos y pagos, aplicando principios de programación orientada a objetos,
manejo de datos mediante JSON y mecanismos de concurrencia.

---

## 3. Objetivos específicos

- Implementar una API REST para el sistema E-Commerce.
- Desarrollar servicios web para las principales funcionalidades.
- Gestionar usuarios y productos.
- Implementar el carrito de compras.
- Gestionar pedidos.
- Implementar el registro y validación de pagos.
- Utilizar JSON para el intercambio de información.
- Aplicar conceptos de programación orientada a objetos.
- Implementar mecanismos de concurrencia.
- Realizar pruebas unitarias, de integración y aceptación.
- Comprobar el funcionamiento del sistema mediante pruebas concurrentes.

---

## 4. Funcionalidades principales

El sistema permite:

- Registrar usuarios.
- Consultar usuarios.
- Actualizar usuarios.
- Eliminar usuarios.
- Registrar productos.
- Consultar productos.
- Actualizar productos.
- Eliminar productos.
- Agregar productos al carrito.
- Calcular el total del carrito.
- Crear pedidos.
- Consultar pedidos.
- Registrar pagos.
- Validar pagos.
- Generar reportes.

---

## 5. Servicios web

El sistema implementa diferentes servicios web mediante endpoints REST.

### Usuarios

- Crear usuario.
- Consultar usuarios.
- Actualizar usuario.
- Eliminar usuario.

### Productos

- Crear producto.
- Consultar productos.
- Actualizar producto.
- Eliminar producto.

### Carrito

- Agregar productos al carrito.
- Consultar carrito.

### Pedidos

- Crear pedido.
- Consultar pedidos.

### Pagos

- Registrar pago.
- Validar pago.

---

## 6. Tecnologías utilizadas

- Go
- API REST
- JSON
- HTTP
- Goroutines
- WaitGroup
- Mutex / RWMutex
- Go testing
- Race Detector
- Git
- GitHub

---

## 7. Arquitectura del proyecto

El proyecto se encuentra organizado mediante diferentes paquetes:

```text
E-COMMERCE/
│
├── data/
│
├── handlers/
│
├── interfaces/
│
├── models/
│
├── services/
│
├── utils/
│
├── .gitattributes
│
├── Diagrama_de_clases.png
│
├── go.mod
│
├── main.go
│
└── README.md

Models

Contiene las estructuras principales del sistema:

Usuario
Producto
CarritoCompras
DetalleCarrito
Pedido
DetallePedido
Pago

Services
Contiene la lógica de negocio:

UsuarioService
ProductoService
CarritoService
PedidoService
PagoService
ReporteService

Handlers
Contiene los manejadores HTTP responsables de recibir las solicitudes
de los clientes y generar las respuestas de la API.

Interfaces

Define los contratos que utilizan los diferentes componentes del sistema.

Utils

Contiene constantes, errores y funciones de validación.

Data

Contiene las estructuras utilizadas para almacenar información durante
la ejecución del sistema.