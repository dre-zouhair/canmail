package main

import "os"

func Up() {
	os.Setenv("SMTP_DOMAIN", "localhost:1025")
	os.Setenv("SMTP_HOST", "localhost")
	os.Setenv("SMTP_USERNAME", "smtp")
	os.Setenv("SMTP_PASSWORD", "password")
	os.Setenv("SMTP_FROM", "from@mailer")
	os.Setenv("SMTP_FROM_PASSWORD", "password")
	os.Setenv("SMTP_FROM_NAME", "from")

	os.Setenv("MONGODB_HOST", "localhost")
	// os.Setenv("MONGODB_USERNAME", "mailer")
	// os.Setenv("MONGODB_PASSWORD", "password")
	os.Setenv("MONGODB_DB_NAME", "mailer")
	os.Setenv("MONGODB_PORT", "27017")
}

func Down() {
	os.Remove("SMTP_DOMAIN")
	os.Remove("SMTP_HOST")
	os.Remove("SMTP_USERNAME")
	os.Remove("SMTP_PASSWORD")
	os.Remove("SMTP_FROM")
	os.Remove("SMTP_FROM_PASSWORD")
	os.Remove("SMTP_FROM_NAME")
	os.Remove("MONGODB_HOST")
	os.Remove("MONGODB_USERNAME")
	os.Remove("MONGODB_PASSWORD")
	os.Remove("MONGODB_DB_NAME")
	os.Remove("MONGODB_PORT")
}
