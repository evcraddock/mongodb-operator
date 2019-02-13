package controller

import (
	"github.com/evcraddock/mongodb-operator/pkg/controller/mondodbbackup"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mondodbbackup.Add)
}
