package utils

import (
	"regexp"
	"strings"
)

func Slugify(title string) string {
	//title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "_", "-")
	reg := regexp.MustCompile(`[^A-Za-z0-9\-]+`)
	title = reg.ReplaceAllString(title, "")
	// collapse multiple dashes if needed
    regDash := regexp.MustCompile(`-+`)
    title = regDash.ReplaceAllString(title, "-")

    return strings.Trim(title, "-")
}