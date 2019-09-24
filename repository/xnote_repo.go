package repository

import (
	"database/sql"
	"fmt"
	"time"

	"yinyushijing.cn/xnote-api/helper"
	"yinyushijing.cn/xnote-api/model"

	//It not be used explicitly
	//This line is to import the sql driver
	_ "github.com/go-sql-driver/mysql"
)

const TABLE_XNOTE = "xnote_xnote"

func conndb() *sql.DB {
	// driver name and username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "bmlist:bmlist_u01@tcp(192.168.56.101:3306)/bmlist")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Create(xnote model.XNote) {

	currentTime := time.Now()
	db := conndb()
	defer db.Close()
	// tx, err := db.Begin()

	insertSQL := fmt.Sprintf("INSERT INTO %s(title,content,create_time,status,uid)VALUES(?,?,?,?,?)", TABLE_XNOTE)
	insertStmt, err := db.Prepare(insertSQL)
	if err != nil {
		panic(err.Error())
	}

	// tx.Commit()

	defer insertStmt.Close()

	_, err = insertStmt.Exec(xnote.Title, xnote.Content, currentTime, xnote.Status, xnote.UID)
	if err != nil {
		panic(err.Error())
	}

}

func Update(xnote model.XNote) {
	db := conndb()
	defer db.Close()
	// tx, err := db.Begin()

	SQL := fmt.Sprintf("UPDATE %s SET title=?,content=?,status=?,uid=? WHERE id=?", TABLE_XNOTE)
	stmt, err := db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}

	// tx.Commit()

	defer stmt.Close()

	_, err = stmt.Exec(xnote.Title, xnote.Content, xnote.Status, xnote.UID, xnote.ID)
	if err != nil {
		panic(err.Error())
	}

}

func Remove(id int) {
	db := conndb()
	defer db.Close()

	SQL := fmt.Sprintf("DELETE FROM %s WHERE id=?", TABLE_XNOTE)
	stmt, err := db.Prepare(SQL)
	if err != nil {
		panic(err.Error())
	}

	// tx.Commit()

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

}
func Search(queryCriteria model.XNoteCriteria) []model.XNote {
	db := conndb()
	defer db.Close()

	querySQL := fmt.Sprintf("SELECT id,title,content,create_time,update_time,status,uid FROM %s", TABLE_XNOTE)
	rows, err := db.Query(querySQL)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var notes []model.XNote

	for rows.Next() {
		var id, status, uid int
		var title, content, createdTime, updatedTime string
		err := rows.Scan(&id, &title, &content, &createdTime, &updatedTime, &status, &uid)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		notes = append(notes, helper.BuildXNote(id, title, content, createdTime, updatedTime, status, uid))
	}

	return notes
}

func QueryByTitle(keywords string) []model.XNote {
	db := conndb()
	defer db.Close()

	likePattern := fmt.Sprintf("%%%s%%", keywords)
	fmt.Println(likePattern)
	rows, err := db.Query("SELECT id,title,content,create_time,update_time,status,uid FROM xnote_xnote WHERE title like ?", likePattern)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer rows.Close()
	//id,title,content,create_time,update_time,status,uid

	var notes []model.XNote

	for rows.Next() {
		var id, status, uid int
		var title, content, createdTime, updatedTime string
		err := rows.Scan(&id, &title, &content, &createdTime, &updatedTime, &status, &uid)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		notes = append(notes, helper.BuildXNote(id, title, content, createdTime, updatedTime, status, uid))
	}

	return notes
}

func GetByID(id int) model.XNote {
	db := conndb()
	stmtOut, err := db.Prepare("SELECT title,content,create_time,update_time,status,uid FROM xnote_xnote WHERE id =?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//id,title,content,create_time,update_time,status,uid
	var status, uid int
	var title, content, createdTime, updatedTime string
	err = stmtOut.QueryRow(id).Scan(&title, &content, &createdTime, &updatedTime, &status, &uid)
	// err = db.QueryRow("SELECT id,title,content,create_time,update_time,status,uid FROM xnote_xnote WHERE id = 1").Scan(&title, &content)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// fmt.Printf("-- the query title is %s;;;%s \n", title, content)
	defer stmtOut.Close()
	defer db.Close()

	return helper.BuildXNote(id, title, content, createdTime, updatedTime, uid, status)
}
