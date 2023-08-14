##  (README)

README ini memberikan gambaran tentang proyek `gobyexample-scraper`, yang mencakup aplikasi Go yang melakukan pengambilan konten dari situs Go By Example dan mengirimkannya melalui API HTTP sederhana.

### Struktur Proyek

Proyek terdiri dari dua paket utama: `service` dan `main`.

#### Paket `service`

Paket `service` berisi fungsi-fungsi untuk melakukan pengambilan konten dari situs Go By Example.

- `ScrapeMain() map[string]string`: Melakukan pengambilan pada halaman utama situs Go By Example untuk mengambil tautan-tautan ke contoh-contoh individual. Mengembalikan peta judul contoh dan URL mereka.

- `ScrapeSubMain(lq string) string`: Melakukan pengambilan pada subhalaman contoh tertentu menggunakan URL yang diberikan. Mengembalikan konten teks dari subhalaman tersebut.

#### Paket `main`

Paket `main` adalah titik masuk dari aplikasi dan mengatur server HTTP untuk melayani konten yang diambil.

- `mainRoute(w http.ResponseWriter, r *http.Request)`: Menangani rute utama API. Ini akan mengambil dan mengirimkan data yang telah di-cache atau melakukan pengambilan data baru dari situs Go By Example berdasarkan path yang diberikan.

- `main()`: Menginisialisasi server HTTP, mengatur rute, dan mulai mendengarkan pada port 8080.

### Cara Menjalankan

Untuk menjalankan aplikasi `gobyexample-scraper`, ikuti langkah-langkah berikut:

1. Pastikan Anda telah menginstal Go pada mesin Anda.

2. Klon repositori ini ke mesin lokal Anda.

3. Buka direktori utama proyek.

4. Jalankan perintah berikut untuk memulai aplikasi:

   ```sh
   go run main.go
   ```

5. Setelah aplikasi berjalan, Anda dapat mengakses API melalui web browser atau menggunakan alat seperti `curl`:
   - Untuk mendapatkan daftar judul contoh dan URL: [http://localhost:8080/](http://localhost:8080/)
   - Untuk mendapatkan konten dari contoh tertentu (ganti `{example_path}` dengan path contoh yang diinginkan): [http://localhost:8080/{example_path}](http://localhost:8080/{example_path})

### Catatan

- Aplikasi menyimpan data yang diambil dalam cache untuk meminimalkan jumlah permintaan ke situs Go By Example. Data yang di-cache disimpan di memori.
- Aplikasi melayani data yang diambil melalui API HTTP sederhana.
- README ini adalah gambaran dasar tentang proyek dan fungsionalitasnya. Anda mungkin perlu untuk lebih mendokumentasikan dan menyesuaikan kode sesuai kebutuhan Anda.

Jangan ragu untuk menghubungi kami jika perlu bantuan lebih lanjut atau penyempurnaan pada proyek ini!
