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
	PodvrsteRadova     []*WingVrstaRadova       `json:"podvrsteradova"`
	NeophodanMaterijal []WingNeophodanMaterijal `json:"neophodanmaterijal"`
}
