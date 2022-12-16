package model

type Employee struct {
	Id     string  `json:"empid" form:"empid" query:"empid"`
	Name   string  `json:"name" form:"name" query:"name"`
	Email  string  `json:"email" form:"email" query:"email"`
	Salary float64 `json:"salary" form:"salary" query:"salary"`
	Cdate  string  `json:"cdate" form:"cdate" query:"cdate"`
	Udate  string  `json:"udate" form:"udate" query:"udate"`
}

// type Employee struct {
// 	Id     string  `json:"id,omitempty" bson:"_id"`
// 	Name   string  `json:"name,omitempty"`
// 	Email  string  `json:"email,omitempty"`
// 	Salary float64 `json:"salary,omitempty"`
// }
