package main

import (
	"fmt"
	"project1/syntax"
)

func main() {
	//int8-128 to 127
	//int16-32.767 to 32.768
	//int32-2.147.483.648 to 2.147.483.647
	//int64-9.223.372.036.854.775.808 to 9.223.372.036.854.775.807

	var name string = "Fikri"
	fmt.Println("Hello : " + name)

	alamat := "Bandung"
	fmt.Println("Alamat : " + alamat)

	var age int = 20
	fmt.Println("Age : ", age)

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke : ", i)
	}

	isActive := true
	isAuthorized := false
	fmt.Println("Is Active : ", isActive)
	fmt.Println("Is Authorized : ", isAuthorized)

	result := "sate"
	if result == "mie ayam" {
		fmt.Println("mie ayam")
	} else if result == "bakso" {
		fmt.Println("bakso")
	} else if result == "sate" {
		fmt.Println("balikin duit gw")
	}

	operasi := "tambah"
	switch operasi {
	case "tambah":
		fmt.Println("hasil tambah")
	case "kurang":
		fmt.Println("hasil kurang")
	default:
		fmt.Println("operasi tidak ditemukan")
	}

	sayHello("Bandung", "davi")
	fmt.Println(selamatPagi("irsyad", 20))
	// fmt.Println(tambah(10, 20))
	total := tambah(10, 20)
	fmt.Println("Total : ", total)

	err := syntax.ShouldBeError()
	if err != nil {
		fmt.Println("Error terjadi : ", err)
	} else {
		fmt.Println("Tidak ada error")
	}
}

func sayHello(alamat string, name string) {
	fmt.Println("Hello " + name + " di " + alamat)
}
func selamatPagi(name string, umur int) string {
	return "Selamat Pagi " + name + ", umur " + fmt.Sprint(umur)
}
func tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func BeError() error {
	return fmt.Errorf("this is an error")
}
