package jsondb

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

func APIimportFS(jsonDBradovi, jsonDBpodradovi embed.FS) (radovi *WingVrstaRadova) {
	radovi = &WingVrstaRadova{
		Id:             0,
		Naziv:          "Radovi",
		Opis:           "Radovi predmer",
		PodvrsteRadova: podvrsteRadova,
	}
	// vrsteFolder, err := jsonDBradovi.ReadDir(filepath.Join("json", "vrste"))
	// if err != nil {
	// }
	// for _, podRadoviRaw := range vrsteFolder {
	// 	podvrstaRadova := WingVrstaRadova{}
	// 	pod, err := jsonDBradovi.ReadFile(filepath.Join("json", "vrste", podRadoviRaw.Name()))
	// 	if err != nil {
	// 		fmt.Println("Err", err)
	// 	}
	// 	jsonErr := json.Unmarshal(pod, &podvrstaRadova)
	// 	if jsonErr != nil {
	// 		log.Fatal(jsonErr)
	// 	}
	// 	radovi.PodvrsteRadova = append(radovi.PodvrsteRadova, podvrstaRadova)
	// }

	radoviFolder, err := jsonDBpodradovi.ReadDir(filepath.Join("json", "radovi"))
	if err != nil {
		fmt.Println("Err", err)
	}
	for _, podradoviFolderRaw := range radoviFolder {
		apiFolder, err := jsonDBpodradovi.ReadDir(filepath.Join("json", "radovi", podradoviFolderRaw.Name()))
		if err != nil {
		}
		fmt.Println("podradoviFolderRaw.Name: ", podradoviFolderRaw.Name())

		folder, err := strconv.Atoi(podradoviFolderRaw.Name())

		for _, radoviRaw := range radovi.PodvrsteRadova {
			fmt.Println("folder.Id : ", folder)

			// fmt.Println("radoviRaw.Id : ", radoviRaw.Id)

			if radoviRaw.Id == folder {
				fmt.Println("radoviRaw.Id : ", radoviRaw.Id)
				for i, podRadoviRaw := range apiFolder {
					podvrstaRadova := WingVrstaRadova{}
					pod, err := jsonDBpodradovi.ReadFile(filepath.Join("json", "radovi", podradoviFolderRaw.Name(), podRadoviRaw.Name()))
					if err != nil {
						fmt.Println("Err", err)
					}
					jsonErr := json.Unmarshal(pod, &podvrstaRadova)
					if jsonErr != nil {
						log.Fatal(jsonErr)
					}
					// fmt.Println("pod: ", podvrstaRadova)
					// radovi.PodvrsteRadova[podradoviFolderRaw.Name()].PodvrsteRadova[strconv.Itoa(podvrstaRadova.Id)] = podvrstaRadova
					// fmt.Println("pod: ", podradoviFolderRaw.Name())
					radovi.PodvrsteRadova[i].PodvrsteRadova = append(radoviRaw.PodvrsteRadova, podvrstaRadova)
				}
			}
		}
	}
	return
}
