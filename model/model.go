package model

type EmailVerifConfig struct{
	CONFIG_SMTP_HOST string
	CONFIG_SMTP_PORT int
	CONFIG_SENDER_NAME string
	CONFIG_AUTH_EMAIL string
	CONFIG_AUTH_PASSWORD string
}