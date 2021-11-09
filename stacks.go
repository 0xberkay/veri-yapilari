package main

import "fmt"

type Yığın struct {
	Elemanlar []int
}

//koy (push)
func (y *Yığın) koy(eleman int) { // yeni bir nesne ekler
	y.Elemanlar = append(y.Elemanlar, eleman) // Yığının elemanlarının üstüne eleman ekler
}

//al (Pop)
func (y *Yığın) al() int { // en üsteki elemanı alır ve siler
	silinecekBoyut := len(y.Elemanlar) - 1     //Silinince olucak boyut
	eleman := y.Elemanlar[silinecekBoyut]      // en üsteki elemanı belirler
	y.Elemanlar = y.Elemanlar[:silinecekBoyut] // en üsteki elemanı yığından çıkarır
	return eleman                              // en üsteki elemanı döndürür
}

//dondur (Top)
func (y *Yığın) dondur() int { // en üsteki elemanı döndürür
	return y.Elemanlar[len(y.Elemanlar)-1]
}

// Boyut (Size)
func (y *Yığın) boyut() int { // yığının eleman sayısını döndürür
	return len(y.Elemanlar)
}

// bosMu (isEmpty)
func (y *Yığın) bosMu() bool { // yığın boş mu kontrolü
	return y.boyut() == 0 // yığının eleman sayısı 0 ise true döndürür
}

func ÇalıştırYığınlar() {
	yığın := Yığın{}
	fmt.Printf("Yığın boş mu ? : %v \n", yığın.bosMu())
	yığın.koy(1)
	yığın.koy(2)
	yığın.koy(3)
	yığın.koy(4)
	yığın.koy(5)
	yığın.koy(6)
	yığın.koy(7)
	fmt.Println(yığın)
	fmt.Printf("Yığının en üsteki elemanı : %d ve silindi\n", yığın.al())
	fmt.Printf("Yığının en üsteki elemanı : %d \n", yığın.dondur())
	fmt.Printf("Yığının boyutu : %d \n", yığın.boyut())
	fmt.Printf("Yığın boş mu ? : %v \n", yığın.bosMu())
	fmt.Println(yığın)

}
