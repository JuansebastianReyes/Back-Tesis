package user_service_test

import (
	m "main/models"
	userService "main/services/user.service"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {

	uid := primitive.NewObjectID()
	userId = uid.Hex()

	user := m.User{
		ID:       uid,
		Name:     "Juan",
		Email:    "juan.reyes@correo.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {
		t.Error(err)
		t.Error("La Prueba de persisitencia de usuario a fallado")
		t.Fail()
	} else {
		t.Log("La Prueba de persisitencia de usuario a pasado")
	}

}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("Se ha presentado un error en la consulta de usuarios")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("No se han encontrado usuarios")
		t.Fail()
	} else {
		t.Log("La pureba a finalizado con exito")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Juan Reyes",
		Email: "juan.reyes@gmail.com",
	}

	err := userService.Update(user, userId)

	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
	} else {
		t.Log("La prueba de actualización se realizo con exito ")
	}

}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)

	if err != nil {
		t.Error("Error al tratar de eliminar el usuario")
		t.Fail()
	} else {
		t.Log("La prueba de eliminación se realizo con exito ")
	}
}
