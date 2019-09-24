package test

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"yinyushijing.cn/xnote-api/repository"
)

const DateTimeFormat = "2006-01-02 15:04:05"

// test filename must be ended with _test.go
func TestGetbyid(t *testing.T) {

	actual := repository.GetByID(1)
	expectedID := 1

	if actual.ID != expectedID {
		t.Error("object get by id not equal")
	}
}

func TestSearch(t *testing.T) {

	fmt.Println("This is to test search logic")

	// repository.Search()

	fmt.Println("This is to test search logic")
}

func TestCreateXNote(t *testing.T) {
	fmt.Println("we gonna saving data")
	currentTime := time.Now()
	currentTimeStr := currentTime.Format(DateTimeFormat)
	log.Println(currentTimeStr)
	// repository.Save(helper.BuildXNote(1, "hello note", "This is content body", currentTimeStr, currentTimeStr, 1, 1))
	// repository.Save(helper.BuildXNote(1, "hello 中文", "中文的亚美尼亚 body", currentTimeStr, currentTimeStr, 1, 1))

}

func TestQueryByTitle(t *testing.T) {
	notes := repository.QueryByTitle("hello")
	actualLen := len(notes)

	if actualLen < 1 {
		t.Errorf("actual length is less than 1:%d", actualLen)
	}

	// for _, note := range notes {
	// 	fmt.Println(note.Title)
	// }
	oneNote := notes[0]
	if !strings.Contains(oneNote.Title, "hello") {
		t.Errorf("the returned object doesn't contain the data expected:%s", oneNote.Title)
	}

}
