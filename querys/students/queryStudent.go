package students

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStudent Query GraphQL para consultar un estudiante en especifico
var QueryStudent = &graphql.Field{
	Type: graphqlTypes.StudentType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		studentFound := &student.Student{}
		motherFound := &student.Mother{}
		fatherFound := &student.Father{}
		representativeFound := &student.Representative{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStudents {
				err := db.Where("id = ?", idQuery).First(&studentFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando el estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				motherFound.ID = studentFound.MotherID
				fatherFound.ID = studentFound.FatherID
				representativeFound.ID = studentFound.RepresentativeID

				err = db.Where("id = ?", motherFound.ID).First(&motherFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando la madre del estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				err = db.Where("id = ?", fatherFound.ID).First(&fatherFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando el padre del estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				err = db.Where("id = ?", representativeFound.ID).First(&representativeFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando el representant del estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				studentFound.MotherFirstName = motherFound.FirstName
				studentFound.MotherLastName = motherFound.LastName
				studentFound.MotherPhoneNumber = motherFound.PhoneNumber
				studentFound.MotherCI = motherFound.CI

				studentFound.FatherFirstName = fatherFound.FirstName
				studentFound.FatherLastName = fatherFound.LastName
				studentFound.FatherPhoneNumber = fatherFound.PhoneNumber
				studentFound.FatherCI = fatherFound.CI

				studentFound.RepresentativeFirstName = representativeFound.FirstName
				studentFound.RepresentativeLastName = representativeFound.LastName
				studentFound.RepresentativePhoneNumber = representativeFound.PhoneNumber
				studentFound.RepresentativeCI = representativeFound.CI
				studentFound.RepresentativeRelationship = representativeFound.Relationship
				studentFound.RepresentativeAddress = representativeFound.Address
			} else {
				studentFound.Message = "#QueryStudent# : No tienes permisos para leer estudiantes."
				studentFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return studentFound, nil
	},
}
