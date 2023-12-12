package url

import "testing"

func TestParse(t *testing.T) {

	const rawurl = "https://foo.com"
	var u *URL
	var err error

	if u, err = Parse(rawurl); err != nil {
		t.Logf("Parse %q err = %q, want nil", rawurl, err)
		t.Fail()
	}
	want := "https"
	got := u.Scheme

	if got != want {
		t.Logf("Parsing (%q).Scheme = %q; want %q", rawurl, got, want)
		t.Fail()
	}
}
