package controllers

import (
	"encoding/json"
	"github.com/CALLlA-74/cashing/internal/domain/Cassette/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (s *Server) changeMoney(context *gin.Context) {
	reqBytes, err := io.ReadAll(context.Request.Body)
	_ = context.Request.Body.Close()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	req := dto.ChangeMoneyReq{}
	if err := json.Unmarshal(reqBytes, &req); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	log.Debug(req)
	resp, e := s.uc.ChangeMoney(req)
	if e != nil {
		context.JSON(http.StatusInternalServerError, nil)
		return
	}

	log.Debug(resp)
	context.JSON(http.StatusOK, resp)
}
