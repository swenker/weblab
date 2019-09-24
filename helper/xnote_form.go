package helper

type XNoteForm struct {
	ID          *int   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CreatedTime string `json:"create_time" binding:"required"`
	UpdatedTime string `json:"update_time" binding:"required"`
	UID         int    `json:"uid" binding:"required"`
	Status      *int   `json:"status" binding:"required"`
}
