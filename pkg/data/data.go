package data

type Project struct {
	ID          int
	Title       string
	Description string
	Image       string
	Link        string
	// Category, Role, Year and Stack come from the matching Pengalaman entry
	// (or the README for the portfolio itself); empty means unknown.
	Category string
	Role     string
	Year     string
	Stack    []string
}

// Projects are ordered newest first; the first two are badged "Recent".
var Projects = []Project{
	{ID: 1, Title: "Website K-EXPO 2026",
		Description: "Portal digital Kemahasiswaan Expo (K-EXPO) ITERA 2026 bertema space exploration. Saya bertanggung jawab sebagai kepala Divisi Implementasi Teknologi.",
		Image:       "/image/kexpo26.png",
		Link:        "",
		Category:    "Event Website", Role: "Kepala Divisi Implementasi Teknologi", Year: "2026",
		Stack: []string{"Next.js", "TypeScript", "Tailwind CSS", "Go", "Gin", "GORM", "PostgreSQL", "Google OAuth"}},
	{ID: 2, Title: "Website PPLK 2026",
		Description: "Website dan game untuk Program Pengenalan Lingkungan Kampus (PPLK) ITERA 2026. Saya bertanggung jawab sebagai kepala sub-divisi back-end: mengatur aliran data, mengelola database, serta memastikan interaksi server dan klien berjalan lancar.",
		Image:       "/image/pplk26.png",
		Link:        "",
		Category:    "Event Website", Role: "Kepala Sub-Divisi Back-End", Year: "2026",
		Stack: []string{"PHP", "Laravel", "MySQL"}},
	{ID: 3, Title: "Website Portofolio",
		Description: "Website portofolio pribadi yang dibuat untuk personal branding saya.",
		Image:       "/image/porto.png",
		Link:        "",
		Category:    "Personal", Role: "Solo", Year: "2025", Stack: []string{"Go", "templ", "HTMX", "Tailwind CSS"}},
	{ID: 4, Title: "Website IGTTPB 2025",
		Description: "Website ini dibuat untuk mencari kelompok, absensi kehadiran peserta, dan Pengumpulan tugas. Saya bertanggung jawab sebagai kepala divisi IT.",
		Image:       "/image/igttpb25.png",
		Link:        "https://www.igttpb.site/",
		Category:    "Event Website", Role: "Kepala Divisi IT", Year: "2025", Stack: []string{"Next.js"}},
	{ID: 5, Title: "Website PPLK 2025",
		Description: "Website ini dibuat untuk Program Pengenalan Lingkungan Kampus (PPLK) Mahasiswa Baru 2025. Saya bertanggung jawab di bidang front-end, yaitu bagian Profile, dan Detail UPA.",
		Image:       "/image/pplk25.png",
		Link:        "",
		Category:    "Event Website", Role: "Front-End", Year: "2025", Stack: []string{"React"}},
	{ID: 6, Title: "Website Pemilihan Ketua Angkatan",
		Description: "Website ini dibuat sebagai sarana pemilihan Ketua Angkatan secara e-voting yang aman, transparan, dan efisien. Saya bertanggung jawab di bidang front-end (halaman utama).",
		Image:       "/image/bumiketupat.png",
		Link:        "https://www.algovista24.site/",
		Category:    "E-Voting", Role: "Front-End", Year: "2025"},
}

// Tech is one stack item; Icon is the file name in public/icons (skillicons.dev).
type Tech struct {
	Name string
	Icon string
}

type StackGroup struct {
	Label string
	Items []Tech
}

var Stack = []StackGroup{
	{Label: "Languages", Items: []Tech{
		{"HTML", "html"}, {"CSS", "css"}, {"JavaScript", "js"}, {"TypeScript", "ts"}, {"PHP", "php"},
		{"Python", "py"}, {"C++", "cpp"}, {"Java", "java"}, {"Go", "go"}, {"Sass", "sass"},
	}},
	{Label: "Frameworks & Libraries", Items: []Tech{
		{"Bootstrap", "bootstrap"}, {"Tailwind CSS", "tailwind"}, {"React", "react"}, {"Next.js", "nextjs"},
		{"Laravel", "laravel"}, {"Prisma", "prisma"}, {"Node.js", "nodejs"},
	}},
	{Label: "Databases", Items: []Tech{
		{"MySQL", "mysql"}, {"PostgreSQL", "postgres"}, {"MongoDB", "mongodb"},
	}},
	{Label: "Tools & Platforms", Items: []Tech{
		{"VS Code", "vscode"}, {"Vite", "vite"}, {"Vercel", "vercel"}, {"Figma", "figma"}, {"Docker", "docker"},
		{"Git", "git"}, {"GitHub", "github"}, {"npm", "npm"}, {"Postman", "postman"},
	}},
}

