package jenkins

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"hash"
)

var _ = Describe("Jenkins", func() {

	var jhash hash.Hash32
	var key []byte

	Describe("New", func() {

		It("returns jenkhash", func() {
			var h *jenkhash
			hash := New()
			Expect(hash).To(BeAssignableToTypeOf(h))
		})
	})

	Describe("Write", func() {

		BeforeEach(func() {
			jhash = New()
			key = []byte("Apple")
		})

		It("returns key length", func() {
			length, _ := jhash.Write(key)
			Expect(length).To(Equal(5))
		})

		It("has no error", func() {
			_, err := jhash.Write(key)
			Expect(err).To(BeNil())
		})

	})

	Describe("Size", func() {

		It("is 4", func() {
			jhash = New()
			Expect(jhash.Size()).To(Equal(4))
		})

	})

	Describe("BlockSize", func() {

		It("is 1", func() {
			jhash = New()
			Expect(jhash.BlockSize()).To(Equal(1))
		})

	})

})
