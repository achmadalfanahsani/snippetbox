package main

import(
	"fmt"
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
	
	w.Write([]byte("Hello from Snippetbox"))
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