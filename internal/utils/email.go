package utils

import (
	"html/template"
	"log"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetEmailName(email string) string {
	// Find the index of the '@' symbol in the email address
	atIndex := strings.Index(email, "@")

	// Extract the substring before the '@' symbol
	name := email[:atIndex]

	// Define the list of separators
	separators := []string{".", "_", "-"}

	// Split the name by separators and capitalize each word
	for _, separator := range separators {
		nameParts := strings.Split(name, separator)
		for i, part := range nameParts {
			caser := cases.Title(language.AmericanEnglish)
			nameParts[i] = caser.String(part)
		}
		name = strings.Join(nameParts, " ")
	}

	return name
}

func GenerateVerificationEmailHtml(name, verificationURL string) (string, error) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Email Verification</title>
	</head>
	<body>
		<h2>Email Verification</h2>
		<p>Dear {{.Name}},</p>
		<p>Thank you for registering. Please click the link below to verify your email:</p>
		<p><a href="{{.VerificationURL}}">{{.VerificationURL}}</a></p>
		<p>Best regards,<br> Your Company</p>
	</body>
	</html>
	`

	data := struct {
		Name            string
		VerificationURL string
	}{
		Name:            name,
		VerificationURL: verificationURL,
	}

	tmpl, err := template.New("verificationEmail").Parse(htmlTemplate)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var resultHTML strings.Builder
	err = tmpl.Execute(&resultHTML, data)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return resultHTML.String(), nil
}