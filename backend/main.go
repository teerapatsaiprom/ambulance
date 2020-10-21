package main

import (
	"context"
	"log"

	"github.com/ambu/app/controllers"
	_ "github.com/ambu/app/docs"
	"github.com/ambu/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Cars struct {
	Car []Car
}

type Car struct {
	CarNo string
}

type Staffs struct {
	Staff []Staff
}

type Staff struct {
	StaffEmail    string
	StaffName     string
	StaffPassword string
}

type Statuscars struct {
	Statuscar []Statuscar
}

type Statuscar struct {
	StatusDetail string
}

type Users struct {
	User []User
}

type User struct {
	UserEmail    string
	UserName     string
	UserPassword string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCarController(v1, client)
	controllers.NewStaffController(v1, client)
	controllers.NewStatuscarController(v1, client)
	controllers.NewUserController(v1, client)
	controllers.NewPredicamentController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"tsaiprom.nam@gmail.com", "Teerapat Saiprom", "123456789"},
			User{"siratee@gmail.com", "Siratee Saiprom", "123456789"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetUserEmail(u.UserEmail).
			SetUserName(u.UserName).
			SetUserPassword(u.UserPassword).
			Save(context.Background())
	}

	// Set Statuscars Data
	statuscars := Statuscars{
		Statuscar: []Statuscar{
			Statuscar{"WORKING"},
			Statuscar{"FIXING"},
			Statuscar{"READY"},
			Statuscar{"LOADING"},
			Statuscar{"DISCHARGE"},
		},
	}

	for _, s := range statuscars.Statuscar {

		client.Statuscar.
			Create().
			SetStatusDetail(s.StatusDetail).
			Save(context.Background())
	}

	// Set Staffs Data
	staffs := Staffs{
		Staff: []Staff{
			Staff{"adam@gmail.com", "Adam Tayler", "123456789"},
			Staff{"alice@gmail.com", "Alice Verga", "123456789"},
		},
	}

	for _, st := range staffs.Staff {

		client.Staff.
			Create().
			SetStaffEmail(st.StaffEmail).
			SetStaffName(st.StaffName).
			SetStaffPassword(st.StaffPassword).
			Save(context.Background())
	}

	// Set Cars Data
	cars := Cars{
		Car: []Car{
			Car{"NO1"},
			Car{"NO2"},
			Car{"NO3"},
			Car{"NO4"},
		},
	}

	for _, cn := range cars.Car {

		client.Car.
			Create().
			SetCarNo(cn.CarNo).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
