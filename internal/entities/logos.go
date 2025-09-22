package entities

type LogoInfo struct {
	TextColor int
	Logo      []string
	BlankRow  string
}

const (
	ArchName   = "arch"
	UbuntuName = "ubuntu"
	TuxName    = "tux"
)

var LogosMap = map[string]LogoInfo{
	ArchName: {
		TextColor: 34,
		Logo: []string{
			"   \033[34m/\\\033[0m   ",
			"  \033[34m/\\ \\\033[0m  ",
			" \033[34m/ .\\ \\\033[0m ",
			"\033[34m/.'  '.\\\033[0m",
		},
		BlankRow: "\t\t",
	},

	UbuntuName: {
		TextColor: 31,
		Logo: []string{
			"  \033[33m/\033[31m--( )\033[0m",
			"\033[31m( )    \033[31m|\033[0m",
			" \033[33m\\     \033[31m/\033[0m",
			"   \033[31m--\033[33m( )\033[0m",
			"\t",
		},
		BlankRow: "\t",
	},

	TuxName: {
		TextColor: 37,
		Logo: []string{
			"   \033[37m.-,\033[0m  ",
			"   \033[37moo\033[37m|\033[0m  ",
			"  /\033[33mv \033[37m\\\033[0m  ",
			" \033[33m(\\\033[37m_^\033[33m/)\033[0m ",
			"",
		},
		BlankRow: "\t",
	},
}

var BasicLogo = LogosMap[TuxName]
