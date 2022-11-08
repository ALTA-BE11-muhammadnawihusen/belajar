package main

import (
	"fmt"
	"io"
	"os"
)

func readonly(path string) {
	// O_RDONLY maksudnya adalah read only
	// parameter ketiga adalah permission
	file, er := os.OpenFile(path, os.O_RDONLY, 0664)
	defer file.Close()
	if isError(er) {
		return
	}
	//baca file dengan membuat variabel penyimpanan terlebih dahulu
	text := make([]byte, 1024)

	// isi penyimpanan dengan isi file
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(string(text))
}

func readwrite(path string, isi string) {
	// 0644 berarti diizinkan melakukan read dan write
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	defer file.Close()
	if isError(err) {
		return
	}

	// tulis data ke file
	_, errf := file.WriteString(isi)
	if isError(errf) {
		return
	}
	// _, errf := file.WriteString(isi) dapat melakukan berkali kali tergantung kebutuhan
	// if isError(errf){return}

	// simpan perubahan
	errs := file.Sync()
	if isError(errs) {
		return
	}
}

func delete(path string) {
	err := os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("File berhasil di delete")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}

// hanya untuk membuat file bukan folder
// ingat hanya untuk membuat file
func MakeFile(Path string) {
	// deteksi file
	test1, test2 := os.Stat(Path)          // jika path sudah ada maka test2 akan menghasilkan error
	fmt.Println("check isi test 1", test1) // jika tidak ada akan menghasilkan nil
	fmt.Println("check isi test 2", test2)

	if os.IsNotExist(test2) {
		file, err := os.Create(Path) // jika file berhasil di buat maka err == nil
		defer file.Close()           // file yang dibuat auto open jadi jangan lupa di close
		if isError(err) {
			return
		}
	} else if test2 == nil {
		fmt.Println("File sudah ada")
		return
	}

	fmt.Println("File berhasil di buat")
}

func main() {
	Path := "/root/git/mystudy"
	// var isi string
	//jika ingin panjang jangan gunakan scan karena scan hanya akan mengambil satu kata
	// fmt.Scan(&isi)
	MakeFile(Path + "/web/temp.html") // membuat file
	// readwrite(Path+"/File/main.txt", isi) menambahkan isi file
	// readonly(Path + "/File/main.txt") read file
	// delete(Path + "/File/mai.txt") delete file
}
