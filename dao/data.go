package dao

import (
	"context"
	"github.com/GrokkingSystemDesign/shortURL/config"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type URLData struct {
	ID    int
	Did   string
	Value string
}

func InsertData(ctx context.Context, data *URLData) (int, error) {
	begin := time.Now()
	_, err := config.ConfDB.Context(ctx).InsertOne(URLData{Did: data.Did, Value: data.Value})
	if err != nil { // 特殊判断1062错误码
		if e, ok := err.(*mysql.MySQLError); ok {
			if e.Number == config.DataAlreadyExist {
				return int(e.Number), nil
			}
		}
		log.Printf("%s|insertMysqlData insertOne error|%s|%s", data.Did, err, time.Since(begin))
		return 0, err
	}
	return 0, nil
}

func GetData(ctx context.Context, did string) (string, error) {
	begin := time.Now()
	var data []*URLData
	err := config.ConfDB.Context(ctx).Where("f_did = ?", did).Find(&data)
	if err != nil {
		log.Printf("%s|getMysqlData selectOne error|%s|%s", did, err, time.Since(begin))
		return "", err
	}
	if len(data) == 0 {
		return "", nil
	}
	return data[0].Value, nil
}
