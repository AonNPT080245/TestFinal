package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Name cannot blank", func(t *testing.T) {
		user := User{
			Name:        "",
			Email:       "aon@mail.com",
			StudentID:   "B6310264",
			PhoneNumber: "0842294356",
			Password:    "12345678",
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot blank"))
	})

	t.Run("check Email invalid format", func(t *testing.T) {
		user := User{
			Name:        "aon",
			Email:       "aoncom",
			StudentID:   "B6310264",
			PhoneNumber: "0842294356",
			Password:    "12345678",
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Email invalid format"))
	})

	t.Run("check StudendID invalid format", func(t *testing.T) {

		studenID := []string{
			"X6310264",
			"BA310264",
			"B10264",
			"B6000310264",
		}

		for _, val := range studenID {
			user := User{
				Name:        "aon",
				Email:       "aon@mail.com",
				StudentID:   val,
				PhoneNumber: "0842294356",
				Password:    "12345678",
			}
			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("StudentID invalid format"))
		}

	})

	t.Run("check PhoneNumber invalid format", func(t *testing.T) {

		phone := []string{
			"1234567890",
			"084229341",     // ขาด 1
			"0423432423234", //เกิน
			"+661234567890",
			"+6612345678",
		}

		for _, val := range phone {
			user := User{
				Name:        "aon",
				Email:       "aon@mail.com",
				StudentID:   "B6310264",
				PhoneNumber: val,
				Password:    "12345678",
			}
			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("PhoneNumber invalid format"))
		}

	})

	t.Run("check password invalid format", func(t *testing.T) {

		pass := []string{
			"0842219", // ขาด 1
			"042343242323412121212121212121212312321", //เกิน
		}

		for _, val := range pass {
			user := User{
				Name:        "aon",
				Email:       "aon@mail.com",
				StudentID:   "B6310264",
				PhoneNumber: "0842293433",
				Password:    val,
			}
			ok, err := govalidator.ValidateStruct(user)
			g.Expect(ok).ToNot(BeTrue())
			g.Expect(err).ToNot(BeNil())
			g.Expect(err.Error()).To(Equal("Password must be between 8 and 20 characters"))
		}

	})

}
