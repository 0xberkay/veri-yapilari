package main

import "fmt"

// Bir sonrakini ve öncekini ifade eden düğüm tipi
type ÇifteDüğüm struct {
	değer   int
	önceki  *ÇifteDüğüm
	sonraki *ÇifteDüğüm
}

//Çift yönlü liste türü
type ÇiftYönlüListe struct {
	baş     *ÇifteDüğüm
	kuyruk  *ÇifteDüğüm
	uzunluk int
}

//Başa ekleme fonksiyonu
func (ç *ÇiftYönlüListe) başaEkle(değer int) {
	yeniDüğüm := &ÇifteDüğüm{
		değer: değer,
	}
	if ç.baş == nil { // Eğer baş boş ise
		ç.baş = yeniDüğüm // listenin başını ve sonunu yeni düğüm olarak ata
		ç.kuyruk = yeniDüğüm
	} else { // Eğer baş boş değil ise
		yeniDüğüm.sonraki = ç.baş // yeni düğümün sonraki düğümü başını ata
		ç.baş.önceki = yeniDüğüm  // başın önceki düğümü yeni düğüm olarak ata
		ç.baş = yeniDüğüm         // başını yeni düğüm olarak ata
	}
	ç.baş = yeniDüğüm
	ç.uzunluk++
}

func (ç *ÇiftYönlüListe) kuyruğaEkle(değer int) { //Kuyruğa ekleme fonksiyonu
	yeniDüğüm := &ÇifteDüğüm{
		değer: değer,
	}
	if ç.kuyruk == nil { // Eğer kuyruk boş ise
		ç.kuyruk = yeniDüğüm // listenin başını ve sonunu yeni düğüm olarak ata
		ç.baş = yeniDüğüm
	} else { // Eğer kuyruk boş değil ise
		yeniDüğüm.önceki = ç.kuyruk  // yeni düğümün önceki düğümü kuyruğun sonuna ata
		ç.kuyruk.sonraki = yeniDüğüm // kuyruğun sonraki düğümü yeni düğüm olarak ata
		ç.kuyruk = yeniDüğüm         // kuyruğun sonraki düğümü yeni düğüm olarak ata
	}
	ç.kuyruk = yeniDüğüm
	ç.uzunluk++
}

func (ç *ÇiftYönlüListe) düğümSil(anahtar int) { //Düğüm silme fonksiyonu
	düğüm := ç.düğümBul(anahtar) // düğümü bul
	if düğüm == nil {            // düğüm boş ise
		return
	} else { // düğüm boş değil ise
		if düğüm.önceki == nil { // düğümün önceki düğümü boş ise
			ç.baş = düğüm.sonraki // başını düğümün sonraki düğümü olarak ata
		} else if düğüm.sonraki == nil { // düğümün sonraki düğümü boş ise
			düğüm.önceki.sonraki = nil // düğümün önceki düğümün sonraki düğümü boş olarak ata
		} else {
			düğüm.önceki.sonraki = düğüm.sonraki // düğümün önceki düğümün sonraki düğümü düğümün sonraki düğümü olarak ata
			düğüm.sonraki.önceki = düğüm.önceki  // düğümün sonraki düğümün önceki düğümü düğümün önceki düğümü olarak ata
		}
		düğüm = nil // düğümü boş yap
		ç.uzunluk-- // uzunluğu 1 azalt
	}
}

func (ç *ÇiftYönlüListe) düğümBul(bulunacak int) *ÇifteDüğüm { //Düğüm bulma fonksiyonu
	düğüm := ç.baş     // başını düğüm olarak ata
	for düğüm != nil { // düğüm boş değil ise
		if düğüm.değer == bulunacak { // düğümün değeri bulunacak değer ile eşit ise
			return düğüm // düğümü döndür
		}
		düğüm = düğüm.sonraki // düğümün sonraki düğümü düğüm olarak ata
	}
	return düğüm // düğümü döndür
}

func (ç *ÇiftYönlüListe) listele() { //listeleme fonksiyonu
	düğüm := ç.baş     // başını düğüm olarak ata
	for düğüm != nil { // düğüm boş değil ise
		fmt.Printf("%d ", düğüm.değer) // düğümün değerini yazdır
		düğüm = düğüm.sonraki          // düğümün sonraki düğümü düğüm olarak ata
	}
}

func ÇifteListeÇalıştır() {
	liste := &ÇiftYönlüListe{}

	liste.başaEkle(1)
	liste.başaEkle(2)
	liste.başaEkle(3)
	liste.kuyruğaEkle(20)
	liste.başaEkle(4)
	liste.başaEkle(5)
	liste.listele()
	fmt.Println()
	liste.düğümBul(3).değer = 13
	if liste.düğümBul(1) != nil {
		fmt.Println("Bulundu")
	} else {
		fmt.Println("Bulunamadı")
	}

	liste.düğümSil(13)
	liste.listele()
}
