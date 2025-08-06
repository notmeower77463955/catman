package utils

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"io"
	"net/http"
)

func Get(path string, key string) (string, error) {
	section := ""
	if strings.Contains(key, ".") {
		parts := strings.SplitN(key, ".", 2)
		section = parts[0]
		key = parts[1]
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inSection := section == ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			current := strings.Trim(line, "[]")
			inSection = (current == section)
			continue
		}

		if inSection && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			v = strings.Trim(v, "\"")

			if k == key {
				return v, nil
			}
		}
	}

	return "", fmt.Errorf("key %s not found", key)
}

func IsValid(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		return false
	}
	defer f.Close()

	foundName, foundURL := false, false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			k := strings.TrimSpace(parts[0])
			if k == "name" {
				foundName = true
			}
			if k == "version" {
				foundURL = true
			}
		}
		if foundName && foundURL {
			return true
		}
	}
	return false
}

func Keys(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	keys := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			k := strings.TrimSpace(parts[0])
			keys = append(keys, k)
		}
	}
	return keys, nil
}

func Values(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	values := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			v := strings.TrimSpace(parts[1])
			values = append(values, v)
		}
	}
	return values, nil
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func ReadAll(file string) (map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			result[k] = v
		}
	}
	return result, nil
}

func GetSection(filePath, section string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	inSection := false
	sectionHeader := fmt.Sprintf("[%s]", section)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			inSection = line == sectionHeader
			continue
		}
		if inSection {
			lines = append(lines, line)
		}
	}

	return lines
}

func DownloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}