package main

import (
	"fmt"
)

type Yığın struct {
	Elemanlar []string
}

//koy (push)  yeni bir nesne ekler
func (y *Yığın) koy(eleman string) {
	y.Elemanlar = append(y.Elemanlar, eleman) // Yığının elemanlarının üstüne eleman ekler
}

//al (Pop) en üsteki elemanı alır ve siler
func (y *Yığın) al() string {
	silinecekBoyut := len(y.Elemanlar) - 1     //Silinince olucak boyut
	eleman := y.Elemanlar[silinecekBoyut]      // en üsteki elemanı belirler
	y.Elemanlar = y.Elemanlar[:silinecekBoyut] // en üsteki elemanı yığından çıkarır
	return eleman                              // en üsteki elemanı döndürür
}

//dondur (Top) en üsteki elemanı döndürür
func (y *Yığın) dondur() string {
	return y.Elemanlar[len(y.Elemanlar)-1]
}

//dondur (Top) girilen miktarı kadar üsten elemanı döndürür
func (y *Yığın) dondurOkadar(sa int) string {
	return y.Elemanlar[len(y.Elemanlar)-sa]
}

// Boyut (Size) yığının eleman sayısını döndürür
func (y *Yığın) boyut() int {
	return len(y.Elemanlar)
}

// bosMu (isEmpty)  yığın boş mu kontrolü
func (y *Yığın) bosMu() bool {
	return y.boyut() == 0 // yığının eleman sayısı 0 ise true döndürür
}

// yığını temizler
func (y *Yığın) kaldır() {
	y.Elemanlar = []string{}
}

func ÇalıştırYığınlar() {
	yığın := Yığın{}
	fmt.Printf("Yığın boş mu ? : %v \n", yığın.bosMu())
	yığın.koy("1")
	yığın.koy("2")
	yığın.koy("3")
	yığın.koy("4")
	yığın.koy("5")
	yığın.koy("6")
	yığın.koy("7")

	fmt.Println(yığın)
	fmt.Printf("Yığının en üsteki elemanı : %s ve silindi\n", yığın.al())
	fmt.Printf("Yığının en üsteki elemanı : %s \n", yığın.dondur())
	fmt.Printf("Yığının boyutu : %v \n", yığın.boyut())
	fmt.Printf("Yığın boş mu ? : %v \n", yığın.bosMu())
	fmt.Println(yığın.dondurOkadar(3))
	fmt.Println(yığın)

}
