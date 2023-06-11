package routers

import (
	"encoding/json"
	"strconv"

	"github.com/cmorrue/discok/bd"
	"github.com/cmorrue/discok/models"
)

func InsertProduct(body string, User string) (int, string) {

	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (Title) del Producto"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InserProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductId: " + strconv.Itoa(int(result)) + "}"
}

func UpdateProduct(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del producto " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Update OK"

}
