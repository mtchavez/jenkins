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
			jhash = New()
			Expect(jhash).To(BeAssignableToTypeOf(h))
		})

		It("initializes offset to 0", func() {
			jhash = New()
			Expect(jhash.Sum32()).To(Equal(uint32(0)))
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

	Describe("Reset", func() {

		It("sets back to 0", func() {
			jhash = New()
			key = []byte("Apple")
			Expect(jhash.Sum32()).To(Equal(uint32(0)))
			jhash.Write(key)
			Expect(jhash.Sum32()).NotTo(Equal(uint32(0)))
			jhash.Reset()
			Expect(jhash.Sum32()).To(Equal(uint32(0)))
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

	Describe("Sum32", func() {

		It("defaults to 0", func() {
			jhash = New()
			Expect(jhash.Sum32()).To(Equal(uint32(0)))
		})

		It("sums hash", func() {
			jhash = New()
			key = []byte("Apple")
			jhash.Write(key)
			Expect(jhash.Sum32()).To(Equal(uint32(884782484)))
		})

	})

})
