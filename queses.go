package main

import "fmt"

type Kuyruk struct {
	Elemanlar []int
}

// sırayaAl (enqueue)

func (k *Kuyruk) sırayaAl(eleman int) { // kuyruğun sonuna eleman ekle
	k.Elemanlar = append(k.Elemanlar, eleman)
}

// sırayıBoşalt (dequeue)

func (k *Kuyruk) sırayıBoşalt() int { // kuyruğun ilk elemanını döndür

	if k.BoşMu() { // Eğer kuyruk boşsa 0 döner(EmptyQueueException)
		fmt.Println("kuyruk boş")
		return 0
	}
	eleman := k.Elemanlar[0]
	k.Elemanlar = k.Elemanlar[1:]
	return eleman
}

// başıDöndür (getFront)

func (k *Kuyruk) başıDöndür() int { // kuyruğun ilk elemanını döndür
	return k.Elemanlar[0]
}

// boyut (size)

func (k *Kuyruk) boyut() int { // kuyruğun eleman sayısını döndür
	return len(k.Elemanlar)
}

// BoşMu (isEmpty)

func (k *Kuyruk) BoşMu() bool { // kuyruğun boş olup olmadığını döndür
	return k.boyut() == 0 // kuyruğun eleman sayısı 0 ise true döndür
}

func ÇalıştırKuyruk() {
	kuyruk := Kuyruk{}

	kuyruk.sırayıBoşalt()
	fmt.Printf("Kuyruğun boyutu : %d \n", kuyruk.boyut())
	kuyruk.sırayaAl(1)
	kuyruk.sırayaAl(2)
	kuyruk.sırayaAl(3)
	kuyruk.sırayaAl(4)
	kuyruk.sırayaAl(5)
	kuyruk.sırayaAl(6)
	kuyruk.sırayaAl(7)
	kuyruk.sırayaAl(8)
	fmt.Println(kuyruk)
	kuyruk.sırayıBoşalt()
	fmt.Println(kuyruk)
	fmt.Printf("Kuyruğun başı : %d \n", kuyruk.başıDöndür())

}
