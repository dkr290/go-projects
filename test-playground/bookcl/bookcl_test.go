package main

import (
	"testing"
)

type testCase struct {
	filePathCl string
	want       []BookReader
	wantErr    bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookReader_Sucess(t *testing.T) {

	var tests = map[string]testCase{
		"file exists": {
			filePathCl: "testdata/bookcl.json",
			want: []BookReader{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			filePathCl: "testdata/no_file_here.json",
			want:       nil,
			wantErr:    true,
		},

		"invalid Json": {
			filePathCl: "testdata/invalid.json",
			want:       nil,
			wantErr:    true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookReades(testCase.filePathCl)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())

			}

			if !equalBookCl(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})
	}
}

func equalBookCl(bookreaders, target []BookReader) bool {
	if len(bookreaders) != len(target) {
		return false
	}

	for i := range bookreaders {
		if bookreaders[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(bookreaders[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {

	tt := map[string]struct {
		input []BookReader
		want  map[Book]uint
	}{
		"normal use case": {
			input: []BookReader{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookreaders": {
			input: []BookReader{},
			want:  map[Book]uint{},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(got, tc.want) {
				t.Fatalf("got different list of books: %v, expected %v", got, tc.want)
			}
		})
	}

}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestFindCommonBooks(t *testing.T) {

	tt := map[string]struct {
		input []BookReader
		want  []Book
	}{
		"no common book": {
			input: []BookReader{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one commmon book": {
			input: []BookReader{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: []Book{handmaidsTale},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !equalBooks(got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}

}
