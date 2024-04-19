package teamsalerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sp-monitoring/models"
)

// SendAlert sends an alert message to Microsoft Teams
func SendAlert(message string) error {

	msg := models.Message{
		Text: message,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
