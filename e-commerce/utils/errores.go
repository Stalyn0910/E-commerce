package utils

import "errors"

//==========================================================
// ERRORES PERSONALIZADOS
//==========================================================

var (

	ErrUsuarioExiste = errors.New("el usuario ya existe")

	ErrUsuarioNoExiste = errors.New("el usuario no existe")

	ErrProductoExiste = errors.New("el producto ya existe")

	ErrProductoNoExiste = errors.New("el producto no existe")

	ErrStockInsuficiente = errors.New("stock insuficiente")

	ErrCantidadInvalida = errors.New("cantidad inválida")

	ErrPagoInvalido = errors.New("el pago no es válido")

)