package users_test

import (
	"genVoice/app/middlewares"
	"genVoice/business"
	"genVoice/business/users"
	usersMock "genVoice/business/users/mocks"
	"genVoice/helper/encrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersMockRepo usersMock.Repository
	jwtAuth       *middlewares.ConfigJWT
	usersService users.Service
	hashedPassword string
	hashedPasswordUpdate string
	usersData users.Domain
	updateData users.UpdateDomain
)

func TestMain(m *testing.M){
	jwtAuth = &middlewares.ConfigJWT{
		SecretJWT:       "rahasiamock123",
		ExpiresDuration: 1,
	}
	hashedPassword, _ = encrypt.HashPassword("mockpassword")
	hashedPasswordUpdate, _ = encrypt.HashPassword("mockpasswordupdate")
	usersData = users.Domain{
		ID			: 1,
		Username	: "Mockusername",
		Name 		: "Mockname",
		Password  	: hashedPassword,
		Email     	: "mockemail@mock.com",
		Address   	: "mock address no.1",
		Token 		: "Mocktoken",
		CreatedAt 	:time.Now(),
		UpdatedAt 	:time.Now(),
	}
	updateData = users.UpdateDomain{
		ID			: 1,
		Name 		: "Mocknameupdate",
		Password  	: hashedPasswordUpdate,
		Email     	: "mockemail@mock.com",
		Address   	: "mock address no.1",
	}
	usersService = users.NewUserService(&usersMockRepo,2,jwtAuth)
	m.Run()
}

