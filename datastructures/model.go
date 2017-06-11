package datastructures

//import "reflect"

type AbstractModelInterface interface {
	GetId() int
}

type AbstractModel struct {
	AbstractModelInterface
	Id int

}

func (instance AbstractModel) GetId() int{
	return instance.Id
}
