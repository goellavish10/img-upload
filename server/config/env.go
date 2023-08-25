package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	pathOfEnv := filepath.Join(dir, ".env")

	err := godotenv.Load(pathOfEnv)
	if err != nil {
		return err
	}
	return nil
}
