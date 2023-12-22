package url

import "testing"

func TestParse(t *testing.T) {

	const rawurl = "https://foo.com/go"
	//const rawurl = "foo.com"
	var u *URL
	var err error

	if u, err = Parse(rawurl); err != nil {
		t.Fatalf("Parse %q err = %q, want nil", rawurl, err)

	}
	want := "https"

	if got := u.Scheme; got != want {
		t.Errorf("Parsing (%q).Scheme = %q; want %q", rawurl, got, want)

	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parsing (%q).Host = %q; want %q", rawurl, got, want)
	}

	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parsing (%q) got Path = %q; want %q", rawurl, got, want)
	}
}

func TestURLPort(t *testing.T) {
	tests := []struct {
		inpt string
		port string
	}{
		{inpt: "foo.com:80", port: "80"}, //host with port
		{inpt: "foo.com", port: ""},      //host without port
		{inpt: "1.2.3.4:90", port: "90"}, //ip with port
		{inpt: "1.2.3.4", port: ""},      //ip without port
	}

	for _, tt := range tests {
		u := &URL{
			Host: tt.inpt,
		}

		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("for host %q; got %q; want %q", tt.inpt, got, want)
		}
	}
}
