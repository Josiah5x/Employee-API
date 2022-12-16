package mongodb

import (
	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/pkg/id"
	"github.com/ellipizle/golang-mvc/repository"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongodb struct {
	db *mgo.Database
}

func New(db *mgo.Database) repository.Repository {
	return &Mongodb{db}
}

func (mgo *Mongodb) AddEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.Id = id.GenerateNewUniqueCode()
	err := mgo.db.C("employee").Insert(employee)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (mgo *Mongodb) EditEmployee(employee *model.Employee) (*model.Employee, error) {
	err := mgo.db.C("employee").Update(bson.M{"_id": employee.Id}, employee)
	if err != nil {
		return employee, err
	}
	return employee, nil

}

func (mgo *Mongodb) GetEmployee(id string) (*model.Employee, error) {
	var employee = new(model.Employee)
	err := mgo.db.C("employee").Find(bson.M{"_id": id}).One(&employee)
	if err != nil {
		return employee, err
	}
	return employee, nil
}
func (mgo *Mongodb) DeleteEmployee(id string) error {
	err := mgo.db.C("employee").Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil

}

func (mgo *Mongodb) GetAllEmployee() ([]*model.Employee, error) {
	var employee []*model.Employee
	err := mgo.db.C("employee").Find(bson.M{}).All(&employee)
	if err != nil {
		return employee, err
	}
	return employee, nil
}
