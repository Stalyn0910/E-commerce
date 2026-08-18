package main

import (
	"net/http"
	"fmt"

	"e-commerce/handlers"
	"e-commerce/services"
)

//==========================================================
//
// SISTEMA DE GESTIÓN E-COMMERCE
//
// API REST
//
// Autor:
// Stalyn Guacales
//
//==========================================================

func main() {

	usuarioService := services.NuevoUsuarioService()

    productoService := services.NuevoProductoService()

    carritoService := services.NuevoCarritoService(1)

    pedidoService := services.NuevoPedidoService()

    pagoService := services.NuevoPagoService()

	usuarioHandler := handlers.NuevoUsuarioHandler(
	    usuarioService,
    )

    productoHandler := handlers.NuevoProductoHandler(
	    productoService,
    )

    carritoHandler := handlers.NuevoCarritoHandler(
	    carritoService,
	    productoService,
    )

    pedidoHandler := handlers.NuevoPedidoHandler(
	    pedidoService,
	    usuarioService,
    )

    pagoHandler := handlers.NuevoPagoHandler(
	    pagoService,
    )

	http.HandleFunc(
	    "/",
	    func(w http.ResponseWriter, r *http.Request) {

		    w.Header().Set(
			    "Content-Type",
			    "text/plain; charset=utf-8",
		    )

		    fmt.Fprintln(
			    w,
			    "Sistema de Gestión E-Commerce - API REST",
		    )

		    fmt.Fprintln(
			    w,
			    "Servidor funcionando correctamente.",
		    )
	    },
    )

	http.HandleFunc(
		"/api/usuarios",
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodPost {
				usuarioHandler.Crear(w, r)
				return
			}

			if r.Method == http.MethodGet {
				usuarioHandler.Listar(w, r)
				return
			}

			http.Error(
				w,
				"Método no permitido",
				http.StatusMethodNotAllowed,
			)
		},
	)

	http.HandleFunc(
		"/api/productos",
		func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodPost {
				productoHandler.Crear(w, r)
				return
			}

			if r.Method == http.MethodGet {
				productoHandler.Listar(w, r)
				return
			}

			http.Error(
				w,
				"Método no permitido",
				http.StatusMethodNotAllowed,
			)
		},
	)

	        http.HandleFunc(
	            "/api/carrito/productos",
	            func(w http.ResponseWriter, r *http.Request) {

		            carritoHandler.AgregarProducto(w, r)

	    },
    )

            http.HandleFunc(
	        "/api/carrito",
	        func(w http.ResponseWriter, r *http.Request) {

		        carritoHandler.ObtenerCarrito(w, r)

	    },
    )

            http.HandleFunc(
	        "/api/pedidos",
	        func(w http.ResponseWriter, r *http.Request) {

		        pedidoHandler.Crear(
			        w,
			        r,
		        )

	    },
    )

            http.HandleFunc(
	        "/api/pagos",
	        func(w http.ResponseWriter, r *http.Request) {

		        pagoHandler.Registrar(
			        w,
			        r,
		        )

	    },
    )

	//==========================================================
// SERVIDOR HTTP
//==========================================================



	fmt.Println("==========================================")
	fmt.Println(" API REST - SISTEMA E-COMMERCE")
	fmt.Println("==========================================")
	fmt.Println("Servidor iniciado en:")
	fmt.Println("http://localhost:8080")
	fmt.Println()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}

}
