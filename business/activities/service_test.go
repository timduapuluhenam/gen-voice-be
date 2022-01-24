package activities_test

import (
	"genVoice/business/activities"
	activitiesMock "genVoice/business/activities/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	activitiesMockRepo activitiesMock.Repository
	activitiesService activities.Service
	activitiesData activities.Domain
)

func TestMain(m *testing.M){
	activitiesService = activities.NewActivityService(&activitiesMockRepo,2)
	m.Run()
}
func TestCreateActivity(t *testing.T){
	t.Run("Test Case 1 | Valid Create Activity", func(t *testing.T){
		activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, nil).Once()

		resp, err := activitiesService.CreateActivity(&activitiesData)
		assert.Nil(t, err)
		assert.Equal(t, activitiesData, resp)
	})
	t.Run("Test Case 2 | Invalid Create Activity", func(t *testing.T){
		activitiesMockRepo.On("CreateActivity",mock.Anything).Return(activitiesData, assert.AnError).Once()

		resp, err := activitiesService.CreateActivity(&activitiesData)
		assert.NotNil(t, err)
		assert.Equal(t, activitiesData, resp)
	})
}