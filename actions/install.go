package actions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"catman/utils"
	"catman/metrics"
)

func InstallModule(moduleName string) {
	start := time.Now()

	metadataURL := fmt.Sprintf("https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/packages/%s/%s.cat", moduleName, moduleName)
	buildURL := fmt.Sprintf("https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/packages/%s/%s.sh", moduleName, moduleName)

	metadataFile := filepath.Join(os.TempDir(), moduleName+".cat")
	buildFile := filepath.Join(os.TempDir(), moduleName+".sh")

	metaStart := time.Now()
	err := utils.DownloadFile(metadataURL, metadataFile)
	if err != nil {
		fmt.Println("* Failed to download metadata:", err)
		return
	}
	fmt.Printf("* Metadata downloaded in %v\n", time.Since(metaStart))

	name, err := utils.Get(metadataFile, "Metadata.name")
	if err != nil {
		fmt.Println("* Error reading name from metadata:", err)
		return
	}
	version, err := utils.Get(metadataFile, "Metadata.version")
	if err != nil {
		fmt.Println("* Error reading version from metadata:", err)
		return
	}
	descLines := utils.GetSection(metadataFile, "Description")

	fmt.Printf("* You are about to install \033[1;36m%s\033[0m version \033[1;33m%s\033[0m\n", name, version)
	for _, line := range descLines {
		fmt.Printf("* %s\n", line)
	}
	fmt.Print("\n* Do you wish to continue? (y/n) > ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if strings.ToLower(input) != "y" {
		fmt.Println("* Installation aborted.")
		return
	}

	buildStart := time.Now()
	err = utils.DownloadFile(buildURL, buildFile)
	if err != nil {
		fmt.Println("* Failed to download build script:", err)
		return
	}
	fmt.Printf("* Build script downloaded in %v\n", time.Since(buildStart))

	fmt.Println("* executing install script...")

	cmd := exec.Command("sh", buildFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("* Error executing install script:", err)
		return
	}

	os.Remove(metadataFile)
	os.Remove(buildFile)

	err = metrics.AddPackage(name, version)
	if err != nil {
		fmt.Println("* uh oh cancer alert: ", err)
	}

	fmt.Printf("* Installed %s into user binary directory.\n", name)
	fmt.Printf("* Total installation time: %v\n", time.Since(start))
}
