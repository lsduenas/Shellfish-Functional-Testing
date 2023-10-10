package server

import (
	"functional/prey"
	"functional/shark"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark     shark.Shark
	prey      prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/prey
/*
1. Se configura la presa para que sea más veloz que el tiburón, y al simular la caza logra escaparse.
*/
func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		// Importante: los datos de Shark y Tuna provienen del handler, es decir se cargan en memoria mientras permanece en ejecución el programa, en este caso para los tests, su creación se realiza en utils.go al crear el server, en cada método del handler se realiza su correspondiente asignación de atributos 

		// Request
		var reqBody request
		context.ShouldBindJSON(&reqBody)

		// Configure Prey
		tuna := h.prey
		tuna.SetSpeed(reqBody.Speed)

		// Configure White Shark
		var shark_position [2]float64
		shark_position[0] = 20
		shark_position[1] = 20
		h.shark.Configure(shark_position, 30.0)

		// llamando a los metodos de Shark
		err, _ := h.shark.Hunt(h.prey)
		log.Println(err)
		if err != nil {
			if err.Error() == "could not catch it" {
				var response response
				response.Success = false
				context.JSON(http.StatusOK, response)
				return
			} else {
				context.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
}

// PUT: /v1/shark
/*
2. Se configura el tiburón más rápido que la presa, pero se encuentra demasiado lejos y no logra cazarla.
*/
func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		// Request
		var reqBody request
		context.ShouldBindJSON(&reqBody)

		// Configure Prey
		tuna := h.prey
		tuna.SetSpeed(30.0)

		// Configure White Shark
		var shark_position [2]float64
		shark_position[0] = reqBody.XPosition
		shark_position[1] = reqBody.YPosition
		h.shark.Configure(shark_position, reqBody.Speed)

		// llamando a los metodos de Shark
		err, _ := h.shark.Hunt(h.prey)
		log.Println(err)
		if err != nil {
			if err.Error() == "could not catch it" {
				var response response
				response.Success = false
				context.JSON(http.StatusOK, response)
				return
			} else {
				context.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
}

// POST: /v1/simulate
/*
3. El tiburón y la presa se configuran de modo que el tiburón logra cazarla luego de 24 segundos (tener en cuenta el algoritmo que usa el simulador).
*/
func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		// Configure Prey
		tuna := h.prey
		tuna.SetSpeed(16.0)

		// Configure White Shark
		var shark_position [2]float64
		shark_position[0] = 20.0
		shark_position[1] = 20.0
		h.shark.Configure(shark_position, 30.5)

		// llamando a los metodos de Shark
		err, timeToCatch := h.shark.Hunt(h.prey)
		log.Println(err)
		if err != nil {
			if err.Error() == "could not catch it" {
				var response response
				response.Message = "could not catch it"
				response.Success = false
				context.JSON(http.StatusOK, response)
				return
			} else {
				context.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
		var response response
		response.Success = true
		response.Message = "White shark catches prey successfully!"
		response.Time = timeToCatch
		context.JSON(http.StatusOK, response)
	}
}
