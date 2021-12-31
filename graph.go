package main

import "fmt"

type Diyagram struct {
	köşeler []*Köşe
}

type Köşe struct {
	anahtar  int
	bitişiği []*Köşe
}

// Köşe ekleme fonksiyonu
func (d *Diyagram) köşeEkle(a int) {
	// köşe var mı kontrol et
	for _, köşe := range d.köşeler { // köşeler dizisinde köşe var mı kontrol et
		if köşe.anahtar == a { // köşe var ise
			fmt.Println(köşe.anahtar, " numaralı Köşe zaten var")
			return //çık
		}
	}
	d.köşeler = append(d.köşeler, &Köşe{anahtar: a}) // köşe yok ise köşe ekle
}

// Kenar ekleme fonksiyonu
func (d *Diyagram) kenarEkle(nerden, nereye int) {

	gelenKöşe := d.köşeDöndür(nerden) // gelen köşe
	gidenKöşe := d.köşeDöndür(nereye) // giden köşe

	// köşelerin bitişi dizisinde köşe var mı kontrol et

	for _, köşe := range gelenKöşe.bitişiği { // köşeler dizisinde köşe var mı kontrol et
		if köşe.anahtar == gidenKöşe.anahtar { // köşe var ise
			fmt.Println(nerden, " numaralı Köşe ", gidenKöşe.anahtar, " numaralı Köşeye zaten bir kenarı var")
			return
		}
	}
	gelenKöşe.bitişiği = append(gelenKöşe.bitişiği, gidenKöşe) // köşe yok ise köşe ekle

}

// Köşe döndürme fonksiyonu
func (d *Diyagram) köşeDöndür(a int) *Köşe {
	for i, v := range d.köşeler { // köşeler dizisinde köşe var mı kontrol et
		if v.anahtar == a { // köşe var ise
			return d.köşeler[i] // köşeyi döndür
		}
	}
	return nil
}

// Yazdır
func (d *Diyagram) Yazdır() {
	for _, köşe := range d.köşeler {
		fmt.Print("[", köşe.anahtar, "] ")
		for _, kenar := range köşe.bitişiği {
			fmt.Print("-> [", kenar.anahtar, "] ")
		}
		fmt.Println()
	}

}

func DiyagramÇalıştır() {
	d := &Diyagram{}

	d.köşeEkle(1)
	d.köşeEkle(2)
	d.köşeEkle(2)
	d.kenarEkle(1, 2)
	d.kenarEkle(1, 2)
	d.Yazdır()
}
