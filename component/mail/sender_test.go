package mail

import (
	"quizen/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {

	if testing.Short() {
		t.Skip("Skip test send email with gmail")
	}

	config, err := config.LoadConfig("../../")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Test subject"
	content := `
	<h1>Test content</h1>
	`
	to := []string{"duongcongcuong02012004@gmail.com"}

	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)

	require.NoError(t, err)
}
