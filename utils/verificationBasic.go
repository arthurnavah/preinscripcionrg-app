package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/arthurnavah/PreInscripcionRG-API/databases"
	"github.com/arthurnavah/PreInscripcionRG-API/models"
	student "github.com/arthurnavah/PreInscripcionRG-API/models/students"
)

//VerificationBasic Hace las verificaciones basicas para el funcionamiento del sistema.
func VerificationBasic() {
	//Comprobacion de Tablas.
	tablesExist := TablesIsExist(
		models.People{},
		models.UserStatus{},
		models.UserKind{},
		models.User{},
		student.Country{},
		student.State{},
		student.Municipality{},
		student.Institution{},
		student.Representative{},
		student.Teacher{},
		student.Parish{},
		student.Sector{},
		student.TypeOfRoad{},
		student.Father{},
		student.Mother{},
		student.Mention{},
		student.Section{},
		student.StudentCondition{},
		student.Student{},
	)
	//-----------------------

	if tablesExist != true {
		var t string
		fmt.Println("Se detecto la ausencia de algunas Tablas necesarias en la Base de Datos, ¿Desea crearlas? (Yes/No)")
		fmt.Scan(&t)
		t = strings.ToUpper(t)
		if t == "YES" {
			databases.CreateTables()
		} else {
			fmt.Println("No se puede continuar sin las tablas necesarias")
			os.Exit(0)
		}

	}
}

//TablesIsExist Comprueba la existencia de las tablas/estructuras que se le pasen como parametros, Devuelve un bool.
func TablesIsExist(modelss ...interface{}) bool {
	allExist := true

	//Se obtiene la conexion a la DB.
	db := databases.GetConnectionDB()
	defer db.Close()

	//Sobreescribe el slice de parametros con un booleano si existe o no.
	for i, v := range modelss {
		existTable := db.HasTable(v)

		modelss[i] = existTable
	}

	//Se recorre el slice sobreescrito, Si se encuentra un valor false significa que una tabla no existe.
	for _, v := range modelss {
		if v == false {
			allExist = false
		}
	}

	//Se devuelve la variable allExist, Si dentro de el slice de parametros habia una tabla que no existia, allExist
	// es false.
	return allExist
}
