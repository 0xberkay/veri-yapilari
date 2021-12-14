package main

import "fmt"

//AVL Ağaç tipi
type AVLAğaç struct {
	veri      int
	yükseklik int
	sol       *AVLAğaç
	sağ       *AVLAğaç
}

// En büyük sayıyı döndüren fonksiyon
func enBüyük(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// En küçük sayıyı döndüren fonksiyon
func enKüçük(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Ekleme
func (a *AVLAğaç) Ekle(veri int) {
	if a == nil {
		a = &AVLAğaç{veri: veri, yükseklik: 1}
	} else if veri < a.veri {
		if a.sol == nil {
			a.sol = &AVLAğaç{veri: veri, yükseklik: 1}
		} else {
			a.sol.Ekle(veri)
		}
	} else if veri > a.veri {
		if a.sağ == nil {
			a.sağ = &AVLAğaç{veri: veri, yükseklik: 1}
		} else {
			a.sağ.Ekle(veri)
		}
	}
	a.yükseklik = 1 + enBüyük(a.sol.yükseklik, a.sağ.yükseklik)

}

func ÇalıştırAVL() {
	a := AVLAğaç{veri: 3}

	fmt.Println(a)
	a.Ekle(2)
	fmt.Println(a)
}
