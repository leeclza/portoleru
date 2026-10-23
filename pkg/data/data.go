package data

type Project struct {
	ID          int
	Title       string
	Description string
	Image       string
	Link        string
	// Category, Role, Year and Stack come from the matching experience entry
	// (or the README for the portfolio itself); empty means unknown.
	Category string
	Role     string
	Year     string
	Stack    []string
}

// Projects are ordered newest first; the first two are badged "Recent".
var Projects = []Project{
	{ID: 1, Title: "Website K-EXPO 2026",
		Description: "A space-exploration themed digital portal for Kemahasiswaan Expo (K-EXPO) ITERA 2026. I led the Technology Implementation Division.",
		Image:       "/image/kexpo26.png",
		Link:        "",
		Category:    "Event Website", Role: "Head of Technology Implementation", Year: "2026",
		Stack: []string{"Next.js", "TypeScript", "Tailwind CSS", "Go", "Gin", "GORM", "PostgreSQL", "Google OAuth"}},
	{ID: 2, Title: "Website PPLK 2026",
		Description: "Website and game for ITERA's 2026 campus orientation program (PPLK). As back-end sub-division lead, I handled the data flow, managed the database, and made sure the server and client worked together smoothly.",
		Image:       "/image/pplk26.png",
		Link:        "",
		Category:    "Event Website", Role: "Back-End Sub-Division Lead", Year: "2026",
		Stack: []string{"PHP", "Laravel", "MySQL"}},
	{ID: 3, Title: "Website Portofolio",
		Description: "My personal portfolio website, built for personal branding.",
		Image:       "/image/porto.png",
		Link:        "",
		Category:    "Personal", Role: "Solo", Year: "2025", Stack: []string{"Go", "templ", "HTMX", "Tailwind CSS"}},
	{ID: 4, Title: "Website IGTTPB 2025",
		Description: "A website for finding groups, tracking participant attendance, and submitting assignments. I was Head of the IT Division.",
		Image:       "/image/igttpb25.png",
		Link:        "https://www.igttpb.site/",
		Category:    "Event Website", Role: "Head of IT Division", Year: "2025", Stack: []string{"Next.js"}},
	{ID: 5, Title: "Website PPLK 2025",
		Description: "Website for ITERA's 2025 new student orientation program (PPLK). I worked on the front-end, building the Profile and UPA Detail pages.",
		Image:       "/image/pplk25.png",
		Link:        "",
		Category:    "Event Website", Role: "Front-End", Year: "2025", Stack: []string{"React"}},
	{ID: 6, Title: "Website Pemilihan Ketua Angkatan",
		Description: "A secure, transparent, and efficient e-voting website for electing the class leader. I worked on the front-end (home page).",
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

// Ongoing marks a date range that is still running, e.g. "Aug 2026 – Present".
const Ongoing = "Present"

type RundownItem struct {
	ID    int
	Date  string
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
	{Key: "kepanitiaan", Label: "Committees"},
	{Key: "magang", Label: "Internships & Work"},
	{Key: "organisasi", Label: "Organizations"},
	{Key: "prestasi", Label: "Achievements & Training"},
}

var DataByTab = map[string][]RundownItem{
	"kepanitiaan": {
		{ID: 1, Date: "Sep 2024", IsIT: false, Role: "Transportation Staff", Event: "First Meet Teknik Informatika 24", Desc: "Coordinated transportation for participants and committee members throughout First Meet."},
		{ID: 2, Date: "Feb 2025 – Mar 2025", IsIT: true, Role: "IT Staff", Event: "Buat Milih Ketua Dua Empat (BUMI KETUPAT)", Desc: "Built the front-end of the website used to run the class leader election digitally."},
		{ID: 3, Date: "Jul 2025", IsIT: false, Role: "Group Mentor", Event: "Penyambutan Maba Teknik Informatika 25", Desc: "Guided and accompanied a group of new students through the entire welcoming program."},
		{ID: 4, Date: "Jun 2025 – Aug 2025", IsIT: true, Role: "Front-End Staff, IMTEK Division", Event: "PPLK ITERA 2025", Desc: "Built the PPLK event website's interface using React."},
		{ID: 5, Date: "Sep 2025 – Oct 2025", IsIT: true, Role: "Head of IT Division", Event: "Informatics Goes to TPB (IGTTPB) 2025", Desc: "Acted as the IT division's project manager and backup developer, making sure the website ran well. Built with Next.js."},
		{ID: 6, Date: "Oct 2025", IsIT: true, Role: "Content Research Staff, IMTEK Division", Event: "PODIUM ITERA", Desc: "Researched and gathered information used as content on the event website."},
		{ID: 7, Date: "Aug 2025 – Nov 2025", IsIT: false, Role: "Sponsorship Staff", Event: "Informatics Festival 2025", Desc: "Found potential sponsors, submitted proposals, and followed up with them to support the event."},
		{ID: 8, Date: "Mar 2026 – May 2026", IsIT: false, Role: "Vice Chair", Event: "Paskah KMK ITERA 2026", Desc: "Helped the chair coordinate the whole event and backed them up on technical decisions."},
		{ID: 10, Date: "Apr 2026 – Present", IsIT: false, Role: "Event Management Staff", Event: "Serasehan HMIF ITERA", Desc: "Prepared the event rundown and SOP as the technical guide for running the event."},
		{ID: 11, Date: "Apr 2026 – Present", IsIT: false, Role: "Timekeeper Sub-Division Lead", Event: "Point Project 4.0 HMIF ITERA", Desc: "Led the timekeeper sub-division, keeping every part of the event on schedule according to the rundown."},
		{ID: 12, Date: "Aug 2026 – Present", IsIT: true, Role: "Back-End Sub-Division Lead, IMTEK Division", Event: "PPLK ITERA 2026", Desc: "Led the IMTEK back-end sub-division building the PPLK 2026 website and game with Laravel and MySQL: handling the data flow, managing the database, and keeping server–client interaction running smoothly."},
		{ID: 13, Date: "Aug 2026 – Present", IsIT: true, Role: "Head of Technology Implementation", Event: "Kemahasiswaan Expo 2026", Desc: "Led the Technology Implementation Division in building the K-EXPO 2026 digital portal, from splitting work across the team to developing the website with Next.js, Go (Gin), PostgreSQL, and Google OAuth login."},
		{ID: 14, Date: "Oct 2026 – Present", IsIT: true, Role: "Back-End Staff", Event: "Neofest 2026", Desc: "Developing the back-end of the Neofest 2026 website, including data management and the server services used by the front-end."},
		{ID: 15, Date: "Oct 2026 – Present", IsIT: true, Role: "Head of Technology Implementation", Event: "Pemilihan Raya KM ITERA 2026", Desc: "Leading the Technology Implementation Division of the KM ITERA 2026 general election, responsible for its technology needs, including the systems and website supporting the election."},
	},
	"magang": {
		{ID: 1, Date: "Sep 2025 – Dec 2025", IsIT: false, Role: "Lab Assistant – TPB 21", Event: "Dasar Teknologi Digital, ITERA", Desc: "Assisted TPB class 21 during Digital Technology Fundamentals labs: explaining material, guiding lab assignments, and helping with grading."},
		{ID: 3, Date: "Feb 2026 – Jun 2026", IsIT: false, Role: "Lab Assistant – T 11", Event: "Pengenalan Komputasi, ITERA", Desc: "Assisted TPB class T 11 during Introduction to Computing labs: explaining material, guiding lab assignments, and helping with grading."},
		{ID: 5, Date: "Jun 2026 – Jul 2026", IsIT: false, Role: "Supervisor Intern", Event: "Chandramart", Desc: "Interned as a supervisor at Chandramart (retail supermarket), helping monitor store operations and coordinate with staff."},
		{ID: 4, Date: "Sep 2026 – Present", IsIT: false, Role: "Lab Assistant – TPB 43", Event: "Dasar Teknologi Digital, ITERA", Desc: "Back as a Digital Technology Fundamentals lab assistant for TPB class 43: explaining material, guiding lab assignments, and helping with grading."},
	},
	"organisasi": {
		{ID: 1, Date: "Jun 2025 – Present", IsIT: false, Role: "Member – IT Division", Event: "Badan Pengurus Angkatan (BPA) Algovista – IF 24", Desc: "Member of the Algovista IF 24 class board's IT Division, supporting the class's technology needs such as the class website."},
		{ID: 5, Date: "Sep 2025 – Dec 2025", IsIT: false, Role: "Intern", Event: "Kementerian Teknologi Informasi KM ITERA (Kabinet Resonara)", Desc: "Interned at the KM ITERA Ministry of Information Technology (Resonara Cabinet), supporting the ministry's IT needs while learning how a student organization works."},
		{ID: 2, Date: "Jan 2026 – Present", IsIT: false, Role: "Junior Member – Technopreneur Division", Event: "Himpunan Mahasiswa Informatika (HMIF) ITERA", Desc: "Organizing the Technopreneur Division within the Professional Development Department."},
		{ID: 3, Date: "Jan 2026 – Present", IsIT: false, Role: "Member – Technology Implementation Division", Event: "Keluarga Mahasiswa Katholik (KMK) St. Thomas Aquinas ITERA", Desc: "Organizing the Technology Implementation Division within the Visual Communication Media Department."},
		{ID: 4, Date: "Oct 2026 – Present", IsIT: false, Role: "Adminfot – Website Division", Event: "Humas Informatika ITERA (Batch 3)", Desc: "Adminfot in the Website Division of Humas Informatika ITERA batch 3, helping manage and update the website's information content."},
	},
	"prestasi": {
		{ID: 1, Date: "Aug 2024", IsIT: false, Role: "Participant", Event: "Pelatihan LKMM Pra-TD I ITERA", Desc: "Took part in pre-basic student leadership and management training to develop leadership skills."},
		{ID: 2, Date: "Nov 2024", IsIT: false, Role: "Participant", Event: "Mini Bootcamp 1.0 HMIF ITERA", Desc: "Joined HMIF ITERA's mini bootcamp, building, releasing, and showcasing a simple website as the final project."},
		{ID: 3, Date: "Aug 2025", IsIT: false, Role: "Volunteer", Event: "HMIF Mengabdi Vol. 3", Desc: "Volunteered in a community service program organized by HMIF ITERA."},
		{ID: 4, Date: "May 2026", IsIT: false, Role: "Participant", Event: "Pelatihan LKMM TD VIII ITERA", Desc: "Took part in basic student leadership and management training to develop leadership skills."},
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
