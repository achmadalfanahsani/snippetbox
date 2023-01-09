package main

import (
	"log"
	"net/http"
)

// Mendefinisikan fungsi home handler yang menulis potongan 
// byte berisi "Hello from Sippetbox" sebagai respone body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	/*
		Gunakan fungsi http.NewServeMux() untuk 
		Menginisialisasi servermux baru, lalu
		daftarkan fungsi home sebagai penanganan
		pola URL "/"
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	/*
		Gunakan fungsi http.ListenAndServe() untuk memulai
		server web baru. Kami lewat 2 parameter: alamat
		jaringan TCP yang akan didengarkan (dalam hal ini
		":4000") dan servemux yang baru di buat. Jika 
		http.ListenAndServe() mengembalikan kesalahan 
		kami menggunakan fungsi log.Fatal() untuk mencatat
		pesan kesalahan dan keluar. Catatan bahwa kesalahan
		apapun yang dikembalikan oleh http.ListenAndServe()
		selalu bukan nil	
	*/
	log.Println("Starting server on : 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}