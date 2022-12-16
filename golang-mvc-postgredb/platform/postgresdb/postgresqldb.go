package mongodb

import (
	"database/sql"
	"time"

	// "github.com/flosch/pongo2/v6"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/pkg/id"
	"github.com/ellipizle/golang-mvc/repository"
	_ "github.com/lib/pq"
	"github.com/siredwin/pongorenderer/renderer"
)

// var db *sql.DB

var (
	paginator = &pagination.Paginator{}
	// data         = pongo2.Context{}
	MainRenderer = renderer.Renderer{Debug: true} // use any renderer
)

type Postgresdb struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Repository {
	return &Postgresdb{db}
}

func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func (pg *Postgresdb) Updateemp(id string) (*model.Employee, error) {
	employee := new(model.Employee)
	data := "SELECT * FROM employees WHERE empid=$1"
	// data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		// , &employee.Cdate, &employee.Udate
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Salary, &employee.Cdate, &employee.Udate)
		if err != nil {
			panic(err)
		}
	}
	return employee, nil
}

func (pg *Postgresdb) Viewemp(id string) (*model.Employee, error) {
	employee := new(model.Employee)
	data := "SELECT * FROM employees WHERE empid=$1"
	// data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		// , &employee.Cdate, &employee.Udate
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Salary, &employee.Cdate, &employee.Udate)
		if err != nil {
			panic(err)
		}
	}
	return employee, nil
}

func (pg *Postgresdb) AddEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.Id = id.GenerateNewUniqueCode()

	employee.Cdate = time.Now().Format("2006-01-02 15:04:05")
	employee.Udate = time.Now().Format("2006-01-02 15:04:05")
	datas := "INSERT INTO employees (empid, name, email, salary, cdate, udate) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := pg.db.Query(datas, employee.Id, employee.Name, employee.Email, employee.Salary, employee.Cdate, employee.Udate)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (pg *Postgresdb) EditEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.Udate = time.Now().Format("2006-01-02 15:04:05")
	data := "UPDATE employees SET name=$2, email=$3, salary=$4, udate=$5 WHERE empid=$1"
	_, err := pg.db.Query(data, employee.Id, employee.Name, employee.Email, employee.Salary, employee.Udate)
	if err != nil {
		return employee, err
	}
	return employee, nil

}

func (pg *Postgresdb) GetEmployee(id string) (*model.Employee, error) {
	employee := new(model.Employee)
	data := "SELECT * FROM employees WHERE empid=$1"
	// data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Cdate, &employee.Udate, &employee.Salary)
		if err != nil {
			panic(err)
		}
	}
	return employee, nil
}
func (pg *Postgresdb) DeleteEmployee(id string) error {
	data := "DELETE FROM employees WHERE empid=$1"
	_, err := pg.db.Query(data, id)
	if err != nil {
		return err
	}
	return nil

}

func (pg *Postgresdb) GetAllEmployee() ([]*model.Employee, error) {
	var employee []*model.Employee
	data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		newEmployee := model.Employee{}
		err = rows.Scan(&newEmployee.Id, &newEmployee.Name, &newEmployee.Email, &newEmployee.Cdate, &newEmployee.Udate, &newEmployee.Salary)
		if err != nil {
			panic(err)
		}
		employee = append(employee, &newEmployee)

	}

	return employee, nil

}
