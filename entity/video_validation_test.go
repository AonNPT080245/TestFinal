package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestVideo(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Name cannot blank", func(t *testing.T) {
		video := Video{
			Name: "",
			Url:  "https://www.youtube.com/watch",
		}

		ok, err := govalidator.ValidateStruct(video)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot blank"))
	})

	t.Run("check Url invalid", func(t *testing.T) {
		video := Video{
			Name: "Name",
			Url:  "wwwyoutubecom/watch",
		}

		ok, err := govalidator.ValidateStruct(video)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url Invalid"))
	})
}
