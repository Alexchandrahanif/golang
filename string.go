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
		// 







	
	// Map

	// IF

	// Switch Expression

	// For Loop

	// Break & Continues

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



