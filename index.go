package main

import (
	"errors"
	"fmt"
)

// function biasa
func biasa(){
	fmt.Println("ini adalah function biasa aja, tanpa return")
}

// function dengan paramter
func denganParameter (angka1 int, angka2 int){
	fmt.Println(angka1 + angka2)
}

// function dengan return biasa tanpa variabel
func returnBiasa(umur int, nama string, hobi string, pekerjaan string) string {
	return fmt.Sprintf("nama saya adalah %s, umur saya %d tahun %s %s", nama, umur,  pekerjaan, hobi)
}		

// fuction dengan return variabel (gila sih ini, mantap)
func returnVariabel(angka int)(hasil string, membuat string){
	membuat = "oke"
	hasil = fmt.Sprintf("Hallo Umur saya %d tahun", angka)
	return 
}

// function variadic (...int, ...string)
	// disebut juga variabel argument, syaratnya harus paling ujung, misal ada 3 paramter maka posisi nya harus paling kanan
	// tidak boleh ada 2 paramter argument didalam 1 fucnction
	// dengan variadic, membuat paramter nya tidak menjadi wajib

	func variadic(number ...int)int{
		total := 0
		for _,number := range number{
			total += number
		}
		return total
	}

// function as paramter
	// dimana function akan menjadi sebuah parameter, gila gak tuh

func filterChatBocilToxic(kata string, filter func(string) string) string{ // artinya ada 2 parameter dimana peram pertama string, kedua ada fuct yg menerima 1 parameter string, dan return string juga
	hasilFilter := filter(kata)
	return hasilFilter	
}

func filterChatToxic(kata string)string{
	if (kata == "anjing" || kata == "Babi" || kata == "Anak Yateam" || kata == "Anak Babi"){
		return "..."
	} else {
		return kata
	}
}



	// kode di atas sama aja dengan for of di js
	// for (let number of numbers) {
  //   total += number;
	// }


// Defer, Panic dan Recover
func contohDefer1(){
	fmt.Println("Ini Defer 1")


	// disini kita menyimpan recover yg berisi data panic, tetapi disini tidaklah wajib untuk menyimpannyha
	// cukup dengan " recover() " saja sebenarnya sudah cukup, tanpa harus di tampilkan
	message := recover()
 
	fmt.Println(message) 
}

func contohDefer2(){
	fmt.Println("Ini Defer 2")
}

func pengunaanDefer(hasil bool){
	defer contohDefer1()
	defer contohDefer2() // ketika ada panic, maka ini dahulu dijalankan dari pada defer 1

 	fmt.Println("Proses Function dijalankan")

	if(hasil){
		panic("APLIKASI EROR")
	}

	fmt.Println("Lanjutan Aplikasi") // ini tidak akan dijalankan ketika panic
	// dan akan melanjutkan proses di dalam function nya ketika ada recover didalan function defer
	
}

// Struct
type Customer struct{
	nama, alamat 	string
	umur 					int
	// bekerja				bool // ketika ini ditambahkan, maka data1 akan error, dan ini kekurangan nya

}
	// cek function methon nya dibaris 103
func (customer Customer) sayHai(){
	fmt.Println("Hallo nama saya " + customer.nama)
}


// Interface
type Pelanggan interface{}



// contoh interface kosong
func InterfaceKosong (i int) interface{}{
	if i == 1 {
		return "Ini Angka 1"
		} else if i == 2 {
			return "Ini Angka 1"
		} else {
			return "Ini Bukan Angka 1 dan 2"
	}
}

