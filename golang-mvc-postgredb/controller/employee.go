package controller

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/pkg/structs"
	"github.com/ellipizle/golang-mvc/service"

	"github.com/flosch/pongo2"
	// "github.com/flosch/pongo2/v6"

	"github.com/labstack/echo/v4"
	"github.com/siredwin/pongorenderer/renderer"
)

var (
	paginator    = &pagination.Paginator{}
	data         = pongo2.Context{}
	MainRenderer = renderer.Renderer{Debug: true} // use any renderer
)

type Controller struct {
	appSvc *service.Service
}

func New(srv *service.Service) *Controller {
	return &Controller{srv}
}

func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func (contrl *Controller) DeleteEmp(c echo.Context) error {
	// dd := " 4c4b0649-5a6e-4a19-a356-9069aed9be06 "
	ids := c.Param("id")
	emps, err := contrl.appSvc.GetEmployee(ids)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "delete_emp.html", emps)
}

func (contrl *Controller) UpdateEmp(c echo.Context) error {
	// dd := "4c4b0649-5a6e-4a19-a356-9069aed9be06"
	ids := c.Param("id")
	fmt.Println(ids)
	emps, err := contrl.appSvc.GetEmployee(ids)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "update_emp.html", emps)
}

func (contrl *Controller) ViewEmp(c echo.Context) error {
	// dd := "4c4b0649-5a6e-4a19-a356-9069aed9be06"
	id := c.Param("id")
	fmt.Println(id)
	emps, err := contrl.appSvc.GetEmployee(id)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "view_emp.html", emps)
}

func (contrl *Controller) AddEmp(c echo.Context) error {
	dd := "Hello"
	return c.Render(http.StatusOK, "add_emp.html", dd)
}

func (contrl *Controller) Home(c echo.Context) error {
	dd := "Hello"
	return c.Render(http.StatusOK, "index.html", dd)
}

var tpl = pongo2.Must(pongo2.FromFile("views/getall_emp.html"))

// GetAllEmployee func gets all exists employee.
func (contrl *Controller) GetAllEmployee(c echo.Context) error {
	emps, err := contrl.appSvc.GetAllEmployee()
	if err != nil {
		return err
	}
	// sets paginator with the current offset (from the url query param)
	postsPerPage := 2
	paginator = pagination.NewPaginator(c.Request(), postsPerPage, len(emps))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)

	//create a new page list that shows up on html
	fmt.Println(idrange)
	var employee []*model.Employee
	for _, num := range idrange {
		//Prevent index out of range errors
		if num <= len(emps)-1 {
			myuser := emps[num]
			employee = append(employee, myuser)
		}
	}

	// set the paginator in context
	// also set the page list in context
	// if you also have more data, set it context
	// data = pongo2.Context{"posts": myusernames}

	return tpl.ExecuteWriter(pongo2.Context{"paginator": paginator, "posts": employee}, c.Response())

	// return c.JSON(http.StatusOK, emps)
	// return c.Render(http.StatusOK, "getall_emp.html", data)
}

// GetEmployee func gets all exists employee. by id
func (contrl *Controller) GetEmployee(c echo.Context) error {
	id := c.Param("id")
	emps, err := contrl.appSvc.GetEmployee(id)
	if err != nil {
		return err
	}
	// return c.JSON(http.StatusOK, emps)
	return c.Render(http.StatusOK, "getsingle_emp.html", emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) AddEmployee(c echo.Context) error {
	newEmp := new(model.Employee)
	if err := c.Bind(newEmp); err != nil {
		return err
	}
	emps, err := contrl.appSvc.AddEmployee(newEmp)
	if err != nil {
		return err
	}
	// fmt.Println(emps)
	// return c.JSON(http.StatusOK, emps)
	return c.Render(http.StatusOK, "message.html", emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) UpdateEmployee(c echo.Context) error {
	r := new(model.Employee)

	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// salary := c.FormValue("salary")

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(r); err != nil {
		return err
	}
	findEm, err := contrl.appSvc.GetEmployee(id)
	if err != nil {
		return err
	}
	structs.Merge(findEm, r)
	emps, err := contrl.appSvc.EditEmployee(findEm)
	if err != nil {
		return err
	}
	// return c.JSON(http.StatusOK, emps)
	return c.Render(http.StatusOK, "message.html", emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	err := contrl.appSvc.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "message.html", id)
}
