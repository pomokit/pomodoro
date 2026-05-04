# pomodoro kit

![image](https://github.com/bukped/pomodoro/assets/11188109/6488ed71-f5fb-459f-8d52-6e839c1dcf22)

![Ilustrasi Pomodoro](./pomodoro_illustration.png)

Command line [pomodoro timer](https://en.wikipedia.org/wiki/Pomodoro_Technique), implemented in Go.

## Installation
Go to [Release Page](https://github.com/pomokit/pomodoro/releases)

## Usage

Aplikasi ini berjalan secara interaktif melalui *command line* (terminal). Berikut cara penggunaannya:

1. Jalankan file *executable* aplikasi (contoh: `./pomodoro` atau `pomodoro.exe`).
2. *(Hanya untuk penggunaan pertama kali)* Akan muncul QR Code WhatsApp di terminal. Lakukan pemindaian (scan) menggunakan aplikasi WhatsApp di ponsel Anda untuk menghubungkan akun.
3. Masukkan ID WhatsApp Group tujuan untuk pelaporan. **(Tips: Ketik `list` lalu Enter untuk menampilkan semua Grup WhatsApp Anda beserta ID-nya)**.
4. Masukkan URL Github terkait dengan tugas yang sedang dikerjakan.
5. Masukkan nama/deskripsi Milestone dari tugas Anda.
6. Timer akan mulai berjalan secara otomatis.

**Alur Aplikasi:**

**Infografis Siklus Pomodoro:**
```mermaid
graph TD
    Start([Mulai Sesi]) --> T1[Task 1: 25 Menit]
    T1 --> B1(Break: 5 Menit)
    B1 --> T2[Task 2: 25 Menit]
    T2 --> B2(Break: 5 Menit)
    B2 --> T3[Task 3: 25 Menit]
    T3 --> B3(Break: 5 Menit)
    B3 --> T4[Task 4: 25 Menit]
    T4 --> LB([Long Break: 25 Menit])
    LB --> End([Selesai 1 Siklus & Lapor WA])
    
    classDef task fill:#ff6b6b,stroke:#c0392b,stroke-width:2px,color:#fff;
    classDef break fill:#4ecdc4,stroke:#16a085,stroke-width:2px,color:#fff;
    classDef startend fill:#2c3e50,stroke:#34495e,stroke-width:2px,color:#fff;
    
    class T1,T2,T3,T4 task;
    class B1,B2,B3,LB break;
    class Start,End startend;
```

- **Siklus Pomodoro:** 1 Sesi (cycle) terdiri dari **4 siklus kerja (masing-masing 25 menit) yang diselingi istirahat pendek (5 menit), dan diakhiri dengan 1 Istirahat Panjang (25 menit)**.
- **Notifikasi Awal:** Saat aplikasi mulai, sebuah pesan notifikasi akan dikirimkan ke WhatsApp Group yang telah ditentukan.
- **Tangkapan Layar Otomatis:** Selama timer pomodoro berjalan, aplikasi akan mengambil beberapa tangkapan layar (*screenshot*) secara *background*.
- **Pelaporan Akhir:** Setelah 1 sesi penuh selesai, aplikasi akan memilih satu *screenshot* secara acak dan mengirimkannya sebagai laporan progres ke WhatsApp Group, beserta informasi *milestone* dan tautan laporan.
- **Notifikasi Pergantian Waktu:** Terdapat peringatan berupa notifikasi desktop atau *beep* setiap kali sesi tugas berakhir dan memasuki waktu istirahat, serta sebaliknya.

## Dev & Maintenance

### Pengembangan Lokal (Local Development)
Untuk mengembangkan atau memodifikasi kode aplikasi ini:
1. Pastikan Golang sudah terinstal di sistem Anda.
2. Clone repositori ini ke komputer lokal.
3. Jalankan `go mod tidy` di terminal untuk mengunduh dan merapikan semua *dependencies* (termasuk modul whatsmeow untuk WhatsApp).
4. Jalankan aplikasi untuk diuji coba menggunakan perintah `go run .`
5. Untuk membuat file binary (*build*), jalankan `go build -o pomodoro .` (atau `pomodoro.exe` di Windows).

### Rilis Versi (Release)
Untuk merilis versi baru, Anda bisa menggunakan perintah *cross-compilation* dan melakukan *tagging* di Git. Berikut adalah contoh perintah rilis menggunakan PowerShell:

```sh
$env:GOOS = 'linux'
$env:GOOS = 'windows'
$env:GOOS = 'darwin'
$env:CGO_ENABLED = '1'

go mod tidy
git tag                                 # melihat versi saat ini
git tag v0.0.3                          # menetapkan versi tag baru
git push origin --tags                  # push tag ke repository untuk trigger rilis
```
