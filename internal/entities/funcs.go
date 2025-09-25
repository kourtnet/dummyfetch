package entities

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Required for mock-ing files in tests
var (
	readFile = os.ReadFile
	getEnv   = os.Getenv
	open     = os.Open
)

func GetDistro() (string, error) {
	return getOSReleaseInfo("ID=")
}

func getPrettyDistro() (string, error) {
	return getOSReleaseInfo("PRETTY_NAME=")
}

func getOSReleaseInfo(prefix string) (string, error) {
	file, err := open("/etc/os-release")
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res string

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, prefix) {
			suffix := text[len(prefix):]
			res = strings.Trim(suffix, `"`)

			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return res, nil
}

func getKernel() (string, error) {
	kernelByte, err := readFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "", err
	}

	kernel := strings.TrimSpace(string(kernelByte))

	return kernel, nil
}

func getShell() (string, error) {
	shellPath := getEnv("SHELL")
	shell := filepath.Base(shellPath)

	return shell, nil
}

func getTerminal() (string, error) {
	term := getEnv("TERM")
	return term, nil
}

func getUptime() (string, error) {
	fileByte, err := readFile("/proc/uptime")
	if err != nil {
		return "", err
	}

	fileFields := strings.Fields(string(fileByte))
	if len(fileFields) == 0 {
		return "", errors.New("uptime file is empty")
	}

	uptimeStr := fileFields[0]

	dur, err := time.ParseDuration(uptimeStr + "s")
	if err != nil {
		return "", err
	}

	days := int(dur.Hours() / 24)
	hours := int(dur.Hours()) - days*24
	minutes := int(dur.Minutes()) - days*24*60 - hours*60

	var res string

	if days > 0 {
		res += strconv.Itoa(days)
		if days > 1 {
			res += " days, "
		} else {
			res += " day, "
		}
	}

	if hours > 0 {
		res += strconv.Itoa(hours)
		if hours > 1 {
			res += " hours, "
		} else {
			res += " hour, "
		}
	}

	if minutes > 0 {
		res += strconv.Itoa(minutes)
		if minutes > 1 {
			res += " mins"
		} else {
			res += " min"
		}
	}

	if res == "" {
		res = "0 mins"
	}

	return strings.TrimSuffix(res, ", "), nil
}

func getPaletteBg() (string, error) {
	output := "\033[2D\033[40m   \033[41m   \033[42m   \033[43m   \033[44m   \033[45m   \033[46m   \033[47m   \033[0m"
	return output, nil
}

func getPaletteFg() (string, error) {
	output := "\033[2D\033[90m███\033[91m███\033[92m███\033[93m███\033[94m███\033[95m███\033[96m███\033[97m███\033[0m"
	return output, nil
}

func getPrint() (string, error) {
	return "\033[2D ", nil
}
