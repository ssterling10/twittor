package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ssterling10/twittor/bd"
	"github.com/ssterling10/twittor/models"
)

/* Registro es la función para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	/* r.Bdoy : tipo stream es un objeto de solo lectura, una vez que se leyó se destruye */
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una constraseña de al menos 6 caracteres", 400)
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)

	}

}
