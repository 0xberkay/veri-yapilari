package main

import (
	"fmt"
)

//İkili Agac Dugumu
type İkiliAğaçDüğümu struct {
	Deger int
	Sol   *İkiliAğaçDüğümu
	Sag   *İkiliAğaçDüğümu
}

// ikili ağaç düğüm Oluşturma Fonksiyonu
func (d *İkiliAğaçDüğümu) Oluştur(veri int) {
	if veri != 0 { // eğer veri 0 değilse
		if d.Deger < veri { // eğer veri düğümün değerinden büyükse
			//sağa taşı

			if d.Sag == nil { // eğer sağ düğüm boşsa
				d.Sag = &İkiliAğaçDüğümu{Deger: veri} // düğümün değerini düğümün sağına ata
			} else { // eğer sağ düğüm boş değilse
				d.Sag.Oluştur(veri) // fonksiyonu tekrar çağır
			}
		} else if d.Deger > veri { // eğer düğümün değeri veri değerinden büyükse
			//sola taşı
			if d.Sol == nil { // eğer sol düğüm boşsa
				d.Sol = &İkiliAğaçDüğümu{Deger: veri} // düğümün değerini düğümün soluna ata
			} else { // eğer sol düğüm boş değilse
				d.Sol.Oluştur(veri) // fonksiyonu tekrar çağır
			}
		}

	} else {
		fmt.Println("Veri 0 olamaz")
	}

}

// ikili ağaç arama fonksiyonu
func (d *İkiliAğaçDüğümu) Ara(veri int) bool {
	if d.Deger == veri { // eğer veri bulunduysa
		return true
	} else if d.Deger > veri { // eğer veri düğümün değerinden büyükse
		if d.Sol == nil { // eğer sol düğümü boşsa
			return false
		} else { // eğer sol düğümü boş değilse
			return d.Sol.Ara(veri)
		}
	} else { // eğer veri düğümün değerinden küçükse
		if d.Sag == nil { // eğer sağ düğümü boşsa
			return false
		} else { // eğer sağ düğümü boş değilse
			return d.Sag.Ara(veri)
		}
	}
}

//Min en küçük elemanı bulma
func (d *İkiliAğaçDüğümu) enKüçük() int {
	if d.Sol == nil { // eğer sol düğüm boşsa
		return d.Deger // düğümün değerini döndür
	} else { // eğer sol düğüm boş değilse
		return d.Sol.enKüçük() // fonksiyonu tekrar çağır
	}
}

//Max en büyük elemanı bulma
func (d *İkiliAğaçDüğümu) enBüyük() int {
	if d.Sag == nil { // eğer sağ düğüm boşsa
		return d.Deger // düğümün değerini döndür
	} else { // eğer sağ düğüm boş değilse
		return d.Sag.enBüyük() // fonksiyonu tekrar çağır
	}
}

// Ekleme fonksiyonu
func (d *İkiliAğaçDüğümu) Ekle(veri int) {
	if d.Deger < veri {
		//sağa taşı
		if d.Sag == nil { // eğer sağ düğüm boşsa
			d.Sag = &İkiliAğaçDüğümu{Deger: veri} // düğümün değerini düğümün sağına ata
		} else { // eğer sağ düğüm boş değilse
			d.Sag.Ekle(veri) // fonksiyonu tekrar çağır
		}
	} else if d.Deger > veri { // eğer düğümün değeri veri değerinden büyükse
		//sola taşı
		if d.Sol == nil { // eğer sol düğüm boşsa
			d.Sol = &İkiliAğaçDüğümu{Deger: veri} // düğümün değerini düğümün soluna ata
		} else { // eğer sol düğüm boş değilse
			d.Sol.Ekle(veri) // fonksiyonu tekrar çağır
		}
	}
}

