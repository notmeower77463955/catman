package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	localVersion = "0.0.1 beta"
	versionURL   = "https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/VERSION"
	timeout      = 5 * time.Second
)

func fetchRemoteVersion(url string) (string, error) {
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("masz raka: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}

func kernelVersion() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func getVersion() string {
	remoteVersion, err := fetchRemoteVersion(versionURL)
	if err != nil || remoteVersion == "" {
		return localVersion
	}
	return remoteVersion
}

func Print() {
	now := time.Now().Format(time.RFC1123Z)
	fmt.Printf("catman version %s (%s)\n", getVersion(), now)
	fmt.Printf("kernel version: %s\n", kernelVersion())
	fmt.Printf("golang version: %s\n", runtime.Version())
	fmt.Println("Copyright Meower Development 2025")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  -h, --help            Show this help")
	fmt.Println("  -i, --install NAME    Install package")
	fmt.Println("  -s, --search NAME     Search package")
	fmt.Println("  -l, --list            List installed packages")
	fmt.Println("  -d, --delete NAME     Delete package")
}

func PrintAndExit() {
	Print()
	os.Exit(0)
}
