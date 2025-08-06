package actions

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const packagesListURL = "https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/packages/package_list"

func Search(query string) {
	start := time.Now()
	fmt.Printf("* Searching for packages: %q\n", query)

	resp, err := http.Get(packagesListURL)
	if err != nil {
		fmt.Println("* Error fetching packages list:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("* Failed to fetch packages list: HTTP", resp.StatusCode)
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	matches := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(query)) {
			matches = append(matches, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("* Error reading packages list:", err)
		return
	}

	if len(matches) == 0 {
		fmt.Println("* No packages found matching", query)
	} else {
		fmt.Printf("* Found %d package(s):\n", len(matches))
		for _, entry := range matches {
			parts := strings.SplitN(entry, "/", 2)
			pkgID := ""
			pkgName := ""
			version := ""

			if len(parts) == 2 {
				pkgID = parts[0]
				pkgName = parts[1]
				idParts := strings.SplitN(pkgID, "@", 2)
				if len(idParts) == 2 {
					version = idParts[1]
				}
			} else {
				pkgName = entry
			}

			if version != "" {
				fmt.Printf("*  - %s (version %s)\n", pkgName, version)
			} else {
				fmt.Printf("*  - %s\n", pkgName)
			}
		}
	}

	fmt.Printf("* Search completed in %v\n", time.Since(start))
}
