package sslcheck

import (
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	expires := time.Now().Local().AddDate(0, 0, 15)
	status := getStatus(expires)
	if status != "healthy" {
		t.Fatalf("expected status to be healthy got %s", status)
	}

	expires = time.Now().Local().AddDate(0, 0, 14)
	status = getStatus(expires)
	if status != "looming" {
		t.Fatalf("expected status to be looming got %s", status)
	}

	expires = time.Now().Local().AddDate(0, 0, -1)
	status = getStatus(expires)
	if status != "expired" {
		t.Fatalf("expected status to be expired got %s", status)
	}

}
