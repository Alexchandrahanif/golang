package main

import "fmt"

func main() {

	
	// string
	fmt.Println("Bismillah")

	// integer
	fmt.Println(10)

	// float 
	fmt.Println(10.5)
	
	// boolean
	fmt.Println(true)
	fmt.Println(false)
	

	x := 10
  y := 5

	// operator aritmatika
	fmt.Println(x + y) // Penjumlahan
	fmt.Println(x - y) // Pengurangan
	fmt.Println(x * y) // Perkalian
	fmt.Println(x / y) // Pembagian
	fmt.Println(x % y) // Modulus
	

	// operator perbandingan
	fmt.Println(x == y) // Sama dengan
	fmt.Println(x != y) // Tidak sama dengan
	fmt.Println(x < y)  // Kurang dari
	fmt.Println(x > y)  // Lebih dari
	fmt.Println(x <= y) // Kurang dari atau sama dengan
	fmt.Println(x >= y) // Lebih dari atau sama dengan

	// operator logika
	p := true
	q := false

	fmt.Println(p && q) // Dan (AND)
	fmt.Println(p || q) // Atau (OR)
	fmt.Println(!p)     // Tidak (NOT)

	// operator penugasan
	 a := 10
	 b := 5

	 // Penugasan nilai
	 fmt.Println("Penugasan Nilai:")
	 fmt.Println("a =", a)

	 // Penugasan penjumlahan
	 a += b
	 fmt.Println("\nPenugasan Penjumlahan:")
	 fmt.Println("a =", a)

	 // Penugasan pengurangan
	 a -= b
	 fmt.Println("\nPenugasan Pengurangan:")
	 fmt.Println("a =", a)

	 // Penugasan perkalian
	 a *= b
	 fmt.Println("\nPenugasan Perkalian:")
	 fmt.Println("a =", a)

	 // Penugasan pembagian
	 a /= b
	 fmt.Println("\nPenugasan Pembagian:")
	 fmt.Println("a =", a)

	 // Penugasan modulus
	 a %= b
	 fmt.Println("\nPenugasan Modulus:")
	 fmt.Println("a =", a)

	// var

	// const

	// type declaration

	// Array

	// cara membuat array terpisah dengan isi nya
	var hobi [2]string // maksimal isi dari array nya 2

	hobi[0] = "Sepakbola"
	hobi[1] = "Membaca"

	fmt.Println(hobi) // menampilkan semua data hobi s
	
	
	// cara membuat array beserta isinya
	var angkaKesukaan = [5]int{
		20,
		30,
		40,
	}
	
	
	fmt.Println(angkaKesukaan) // [20, 30, 40, 0, 0]
	
	// untuk mendapatkan banyak data
	fmt.Println(len(angkaKesukaan)) // 5, walaupun cuman ada 3 data, tapi data ke 4, dan 5 di defind defaul = 0
	
	// menampilkan data array per index 
	fmt.Println(hobi[1]) // menampilkan data hobi index 1
	
	// mengubah isi data per index

	angkaKesukaan[3] = 40
	fmt.Println(angkaKesukaan) // [20, 30, 40, 40, 0]

	// Slice
		// Hampir mirip dengan Array
		// Punya Length, Capacity, Pointer
		// Pointer => penujuk data pertama pada array
		// Length => panjang dari slice
		// Capacity => kapasitas slice, length tidak boleh lebih dari kapasitas
		// Slice bisa di ubah, append, dan copy








	
	// Map
		// Map => sama aja dengan Object Pada JavaScript
		// Memiliki Key dan Value

	contohMapString := map[string]string{
		"Barcelona" : "Messi",
		"Real Madrid" : "Ronaldo",
		"Manchester United" : "Rooney",
		"Manchester City" : "Dzeko",
	} 
	contohMapInteger := map[string]int{
		"Barcelona" : 10,
		"Real Madrid" : 20,
	} 

	// mendapatkan banyak data didalam map
	fmt.Println(len(contohMapString))

	// mengambil data didalam map, dengan key
	fmt.Println(contohMapString["Barcelona"])
	
	// mengubah data didalam map, dengan key
	contohMapString["Barcelona"] = "David Villa"

	// menghapus data didalam map dengan key
	delete(contohMapString, "Barcelona")

	// mau buat map baru
	mapBaru := make(map[string]string)

	// mengisi map dengan data baru 
	mapBaru["datasatu"] = "dataSatu"
	
	fmt.Println(mapBaru)

	fmt.Println(contohMapString, contohMapInteger)



	// IF
		// sama aja sih dengan js, jadi ezpz
	if (10 < 1) {
		fmt.Println("kocak")
	} else {
		fmt.Println("Kocag")
	}

	if(10 > 20){
		fmt.Println("Mantap")
	} else if(10 > 5){
		fmt.Println("Mantap Juga")
	} else {
		fmt.Println("KOCAKKS")
	}

	// Switch Expression
		// sama juga dengan switch di JS ejeed
	day := 4

  switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
  }

	// kalau multicase nya gini

  switch day {
   	case 1,3,5:
    	fmt.Println("Odd weekday")
   	case 2,4:
     	fmt.Println("Even weekday")
   	case 6,7:
    	fmt.Println("Weekend")
   	default:
    	fmt.Println("Invalid day of day number")
  }
	// For Loop
		// sama juga dengan JS gak terlalu beda
	for i:= 1; i<10; i++{
		fmt.Println(i)
	}



	// Break & Continues
		// sama juga dengan js

	for i:=0; i < 5; i++ {
		if i == 3 {
			continue
		}
	 fmt.Println(i, "continue")
	}	

	for i:=0; i < 5; i++ {
		if i == 3 {
			break
		}
	 fmt.Println(i, "break")
	}	

	// Function
		// Function Parameter
		// Function Return Value
		// Returning Multiple Values
		// Named Multiple Values
		// Variadic Function
		// Function Value
		// Function as Parameter
		// Anonymous Function
		// Recursive Function

	// Closure (Part 30)
}