func TestLogin(t *testing.T){
	t.Run("Test Case 1 | Valid Login User", func(t *testing.T){
		usersMockRepo.On("Login",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(usersData, nil).Once()

		res,err := usersService.Login("Mockusername","mockpassword")
		assert.NotNil(t, res)
		assert.Nil(t,err)
	})
	t.Run("Test Case 2 | Invalid Login User", func(t *testing.T){
		usersMockRepo.On("Login",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(users.Domain{}, assert.AnError).Once()

		res,err := usersService.Login("Mockusername","mockpassword")
		assert.Equal(t, users.Domain{}, res)
		assert.NotNil(t,err)
	})
	t.Run("Test Case 3 | Invalid Password User", func(t *testing.T){
		usersMockRepo.On("Login",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(users.Domain{}, business.ErrUser).Once()

		res,err := usersService.Login("Mockusername","notmock")
		valid := encrypt.CheckPasswordHash("notmock","notmock")
		assert.Equal(t,valid,false)
		assert.Equal(t, users.Domain{}, res)
		assert.NotNil(t,err)
	})
	t.Run("Test Case 4 | Empty Username Login User", func(t *testing.T){
		usersMockRepo.On("Login",mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()

		res,err := usersService.Login("","mockpassword")
		assert.Empty(t, res)
		assert.NotNil(t,err)
	})
	t.Run("Test Case 5 | Empty Password Login User", func(t *testing.T){
		usersMockRepo.On("Login",mock.AnythingOfType("string"),"").Return(nil, assert.AnError).Once()

		res,err := usersService.Login("Mockusername","")
		assert.Empty(t, res)
		assert.NotNil(t,err)
	})
}

func TestRegister(t *testing.T){
	t.Run("Test Case 1 | Valid Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(usersData, nil).Once()

		res,err := usersService.Register(&usersData)
		assert.NotNil(t, res)
		assert.Nil(t,err)
	})
	t.Run("Test Case 2 | Invalid Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		res,err := usersService.Register(&usersData)
		assert.Equal(t, users.Domain{}, res)
		assert.NotNil(t,err)
	})
	t.Run("Test Case 3 | Empty Name Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		usersData.Name = ""
		res,err := usersService.Register(&users.Domain{ID			: 1,
			Username	: "Mockusername",
			Name 		: "",
			Password  	: hashedPassword,
			Email     	: "mockemail@mock.com",
			Address   	: "mock address no.1",
			Token 		: "Mocktoken",
			CreatedAt 	:time.Now(),
			UpdatedAt 	:time.Now(),})
		assert.Equal(t, res,users.Domain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 4 | Empty Username Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		res,err := usersService.Register(&users.Domain{ID			: 1,
			Username	: "",
			Name 		: "Mockname",
			Password  	: hashedPassword,
			Email     	: "mockemail@mock.com",
			Address   	: "mock address no.1",
			Token 		: "Mocktoken",
			CreatedAt 	:time.Now(),
			UpdatedAt 	:time.Now(),})
		assert.Equal(t, res,users.Domain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 5 | Empty Email Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		res,err := usersService.Register(&users.Domain{ID			: 1,
			Username	: "Mockusername",
			Name 		: "Mockname",
			Password  	: hashedPassword,
			Email     	: "",
			Address   	: "mock address no.1",
			Token 		: "Mocktoken",
			CreatedAt 	:time.Now(),
			UpdatedAt 	:time.Now(),})
		assert.Equal(t, res,users.Domain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 6 | Empty Address Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		res,err := usersService.Register(&users.Domain{ID			: 1,
			Username	: "Mockusername",
			Name 		: "Mockname",
			Password  	: hashedPassword,
			Email     	: "mockemail@mock.com",
			Address   	: "",
			Token 		: "Mocktoken",
			CreatedAt 	:time.Now(),
			UpdatedAt 	:time.Now(),})
		assert.Equal(t, res,users.Domain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 7 | Empty Password Register User", func(t *testing.T){
		usersMockRepo.On("Register",mock.Anything).Return(users.Domain{}, assert.AnError).Once()
		res,err := usersService.Register(&users.Domain{ID			: 1,
			Username	: "Mockusername",
			Name 		: "Mockname",
			Password  	: "",
			Email     	: "mockemail@mock.com",
			Address   	: "mock address no.1",
			Token 		: "Mocktoken",
			CreatedAt 	:time.Now(),
			UpdatedAt 	:time.Now(),})
		assert.Equal(t, res,users.Domain{})
		assert.NotNil(t,err)
	})
}

func TestUpdate(t *testing.T){
	t.Run("Test Case 1 | Valid Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(updateData, nil).Once()

		res,err := usersService.Update(&updateData)
		assert.NotNil(t, res)
		assert.Nil(t,err)
	})
	t.Run("Test Case 2 | Invalid Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(users.UpdateDomain{}, assert.AnError).Once()
		res,err := usersService.Update(&updateData)
		assert.Equal(t, users.UpdateDomain{}, res)
		assert.NotNil(t,err)
	})
	t.Run("Test Case 3 | Empty Name Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(users.UpdateDomain{}, assert.AnError).Once()
		usersData.Name = ""
		res,err := usersService.Update(&users.UpdateDomain{
			ID			: 1,
			Name 		: "",
			Password  	: hashedPasswordUpdate,
			Email     	: "mockemail@mock.com",
			Address   	: "mock address no.1",
		})
		assert.Equal(t, res,users.UpdateDomain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 4 | Empty Email Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(users.UpdateDomain{}, assert.AnError).Once()
		res,err := usersService.Update(&users.UpdateDomain{
			ID			: 1,
			Name 		: "Mocknameupdate",
			Password  	: hashedPasswordUpdate,
			Email     	: "",
			Address   	: "mock address no.1",
		})
		assert.Equal(t, res,users.UpdateDomain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 5 | Empty Address Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(users.UpdateDomain{}, assert.AnError).Once()
		res,err := usersService.Update(&users.UpdateDomain{
			ID			: 1,
			Name 		: "Mocknameupdate",
			Password  	: hashedPasswordUpdate,
			Email     	: "mockemail@mock.com",
			Address   	: "",
		})
		assert.Equal(t, res,users.UpdateDomain{})
		assert.NotNil(t,err)
	})
	t.Run("Test Case 6 | Empty Password Update User", func(t *testing.T){
		usersMockRepo.On("Update",mock.Anything).Return(users.UpdateDomain{}, assert.AnError).Once()
		res,err := usersService.Update(&users.UpdateDomain{
			ID			: 1,
			Name 		: "Mocknameupdate",
			Password  	: "",
			Email     	: "mockemail@mock.com",
			Address   	: "mock address no.1",
		})
		assert.Equal(t, res,users.UpdateDomain{})
		assert.NotNil(t,err)
	})
}
