package provider

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

type environmentKeys struct {
	Version string
}

var keys = environmentKeys{
	Version: "VERSION",
}

func LoadEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return err
}

func CheckForUpdates() {
	release, err := GetLatestRelease()
	if err != nil {
		fmt.Println(err)
		return
	}

	localVersion, err := getVersionFromString(os.Getenv(keys.Version))
	if err != nil {
		fmt.Println(`Error getting local version. Make sure you have a ".env" file with a "VERSION" key, you can get it from the releases page on GitHub.`)
		fmt.Println("------------------------------------------------------------")
		return
	}

	remoteVersion, err := getVersionFromString(release.GetTagName())
	if err != nil {
		fmt.Println("Error getting remote version")
		fmt.Println("------------------------------------------------------------")
		return
	}

	if localVersion < remoteVersion {
		fmt.Println("There is a new version available:", release.GetTagName())
		fmt.Println("You can download it at:", release.GetHTMLURL())
		fmt.Println("------------------------------------------------------------")
	}
}

func getVersionFromString(version string) (float64, error) {
	pattern := regexp.MustCompile(`v(\d+(\.\d+)?)`)
	allMatches := pattern.FindAllStringSubmatch(version, -1)
	if len(allMatches) == 0 {
		return 0.0, fmt.Errorf("version not found")
	}

	matches := allMatches[0]
	if len(matches) < 2 {
		return 0.0, fmt.Errorf("version not found")
	}

	match := matches[1]
	if match == "" {
		return 0.0, fmt.Errorf("version not found")
	}

	number, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0.0, fmt.Errorf("version not found")
	}

	return number, nil
}
