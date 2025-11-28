package utilidades

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func EnviarCorreo(correo, asunto, mensaje string) error {
	// Obtener configuración SMTP desde variables de entorno
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	fromEmail := os.Getenv("SMTP_FROM_EMAIL")

	// Validar que las variables necesarias estén configuradas
	if smtpHost == "" || smtpPortStr == "" || smtpUser == "" || smtpPassword == "" {
		return fmt.Errorf("configuración SMTP incompleta en variables de entorno")
	}

	// Usar email por defecto si no está configurado
	if fromEmail == "" {
		fromEmail = "noreply@example.com"
	}

	// Convertir puerto a entero
	port, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("puerto SMTP inválido: %v", err)
	}

	// Crear mensaje
	msg := gomail.NewMessage()
	msg.SetHeader("From", fromEmail)
	msg.SetHeader("To", correo)
	msg.SetHeader("Subject", asunto)
	msg.SetBody("text/html", mensaje)

	// Configurar conexión SMTP
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPassword)

	// Enviar correo
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error al enviar correo: %v", err)
	}

	return nil
}