// Silme fonksiyonu
func (d *İkiliAğaçDüğümu) Sil(veri int) {

	if d.Deger == veri { // eğer veri bulunduysa
		if d.Sol == nil && d.Sag == nil { // eğer sol ve sağ düğümler boşsa
			d.Deger = 0 // düğümün değerini sıfırla
		} else if d.Sol == nil { // eğer sol düğüm boşsa

			d.Deger = d.Sag.enKüçük() // düğümün değerini sağ düğümün en küçük elemanına ata
			d.Sag.Sil(d.Deger)        // fonksiyonu tekrar çağır
		} else if d.Sag == nil { // eğer sağ düğüm boşsa
			d.Deger = d.Sol.enBüyük() // düğümün değerini sol düğümün en büyük elemanına ata
			d.Sol.Sil(d.Deger)        // fonksiyonu tekrar çağır
		} else { // eğer sol ve sağ düğümler boş değilse

			d.Deger = d.Sag.enKüçük() // düğümün değerini sağ düğümün en küçük elemanına ata
			d.Sag.Sil(d.Deger)        // fonksiyonu tekrar çağır
		}
	} else if d.Deger > veri { // eğer veri düğümün değerinden büyükse
		if d.Sol == nil { // eğer sol düğüm boşsa
			fmt.Println("Veri bulunamadı") // veri bulunamadı mesajı yazdır
		} else { // eğer sol düğüm boş değilse
			d.Sol.Sil(veri) // fonksiyonu tekrar çağır
		}
	} else { // eğer veri düğümün değerinden küçükse
		if d.Sag == nil { // eğer sağ düğüm boşsa
			fmt.Println("Veri bulunamadı") // veri bulunamadı mesajı yazdır
		} else { // eğer sağ düğüm boş değilse
			d.Sag.Sil(veri) // fonksiyonu tekrar çağır
		}
	}
}

// Preorder fonksiyonu
// bir düğüm onun neslinden önce ziyaret edilir.
func (d *İkiliAğaçDüğümu) Preorder() {
	if d != nil {
		println(d.Deger)
		d.Sol.Preorder()
		d.Sag.Preorder()
	}
}

// Postoreder fonksiyonu
// bir düğüm onun neslinden sonra ziyaret edilir.
func (d *İkiliAğaçDüğümu) Postorder() {
	if d != nil {
		d.Sol.Postorder()
		d.Sag.Postorder()
		println(d.Deger)
	}
}

// Inorder fonksiyonu
// bir düğüm onun neslinden arasında ziyaret edilir.
func (d *İkiliAğaçDüğümu) Inorder() {
	if d != nil {
		d.Sol.Inorder()
		println(d.Deger)
		d.Sag.Inorder()
	}
}

// main fonksiyonu
func ÇalıştırAğaç() {
	İkiliAğaçDüğümu := &İkiliAğaçDüğümu{Deger: 15} // düğümün kök değeri 10 olsun

	İkiliAğaçDüğümu.Oluştur(6)
	İkiliAğaçDüğümu.Oluştur(13)
	İkiliAğaçDüğümu.Oluştur(18)
	İkiliAğaçDüğümu.Oluştur(7)
	İkiliAğaçDüğümu.Oluştur(30)
	İkiliAğaçDüğümu.Oluştur(3)
	İkiliAğaçDüğümu.Oluştur(2)
	İkiliAğaçDüğümu.Oluştur(4)
	if İkiliAğaçDüğümu.Ara(4) {
		println("4 bulundu")
	} else {
		println("4 bulunamadı")
	}
	fmt.Println("En küçük eleman:", İkiliAğaçDüğümu.enKüçük())
	fmt.Println("En büyük eleman:", İkiliAğaçDüğümu.enBüyük())
	İkiliAğaçDüğümu.Postorder()
	İkiliAğaçDüğümu.Sil(13)
	fmt.Println("--")
	İkiliAğaçDüğümu.Postorder()
	if İkiliAğaçDüğümu.Ara(13) {
		println("13 bulundu")
	} else {
		println("13 bulunamadı")

	}

}
