package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string
	Edad      int
	Dirección string
	Teléfono  string
	Activo    bool
}

func main() {
	r := gin.Default()

	persona := `{"nombre":"Vincent","apellido":"Conace","edad": 25,"direccion":"independencia","telefono":"23421435","activo":true}`

	var p Persona

	err := json.Unmarshal([]byte(persona), &p)
	if err != nil {
		return
	}

	fmt.Println(p)
	fmt.Printf("%T\n", p)

	per := &Persona{
		Nombre:   "vincent",
		Apellido: "conace",
	}

	jsonData, err := json.Marshal(per)
	if err != nil {
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/persona", func(c *gin.Context) {
		c.JSON(http.StatusOK, persona)
	})

	r.GET("/persona2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": string(jsonData),
		})
	})

	r.Run()
}
