package utils

import "strings"

//==========================================================
// VALIDACIONES
//==========================================================

// Valida que un texto no esté vacío.
func TextoVacio(texto string) bool {

	return strings.TrimSpace(texto) == ""

}

// Valida que el precio sea mayor a cero.
func PrecioValido(precio float64) bool {

	return precio > 0

}

// Valida que el stock sea positivo.
func StockValido(stock int) bool {

	return stock >= 0

}

// Valida un correo de forma básica.
func EmailValido(email string) bool {

	return strings.Contains(email, "@")

}