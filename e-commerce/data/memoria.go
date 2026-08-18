package data

import "e-commerce/models"

//==========================================================
// MEMORIA DEL SISTEMA
//
// Este archivo simula una base de datos utilizando
// estructuras en memoria.
//
// Todas las entidades del sistema se almacenarán aquí
// mientras el programa esté en ejecución.
//
//==========================================================

// Usuarios registrados.
var Usuarios = make(map[int]*models.Usuario)

// Productos registrados.
var Productos = make(map[int]*models.Producto)

// Pedidos registrados.
var Pedidos = make(map[int]*models.Pedido)

// Pagos registrados.
var Pagos = make(map[int]*models.Pago)