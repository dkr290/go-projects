package store

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Link struct {
	Key string
	URL string
}

// validateNewLink checks a new link's validity.
func validateNewLink(link Link) error {
	if err := validateLinkKey(link.Key); err != nil {
		return err
	}
	u, err := url.Parse(link.URL)
	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("empty host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("scheme must be http or https")
	}
	return nil
}

// MaxKeyLen is the maximum length of a key.
const MaxKeyLen = 16

// validateLinkKey checks the key's validity.
func validateLinkKey(key string) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("empty key")
	}
	if len(key) > MaxKeyLen {
		return fmt.Errorf("key too long (max %d)", MaxKeyLen)
	}
	return nil
}
