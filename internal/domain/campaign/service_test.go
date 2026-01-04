package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test@example.com"},
	}	

	_, err := service.Create(newCampaign)
	
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test@example.com"},
	}	

	_, err := service.Create(newCampaign)
	
	assert.Nil(err)
}