package entities

type LogoInfo struct {
	TextColor string // field is string type in order not to waste time for strconv.Itoa at runtime
	Logo      []string
	BlankRow  string
}

const (
	ArchName   = "arch"
	UbuntuName = "ubuntu"
	TuxName    = "tux"
	CustomName = "custom"
	AutoName   = "auto"
)

const DefaultLogoName = TuxName

var LogosMap = map[string]LogoInfo{
	ArchName: {
		TextColor: "34",
		Logo: []string{
			"   \033[34m/\\   ",
			"  /\\ \\  ",
			" / .\\ \\ ",
			"/.'  '.\\",
		},
		BlankRow: "\x1b[8C",
	},

	UbuntuName: {
		TextColor: "31",
		Logo: []string{
			"  \033[33m/\033[31m--( )",
			"( )    |",
			" \033[33m\\     \033[31m/",
			"   --\033[33m( )",
		},
		BlankRow: "\x1b[8C",
	},

	TuxName: {
		TextColor: "97",
		Logo: []string{
			"   \033[37m.-,  ",
			"   oo|  ",
			"  /\033[33mv \033[37m\\  ",
			" \033[33m(\\\033[37m_^\033[33m/) ",
		},
		BlankRow: "\x1b[8C",
	},
	CustomName: {
		TextColor: "97",
	},
}
