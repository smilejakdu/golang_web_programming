package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		app.Create(CreateRequest{req.UserName, req.MembershipType})
		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})
		assert.Error(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"", "naver"}
		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})
		assert.Error(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "kakao"}
		_, err := app.Create(CreateRequest{req.UserName, req.MembershipType})
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {
		newName := "ash_update"
		newMemberShip := "naver"

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		createdUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		updatedUser, _ := app.Update(UpdateRequest{createdUser.ID, newName, newMemberShip})
		assert.Equal(t, req.UserName, updatedUser.UserName)
		assert.Equal(t, req.MembershipType, updatedUser.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		newName := "ash_update"
		newMemberShip := "naver"

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		createdUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		updatedUser, _ := app.Update(UpdateRequest{createdUser.ID, newName, newMemberShip})
		assert.Equal(t, req.UserName, updatedUser.UserName)
		assert.Equal(t, req.MembershipType, updatedUser.MembershipType)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		newName := "ash_update"
		newMemberShip := "naver"

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Update(UpdateRequest{"", newName, newMemberShip})
		assert.Nil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		newName := ""
		newMemberShip := "naver"

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		createdUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Update(UpdateRequest{createdUser.ID, newName, newMemberShip})
		assert.Nil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		newName := "ash"
		newMemberShip := ""

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		createdUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Update(UpdateRequest{createdUser.ID, newName, newMemberShip})
		assert.Nil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {

		newName := "ash"
		newMemberShip := "kakao"

		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		createdUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Update(UpdateRequest{createdUser.ID, newName, newMemberShip})
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		newUser, _ := app.Create(CreateRequest{req.UserName, req.MembershipType})

		res, err := app.Delete(newUser.ID)
		if err != nil {
			assert.Nil(t, err)
		}
		assert.Equal(t, res.ID, newUser.ID)
	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Delete("")
		assert.Error(t, err)
	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "toss"}
		app.Create(CreateRequest{req.UserName, req.MembershipType})

		_, err := app.Delete("351")
		assert.Error(t, err)
	})
}
