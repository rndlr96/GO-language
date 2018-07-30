package triangle

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Test Cases", func() {
	It("should return true", func() {
		Expect(EquableTriangle(5, 12, 13)).To(Equal(true))
	})
	It("should return false", func() {
		Expect(EquableTriangle(2, 3, 4)).To(Equal(false))
	})
})
