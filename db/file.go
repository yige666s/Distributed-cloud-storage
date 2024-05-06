package db

import (
	"Distributed-cloud-storage/db/mysql"
	"database/sql"
	"fmt"
	"log"
)

// 文件上传
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mysql.BDConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("failed to prepare statment,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if rf, err := ret.RowsAffected(); nil == err { // 判断是否有重复hash
		if rf < 0 {
			fmt.Printf("File with hash: %s has been uploaded before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// 从mysql获取文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mysql.BDConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1= ? and status = 1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile, &tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}

// UpdateFileLocation : 更新文件的存储地址(如文件被转移了)
func UpdateFileLocation(filehash string, fileaddr string) (res ExecResult) {
	stmt, err := mysql.BDConn().Prepare(
		"update tbl_file set`file_addr`=? where  `file_sha1`=? limit 1")
	if err != nil {
		log.Println("预编译sql失败, err:" + err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	ret, err := stmt.Exec(fileaddr, filehash)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}

	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			log.Printf("更新文件location失败, filehash:%s", filehash)
			res.Suc = false
			res.Msg = "无记录更新"
			return
		}
		res.Suc = true
		return
	} else {
		res.Suc = false
		res.Msg = err.Error()
		return
	}
}
