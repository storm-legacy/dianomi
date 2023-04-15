package mail

import "github.com/storm-legacy/dianomi/pkg/config"

var (
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
	smtpTLS      bool
)

func init() {
	smtpHost = config.GetString("APP_SMTP_HOST")
	smtpPort = config.GetInt("APP_SMTP_PORT")
	smtpUsername = config.GetString("APP_SMTP_USER")
	smtpPassword = config.GetString("APP_SMTP_PASSWORD")
	smtpTLS = config.GetBool("APP_SMTP_TLS")
}
