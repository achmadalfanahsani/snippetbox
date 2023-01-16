package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

/*
	Pembantu serverError menulis pesan kesalahan
	dan pelacakan tumpukan ke errorLog, kemudian
	mengirimkan response 500 Internal Server
	Error generik kepada pengguna.
*/
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/*
	Pembantu clientError mengirimkan kode status spesifik 
	dan deskripsi terkait kepada pengguna. Kami akan 
	mengembalikan ini nanti di buku untuk mengirimkan
	tanggapan seperti 400 "Bad Request" ketika ada 
	masalah dengan permintaan yang dikirim pengguna  
*/
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

/*
	Untuk konsistensi, kami juga akan mengimplementasikan 
	helper notFound. Ini hanyalah sebuah pembungkus 
	kenyamanan di sekitar clientError yang mengirimkan 
	response 404 "Not Found"  kepada pengguna.
*/
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}