package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	/*
		Pernyataan import:
		"{jalur-modul}/internal/models"
	*/
	"github.com/achmadalfanahsani/sippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql" // New import
)

/*
	Mendefinisikan struct application untuk menyimpan
	depedensi seluruh aplikasi untuk aplikasi web.
	Untuk saat ini kami hanya akan menyertakan bidang
	untuk dua pencatat, tetapi kami akan menambah lebih
	banyak lagi seiring pembangunan berlangsung.
*/
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	/*
		Tambahkan bidang cuplikan ke struct application.
		sehingga SnippetModel tersedia di penanganan.
	*/
	snippets *models.SnippetModel
}



func main() {
	/*
		Tentukan flag baris perintah baru dengan nama
		"addr", nilai default  ":4000" dan beberapa 
		teks bantuan singkat yang menjelaskan apa yang
		di kontrol flag. Nilai dari flag akan di simpan
		dalam variabel addr saat runtime 
	*/
	addr := flag.String("addr", ":4000", "HTTP network address")

	/*
		Menentukan flag baris perintah baru untuk string
		MySQL DSN.
	*/
	dsn := flag.String("dsn", "web:achmadalfanahsani@/snippetbox?parseTime=true", "MySQL data source name")

	/*
		Yang penting, kita menggunakan fungsi flag.Parse()
		untuk mengurai flag baris perintah. Ini membaca 
		nilai bendera baris perintah dan menugaskannya ke
		addr variabel. Anda perlu memanggil ini *Sebelum*
		Anda menggunakan variabel addr, selain itu akan 
		berisi nilai default ":4000". jika ada kesalahan
		ditemui selama penguraian aplikasi dihentikan.
	*/
	flag.Parse()

	/*
		Guanakan log.New() untuk membuat logger untuk 
		menulis pesan informasi. Ini membutuhkan tiga
		parameter: tujuan untuk menulis log ke (os.Stdout),
		sebuah string awalan untuk pesan (INFO diikuti
		dengan tab), dan flag untuk menunjukkan apa 
		informasi tambahan untuk disertakan (tanggal dan 
		waktu lokal). Perhatikan bahwa bendera digabungkan
		menggunakan operator bitwise OR |.
	*/
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	/*
		Buat logger untuk menulis pesan kesalahan dengan 
		cara yang sama, tetapi gunakan stderr sebagai 
		tujuan dan gunakan flag log.LshortFile untuk 
		masukkan yang relevan nama file dan nomor baris
	*/
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	/*
		Agar fungsi main() tetap rapi, saya telah memasukkan 
		kode untuk membuat koneksi kumpulkan ke dalam fungsi
		opeDB() terpisah di bawah ini. Kami melewati openDB()
		DSN dari flag baris perinth. 
	*/
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	/*
		Kami juga menunda panggilan ke db.Close(), sehingga
		kumpulan koneksi ditutup sebelum fungsi main() keluar
	*/
	defer db.Close()

	/*
		Initialize a new instance of our application struct, 
		containing the dependencies.
	*/
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		// inisialisasi instance models.SnippetModel
		snippets: &models.SnippetModel{DB: db},
	}

	/*
		Menginisialisasi struct http.Server baru. Kami
		mengatur bidang Addr dan Handler jadi bahwa 
		server menggunakan alamat dan rute jaringan 
		yang sama seperti sebelumnya,  dan atur kolom
		errorLog sehingga server sekarang menggunakan 
		errorLogger khusus jika ada masalah.
	*/
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		// Memanggil metode app.routes() untuk mendapatkan 
		//servermux yang berisi rute yang telah dibuat.
		Handler: app.routes(),
	}

	/*
		Nilai yang dikembalikan dari fungsi flag.String()
		adalah pointer ke flag nilai, bukan nilai itu
		sendiri. Jika kita perlu melakukan dereferensi 
		penunjuk (ex: awali dengan simbol *) sebelum
		menggunakannya. Perhatikan bahwa kami 
		menggunakan fungsi log.Printf() untuk 
		menginterpolasi dengan pesan log.
	*/

	// tulis pesan dengan 2 logger baru
	// bukan logger standar
	infoLog.Printf("Starting server on %s", *addr)

	/*
		Karena variable err sudah dideklarasikan dalam 
		kode di atas, kita perlu untuk menggunakan =
		di sini, bukan := "deklarasikan dan tetapkan"
		operator.
	*/
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

/*
	Fungsi openDB() membungkus sql.Open() dan mengembalikan
	kumpulan koneksi sql.DB untuk DSN tertentu.
*/
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}