package datastructures

import (
	"fmt"

)

type AbtractCollectionInterface interface {
	Add(instance AbstractModelInterface)
}

type AbstractCollection struct {
	AbtractCollectionInterface
	models []AbstractModelInterface
}


func (collection *AbstractCollection) Get(id int) AbstractModel {
	for _, item := range collection.models {

		_, ok := item.(AbstractModelInterface) // Kann ich das Item in ein AbstractModel casten?
		if !ok {
			fmt.Println("DAMN")
		} else {
			fmt.Println("YEAH")
			//if model.Id == id {
			//	return model
			//}
		}

	}
	return AbstractModel{}
}


func (collection *AbstractCollection) Add(instance interface{}) {
	if abstraceModelLike, ok := instance.(AbstractModelInterface); ok {
		collection.models = append(collection.models, abstraceModelLike)
	} else {
		fmt.Println("DAMN2")
	}
}


func (c *AbstractCollection) Length() int {
	i := len(c.models)
	return i
}

