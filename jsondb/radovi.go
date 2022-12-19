package jsondb

import "github.com/w-ingsolutions/predmer-api/models"

// type WingVrstaRadova struct {
// 	Id                 int                      `json:"id"`
// 	Naziv              string                   `json:"naziv"`
// 	Opis               string                   `json:"opis"`
// 	Obracun            string                   `json:"obracun"`
// 	Jedinica           string                   `json:"jedinica"`
// 	Cena               float64                  `json:"cena"`
// 	Slug               string                   `json:"slug"`
// 	Omogucen           bool                     `json:"omogucen"`
// 	Baza               bool                     `json:"baza"`
// 	Element            bool                     `json:"element"`
// 	PodvrsteRadova     []WingVrstaRadova        `json:"podvrsteradova"`
// 	NeophodanMaterijal []WingNeophodanMaterijal `json:"neophodanmaterijal"`
// }

var podvrsteRadova []models.WingVrstaRadova = []models.WingVrstaRadova{
	{
		Id:    1,
		Naziv: "Pripremni",
		Opis:  "xxxxxxxxxx",
		Slug:  "pripremni",
	},
	{
		Id:    2,
		Naziv: "Istraživački",
		Opis:  "xxxxxxxxxx",
		Slug:  "istrazivacki",
	},
	{
		Id:    3,
		Naziv: "Rušenja",
		Opis:  "xxxxxxxxxx",
		Slug:  "rusenja",
	},
	{
		Id:    4,
		Naziv: "Zemljani",
		Opis:  "xxxxxxxxxx",
		Slug:  "zemljani",
	},
	{
		Id:    5,
		Naziv: "Zidarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "zidarski",
	},
	{
		Id:    6,
		Naziv: "Betonski",
		Opis:  "xxxxxxxxxx",
		Slug:  "betonski",
	},
	{
		Id:    7,
		Naziv: "Tesarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "tesarski",
	},
	{
		Id:    8,
		Naziv: "Pokrivački",
		Opis:  "xxxxxxxxxx",
		Slug:  "pokrivacki",
	},
	{
		Id:    9,
		Naziv: "Izolaterski",
		Opis:  "xxxxxxxxxx",
		Slug:  "izolaterski",
	},
	{
		Id:    10,
		Naziv: "Stolarija",
		Opis:  "xxxxxxxxxx",
		Slug:  "stolarija",
	},
	{
		Id:    11,
		Naziv: "Stolarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "stolarski",
	},
	{
		Id:    12,
		Naziv: "Bravarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "bravarski",
	},
	{
		Id:    13,
		Naziv: "Limarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "limarski",
	},
	{
		Id:    14,
		Naziv: "Staklorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "staklorezacki",
	},
	{
		Id:    15,
		Naziv: "Keramičarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "keramicarski",
	},
	{
		Id:    16,
		Naziv: "Teracerski",
		Opis:  "xxxxxxxxxx",
		Slug:  "teracerski",
	},
	{
		Id:    17,
		Naziv: "Kamenorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "kamenorezacki",
	},
	{
		Id:    18,
		Naziv: "Parketarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "parketarski",
	},
	{
		Id:    19,
		Naziv: "Podopolagački",
		Opis:  "xxxxxxxxxx",
		Slug:  "podopolagacki",
	},
	{
		Id:    20,
		Naziv: "Tapetarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "tapetarski",
	},
	{
		Id:    21,
		Naziv: "Roletnarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "roletnarski",
	},
	{
		Id:    22,
		Naziv: "Suvomontažni",
		Opis:  "xxxxxxxxxx",
		Slug:  "suvomontazni",
	},
	{
		Id:    23,
		Naziv: "Gipsarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "gipsarski",
	},
	{
		Id:    24,
		Naziv: "Fasaderski",
		Opis:  "xxxxxxxxxx",
		Slug:  "fasaderski",
	},
	{
		Id:    25,
		Naziv: "Likorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "likorezacki",
	},
	{
		Id:    26,
		Naziv: "Molerski",
		Opis:  "xxxxxxxxxx",
		Slug:  "molerski",
	},
	{
		Id:    27,
		Naziv: "Livački",
		Opis:  "xxxxxxxxxx",
		Slug:  "livacki",
	},
	{
		Id:    28,
		Naziv: "Razni",
		Opis:  "xxxxxxxxxx",
		Slug:  "razni",
	},
	{
		Id:    29,
		Naziv: "Vodovod",
		Opis:  "xxxxxxxxxx",
		Slug:  "vodovod",
	},
	{
		Id:    30,
		Naziv: "Kanalizacija",
		Opis:  "xxxxxxxxxx",
		Slug:  "kanalizacija",
	},
	{
		Id:    31,
		Naziv: "Sanitarije",
		Opis:  "xxxxxxxxxx",
		Slug:  "sanitarije",
	},
}
