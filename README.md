# Portofolio Leon

Website portofolio pribadi, dibangun dengan **Go + [templ](https://templ.guide) + [HTMX](https://htmx.org)** dan Tailwind CSS. Deploy di Vercel (Go Function).

## Struktur

```
api/index.go          entry point Vercel Function
cmd/dev/main.go       server lokal (port 3000)
internal/app/         routing
internal/data/        data project, tools, pengalaman
internal/views/       komponen .templ (+ hasil generate *_templ.go)
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
