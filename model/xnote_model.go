package model

type XNote struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedTime string `json:"create_time"`
	UpdatedTime string `json:"update_time"`
	UID         int    `json:"uid"`
	Status      int    `json:"status"`
}
