package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Mendefinisikan fungsi home handler yang menulis potongan
// byte berisi "Hello from Sippetbox" sebagai respone body
func home(w http.ResponseWriter, r *http.Request) {
	/*
		Periksa apakah jalur URL permintaan saat ini sama persis dengan "/".
		Jika tidak, gunakan fungsi http.NotFound() untuk mengirim response
		404 ke klien. Yang penting kembalian (return)
	*/
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	/*
		Inisialisasi sebuah slice yang berisi path kedua file. Ini penting
		perhatikan bahwa file yang berisi template dasar kita harus 
		*Pertama* file dalam irisan.
	*/

	file := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	/*
		Gunakan fungsi template.ParseFiles() untuk membaca file template 
		menjadi set template. Jika ada kesalahan, kami mencatat pesan 
		kesalahan mendetail dan menggunakannya fungsi http.Error() untuk
		mengirim 500 Ineternal Server Error generik respon ke pengguna.
		Perhatikan bahwa kita dapat melewatkan potongan file path sebagai
		parameter variadik?
	*/
	ts, err := template.ParseFiles(file...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	/*
		Kami kemudian menggunakan metode Execute() pada set template
		untuk menulis konten template sebagai isi respon. Parameter 
		terakhir untuk Execute() mewakili data dinamis apapun yang
		kita berikan, yang untuk saat ini kita lakukan biarkan 
		kosong.

	*/
	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }	

	/*
		Gunakan metode ExecuteTemplate() untuk menulis kontent
		"base" template sebagai badan respons.
	*/
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

// Menambah fungsi handler snippetView
func snippetView(w http.ResponseWriter, r *http.Request) {
	/*
		Ekstrak nilai parameter id dari string kueri dan coba
		mengubah menjadi bilangan bulat menggunakan fungsi 
		strconv.Atoi().Jika tidak bisa dikonversi menjadi 
		bilangan bulat, atau nilainya kurang dari 1, kami
		mengembalikan halaman 404 (not found response)
	*/
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	/*
		Gunakan fungsi fmt.Printf() untuk menginterpolasi 
		nilai id dengan respon kita dan tulis ke 
		http.ResponseWriter 
	*/
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Menambah fungsi handler snippetCreate
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	/*
		Gunakan r.Method untuk memeriksa apakah request 
		menggunakan POST atau tidak. 
	*/
	if r.Method != "POST" {
		/*
			Jika tidak, gunakan metode w.writeHeader() untuk
			mengirim pesan status 405 kode dan metode 
			w.Write() untuk menulis "Method not allowed"
			badan tanggapan. Kemudian kami menulis return
			sehingga kode selanjutnya tidak dieksekusi.

			Gunakan metode Header().Set() untuk menambahkan
			header "Allow : POST" ke peta untuk tajuk
			response. Parameter pertama adalah nama header,
			dan parameter kedua adalah nilai header. 
		*/

		w.Header().Set("Allow", "POST")

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		// or
		
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		
		return
	}

	w.Write([]byte("Create a new snippet...."))
}