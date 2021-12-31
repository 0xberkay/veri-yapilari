package main

import (
	"fmt"
)

type AVL struct {
	Deger     int
	Yukseklik int
	Sol       *AVL
	Sag       *AVL
}

// AVL ağacı ekleme fonksiyonu
func (a *AVL) Ekle(veri int) {

	if a.Deger < veri { // eğer veri düğümün değerinden büyükse
		//sağa taşı
		if a.Sag == nil { // eğer sağ düğüm boşsa
			a.Sag = &AVL{Deger: veri} // düğümün değerini düğümün sağına ata
		} else { // eğer sağ düğüm boş değilse
			a.Sag.Ekle(veri) // fonksiyonu tekrar çağır
		}
	} else if a.Deger > veri { // eğer düğümün değeri veri değerinden büyükse
		//sola taşı
		if a.Sol == nil { // eğer sol düğüm boşsa
			a.Sol = &AVL{Deger: veri} // düğümün değerini düğümün soluna ata
		} else { // eğer sol düğüm boş değilse
			a.Sol.Ekle(veri) // fonksiyonu tekrar çağır
		}
	}
	a.Yukseklik = a.YukseklikFonksiyonu() // yukseklik fonksiyonu çağır
}

// AVL ağacının yukseklik fonksiyonu
func (a *AVL) YukseklikFonksiyonu() int {
	if a == nil { // eğer düğüm boşsa
		return 0 // yukseklik 0 döndür
	}
	sol := a.Sol.YukseklikFonksiyonu() // solun yukseklik değerini al
	sag := a.Sag.YukseklikFonksiyonu() // sagın yukseklik değerini al
	if sol > sag {                     // eğer solun yukseklik değeri sagının yukseklik değerinden büyükse
		return sol + 1 // solun yukseklik değerini 1 artır
	} else { // eğer sagının yukseklik değerinden küçükse
		return sag + 1 // sagının yukseklik değerini 1 artır
	}
}

// AVL ağacının sola dengeleme fonksiyonu (LL)
func (a *AVL) SolaDöndürme() *AVL {
	gecici := a       // gecici düğümünü a'ya ata
	a = a.Sol         // a'nın solunu a'ya ata
	if a.Sag != nil { // eğer a'nın sagı boş değilse
		if a.Sag.Deger > gecici.Deger { // eğer a'nın sagının değeri gecici düğümün değerinden büyükse
			gecici2 := a.Sag    // gecici2 düğümünü a'nın sagının sağına ata
			a.Sag = gecici      // a'nın sagını gecici düğümünü ata
			a.Sag.Sag = gecici2 // a'nın sagının sağını gecici2 düğümünü ata
			return a
		} else { // eğer a'nın sagının değeri gecici düğümün değerinden küçükse
			a.Sag.Sag = gecici // a'nın sagının sağını gecici düğümünü ata
			return a
		}
	} else { // eğer a'nın sagı boşsa
		a.Sag = gecici // a'nın sagını gecici düğümünü ata
		return a
	}

}

// AVL ağacının sağa dengeleme fonksiyonu (RR)
func (a *AVL) SagaDöndürme() *AVL {
	gecici := a       // gecici düğümünü a'ya ata
	a = a.Sag         // a'nın sagını a'ya ata
	if a.Sol != nil { // eğer a'nın solu boş değilse
		if a.Sol.Deger < gecici.Deger { // eğer a'nın solunun değeri gecici düğümün değerinden küçükse
			gecici2 := a.Sol    // gecici2 düğümünü a'nın solunun soluna ata
			a.Sol = gecici      // a'nın solunu gecici düğümünü ata
			a.Sol.Sol = gecici2 // a'nın solunun soluna gecici2 düğümünü ata
			return a
		} else { // eğer a'nın solunun değeri gecici düğümün değerinden büyükse
			a.Sol.Sol = gecici // a'nın solunun soluna gecici düğümünü ata
			return a
		}
	} else { // eğer a'nın solu boşsa
		a.Sol = gecici // a'nın solunu gecici düğümünü ata
		return a
	}
}

func ÇalıştırAVL() {
	a := &AVL{Deger: 8}
	a.Ekle(5)
	a.Ekle(3)

	fmt.Println(a.Deger)

	fmt.Println("a ", a.Deger)
	fmt.Println("a.Sol ", a.Sol.Deger)
	fmt.Println("a.Sol.Sol ", a.Sol.Sol.Deger)

	fmt.Println("______")
	a = a.SolaDöndürme()
	fmt.Println("a ", a.Deger)
	fmt.Println("a.Sağ", a.Sag.Deger)
	fmt.Println("a.Sol ", a.Sol.Deger)

}
