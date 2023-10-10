package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
1. Se configura la presa para que sea más veloz que el tiburón, y al simular la caza logra escaparse.
*/
func TestPreyIsFasterThanSharkButPrey(t *testing.T) {
	// Arrange
	// crear el Server y definir las Rutas
	r := createServer()
	responseExpected := `{
		"success": false
	}`
	// crear Request del tipo PUT y Response para obtener el resultado
	requestExpected, responseObtained := createRequestTest(http.MethodPut, "/v1/prey", `{
        "speed": 34.0
    }`)

	// Act
	// Ejecutando el handler
	r.ServeHTTP(responseObtained, requestExpected)

	// Assert
	assert.Equal(t, http.StatusOK, responseObtained.Code)
	assert.JSONEq(t, responseExpected, responseObtained.Body.String())
}

/*
2. Se configura el tiburón más rápido que la presa, pero se encuentra demasiado lejos y no logra cazarla.
*/
func TestSharkIsFasterThanPreyButCanNotHuntIt(t *testing.T) {
	// Arrange
	// crear el Server y definir las Rutas
	r := createServer()
	responseExpected := `{
		"success": false
	}`
	// crear Request del tipo PUT y Response para obtener el resultado
	requestExpected, responseObtained := createRequestTest(http.MethodPut, "/v1/shark", `{
        "x_position": 20.0, "y_position":20.0, "speed": 54.0,
    }`)

	// Act
	// Ejecutando el handler
	r.ServeHTTP(responseObtained, requestExpected)

	// Assert
	assert.Equal(t, http.StatusOK, responseObtained.Code)
	assert.JSONEq(t, responseExpected, responseObtained.Body.String())
}

/*
3. El tiburón y la presa se configuran de modo que el tiburón logra cazarla luego de 24 segundos (tener en cuenta el algoritmo que usa el simulador).
*/
func TestSharkHuntsPreySuccessfully(t *testing.T) {
	// Arrange
	// crear el Server y definir las Rutas
	r := createServer()
	responseExpected := `{
		"success": true,
		"message": "White shark catches prey successfully!",
		"time": 1.0
	}`
	// crear Request del tipo POST y Response para obtener el resultado
	requestExpected, responseObtained := createRequestTest(http.MethodPost, "/v1/simulate", "")

	// Act
	// Ejecutando el handler
	r.ServeHTTP(responseObtained, requestExpected)

	// Assert
	assert.Equal(t, http.StatusOK, responseObtained.Code)
	assert.JSONEq(t, responseExpected, responseObtained.Body.String())
}

/*
4. Testear, para todos los endpoints de configuración, casos donde los tipos de los campos no son esperados.
*/