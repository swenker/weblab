package model

//XNoteCriteria For query
type XNoteCriteria struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedDate string `json:"create_time"`
	UID         int    `json:"uid"`
	Status      int    `json:"status"`
}
