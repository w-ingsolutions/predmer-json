package jsondb

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

func APIimportFS(jsonDBradovi, jsonDBpodradovi embed.FS) (radovi *WingVrstaRadova) {
	radovi = &WingVrstaRadova{
		Id:    0,
		Naziv: "Radovi",
		Opis:  "Radovi predmer",
		// PodvrsteRadova: []*WingVrstaRadova,
	}
	radoviFolder, err := jsonDBradovi.ReadDir(filepath.Join("json", "radovi"))
	if err != nil {
	}
	for _, podRadoviRaw := range radoviFolder {
		podvrstaRadova := &WingVrstaRadova{}
		pod, err := jsonDBradovi.ReadFile(filepath.Join("json", "radovi", podRadoviRaw.Name()))
		if err != nil {
			fmt.Println("Err", err)
		}
		jsonErr := json.Unmarshal(pod, &podvrstaRadova)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		// radovi.PodvrsteRadova[strconv.Itoa(podvrstaRadova.Id)] = podvrstaRadova
	}

	podradoviFolder, err := jsonDBpodradovi.ReadDir(filepath.Join("json", "podradovi"))

	for _, podradoviFolderRaw := range podradoviFolder {
		//vrstaRadova := &model.WingVrstaRadova{
		//	Id:             0,
		//	Naziv:          apiFolderRaw.Name(),
		//	PodvrsteRadova: make(map[string]*model.WingVrstaRadova),
		//}
		// radovi.PodvrsteRadova[podradoviFolderRaw.Name()].PodvrsteRadova = make(map[string]*WingVrstaRadova)
		apiFolder, err := jsonDBpodradovi.ReadDir(filepath.Join("json", "podradovi", podradoviFolderRaw.Name()))
		if err != nil {
		}
		fmt.Println("podradoviFolderRaw.Name: ", podradoviFolderRaw.Name())

		for _, podRadoviRaw := range apiFolder {
			podvrstaRadova := &WingVrstaRadova{}
			pod, err := jsonDBpodradovi.ReadFile(filepath.Join("json", "podradovi", podradoviFolderRaw.Name(), podRadoviRaw.Name()))
			if err != nil {
				fmt.Println("Err", err)
			}
			jsonErr := json.Unmarshal(pod, &podvrstaRadova)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			//fmt.Println("pod: ", pod)
			// radovi.PodvrsteRadova[podradoviFolderRaw.Name()].PodvrsteRadova[strconv.Itoa(podvrstaRadova.Id)] = podvrstaRadova
		}
	}
	return
}
