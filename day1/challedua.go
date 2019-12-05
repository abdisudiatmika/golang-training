package soal

import (
	"strings"
)

//Siswa struk simpan siswa
type Siswa struct {
	Name       string
	NameEncode string
	Score      int
}

//Chiper sasasaasa
type Chiper interface {
	Encode() string
	Decode() string
}

//Encode sdada
func (s *Siswa) Encode() string {
	//var enkrip string
	for i := 0; i < len(s.Name); i++ {
		if i%2 == 0 {
			s.NameEncode += string(s.Name[i]) + "DE"
		} else {
			s.NameEncode += string(s.Name[i])

		}

	}
	return s.NameEncode
}

//Decode sdadadada
func (s Siswa) Decode() string {
	//var dekrip string
	//fmt.Println(s.NameEncode)
	Dekrip := strings.ReplaceAll(s.NameEncode, "DE", "")
	return Dekrip
}
