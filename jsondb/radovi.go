package jsondb

type WingVrstaRadova struct {
	Id                 int                      `json:"id"`
	Naziv              string                   `json:"naziv"`
	Opis               string                   `json:"opis"`
	Obracun            string                   `json:"obracun"`
	Jedinica           string                   `json:"jedinica"`
	Cena               float64                  `json:"cena"`
	Slug               string                   `json:"slug"`
	Omogucen           bool                     `json:"omogucen"`
	Baza               bool                     `json:"baza"`
	Element            bool                     `json:"element"`
	PodvrsteRadova     []WingVrstaRadova        `json:"podvrsteradova"`
	NeophodanMaterijal []WingNeophodanMaterijal `json:"neophodanmaterijal"`
}

var podvrsteRadova []WingVrstaRadova = []WingVrstaRadova{
	WingVrstaRadova{
		Id:    1,
		Naziv: "Pripremni",
		Opis:  "xxxxxxxxxx",
		Slug:  "pripremni",
	},
	WingVrstaRadova{
		Id:    2,
		Naziv: "Istraživački",
		Opis:  "xxxxxxxxxx",
		Slug:  "istrazivacki",
	},
	WingVrstaRadova{
		Id:    3,
		Naziv: "Rušenja",
		Opis:  "xxxxxxxxxx",
		Slug:  "rusenja",
	},
	WingVrstaRadova{
		Id:    4,
		Naziv: "Zemljani",
		Opis:  "xxxxxxxxxx",
		Slug:  "zemljani",
	},
	WingVrstaRadova{
		Id:    5,
		Naziv: "Zidarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "zidarski",
	},
	WingVrstaRadova{
		Id:    6,
		Naziv: "Betonski",
		Opis:  "xxxxxxxxxx",
		Slug:  "betonski",
	},
	WingVrstaRadova{
		Id:    7,
		Naziv: "Tesarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "tesarski",
	},
	WingVrstaRadova{
		Id:    8,
		Naziv: "Pokrivački",
		Opis:  "xxxxxxxxxx",
		Slug:  "pokrivacki",
	},
	WingVrstaRadova{
		Id:    9,
		Naziv: "Izolaterski",
		Opis:  "xxxxxxxxxx",
		Slug:  "izolaterski",
	},
	WingVrstaRadova{
		Id:    10,
		Naziv: "Stolarija",
		Opis:  "xxxxxxxxxx",
		Slug:  "stolarija",
	},
	WingVrstaRadova{
		Id:    11,
		Naziv: "Stolarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "stolarski",
	},
	WingVrstaRadova{
		Id:    12,
		Naziv: "Bravarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "bravarski",
	},
	WingVrstaRadova{
		Id:    13,
		Naziv: "Limarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "limarski",
	},
	WingVrstaRadova{
		Id:    14,
		Naziv: "Staklorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "staklorezacki",
	},
	WingVrstaRadova{
		Id:    15,
		Naziv: "Keramičarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "keramicarski",
	},
	WingVrstaRadova{
		Id:    16,
		Naziv: "Teracerski",
		Opis:  "xxxxxxxxxx",
		Slug:  "teracerski",
	},
	WingVrstaRadova{
		Id:    17,
		Naziv: "Kamenorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "kamenorezacki",
	},
	WingVrstaRadova{
		Id:    18,
		Naziv: "Parketarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "parketarski",
	},
	WingVrstaRadova{
		Id:    19,
		Naziv: "Podopolagački",
		Opis:  "xxxxxxxxxx",
		Slug:  "podopolagacki",
	},
	WingVrstaRadova{
		Id:    20,
		Naziv: "Tapetarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "tapetarski",
	},
	WingVrstaRadova{
		Id:    21,
		Naziv: "Roletnarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "roletnarski",
	},
	WingVrstaRadova{
		Id:    22,
		Naziv: "Suvomontažni",
		Opis:  "xxxxxxxxxx",
		Slug:  "suvomontazni",
	},
	WingVrstaRadova{
		Id:    23,
		Naziv: "Gipsarski",
		Opis:  "xxxxxxxxxx",
		Slug:  "gipsarski",
	},
	WingVrstaRadova{
		Id:    24,
		Naziv: "Fasaderski",
		Opis:  "xxxxxxxxxx",
		Slug:  "fasaderski",
	},
	WingVrstaRadova{
		Id:    25,
		Naziv: "Likorezački",
		Opis:  "xxxxxxxxxx",
		Slug:  "likorezacki",
	},
	WingVrstaRadova{
		Id:    26,
		Naziv: "Molerski",
		Opis:  "xxxxxxxxxx",
		Slug:  "molerski",
	},
	WingVrstaRadova{
		Id:    27,
		Naziv: "Livački",
		Opis:  "xxxxxxxxxx",
		Slug:  "livacki",
	},
	WingVrstaRadova{
		Id:    28,
		Naziv: "Razni",
		Opis:  "xxxxxxxxxx",
		Slug:  "razni",
	},
	WingVrstaRadova{
		Id:    29,
		Naziv: "Vodovod",
		Opis:  "xxxxxxxxxx",
		Slug:  "vodovod",
	},
	WingVrstaRadova{
		Id:    30,
		Naziv: "Kanalizacija",
		Opis:  "xxxxxxxxxx",
		Slug:  "kanalizacija",
	},
	WingVrstaRadova{
		Id:    31,
		Naziv: "Sanitarije",
		Opis:  "xxxxxxxxxx",
		Slug:  "sanitarije",
	},
}
