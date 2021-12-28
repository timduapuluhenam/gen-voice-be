package notifications

import (
	"genVoice/business/notifications"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

type NotifController struct {
	NotifService notifications.Service
}

func NewNotifController(service notifications.Service) *NotifController {
	return &NotifController{
		NotifService: service,
	}
}

func (ctrl *NotifController) GetNotif(c echo.Context) error {
	var reqPayload map[string]interface{}

	json.NewDecoder(c.Request().Body).Decode(&reqPayload)

	status := reqPayload["transaction_status"]
	siganture_key := reqPayload["signature_key"]

	ctrl.NotifService.GetNotif(status.(string), siganture_key.(string))

	return nil
}
