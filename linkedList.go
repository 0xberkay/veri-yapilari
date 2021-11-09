package main

import (
	"fmt"
)

type Düğüm struct { // Bir sonrakini ifade eden düğüm tipi
	değer int

	sonraki *Düğüm
}

type BağlıListe struct { // Bağlı liste tipi
	baş     *Düğüm
	uzunluk int
}

func (l *BağlıListe) başaEkle(n *Düğüm) { // Bağlı liste öne geçirerek ekleme fonksiyonu
	ikinci := l.baş        // ikinci düğümün baş düğüme atandı
	l.baş = n              // baş düğümü n düğümüne atandı
	l.baş.sonraki = ikinci // baş düğümün sonraki düğümü ikinci düğümüne atadı
	l.uzunluk++            // liste uzunluğu 1 arttırıldı
}

func (l *BağlıListe) değeriSil(silinecekDeğer int) {
	if l.uzunluk == 0 { // bağlı listenin uzunluğu 0 ise
		fmt.Println("Liste boş") // liste boş yazdırıldı
		return
	}
	if l.baş.değer == silinecekDeğer { // baş düğümün değeri silinecek değer ile eşitse
		l.baş = l.baş.sonraki // baş düğümün sonraki düğümü baş düğümüne atandı
		l.uzunluk--           // liste uzunluğu 1 azaltıldı
		return
	}

	silinecek := l.baş                              // silinecek düğüm baş düğümüne atandı
	for silinecek.sonraki.değer != silinecekDeğer { // silinecek düğümün sonraki düğümün değeri silinecek değer ile eşit değilse
		if silinecek.sonraki.sonraki == nil { // silinecek düğümün sonraki düğümün sonraki düğümü boş ise
			fmt.Println("Silinecek değer bulunamadı")
			return
		}
		silinecek = silinecek.sonraki // silinecek düğüm kendinden sonra gelen düğüme atandı bu şekilde değer ortadan kayboldu
	}
	silinecek.sonraki = silinecek.sonraki.sonraki // silinecek düğümün sonrakide güncellendi
	l.uzunluk--                                   // liste uzunluğu 1 azaltıldı
}

func (l BağlıListe) listeyiYazdır() { //Bağlı listeyi yazdırma fonksiyonu
	yazılacak := l.baş   // baş düğüm yazılacak düğüme atandı
	for l.uzunluk != 0 { // bağlı listenin sonuna gelene kadar
		fmt.Printf("%d", yazılacak.değer) // yazılacak düğümün değeri yazdırıldı
		yazılacak = yazılacak.sonraki     // yazılacak düğüm yazılacağın sonrakine atandı
		l.uzunluk--                       // liste uzunluğu 1 azaltıldı
		fmt.Printf(" ")

	}
}

func ÇalıştırBağlıListe() {
	BağlıListem := BağlıListe{} // Bağlı liste nesnesi oluşturuldu
	Düğüm1 := &Düğüm{değer: 1}  // Düğüm1 nesnesi oluşturuldu
	Düğüm2 := &Düğüm{değer: 23} // Düğüm2 nesnesi oluşturuldu
	Düğüm3 := &Düğüm{değer: 35} // Düğüm3 nesnesi oluşturuldu
	Düğüm4 := &Düğüm{değer: 44} // Düğüm4 nesnesi oluşturuldu

	BağlıListem.başaEkle(Düğüm1) // Düğüm1 bağlı listeye eklendi
	BağlıListem.başaEkle(Düğüm2) // Düğüm2 bağlı listeye eklendi
	BağlıListem.başaEkle(Düğüm3) // Düğüm3 bağlı listeye eklendi
	BağlıListem.başaEkle(Düğüm4) // Düğüm4 bağlı listeye eklendi
	BağlıListem.değeriSil(44)    // Düğüm2 silindi
	BağlıListem.listeyiYazdır()  // Bağlı listeyi yazdırma fonksiyonu çalıştırıldı
}
