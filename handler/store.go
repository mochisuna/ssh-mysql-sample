package handler

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mochisuna/ssh-mysql-sample/domain"
)

// response format
type storeResponse struct {
	ID        domain.StoreID `json:"id"`
	UID       string         `json:"uid"`
	Name      string         `json:"name"`
	Status    int            `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// csv format
type csvFormat struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
	UID  string `csv:"uid"`
}

// IDを指定してストア情報を取得
func (h *Handler) GetStore(storeID domain.StoreID) (string, error) {
	res, err := h.StoreService.Get(storeID)
	if err != nil {
		return "", err
	}
	sr := storeResponse{
		ID:        res.ID,
		UID:       res.UID,
		Name:      res.Name,
		Status:    res.Status,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	jsonResp, err := json.Marshal(sr)
	return string(jsonResp), err
}

// 全てのストア情報を取得
func (h *Handler) GetStores() (string, error) {
	resposes, err := h.StoreService.GetList()
	if err != nil {
		return "", err
	}
	srs := []storeResponse{}
	for _, res := range resposes {
		sr := storeResponse{
			ID:        res.ID,
			UID:       res.UID,
			Name:      res.Name,
			Status:    res.Status,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		}
		srs = append(srs, sr)
	}
	jsonResp, err := json.Marshal(srs)
	return string(jsonResp), err
}

// 全てのストア情報を取得し、CSVに書き出す
func (h *Handler) OutputCSV(fileName string) error {
	out := []csvFormat{}
	store, err := h.StoreService.GetList()
	if err != nil {
		return err
	}

	for _, s := range store {
		out = append(out, csvFormat{
			ID:   int(s.ID),
			Name: s.Name,
			UID:  s.UID,
		})
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error %#v\n", err)
		return err
	}
	defer file.Close()
	gocsv.MarshalFile(out, file)
	return nil
}
