package handler

import (
	"errors"
	"github.com/dre-zouhair/mailer/internal/service"
)

func mailerHandler(templateName string) []error {
	template, targets := service.RetrieveBulkData(templateName)
	if template == nil || targets == nil {
		return []error{errors.New("Unable to retrieve " + templateName + " data")}
	}
	return service.BulkMails(template, targets)

}
