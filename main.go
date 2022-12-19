package main

import (
	"embed"

	"github.com/davecgh/go-spew/spew"
	"github.com/w-ingsolutions/predmer-json/jsondb"
)

//go:embed json/vrste/*
var jsonDBradovi embed.FS

//go:embed json/radovi/*
var jsonDBpodradovi embed.FS

func main() {
	// app.NewWingRadovi()
	// w := calc.NewWingCal()
	// if cfg.Initial {
	// 	fmt.Println("running initial sync")
	// }
	// in.Init(w.Podesavanja.File)
	// w.APIimportFS(jsonDBpodradovi,jsonDBradovi)
	// jsondb.APIimportFS(jsonDBradovi, jsonDBpodradovi)

	spew.Dump(jsondb.APIimportFS(jsonDBradovi, jsonDBpodradovi))

	//fmt.Println(" PodvrsteRadovaPodvrsteRadova::::::::::: ", w.Radovi.PodvrsteRadova)

	// w.APIpozivIzbornik(w.Radovi.PodvrsteRadova)

}
