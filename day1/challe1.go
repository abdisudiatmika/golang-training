package soal

//Student untuk menyimpan sebuah siswa
type Student struct {
	Name  []string
	Score []int
}

//Average sdadada
func (s Student) Average() (float64, string, string) {
	var jumlah int
	min := 0
	namamin := ""
	namamax := ""
	max := 0
	for index, value := range s.Score {
		jumlah += value

		if min > value || min == 0 {
			min = value
			namamin = s.Name[index]
			//fmt.Println(s.Name)
		}
		if max < value {
			max = value
			namamax = s.Name[index]
			//fmt.Println(s.Name)
		}

	}
	Average := float64(jumlah) / float64(len(s.Score))

	return Average, namamin, namamax
}
