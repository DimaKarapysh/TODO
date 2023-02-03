package app

import "github.com/joho/godotenv"

func InitApp() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = InitLogs()
	if err != nil {
		return err
	}

	return nil
}
