package model



type EmailVerifConfig struct{
	CONFIG_SMTP_HOST string
	CONFIG_SMTP_PORT int
	CONFIG_SENDER_NAME string
	CONFIG_AUTH_EMAIL string
	CONFIG_AUTH_PASSWORD string
}

type GoogleOauthConfig struct{
	CLIENT_ID string
	CLIENT_SECRET string
}

type LoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,contains_number,contains_special"`
	
}