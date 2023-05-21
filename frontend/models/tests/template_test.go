package tests

import (
	"pdv-example/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateStoreStatic(t *testing.T) {
	template := models.Template{}
	err := template.GetTemplate()
	assert.NoError(t, err)
	// TEST NAVBAR
	assert.Equal(t, true, len(template.Navbar.Logo) >= 1)
	assert.Equal(t, 4, len(template.Navbar.Actions))
	// TEST FOOTER
	assert.Equal(t, true, len(template.Footer.Contact.LinkLocation) >= 1)
	assert.Equal(t, true, len(template.Footer.Contact.LocationTitle) >= 1)
	assert.Equal(t, true, len(template.Footer.Contact.Telephone) >= 1)
	assert.Equal(t, true, len(template.Footer.Contact.Mailer) >= 1)
	assert.Equal(t, true, len(template.Footer.Operation.Title) > 1)
	assert.Equal(t, true, len(template.Footer.Operation.Hours) > 0)
	assert.Equal(t, true, len(template.Footer.Operation.Hours) <= 2)
	for index := range template.Footer.Operation.Hours {
		assert.Equal(t, true, len(template.Footer.Operation.Hours[index].Hour) > 1)
		assert.Equal(t, true, len(template.Footer.Operation.Hours[index].Subtile) > 1)
	}
	// TEST PAGES
	// assert.Equal(t, len(template.Pages), 4)
}
