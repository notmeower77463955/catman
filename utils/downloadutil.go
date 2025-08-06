package utils

// import (
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// func DownloadFile(url string) (string, error) {
// 	tmpDir := os.TempDir()
// 	fileName := filepath.Base(url)
// 	dest := filepath.Join(tmpDir, fileName)

// 	out, err := os.Create(dest)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer out.Close()

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return "", fmt.Errorf("failed to download: %s (status %d)", url, resp.StatusCode)
// 	}

// 	_, err = io.Copy(out, resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return dest, nil
// }
