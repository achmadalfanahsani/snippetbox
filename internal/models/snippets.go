package models

import (
	"database/sql"
	"errors"
	"time"
)

/*
	Tentukan jenis cuplikan menyimpan data untuk
	cuplikan individual. Perhatikan caranya bidang
	struct sesui dengan bidang di cuplikan table
	MySQL yang telah di buat.
*/
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

/*
	Tentukan jenis SnippetModel yang membungkus 
	kumpulan koneksi sql.DB
*/
type SnippetModel struct {
	DB *sql.DB
}

// Ini akan memasukkan potongan baru ke dalam database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Tulis pernyataan SQL yang ingin kita jalankan.
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	/*
		Menggunakan metode Exec() pada koneksi tersemat 
		untuk mengeksekusi pernyataan. Parameter pertama
		adalah pernyataan SQL, diikuti oleh nilai title,
		content, expires untuk parameter placeholder.
	*/
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	/*
		Menggunakan motode LastInsertId() pada result untuk 
		mendapatkan ID dari record yang baru disisipkan
		di table snippet.
	*/
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}


	return int(id), nil
}

// Ini akan mengembalikan potongan spesifik berdasarkan idnya
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// Query SQL
	stmt := `SELECT id, title, content, created, expires from snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	//  QueryRow() digunakan untuk mengembalikan 1 baris data
	row := m.DB.QueryRow(stmt, id)

	// inisialisasi struct sebagai penampung data
	s := &Snippet{}

	// Scan() digunakan untuk membaca row data
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// errors.Is() digunakan untuk penanganan error khusus
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// kembalikan objek jika semua telah berjalan dengan baik
	return s, nil
}

// Ini akan mengembalikan 10 snippet yang paling baru di buat
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM
	snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Get row data lebih dari satu
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	// row.Next() digunakan untuk melakukan iterasi hasil queri
	// data lebih dari 1 dan akan menutup perulangan secara 
	// otomatis
	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}