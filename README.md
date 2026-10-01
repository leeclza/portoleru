# Portofolio Leon

Website portofolio pribadi, dibangun dengan **Go + [templ](https://templ.guide) + [HTMX](https://htmx.org)** dan Tailwind CSS. Deploy di Vercel (Go Function).

## Tentang Website

Website ini adalah portofolio pribadi **Christopher Leon Saputra**, mahasiswa Teknik Informatika Institut Teknologi Sumatera (ITERA). Isinya perkenalan diri, project yang pernah dikerjakan, tools/skill yang dikuasai, riwayat pengalaman (kepanitiaan, magang & kerja, organisasi, prestasi & pelatihan), dan kontak media sosial. Tujuannya sebagai media *personal branding* yang bisa dilihat oleh rekruter, dosen, maupun rekan organisasi.

## SKPL Singkat

Spesifikasi Kebutuhan Perangkat Lunak (ringkas).

### Tujuan
Menyediakan halaman web yang menampilkan profil, karya, dan pengalaman Leon secara ringkas, menarik, dan mudah diakses dari perangkat apa pun.

### Pengguna
| Pengguna | Kebutuhan |
|---|---|
| Pengunjung (rekruter, dosen, teman) | Melihat profil, project, skill, pengalaman, dan menghubungi pemilik |
| Pemilik (Leon) | Memperbarui konten (data project, tools, pengalaman) lewat kode |

### Kebutuhan Fungsional
| Kode | Kebutuhan |
|---|---|
| F-01 | Sistem menampilkan halaman utama berisi intro, foto profil, daftar project, tools, dan kontak. |
| F-02 | Sistem menampilkan 2 project pertama dan tombol **Show More / Show Less** untuk menampilkan semua project. |
| F-03 | Setiap project menampilkan gambar, judul, deskripsi, dan link (jika ada). |
| F-04 | Sistem menampilkan halaman **Pengalaman** dalam bentuk timeline, urut dari yang terbaru. |
| F-05 | Pengunjung dapat memilih kategori pengalaman: Kepanitiaan, Magang & Kerja, Organisasi, Prestasi & Pelatihan. |
| F-06 | Pada kategori Kepanitiaan, pengunjung dapat memfilter pengalaman di bidang **IT**. |
| F-07 | Pengunjung dapat mengklik item pengalaman untuk melihat detail (peran, kegiatan, tanggal, deskripsi) dalam modal. |
| F-08 | Sistem menyediakan link ke LinkedIn, Instagram, dan GitHub pemilik. |
| F-09 | Kategori dan filter yang dipilih tersimpan di URL, sehingga bisa di-refresh atau dibagikan. |

### Kebutuhan Non-Fungsional
| Kode | Kebutuhan |
|---|---|
| NF-01 | **Responsif**: tampilan menyesuaikan layar HP, tablet, dan desktop. |
| NF-02 | **Performa**: halaman dirender di server (SSR) dan hanya memuat JS kecil (HTMX), sehingga cepat dibuka. |
| NF-03 | **Ketersediaan**: di-host di Vercel dan dapat diakses 24/7. |
| NF-04 | **Aksesibilitas**: gambar memiliki teks alternatif, tombol memiliki label, dan animasi dinonaktifkan jika pengguna memilih *reduced motion*. |
| NF-05 | **Pemeliharaan**: konten dipisah di `pkg/data` sehingga mudah diperbarui tanpa mengubah tampilan. |

### Batasan
- Belum ada panel admin, sehingga konten diperbarui langsung lewat kode.
- Belum menggunakan database; [Neon](https://neon.tech) (PostgreSQL) disiapkan jika nantinya dibutuhkan.

## Struktur

```
api/index.go          entry point Vercel Function
cmd/dev/main.go       server lokal (port 3000)
pkg/app/              routing
pkg/data/             data project, tools, pengalaman
pkg/views/            komponen .templ (+ hasil generate *_templ.go)
styles/input.css      sumber Tailwind + animasi
public/               file statis (gambar, favicon, styles.css hasil build)
```

## Setup

1. Install Go 1.25+.
2. Install templ: `go install github.com/a-h/templ/cmd/templ@latest`
3. Download [Tailwind CLI standalone](https://github.com/tailwindlabs/tailwindcss/releases/latest) ke `bin/tailwindcss.exe` (Windows) atau `bin/tailwindcss`.

## Development

```sh
templ generate                                                  # .templ -> *_templ.go
./bin/tailwindcss -i styles/input.css -o public/styles.css --minify
go build -o bin/dev.exe ./cmd/dev && ./bin/dev.exe              # http://localhost:3000
```

Mode watch (jalankan di terminal terpisah):

```sh
templ generate --watch
./bin/tailwindcss -i styles/input.css -o public/styles.css --watch
```

> Hasil `templ generate` (`*_templ.go`) dan `public/styles.css` **ikut di-commit**, karena Vercel cuma meng-compile Go tanpa menjalankan templ/Tailwind. Jalankan dua perintah generate di atas sebelum commit.

## Deploy

Push ke GitHub, import repo di Vercel (Framework Preset: **Other**). `vercel.json` sudah mengatur:
- file di `public/` disajikan sebagai file statis,
- request lain diarahkan ke `api/index.go`.
