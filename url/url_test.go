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

// func TestURLPort(t *testing.T) {
// 	tests := []struct {
// 		name string // adds new field to name the test cases
// 		inpt string
// 		port string
// 	}{
// 		{
// 			name: "with port",
// 			inpt: "foo.com:80", port: "80",
// 		}, //host with port
// 		{
// 			name: "with empty port",
// 			inpt: "foo.com:", port: "",
// 		}, //host without or empty port
// 		{

// 			name: "without port",
// 			inpt: "foo.com", port: "", // without port
// 		},
// 		{
// 			name: "ip with port",
// 			inpt: "1.2.3.4:90", port: "90",
// 		}, //ip with port

// 		{
// 			name: "ip without port",
// 			inpt: "1.2.3.4", port: "",
// 		}, //ip without port
// 	}

// 	for _, tt := range tests {
// 		u := &URL{
// 			Host: tt.inpt,
// 		}

// 		if got, want := u.Port(), tt.port; got != want {
// 			t.Errorf("%s: for host %q; got %q; want %q", tt.name, tt.inpt, got, want)
// 		}
// 	}
// }

// func TestURLPort(t *testing.T) {
// 	tests := map[string]struct { //using map type for the table tests
// 		inpt string // removuibng the name strring
// 		port string
// 	}{
// 		"with port": {
// 			inpt: "foo.com:80", port: "80",
// 		}, //host with port
// 		"with empty port": {
// 			inpt: "foo.com:", port: "",
// 		}, //host without or empty port
// 		"without port": {

// 			inpt: "foo.com", port: "", // without port
// 		},
// 		"ip with port": {
// 			inpt: "1.2.3.4:90", port: "90",
// 		}, //ip with port

// 		"ip without port": {
// 			inpt: "1.2.3.4", port: "",
// 		}, //ip without port
// 	}

// 	for name, tt := range tests {
// 		u := &URL{
// 			Host: tt.inpt,
// 		}

//			if got, want := u.Port(), tt.port; got != want {
//				t.Errorf("%s: for host %q; got %q; want %q", name, tt.inpt, got, want)
//			}
//		}
//	}
func TestURLPort(t *testing.T) {
	t.Run("with port", func(t *testing.T) {

		const inpt = "foo.com:80"

		u := &URL{Host: inpt}

		if got, want := u.Port(), "80"; got != want {
			t.Errorf("for host %s, got %s, want %s", u.Host, got, want)
		}

	})

	t.Run("empty port", func(t *testing.T) {

		const inpt = "foo.com:"

		u := &URL{Host: inpt}

		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %s, got %s, want %s", u.Host, got, want)
		}

	})

	t.Run("without port", func(t *testing.T) {

		const inpt = "foo.com"

		u := &URL{Host: inpt}

		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %s, got %s, want %s", u.Host, got, want)
		}

	})

	t.Run("ip with port", func(t *testing.T) {

		const inpt = "1.2.3.4:90"

		u := &URL{Host: inpt}

		if got, want := u.Port(), "90"; got != want {
			t.Errorf("for host %s, got %s, want %s", u.Host, got, want)
		}

	})

	t.Run("ip without port", func(t *testing.T) {

		const inpt = "1.2.3.4"

		u := &URL{Host: inpt}

		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %s, got %s, want %s", u.Host, got, want)
		}

	})

}
