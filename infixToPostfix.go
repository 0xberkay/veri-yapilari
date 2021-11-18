package main

import (
	"fmt"
	"strings"
)

// İnfix ifadeyi postfix ifadeye dönüştürür
func İnfixToPrefix(a string) string {
	yığın := Yığın{}               // Yığını oluştur
	postfix := []string{}          // Postfix ifadeyi oluştur
	split := strings.Split(a, " ") // İnfix ifadeyi parçalayıp parçaları bir diziye atıyoruz

	for _, v := range split { // Dizi içindeki her bir elemanı kontrol ediyoruz
		if kontrol(v) == 0 { // Eğer parantez açıyorsa yığına ekliyoruz
			yığın.koy(v)
		}
		if kontrol(v) == 3 { // Eğer sayıysa postfix ifadeye ekliyoruz
			postfix = append(postfix, v)

		}
		if kontrol(v) == 2 { // Eğer işlemse
			if yığın.bosMu() || yığın.dondur() == "(" { // Eğer yığın boş ise veya yığının en üstündeki eleman parantez ise
				yığın.koy(v) // Yığına ekliyoruz
			} else { // Eğer değilse
				if v == "+" || v == "-" { // Eğer işlem + veya - ise
					if yığın.bosMu() { // Eğer yığın boş ise
						postfix = append(postfix, v) // Postfix ifadeye ekliyoruz
					}
					if yığın.dondur() == "*" || yığın.dondur() == "/" { // Eğer yığının en üstündeki eleman * veya / ise
						postfix = append(postfix, yığın.dondur()) // Postfix ifadeye ekliyoruz
						yığın.kaldır()                            // Yığından elemanı kaldırıyoruz
						yığın.koy(v)                              // Yığına ekliyoruz
					} else {
						yığın.koy(v) // Yığına ekliyoruz
					}

				}
				if v == "*" || v == "/" { // Eğer işlem * veya / ise
					yığın.koy(v) // Yığına ekliyoruz
				}

			}
		}
		if kontrol(v) == 1 { // Eğer parantez kapatıyorsa

			for i := 0; i < yığın.boyut(); i++ { // Yığının boyutu kadar döngüye gireceğiz
				if yığın.dondur() != "(" { // Eğer yığının en üstündeki eleman parantez değilse
					postfix = append(postfix, yığın.al()) // Postfix ifadeye ekliyoruz
				}
			}
			if !yığın.bosMu() { // Eğer yığın boş değilse
				yığın.al() // Yığından elemanı siliyoruz
			}

		}

	}
	for !yığın.bosMu() { // Eğer yığın boş değilse
		postfix = append(postfix, yığın.al()) // Postfix ifadeye ekliyoruz
	}

	return fmt.Sprintf("%v", postfix) // Postfix ifadeyi döndürüyoruz

}

//Kontrol gurubu
// 0 ==> aç parantez
// 1 ==> kapa parantez
// 2 ==> işlem
// 3 ==> sayı
//( 3 + 2 ) / 8
func kontrol(s string) int {
	if s == "*" || s == "/" || s == "+" || s == "-" {
		return 2
	} else if s == "(" {
		return 0
	} else if s == ")" {
		return 1
	} else {
		return 3
	}

}
