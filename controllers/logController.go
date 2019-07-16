package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"../models"
	"../repository"
	"github.com/gin-gonic/gin"
)

const logPath = "/logs/"

func AddLog(log models.LogModel) {
	if !repository.IsInitialized {
		panic("DB is not initialized")
	}
	repository.GetDb().Save(log)
}

func GetSessions(c *gin.Context) {
	var response []models.LogModel

	repository.GetDb().Find(&response)

	var reversedResponse []models.LogModel

	for i := len(response) - 1; i >= 0; i-- {
		reversedResponse = append(reversedResponse, response[i])
	}

	c.JSON(http.StatusOK, reversedResponse)

}

func GetLog(c *gin.Context) {
	var sessionId = c.Param("id")

	var session models.LogModel

	repository.GetDb().Where("session_id = ?", sessionId).First(&session)

	if session.ID == 0 {
		c.Data(http.StatusBadRequest, "text/plain; charset=utf-8", []byte("No session found"))
		return
	}

	var path = session.SessionPath

	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.Data(http.StatusBadRequest, "text/plain; charset=utf-8", []byte("No logs found"))
		return
	}

	file, err := os.Open(path)
	if err != nil {
		c.Data(http.StatusBadRequest, "text/plain; charset=utf-8", []byte("Error while reading log"))
		return
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	c.Data(http.StatusOK, "text/plain; charset=utf-8", b)
}

func AppendToLog(c *gin.Context) {
	var response models.AppendToLog

	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusCreated,
			gin.H{
				"status:": http.StatusBadRequest,
				"message": "Some fields are missing"})
		return
	}

	var session models.LogModel

	var sessionId = response.SessionId

	repository.GetDb().Where("session_id = ?", sessionId).First(&session)

	if session.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No session with this ID found"})
		return
	}

	repository.GetDb().Model(&session).Update("session_last_active", time.Now().Unix())

	var message = response.LogMessage

	var err = AppendMessageToLog(session.SessionPath, message)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Log succesfully added"})

}

func AddNewLog(c *gin.Context) {
	var response models.LogModel
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusCreated, gin.H{"status:": http.StatusBadRequest, "message": "Some fields are missing"})
		return
	}

	var appId = response.AppId
	var deviceId = response.DeviceId
	var sessionLastActive = time.Now().Unix()
	var sessionPath, sessionId = GenerateSessionPath(appId, deviceId)

	log := models.LogModel{
		AppId:             appId,
		DeviceId:          deviceId,
		SessionId:         sessionId,
		SessionLastActive: sessionLastActive,
		SessionPath:       sessionPath,
	}

	if appId == "" || deviceId == "" {
		c.JSON(http.StatusCreated, gin.H{"status:": http.StatusBadRequest, "message": "Some fields are missing"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status:": http.StatusOK, "sessionId": sessionId})
		repository.GetDb().Save(&log)
	}
}

func AppendMessageToLog(path string, message string) error {

	var directory = path[0:strings.LastIndex(path, "/")]

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return errors.New("Path for this session does not exist.")
	}

	sessionFile, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer sessionFile.Close()

	if _, err := sessionFile.WriteString(message + "\n"); err != nil {
		return err
	}
	return nil
}

func GenerateSessionPath(appId string, deviceId string) (string, uint) {
	var path = logPath + appId + "/" + deviceId + "/"

	dir, _ := os.Getwd()

	path = dir + path

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
	var timeStamp = int(time.Now().Unix())
	path += strconv.Itoa(timeStamp)
	return path, uint(timeStamp)
}
