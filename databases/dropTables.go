package databases

import (
	"log"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models/student"
)

//DropTables elimina las tablas de la base de datos.
func DropTables() {
	log.Println("---------------------------------------")
	log.Println("-         Eliminando tablas..		   -")

	//Se obtiene la conexion a la DB.
	db := GetConnectionDB()
	defer db.Close()

	//Se eliminan las tablas.
	db.DropTable(&models.User{})
	db.DropTable(&models.People{})
	db.DropTable(&models.UserStatus{})
	db.DropTable(&models.UserKind{})
	db.DropTable(&student.Country{})
	db.DropTable(&student.State{})
	db.DropTable(&student.Municipality{})
	db.DropTable(&student.Institution{})
	db.DropTable(&student.Representative{})
	db.DropTable(&student.Teacher{})
	db.DropTable(&student.Parish{})
	db.DropTable(&student.Sector{})
	db.DropTable(&student.TypeOfRoad{})
	db.DropTable(&student.Father{})
	db.DropTable(&student.Mother{})
	db.DropTable(&student.Mention{})
	db.DropTable(&student.Section{})
	db.DropTable(&student.StudentCondition{})
	db.DropTable(&student.Student{})
	//------------------------

	log.Println("-                 -                    -")
	log.Println("- ...Finalizo la eliminacion de Tablas -")
	log.Println("----------------------------------------")
}