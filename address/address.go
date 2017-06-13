package address

import (
	"learning_go/datastructures"
)

/*
Ich brauche eine struct Collection / Instance deren
models properties ursprünglich per construct reingegebenw erden

new AbstractInstance({ id: int, street: string })
 */
type AddressInstance struct {
	datastructures.AbstractModel

	Recipient_additional string
	Street string
	Street_additional string
	Postcode string
	City string
	Company_additional string
	Lat float64
	Lng float64
	Company_label string
	Region string
}

type AddressCollection struct {
	datastructures.AbstractCollection
}


