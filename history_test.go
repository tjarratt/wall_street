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

		Context("with history", func() {
			BeforeEach(func() {
				AddHistory("This is not a test...")
				AddHistory("There is nothing wrong with your television")
			})

			It("should navigate to older history", func() {
				Expect(CurrentHistory().Line).To(Equal("There is nothing wrong with your television"))
				Expect(PreviousHistory().Line).To(Equal("This is not a test..."))
			})

			It("should navigate to more recent history", func() {
				AddHistory("We are controlling transmission")
				PreviousHistory()

				Expect(NextHistory().Line).To(Equal("We are controlling transmission"))
			})
		})
	})

	Describe("getting history by index", func() {
		BeforeEach(func() {
			AddHistory("if we wish to make it louder")
			AddHistory("we will bring up the volume")
			PreviousHistory() // back to the first
		})

		It("should return the line with that index", func() {
			Expect(HistoryGet(0).Line).To(Equal("if we wish to make it louder"))
		})

		It("should be relative to current history", func() {
			NextHistory()
			Expect(HistoryGet(0).Line).To(Equal("we will bring up the volume"))
		})
	})
})
