package helper

// XNoteForm is used for bind data from user input
type XNoteForm struct {
	ID          *int   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CreatedTime string `json:"create_time"`
	UpdatedTime string `json:"update_time" `
	UID         int    `json:"uid" binding:"required"`
	Status      *int   `json:"status"`
}
