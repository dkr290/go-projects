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
			got, err := loadBookCl(testCase.filePathCl)
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
