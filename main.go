package main

import "fmt"

func main() {

	//ÇalıştırBağlıListe()
	//ÇifteListeÇalıştır()
	//ÇalıştırKuyruk()
	//ÇalıştırYığınlar()
	//DaireselListeÇalıştır()
	makine()
}

func makine() {
	fmt.Println("Çalıştırılması istediğiniz değeri girin : ")
	var a int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Println("1. Çalıştırılmıştır.")
		ÇalıştırBağlıListe()

	} else if a == 2 {
		fmt.Println("2. Çalıştırılmıştır.")
		ÇifteListeÇalıştır()
	} else if a == 3 {
		fmt.Println("3. Çalıştırılmıştır.")
		ÇalıştırYığınlar()
	}
}
