package service

import (
	"yinyushijing.cn/xnote-api/model"
	"yinyushijing.cn/xnote-api/repository"
)

func GetNoteByID(id int) model.XNote {
	return repository.GetByID(id)
}

func QueryNotesByTitle(keywords string) []model.XNote {
	return repository.QueryByTitle(keywords)
}

func CreateXNote(note *model.XNote) {
	repository.Create(*note)
}

func UpdateXNote(note *model.XNote) {
	repository.Update(*note)
}

func RemoveXNoteByID(id int) {
	repository.Remove(id)
}
