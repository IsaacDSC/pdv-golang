package settings

import (
	"log"

	"sync"

	"github.com/IsaacDSC/pdv-golang/src/models"
)

var lock = &sync.Mutex{}
var template *models.Template

func GetTemplateSingleton() *models.Template {
	lock.Lock()
	defer lock.Unlock()
	if template == nil {
		template = &models.Template{}
		err := template.GetTemplate()
		if err != nil {
			log.Fatalf(err.Error())
		}
		return template
	}
	return template
}
