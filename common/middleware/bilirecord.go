package middleware

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	UserInfo   UserInfo   `json:"user_info"`
	RoomInfo   RoomInfo   `json:"room_info"`
	TaskStatus TaskStatus `json:"task_status"`
}
type UserInfo struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Face   string `json:"face"`
	UID    int    `json:"uid"`
	Level  int    `json:"level"`
	Sign   string `json:"sign"`
}

type RoomInfo struct {
	UID            int    `json:"uid"`
	RoomID         int    `json:"room_id"`
	ShortRoomID    int    `json:"short_room_id"`
	AreaID         int    `json:"area_id"`
	AreaName       string `json:"area_name"`
	ParentAreaID   int    `json:"parent_area_id"`
	ParentAreaName string `json:"parent_area_name"`
	LiveStatus     int    `json:"live_status"`
	LiveStartTime  int    `json:"live_start_time"`
	Online         int    `json:"online"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	Tags           string `json:"tags"`
	Description    string `json:"description"`
}

type TaskStatus struct {
	MonitorEnabled         bool            `json:"monitor_enabled"`
	RecorderEnabled        bool            `json:"recorder_enabled"`
	RunningStatus          string          `json:"running_status"`
	StreamURL              string          `json:"stream_url"`
	StreamHost             string          `json:"stream_host"`
	DLTotal                int             `json:"dl_total"`
	DLRate                 float64         `json:"dl_rate"`
	RecElapsed             float64         `json:"rec_elapsed"`
	RecTotal               int             `json:"rec_total"`
	RecRate                float64         `json:"rec_rate"`
	DanmuTotal             int             `json:"danmu_total"`
	DanmuRate              float64         `json:"danmu_rate"`
	RealStreamFormat       json.RawMessage `json:"real_stream_format"`
	RealQualityNumber      json.RawMessage `json:"real_quality_number"`
	RecordingPath          string          `json:"recording_path"`
	PostprocessorStatus    string          `json:"postprocessor_status"`
	PostprocessingPath     json.RawMessage `json:"postprocessing_path"`
	PostprocessingProgress json.RawMessage `json:"postprocessing_progress"`
}

func RecRequest(method string, url string, key string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("x-api-key", key)
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	respbody, err := ioutil.ReadAll(resp.Body)
	return respbody, err
}

func decodeCharset(data []byte, charsets string) ([]byte, error) {
	r, err := charset.NewReaderLabel(charsets, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
