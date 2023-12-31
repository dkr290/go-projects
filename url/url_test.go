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
// func TestURLPort(t *testing.T) {

// 	//helper function

// 	testPort := func(inpt, wantPort string) {

// 		t.Helper()
// 		u := &URL{Host: inpt}
// 		if got := u.Port(); got != wantPort {
// 			t.Errorf("for host %s, got %s, want %s", inpt, got, wantPort)
// 		}

// 	}

// 	t.Run("with port", func(t *testing.T) {

// 		testPort("foo.com:80", "80")

// 	})

// 	t.Run("empty port", func(t *testing.T) {

// 		testPort("foo.com:", "")

// 	})

// 	t.Run("without port", func(t *testing.T) {

// 		testPort("foo.com", "")

// 	})

// 	t.Run("ip with port", func(t *testing.T) {

// 		testPort("1.2.3.4:90", "90")

// 	})

// 	t.Run("ip without port", func(t *testing.T) {

// 		testPort("1.2.3.4", "")

// 	})

// }
// another way but is a bit more complex

// func TestURLPort(t *testing.T) {

// 	//helper function

// 	testPort := func(inpt, wantPort string) func(*testing.T) {

// 		return func(t *testing.T) {

// 			t.Helper()
// 			u := &URL{Host: inpt}
// 			if got := u.Port(); got != wantPort {
// 				t.Errorf("for host %s, got %s, want %s", inpt, got, wantPort)
// 			}

// 		}

// 	}

// 	t.Run("with port", testPort("foo.com:80", "80"))
// 	t.Run("empty port", testPort("foo.com:", ""))

// 	t.Run("without port", testPort("foo.com", ""))

// 	t.Run("ip with port", testPort("1.2.3.4:90", "90"))

// 	t.Run("ip without port", testPort("1.2.3.4", ""))

// }

//combine table driven tests with subtests

func TestURLPort(t *testing.T) {
	tests := map[string]struct {
		inpt string //URL.host field
		port string
	}{
		"with port":       {inpt: "foo.com:80", port: "80"},
		"empty port":      {inpt: "foo.com:", port: ""},
		"without port":    {inpt: "foo.com", port: ""},
		"ip with port":    {inpt: "1.2.3.4:90", port: "90"},
		"ip without port": {inpt: "1.2.3.4", port: ""},
	}

	for name, tt := range tests {

		t.Run(name, func(t *testing.T) {
			u := &URL{Host: tt.inpt}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("for host %s, got %s, want %s", tt.inpt, got, want)
			}
		})
	}

}
