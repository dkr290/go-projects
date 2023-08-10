package sslcheck

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/dkr290/go-projects/httpchecker/data"
)

func Inspect(ctx context.Context, ssldomain string) (*data.SSLTracking, error) {

	errch := make(chan error, 1)
	result := make(chan data.SSLTracking, 1)

	go func() {

		config := &tls.Config{}
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", ssldomain), config)
		if err != nil {
			errch <- err
		}
		defer conn.Close()
		state := conn.ConnectionState()
		cert := state.PeerCertificates[0]
		result <- data.SSLTracking{
			Expires:    cert.NotAfter,
			DomainName: ssldomain,
			Issuer:     cert.Issuer.Organization[0],
			Status:     getStatus(cert.NotAfter),
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-result:
		return &result, nil
	}

}

var loomingTreshold = time.Hour * 24 * 7 * 2 // 2 weeks

func getStatus(expires time.Time) string {
	// 	if expires.Before(time.Now()) {

	// 	}

	// 1. healthy expires > now
	//2. expores > now but in x amount of time
	//3 .Expired < now

	if expires.Before(time.Now()) {
		return "expired"
	}

	timeLeft := time.Until(expires)

	if timeLeft < loomingTreshold {
		return "looming"
	}
	fmt.Println(timeLeft)
	return "healthy"
}
