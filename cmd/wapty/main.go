package main

import (
	"github.com/blogtitle/wapty/pkg/backend"
	"github.com/blogtitle/wapty/pkg/db"
	"github.com/blogtitle/wapty/pkg/intercept"
	"github.com/blogtitle/wapty/pkg/proxy"
	"github.com/blogtitle/wapty/pkg/structures"
)

func main() {
	prdb := make(chan structures.Request)
	dbint := make(chan structures.Request)
	intba := make(chan structures.Request)

	go proxy.Handle(prdb)
	go db.Save(dbint, prdb)
	go intercept.Dispatch(intba, dbint)
	go backend.Interact(intba)

	// TODO(*): detect cancelation, close gracefully
	var termination chan error
	<-termination
}
