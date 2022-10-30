package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"sbercloud_test/pkg/models"
	"sbercloud_test/pkg/services"
	"strconv"
)

type Handler struct {
	service *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{service: s}
}

func writeData(writer http.ResponseWriter, data []models.Data) {
	writer.Write([]byte{'{'})
	for i, _ := range data {
		writer.Write([]byte("\"" + data[i].Key + "\""))
		writer.Write([]byte(": "))
		writer.Write([]byte("\"" + data[i].Value + "\""))
		if i != (len(data) - 1) {
			writer.Write([]byte(", "))
		}
	}
	writer.Write([]byte{'}'})
}

func (h *Handler) createConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	bytes, _ := ioutil.ReadAll(request.Body)
	var rawConfig map[string]interface{}
	json.Unmarshal(bytes, &rawConfig)

	serviceName := fmt.Sprintf("%v", rawConfig["serviceName"])
	var data []models.Data
	v := reflect.ValueOf(rawConfig["data"])
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			in := v.Index(i).Interface()
			m := reflect.ValueOf(in)
			for _, key := range m.MapKeys() {
				data = append(data, models.Data{Key: key.String(), Value: m.MapIndex(key).Elem().String()})
			}
		}
	}

	s := *h.service
	config := s.CreateServiceConfig(serviceName, data)
	json.NewEncoder(writer).Encode(config)
}

func (h *Handler) getConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.URL.Query().Has("service") {
		serviceName := request.URL.Query().Get("service")

		var data []models.Data
		s := *h.service
		if request.URL.Query().Has("version") {
			version, _ := strconv.Atoi(request.URL.Query().Get("version"))
			data = s.GetServiceConfigByVersion(serviceName, uint(version))
		} else {
			data = s.GetServiceConfig(serviceName)
		}
		writeData(writer, data)
	}
}

func (h *Handler) deleteConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.URL.Query().Has("service") {
		serviceName := request.URL.Query().Get("service")

		s := *h.service
		deleted := &models.Config{}
		if request.URL.Query().Has("version") {
			version, _ := strconv.Atoi(request.URL.Query().Get("version"))
			deleted = s.DeleteServiceConfigByVersion(serviceName, uint(version))
		} else {
			deleted = s.DeleteServiceConfig(serviceName)
		}

		json.NewEncoder(writer).Encode(deleted)
	}
}

func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/config", h.createConfigHandler).Methods("POST", "PUT")
	router.HandleFunc("/config", h.getConfigHandler).Methods("GET")
	router.HandleFunc("/config", h.deleteConfigHandler).Methods("DELETE")
}
