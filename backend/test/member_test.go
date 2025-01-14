package test

import(
	"testing"

	entity "example.com/sa-67-example/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestMemberValidetion(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`true`, func(t *testing.T) {
		member := entity.Member{
			Username: "test",
			Password: "123456",
			Email:    "test@gmail.com",
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).To(BeTrue()) // ok ต้องเป็นค่า true
		g.Expect(err).To(BeNil()) // err ต้องเป็นค่า nil
	})
}

func TestUserNameValidetion(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`user_name is required`, func(t *testing.T) {
		member := entity.Member{
			Username: "", // ผิดตรงนี้
			Password: "123456",
			Email:    "test@gmail.com",
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(member)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).NotTo(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Username is required."))
	})
}

func TestEmailValidetion(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`email is required`, func(t *testing.T) {
		member := entity.Member{
			Username: "test",
			Password: "123456",
			Email:    "test", //ผิดตรงนี้
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(member)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).NotTo(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Email is invalid."))
	})
}