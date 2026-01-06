package testo

import (
	"test_se_final/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStudent(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		FirstName:  "Sippakorn",
		LastName:   "Bunyu",
		Email:      "JoeMaMa@gmail.com",
		Age:        22,
		Student_id: "B6504472",
	}
	t.Run("Case 1 : All valid", func(t *testing.T) {
		v := student
		ok, _ := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeTrue())
	})

	t.Run("Case 2 : FirstName invalid", func(t *testing.T) {
		v := student
		v.FirstName = "Chrliez2xxx"
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("must be alphabet"))
	})

	t.Run("Case 3 :Lastname invalid", func(t *testing.T) {
		v := student
		v.LastName = "23213ads"
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("must be alphabet"))
	})
}
