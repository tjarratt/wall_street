package wall_street_test

import (
	. "wall_street"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("History", func() {
	BeforeEach(func() {
		ResetHistory()
	})

	Describe("adding history", func() {
		BeforeEach(func() {
			AddHistory("do not adjust your television")
		})

		It("should add a record to the history", func() {
			Expect(len(HistoryList())).To(Equal(1))
		})

		It("should store the string provided", func() {
			Expect(CurrentHistory().Line).To(Equal("do not adjust your television"))
		})

		It("should store a date", func() {
			Expect(CurrentHistory().Timestamp).ToNot(Equal(0))
		})
	})

	Describe("navigating history", func() {
		It("should initially be nil", func() {
			Expect(CurrentHistory()).To(BeNil())
		})
	})
})
