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
