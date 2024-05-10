package db

import (
	"api/config"
	"api/model"
	"errors"

	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataJadwal(requestData model.JadwalMatkul) error {
	db := mongo.MongoConnect(DBATS)
	insertedID := mongo.InsertOneDoc(db, config.JadwalColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func GetDataJadwalFilter(filter bson.M) ([]model.JadwalMatkul, error) {
	db := mongo.MongoConnect(DBATS)
	var datas []model.JadwalMatkul
	err := mongo.GetAllDocByFilter[model.JadwalMatkul](db, config.JadwalColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneDataJadwalFilter(filter bson.M) (model.JadwalMatkul, error) {
	db := mongo.MongoConnect(DBATS)
	var data model.JadwalMatkul
	err := mongo.GetOneDocByFilter[model.JadwalMatkul](db, config.JadwalColl, filter, &data)
	if err != nil {
		return model.JadwalMatkul{}, err
	}
	return data, nil
}
