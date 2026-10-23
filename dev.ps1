# Run the portfolio locally at http://localhost:3000.
# The Go server runs inside Docker because Windows Smart App Control blocks
# locally built .exe files. Needs Docker Desktop open. Stop with Ctrl+C.

docker info *> $null
if ($LASTEXITCODE -ne 0) {
    Write-Host "Docker Desktop belum jalan. Buka Docker Desktop dulu, tunggu 'Engine running', lalu jalankan lagi." -ForegroundColor Red
    exit 1
}

templ generate
if ($LASTEXITCODE -ne 0) { exit 1 }

.\bin\tailwindcss.exe -i styles/input.css -o public/styles.css --minify
if ($LASTEXITCODE -ne 0) { exit 1 }

Write-Host "Buka http://localhost:3000 setelah muncul 'listening'. Stop: Ctrl+C" -ForegroundColor Cyan
docker run --rm -it -p 3000:3000 -v "${PWD}:/app" -w /app golang:1.25 go run ./cmd/dev
