package main

import (
	"github.com/ellipizle/golang-mvc/config"
	"github.com/ellipizle/golang-mvc/controller"
	"github.com/ellipizle/golang-mvc/pkg/postgresdb"
	postgresRepo "github.com/ellipizle/golang-mvc/platform/postgresdb"
	"github.com/ellipizle/golang-mvc/service"
	"github.com/ellipizle/golang-mvc/templates"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/siredwin/pongorenderer/renderer"
)

var (
	MainRenderer = renderer.Renderer{Debug: true} // use any renderer
)

func main() {
	//connect mongodb
	// db := mongodb.New(config.Config("MONGO_DB"))
	//connect Postgresqldb
	pgdb := postgresdb.New(config.Config("POSTGRES_DB"))
	// initialize repository
	// repo := mongoRepo.New(pdb)
	// initialize repository
	repo := postgresRepo.New(pgdb)
	//initialize service
	svc := service.New(repo)
	//initialize controller
	employeeController := controller.New(svc)
	// initialize new echo
	e := echo.New()
	// e.Renderer = template.New()
	// Files are provided as a slice of strings.

	e.Renderer = templates.NewRenderer("view/*.html", true)

	// Middleware
	// e.Static("/view", "getall_emp.html")
	// e.File("/", "assets")
	e.Static("/", "assets")
	e.Use(m.Logger())  // Logger
	e.Use(m.Recover()) // Recover
	// create a group v1

	v1 := e.Group("/v1")
	//create employee group route
	employee := v1.Group("/employee")
	employee.GET("/home", employeeController.Home)            //added
	employee.GET("/add", employeeController.AddEmp)           //added
	employee.GET("/update/:id", employeeController.UpdateEmp) //added
	employee.GET("/view/:id", employeeController.ViewEmp)     //added
	// employee.GET("/delete/:id", employeeController.DeleteEmp) //added
	////////////////////////////////////////////////////////////////////
	employee.GET("/allemployee", employeeController.GetAllEmployee)
	employee.POST("/addemployee", employeeController.AddEmployee)
	employee.POST("/updates/:id", employeeController.UpdateEmployee)
	employee.GET("/delete/:id", employeeController.DeleteEmployee)
	employee.GET("/singleemployee/:id", employeeController.GetEmployee)
	////////////////////////////////////////////////////////////////////
	e.Logger.Fatal(e.Start(":8080"))
}
