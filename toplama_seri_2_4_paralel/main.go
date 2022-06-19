package main

import (
	"fmt"
	"time"
)

func main() {
	array_size := 100000000

	CriticalSection := make(chan bool, 1)
	//NumberArray int dizisi 10 sayi tutacak.
	NumberArray := make([]uint64, array_size)

	//NumberArray icinde donuyoruz verileri setliyoruz. (Seri Kısım)
	for i := 0; i < len(NumberArray); i++ {
		NumberArray[i] = uint64(i)
	}

	//Kanala true boolean deger gonderiyoruz.
	CriticalSection <- true

	//Değerler setlenmeden geçilemez. "Barrier" with Critical Section.
	<-CriticalSection

	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//Paralell Toplama icin 2 adet kanal olusturduk.
	k := make(chan uint64, 2)
	baslangic_paralel := time.Now() //Time Start.
	//Anonim Fonksiyon1.
	go func() {
		var toplam uint64 = 0
		for i := 0; i < array_size/2; i++ {
			toplam += NumberArray[i]

		}
		k <- toplam
	}()

	//Anonim Fonksiyon2.
	go func() {
		var toplam uint64 = 0
		for i := array_size / 2; i < array_size; i++ {
			toplam += NumberArray[i]

		}
		k <- toplam
	}()

	//İki adet kanalımızın değerini dopluyoruz (1-100 arası için)->{ilk kanal 0-50, ikinci kanal 50-100}.
	toplam := <-k + <-k
	gecenSure_paralel := time.Since(baslangic_paralel) // Time End.
	fmt.Println("[2] Paralel Toplam: [ ", toplam, " ] ", "GecenSure: [", gecenSure_paralel, "]")
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//Seri fonksiyon için 2 saniye bekleyelim (Burası önemsiz-kısımdır sadece geçişi farkedelim diye eklendi.)
	time.Sleep(2 * time.Second)

	baslangic := time.Now() //Time Start.
	var toplam_seri uint64 = 0
	for i := 0; i < array_size; i++ {
		toplam_seri += NumberArray[i]

	}
	gecenSure := time.Since(baslangic) // Time End.

	fmt.Println("Seri Toplam: [ ", toplam_seri, " ] ", "GecenSure: [", gecenSure, "]")
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------
	//Paralell Toplama icin 2 adet kanal olusturduk.
	f := make(chan uint64, 4)
	baslangic_paralel2 := time.Now() //Time Start.
	//Anonim Fonksiyon1.
	go func() {
		var toplam uint64 = 0
		for i := 0; i < array_size/4; i++ {
			toplam += NumberArray[i]

		}
		f <- toplam
	}()

	//Anonim Fonksiyon2.
	go func() {
		var toplam uint64 = 0
		for i := array_size / 4; i < (array_size/4)*2; i++ {
			toplam += NumberArray[i]

		}
		f <- toplam
	}()

	//Anonim Fonksiyon3.
	go func() {
		var toplam uint64 = 0
		for i := (array_size / 4) * 2; i < (array_size/4)*3; i++ {
			toplam += NumberArray[i]

		}
		f <- toplam
	}()

	//Anonim Fonksiyon4.
	go func() {
		var toplam uint64 = 0
		for i := (array_size / 4) * 3; i < (array_size/4)*4; i++ {
			toplam += NumberArray[i]

		}
		f <- toplam
	}()

	//İki adet kanalımızın değerini dopluyoruz (1-100 arası için)->{0-25-50-100}.
	toplam2 := <-f + <-f + <-f + <-f
	gecenSure_paralel2 := time.Since(baslangic_paralel2) // Time End.
	fmt.Println("[4] Paralel Toplam: [ ", toplam2, " ] ", "GecenSure: [", gecenSure_paralel2, "]")

}
