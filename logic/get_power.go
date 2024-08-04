package logic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"power_warning/conf"
)

type GetPowerResp struct {
	E int    `json:"e"`
	M string `json:"m"`
	D struct {
		Data struct {
			Phone           string  `json:"phone"`
			FloorName       int     `json:"floorName"`
			Model           string  `json:"model"`
			Time            string  `json:"time"`
			VTotal          string  `json:"vTotal"`
			Price           string  `json:"price"`
			ITotal          string  `json:"iTotal"`
			ParName         string  `json:"parName"`
			FreeEnd         int     `json:"freeEnd"`
			CosTotal        string  `json:"cosTotal"`
			PTotal          string  `json:"pTotal"`
			Surplus         float64 `json:"surplus"`
			TotalActiveDisp string  `json:"totalActiveDisp"`
		} `json:"data"`
	} `json:"d"`
}

func GetPower(roomConfig conf.RoomConfig) (*GetPowerResp, error) {

	url := "https://app.bupt.edu.cn/buptdf/wap/default/search"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("partmentId", roomConfig.PartmentId)
	_ = writer.WriteField("floorId", roomConfig.FloorId)
	_ = writer.WriteField("dromNumber", roomConfig.DromNumber)
	_ = writer.WriteField("areaid", roomConfig.AreaId)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", roomConfig.Cookie)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "app.bupt.edu.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------692284755493491450325826")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	power := &GetPowerResp{}
	json.Unmarshal(body, &power)
	return power, nil
}