// Repeat returns items repeated until there are at least min of them, so a
// short group still fills a marquee row.
func Repeat(items []Tech, min int) []Tech {
	out := items
	for len(out) < min {
		out = append(out, items...)
	}
	return out
}

// TabCount is the number of entries in a tab.
func TabCount(tab string) int { return len(DataByTab[tab]) }

type RundownItem struct {
	ID    int
	Date  string
	About string
	IsIT  bool
	Role  string
	Event string
	Desc  string
}

type Tab struct {
	Key   string
	Label string
}

var Tabs = []Tab{
	{Key: "kepanitiaan", Label: "Kepanitiaan"},
	{Key: "magang", Label: "Magang & Kerja"},
	{Key: "organisasi", Label: "Organisasi"},
	{Key: "prestasi", Label: "Prestasi & Pelatihan"},
}

var DataByTab = map[string][]RundownItem{
	"kepanitiaan": {
		{ID: 1, Date: "Sep 2024", About: "First Meet Teknik Informatika 24, Staff Divisi Transportasi", IsIT: false, Role: "Staff Divisi Transportasi", Event: "First Meet Teknik Informatika 24", Desc: "Mengoordinasikan kebutuhan transportasi peserta maupun panitia selama kegiatan First Meet berlangsung."},
		{ID: 2, Date: "Feb 2025 – Mar 2025", About: "BUMI KETUPAT, Staff Divisi IT", IsIT: true, Role: "Staff Divisi IT", Event: "Buat Milih Ketua Dua Empat (BUMI KETUPAT)", Desc: "Mengembangkan tampilan front-end website sebagai media pelaksanaan pemilihan ketua angkatan berbasis digital."},
		{ID: 3, Date: "Jul 2025", About: "Penyambutan Maba Teknik Informatika 25, Mentor Kelompok", IsIT: false, Role: "Mentor Kelompok", Event: "Penyambutan Maba Teknik Informatika 25", Desc: "Membimbing dan mendampingi kelompok mahasiswa baru dalam mengikuti seluruh rangkaian kegiatan penyambutan."},
		{ID: 4, Date: "Jun 2025 – Aug 2025", About: "PPLK 2025, Staff Front-End Divisi IMTEK", IsIT: true, Role: "Staff Front-End Divisi IMTEK", Event: "PPLK ITERA 2025", Desc: "Mengembangkan tampilan website PPLK kegiatan menggunakan framework ReactJS."},
		{ID: 5, Date: "Sep 2025 – Oct 2025", About: "IGTTPB 2025, Kepala Divisi IT", IsIT: true, Role: "Kepala Divisi IT", Event: "Informatics Goes to TPB (IGTTPB) 2025", Desc: "Berperan sebagai project manager divisi IT sekaligus backup staff dan memastikan website berjalan dengan baik. Framework yang digunakan adalah NextJS."},
		{ID: 6, Date: "Oct 2025", About: "PODIUM ITERA, Staff Content Research Divisi IMTEK", IsIT: true, Role: "Staff Content Research Divisi IMTEK", Event: "PODIUM ITERA", Desc: "Melakukan riset dan pengumpulan informasi sebagai bahan konten yang ditampilkan pada website kegiatan."},
		{ID: 7, Date: "Aug 2025 – Nov 2025", About: "Informatics Festival 2025, Staff Sponsorship", IsIT: false, Role: "Staff Sponsorship", Event: "Informatics Festival 2025", Desc: "Mencari, mengajukan proposal, dan melakukan follow up kepada calon sponsor untuk mendukung kegiatan."},
		{ID: 8, Date: "Mar 2026 – May 2026", About: "Paskah KMK ITERA 2026, Wakil Ketua Pelaksana", IsIT: false, Role: "Wakil Ketua Pelaksana", Event: "Paskah KMK ITERA 2026", Desc: "Membantu ketua pelaksana dalam mengoordinasikan seluruh rangkaian kegiatan dan menjadi backup dalam pengambilan keputusan teknis acara."},
		{ID: 9, Date: "Apr 2026 – Sekarang", About: "LDOP Arithmatic 8.0, Staff Komisi Disiplin", IsIT: false, Role: "Staff Komisi Disiplin", Event: "LDOP Arithmatic 8.0", Desc: "Mengawasi dan menegakkan ketertiban peserta selama kegiatan berlangsung sesuai SOP yang telah ditetapkan."},
		{ID: 10, Date: "Apr 2026 – Sekarang", About: "Serasehan HMIF ITERA, Staff Manajemen Acara", IsIT: false, Role: "Staff Manajemen Acara", Event: "Serasehan HMIF ITERA", Desc: "Menyusun rundown acara dan SOP sebagai panduan teknis pelaksanaan kegiatan."},
		{ID: 11, Date: "Apr 2026 – Sekarang", About: "Point Project 4.0 HMIF ITERA, Kepala Sub-Divisi Time Keeper", IsIT: false, Role: "Kepala Sub-Divisi Time Keeper", Event: "Point Project 4.0 HMIF ITERA", Desc: "Menjadi kepala sub-divisi time keeper yang memastikan ketepatan waktu seluruh rangkaian acara sesuai rundown yang telah ditetapkan."},
		{ID: 12, Date: "Aug 2026 – Sekarang", About: "PPLK ITERA, Kepala Sub-Divisi Back-End Divisi IMTEK", IsIT: true, Role: "Kepala Sub-Divisi Back-End Divisi IMTEK", Event: "PPLK ITERA 2026", Desc: "Bertanggung jawab dalam mengatur aliran data, mengelola database, serta memastikan interaksi server dan klien berjalan lancar pada website dan game PPLK 2026."},
		{ID: 13, Date: "Aug 2026 – Sekarang", About: "Kemahasiswaan Expo 2026, Kepala Divisi Implementasi Teknologi", IsIT: true, Role: "Kepala Divisi Implementasi Teknologi", Event: "Kemahasiswaan Expo 2026", Desc: "Memimpin Divisi Implementasi Teknologi dalam kepanitiaan Kemahasiswaan Expo 2026."},
		{ID: 14, Date: "Oct 2026 – Sekarang", About: "Neofest 2026, Staff Back-End", IsIT: true, Role: "Staff Back-End", Event: "Neofest 2026", Desc: "Menjadi staff back-end dalam pengembangan website kegiatan Neofest 2026."},
		{ID: 15, Date: "Oct 2026 – Sekarang", About: "Pemilihan Raya KM ITERA 2026, Kepala Divisi", IsIT: false, Role: "Kepala Divisi", Event: "Pemilihan Raya KM ITERA 2026", Desc: "Menjadi kepala divisi dalam kepanitiaan Pemilihan Raya KM ITERA 2026."},
	},
	"magang": {
		{ID: 1, Date: "Sep 2025 – Dec 2025", About: "Asisten Praktikum, Dasar Teknologi Digital (TPB 21)", IsIT: false, Role: "Asisten Praktikum – TPB 21", Event: "Dasar Teknologi Digital, ITERA", Desc: "Menjadi asisten praktikum mata kuliah dasar teknologi digital untuk TPB 21."},
		{ID: 2, Date: "Sep 2025 – Dec 2025", About: "Staff Magang, KM ITERA", IsIT: false, Role: "Staff Magang", Event: "Kementerian Teknologi Informasi KM ITERA (Kabinet Resonara)", Desc: "Menjadi salah satu anggota magang pada Teknologi Informasi KM Kabinet Resonara."},
		{ID: 3, Date: "Feb 2026 – Jun 2026", About: "Asisten Praktikum, Pengenalan Komputasi (T 11)", IsIT: false, Role: "Asisten Praktikum – T 11", Event: "Pengenalan Komputasi, ITERA", Desc: "Menjadi asisten praktikum mata kuliah pengenalan komputasi untuk TPB kelas T 11."},
		{ID: 5, Date: "Jun 2026 – Jul 2026", About: "Magang Supervisor, Chandramart", IsIT: false, Role: "Magang – Supervisor", Event: "Chandramart", Desc: "Magang di bidang retail supermarket pada posisi supervisor."},
		{ID: 4, Date: "Sep 2026 – Sekarang", About: "Asisten Praktikum, Dasar Teknologi Digital (TPB 43)", IsIT: false, Role: "Asisten Praktikum – TPB 43", Event: "Dasar Teknologi Digital, ITERA", Desc: "Kembali menjadi asisten praktikum mata kuliah dasar teknologi digital untuk TPB 43."},
	},
	"organisasi": {
		{ID: 1, Date: "Jun 2025 – Sekarang", About: "Anggota, Badan Pengurus Angkatan Algovista", IsIT: false, Role: "Anggota – Divisi IT", Event: "Badan Pengurus Angkatan (BPA) Algovista – IF 24", Desc: "Mengorganisir Divisi IT dalam Kepengurusan Angkatan."},
		{ID: 2, Date: "Jan 2026 – Sekarang", About: "Anggota Muda, Himpunan Mahasiswa Informatika ITERA", IsIT: false, Role: "Anggota Muda – Divisi Technopreneur", Event: "Himpunan Mahasiswa Informatika (HMIF) ITERA", Desc: "Mengorganisir Divisi Technopreneur dalam Departemen Keprofesian."},
		{ID: 3, Date: "Jan 2026 – Sekarang", About: "Anggota, Keluarga Mahasiswa Katholik St. Thomas Aquinas ITERA", IsIT: false, Role: "Anggota – Divisi Implementasi Teknologi", Event: "Keluarga Mahasiswa Katholik (KMK) St. Thomas Aquinas ITERA", Desc: "Mengorganisir Divisi Implementasi Teknologi dalam Departemen Media Komunikasi Visual."},
		{ID: 4, Date: "Oct 2026 – Sekarang", About: "Humas Informatika, Divisi Website (Batch 3)", IsIT: false, Role: "Adminfot – Divisi Website", Event: "Humas Informatika ITERA (Batch 3)", Desc: "Bergabung di Divisi Website Humas Informatika batch 3."},
	},
	"prestasi": {
		{ID: 1, Date: "Aug 2024", About: "Peserta Pelatihan LKMM Pra-TD I, ITERA", IsIT: false, Role: "Peserta", Event: "Pelatihan LKMM Pra-TD I ITERA", Desc: "Mengikuti pelatihan kepemimpinan dan manajemen mahasiswa tingkat pra-dasar untuk mengembangkan jiwa kepemimpinan."},
		{ID: 2, Date: "Nov 2024", About: "Peserta Mini Bootcamp 1.0 HMIF, ITERA", IsIT: false, Role: "Peserta", Event: "Mini Bootcamp 1.0 HMIF ITERA", Desc: "Mengikuti mini bootcamp yang diselenggarakan oleh HMIF ITERA dengan mengerjakan, merilis, dan showcase sebuah website sederhana sebagai hasil akhir dari rangkaian bootcamp."},
		{ID: 3, Date: "Aug 2025", About: "Volunteer HMIF Mengabdi Vol. 3, ITERA", IsIT: false, Role: "Volunteer", Event: "HMIF Mengabdi Vol. 3", Desc: "Berpartisipasi sebagai volunteer dalam kegiatan pengabdian masyarakat yang diselenggarakan oleh HMIF ITERA."},
		{ID: 4, Date: "May 2026", About: "Peserta Pelatihan LKMM TD VIII, ITERA", IsIT: false, Role: "Peserta", Event: "Pelatihan LKMM TD VIII ITERA", Desc: "Mengikuti pelatihan kepemimpinan dan manajemen mahasiswa tingkat dasar untuk mengembangkan jiwa kepemimpinan."},
	},
}

// Items returns the tab's items, oldest at the bottom (newest first),
// optionally filtered to IT-related ones.
func Items(tab string, itOnly bool) []RundownItem {
	src := DataByTab[tab]
	out := make([]RundownItem, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		if itOnly && !src[i].IsIT {
			continue
		}
		out = append(out, src[i])
	}
	return out
}

func FindItem(tab string, id int) (RundownItem, bool) {
	for _, it := range DataByTab[tab] {
		if it.ID == id {
			return it, true
		}
	}
	return RundownItem{}, false
}

func ValidTab(tab string) bool {
	_, ok := DataByTab[tab]
	return ok
}
