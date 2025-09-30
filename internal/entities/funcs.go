package entities

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"os/exec"
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

func getHost() (string, error) {
	host, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		return "", err
	}

	ver, err := os.ReadFile("/sys/class/dmi/id/product_version")
	if err != nil {
		return "", err
	}

	res := strings.TrimSpace(string(host)) + " (" + strings.TrimSpace(string(ver)) + ")"

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
	shell := shellPath[strings.LastIndex(shellPath, "/")+1:]

	return shell, nil
}

func getTerminal() (string, error) {
	term := os.Getenv("TERM")
	term = term[strings.Index(term, "-")+1:]

	getVer := exec.Command(term, "--version")

	output, err := getVer.Output()
	if err != nil {
		return "", err
	}

	version := string(bytes.Fields(output)[1])

	res := term + " " + version

	return res, nil
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

// WARNING: works only with pacman (I use Arch, btw)
func getPackages() (string, error) {
	entries, err := os.ReadDir("/var/lib/pacman/local/")
	if err != nil {
		return "", err
	}

	dirsNum := 0
	for _, entry := range entries {
		if entry.IsDir() {
			dirsNum++
		}
	}

	res := strconv.Itoa(dirsNum) + " (pacman)"

	return res, nil
}

func getMemory() (string, error) {
	file, err := open("/proc/meminfo")
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	var memAvailable, memTotal int
	for scanner.Scan() && counter < 2 {
		strs := strings.Fields(scanner.Text())
		if strs[0] == "MemTotal:" {
			counter++

			var err error
			memTotal, err = strconv.Atoi(strs[1])
			if err != nil {
				return "", err
			}
		} else if strs[0] == "MemAvailable:" {
			counter++

			var err error
			memAvailable, err = strconv.Atoi(strs[1])
			if err != nil {
				return "", err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	memTotal /= 1024
	memUsed := memTotal - memAvailable/1024

	res := strconv.Itoa(memUsed) + " MiB / " + strconv.Itoa(memTotal) + " MiB"

	return res, nil
}

func getCPU() (string, error) {
	file, err := open("/proc/cpuinfo")
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res string
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ":")

		if strings.TrimSpace(strs[0]) == "model name" {
			res = strings.TrimSpace(strs[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return res, nil
}

func getPaletteBg() (string, error) {
	output := "\033[2D\033[40m   \033[41m   \033[42m   \033[43m   \033[44m   \033[45m   \033[46m   \033[47m   \033[0m"
	return output, nil
}

func getPaletteFg() (string, error) {
	output := "\033[2D\033[90m███\033[91m███\033[92m███\033[93m███\033[94m███\033[95m███\033[96m███\033[97m███\033[0m"
	return output, nil
}
