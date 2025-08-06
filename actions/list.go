package actions

import (
	"fmt"
	"time"

	"catman/metrics"
)

func ListPackages() {
	pkgs, err := metrics.ListPackages()
	if err != nil {
		fmt.Println("Error reading installed packages:", err)
		return
	}

	fmt.Printf("Total installed packages: %d\n\n", len(pkgs))

	for _, pkg := range pkgs {
		t := time.Unix(pkg.InstalledAt, 0)
		fmt.Printf("- %s (version: %s) installed at %s\n", pkg.Name, pkg.Version, t.Format("2006-01-02 15:04:05"))
	}
}
