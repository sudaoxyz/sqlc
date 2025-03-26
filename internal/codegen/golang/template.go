package golang

import "os"

func templatePattern() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return wd + "/templates/*.tmpl", nil
}
