package metrics

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type InstalledPackage struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	InstalledAt int64  `json:"installed_at"`
}

var dbPath = filepath.Join(os.Getenv("HOME"), ".catman", "installed.json")

func loadInstalled() ([]InstalledPackage, error) {
	data, err := ioutil.ReadFile(dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []InstalledPackage{}, nil
		}
		return nil, err
	}

	var pkgs []InstalledPackage
	if err := json.Unmarshal(data, &pkgs); err != nil {
		return nil, err
	}
	return pkgs, nil
}

func saveInstalled(pkgs []InstalledPackage) error {
	data, err := json.MarshalIndent(pkgs, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(dbPath, data, 0644)
}

func AddPackage(name, version string) error {
	pkgs, err := loadInstalled()
	if err != nil {
		return err
	}

	for _, p := range pkgs {
		if p.Name == name {
			return nil
		}
	}

	pkgs = append(pkgs, InstalledPackage{
		Name:        name,
		Version:     version,
		InstalledAt: time.Now().Unix(),
	})

	return saveInstalled(pkgs)
}

func RemovePackage(name string) error {
	pkgs, err := loadInstalled()
	if err != nil {
		return err
	}

	filtered := []InstalledPackage{}
	for _, p := range pkgs {
		if p.Name != name {
			filtered = append(filtered, p)
		}
	}

	return saveInstalled(filtered)
}

func ListPackages() ([]InstalledPackage, error) {
	return loadInstalled()
}
