package mapstore_test

import (
	"testing"

	"github.com/artemxgod/11-go-projects/web-server/internal/app/model"
	"github.com/artemxgod/11-go-projects/web-server/internal/app/store/mapstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateAndFind(t *testing.T) {
	s := mapstore.New()
	u := model.TestUser(t)

	s.User().Create(u)
	res, err := s.User().Find(1)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, res.Name, u.Name)

}