// contoh interface error
func Pembagi(nilai int, pembagi int) (int, error) { // pada contoh disini func mereturn int, dan error
	if(pembagi == 0) {
		return 0,errors.New("Pembagi Tidak Boleh 0") 	// itu kenapa disini pertama return 0, dan pesan error,
	} else {																				// pesan error nya dibuat dengan .New("pesan errro")
		return nilai/pembagi, nil	// dan disini, karena tidak eror, maka dikembalikam nil diparamter return ke 2
	}
}


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
		// model biasa
		biasa()
		
		// Function Parameter
		denganParameter(10, 20)

		// function dengan return biasa tanpa variabel
		hasilreturnbiasa := returnBiasa(20, "alex", "Ngoding", "Tukang Ketik" )
		fmt.Println(hasilreturnbiasa)

		// fuction dengan return variabel (gila sih ini, mantap)
		hasil, membuat := returnVariabel(10)
		fmt.Println(hasil, membuat)

		// Variadic Function (menambahkan parameter lebih dari 1)

		hasilVariadicBiasa := variadic(1,2,3,4,5)
		fmt.Println(hasilVariadicBiasa)

		hasilVariadicKosongan := variadic()
		fmt.Println(hasilVariadicKosongan)



		// kalau kita mengisi parameter dengan slice, maka ada tambahan 3 titik (...) didepan
		varibelSlice := []int{2,2,2,2}
		hasilVariadicDenganSlice := variadic(varibelSlice...)
		fmt.Println(hasilVariadicDenganSlice)


		// Function Value
			// maksudnya disini kita menyimpan function ke dalam sebuah vaiabel dahulu
			// syaratanya mesti mereturn sesuatu, bukan yg fmt.PrintLn aja

		contohVariabelValue := returnBiasa

		hasilVariabelValue := contohVariabelValue(1, "ini", "merupakan function value", "gokil banger")
		fmt.Println(hasilVariabelValue)

		// Function as Parameter
			// disini lebih gila lagi, kita bisa membuat function sebagai parameter, gokil banget
		
		hasilDariFunctionFilterToxic := filterChatBocilToxic("Babi", filterChatToxic)
		fmt.Println(hasilDariFunctionFilterToxic)

	// Anonymous Function
		// sama aja kek js
		
		
	// Recursive Function
		// menjalankan function didalam function sendiri
		

	// Closure
		// mengubah variabel diluar function, didalam function, sama aja kek js
		
	// Defer Panic Recover
		// Defer => menjalan kan function A, ketika function B selesai, mau error atau engga tetap muncul
		// Panic => untuk mentriger ketika ada eror didalam funtion A, setelah function B dijalankan
		// Recover => untuk menangkap data panic, dengan recover => function nya tetap berjalan sampai bawah walaupun ada panic
	

	pengunaanDefer(true)

	fmt.Println("Lanjutan Panic, ketika ada recover")


	//! Struct
		// Template untuk menggabungkan lebih dari 1 tipe data
		// Merepresentasi data
		// Kumpulan dari Field

		// type Customer struct{
		// 	nama, alamat 	string
		// 	umur 					int
		// 	bekerja				bool // ketika ini ditambahkan, maka data1 akan error, dan ini kekurangan nya
		
		// }


		
	// bisa cara ini 
	var data Customer
	data.nama = "Alex Chandra Hanif"
	data.alamat = "Jakarta Pusat"
	data.umur = 24 
	
	fmt.Println(data) // semua data
	fmt.Println(data.nama) // data nama saja


	// atau ini
	data1 := Customer{"Alex", "Senen", 10}
	fmt.Println(data1)
	// tetapi disini, untuk posisi data harus sesuai dengan struct nya, 
	// tidak boleh kebalik, misal int di nomor 2, maka harus int, gak boleh string

	// atau ini
	data2 := Customer{
		nama: "Chandra",
		alamat: "Pekanbaru", 
		umur: 20,
	}

	fmt.Println(data2)


	// penggunaan Struct Method
		// cek function methon nya dibaris 103
		// jadi dia memanggil menjalankan function yg mana data nya dari struct



	data.sayHai() // Hallo nama saya Alex Chandra Hanif
	data1.sayHai() // Hallo nama saya Alex
	data2.sayHai()	// Hallo Nama saya Chandra


	// dibawah ini contoh lain

	// type Person struct {
  //   Name    string
  //   Age     int
  //   Address string
	// }		

	// func (p *Person) sayHello() {
	// 		fmt.Println("Hello, I'm", p.Name)
	// }

	// func main() {
  //   person := Person{"John Doe", 30, "123 Main Street"}
  //   fmt.Println(person.Name, person.Age, person.Address)

  //   person.sayHello()
	// }

	


	//! Interface
		// Merupakan tipe data Abstrack, yang tidak memiliki implementasi langsung
		// biasanya digunakan sebagai kontrak
		// masih belum mengerti


	//! Interface Kosong
	// ada juga interface kosong (baris 114)
	hasilInterfaceKosong := InterfaceKosong(3)
	fmt.Println(hasilInterfaceKosong)

	//! Interface Error ( baris 127)
		// untuk return error seperti dijavascript

	hasilErorInterface, eror := Pembagi(10, 0)

	if(eror == nil){
		fmt.Println(hasilErorInterface)
	} else {
		fmt.Println(eror)
	}

	//! Type Assertions
		// mengubah function yg awal nya interface kosong (any), menjadi type suatu type data yg kira yakin return nya
	// contoh 

	// func contoh() interface{}{
	// 	retun "OKE"
	// }

	// ketika di main
	// hasilContoh := contoh()  => disini hasilnya adalah OKE
	// typebaru := hasilContoh.(string) => disini menyimpan divariabel baru dan diubah type data interface menjadi string


	// catatan 
	// tapi ingat, untuk melakukan di atas, harus yakin bahwa return nya ada lah string atau yg lain
	// ketika kita buat string, dan rupanya return nya adalah bool, atau int, maka akan timbul panic


	// tapi ketika mau memaksakan, maka bisa membuat switch case

	// switch value := typebaru.(type) {
		// case string : 
		// fmt.PrintLn("string", value)
		// case int : 
		// fmt.PrintLn("integer", value)
		// case bool : 
		// fmt.PrintLn("boolean", value)
		// default : 
		// fmt.PrintLn("Unknown", value)
	//}

	// dengan witch di atas, maka kita bisa set sesuai kemungkinan yg akan terjadi


	//! Pointer
		// intinya seperti di js, ketika data variabel A, dipindahkan di variabel B
		// ketika data variabel B di ubah, maka data Varibel A tidak ikut ke ubah
		
	
	

	





}




