package helper

import "yinyushijing.cn/xnote-api/model"

// BuildXNote builds a XNote object with the fields given
func BuildXNote(id int, title, content, createdTime, updatedTime string, status, uid int) model.XNote {
	return model.XNote{
		ID:          id,
		Title:       title,
		Content:     content,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		Status:      status,
		UID:         uid,
	}
}

// Form2Domain converts form object to domain object
func Form2Domain(xnoteForm *XNoteForm) *model.XNote {
	var domainNote = model.XNote{
		ID:          *(xnoteForm.ID),
		Title:       xnoteForm.Title,
		Content:     xnoteForm.Content,
		CreatedTime: xnoteForm.CreatedTime,
		UpdatedTime: xnoteForm.UpdatedTime,
		Status:      *(xnoteForm.Status),
		UID:         xnoteForm.UID,
	}

	return &domainNote
}
