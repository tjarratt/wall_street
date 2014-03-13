package wall_street_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "wall_street"
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

	Describe("reading history as a slice", func() {
		var expectedHistory []string

		BeforeEach(func() {
			AddHistory("The Devil in the Dark")
			AddHistory("The Sixth Finger")
			AddHistory("The Conscience of the King")

			expectedHistory = []string{
				"The Devil in the Dark",
				"The Sixth Finger",
				"The Conscience of the King",
			}
		})

		It("should match the history list", func() {
			// navigate forward one step to get back to the beginning
			NextHistory()

			Expect(historyListToStrings()).To(Equal(expectedHistory))
		})

		It("should match the history as you walk over it", func() {
			readHistory := []string{}
			for i := 0; i < len(expectedHistory); i++ {
				readHistory = append(readHistory, NextHistory().Line)
			}

			Expect(readHistory).To(Equal(expectedHistory))
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

	Describe("overwriting history", func() {
		BeforeEach(func() {
			AddHistory("if we wish to make it softer")
			ReplaceHistoryEntry(0, "we will tune it to a whisper", nil)
		})

		It("should replace the line", func() {
			Expect(CurrentHistory().Line).To(Equal("we will tune it to a whisper"))
		})

		It("should be relative to current history", func() {
			AddHistory("We will control the horizontal")
			ReplaceHistoryEntry(0, "We will control the vertical", nil)
			Expect(CurrentHistory().Line).To(Equal("We will control the vertical"))
		})
	})

	Describe("removing history", func() {
		BeforeEach(func() {
			AddHistory("We will now return control of your television set to you")
			AddHistory("Until next week at the same time, when the control voice will take you to...")
			AddHistory("The Outer Limits")
			NextHistory()

			Expect(len(HistoryList())).To(Equal(3))
			Expect(CurrentHistory().Line).To(Equal("We will now return control of your television set to you"))
		})

		It("removes the history element at the given index", func() {
			RemoveHistory(1)
			Expect(len(HistoryList())).To(Equal(2))
			Expect(CurrentHistory().Line).To(Equal("We will now return control of your television set to you"))
			Expect(NextHistory().Line).To(Equal("The Outer Limits"))
		})

		It("is relative to current history", func() {
			RemoveHistory(-1)
			Expect(historyListToStrings()).To(Equal([]string{
				"We will now return control of your television set to you",
				"Until next week at the same time, when the control voice will take you to...",
			}))
		})
	})
})

func historyListToStrings() []string {
	list := []string{}
	history := HistoryList()
	for i := 0; i < len(history); i++ {
		list = append(list, history[i].Line)
	}
	return list
}
