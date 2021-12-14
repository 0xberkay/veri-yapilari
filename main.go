package main

import "fmt"

func main() {

	makine()
}

func makine() {
	fmt.Println("Çalıştırılması istediğiniz değeri girin : ")
	var a int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Println("1. ÇalıştırBağlıListe.")
		ÇalıştırBağlıListe()

	} else if a == 2 {
		fmt.Println("2. ÇifteListeÇalıştır.")
		ÇifteListeÇalıştır()
	} else if a == 3 {
		fmt.Println("3. ÇalıştırYığınlar.")
		ÇalıştırYığınlar()
	} else if a == 4 {
		fmt.Println("4. ÇalıştırKuyruk.")
		ÇalıştırKuyruk()
	} else if a == 5 {
		fmt.Println("5. DaireselListeÇalıştır.")
		DaireselListeÇalıştır()
	} else if a == 6 {
		stringim := "( a + b ) / c + ( d - c ) * f"
		fmt.Println(İnfixToPrefix(stringim))
	} else if a == 7 {
		ÇalıştırAğaç()
	} else if a == 8 {
		ÇalıştırAVL()
	}

}
