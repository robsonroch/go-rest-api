package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func TestNewCampaign(t *testing.T) {

	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(2, len(campaign.Contacts))
	assert.Equal("email1@e.com", campaign.Contacts[0].Email)
	assert.Equal("email2@e.com", campaign.Contacts[1].Email)
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {

	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNill(t *testing.T) {

	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

