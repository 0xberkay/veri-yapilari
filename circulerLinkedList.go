package main

import "fmt"

func (ç *ÇiftYönlüListe) DaireselListeDönüştür() *ÇiftYönlüListe { // ÇiftYönlüListeyi dairesel listeye dönüştürür
	şuanki := ç.baş
	for şuanki.sonraki != nil { // Şuankinin sonrasındaki boş olana kadar çalışır
		şuanki = şuanki.sonraki // Şuanki sonrakiye atar
	}
	şuanki.sonraki = ç.baş // sonla ilk birleştirir
	return ç               // Dairesel listeyi döndürür
}

func DaireselListeÇalıştır() {
	liste := &ÇiftYönlüListe{}

	liste.başaEkle(1)
	liste.başaEkle(2)
	liste.başaEkle(3)
	liste.başaEkle(4)
	liste.başaEkle(5)
	liste.başaEkle(23)
	liste.başaEkle(24)
	liste.başaEkle(25)
	liste.listele()
	dairesel := liste.DaireselListeDönüştür()
	fmt.Println("\nDairesel Liste")
	fmt.Printf("Kuyruk Değeri %d\n", dairesel.kuyruk.değer)
	fmt.Printf("Baş Değeri %d\n", dairesel.baş.değer)
	fmt.Printf("Kuyruktan sonra gelen değer %d\n", dairesel.kuyruk.sonraki.değer)
}
