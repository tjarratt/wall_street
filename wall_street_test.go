package wall_street_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "wall_street"
	. "wall_street/testhelpers"
)

var _ = Describe("Wall Street", func() {
	var (
		reader *ReadlineReader
	)

	BeforeEach(func() {
		reader = NewReadline()
	})

	Describe("Readline", func() {
		BeforeEach(func() {
			reader.DisablePrompt()
		})

		It("reads from a pipe and returns a string", func() {
			SimulatePipes(reader, "The return of the Archons", func() {
				readline := reader.Readline("Tonight on The Outer Limits:")
				Expect(readline).To(Equal("The return of the Archons"))
			})
		})

		Describe("echoing output back to the terminal", func() {
			It("writes to a pipe", func() {
				out := SimulatePipes(reader, "Where No Man Has Gone Before", func() {
					reader.Readline("Tonight on The Outer Limits")
				})

				Expect(out).To(Equal([]string{"Where No Man Has Gone Before"}))
			})

			It("can be suppressed", func() {
				reader.DisableEcho()

				output := SimulatePipes(reader, "Where No Man Has Gone Before", func() {
					reader.Readline("Tonight on The Outer Limits")
				})

				Expect(output).To(Equal([]string{}))
			})
		})

		Describe("The prompt", func() {
			BeforeEach(func() {
				reader.EnablePrompt()
			})

			It("is printed to stdout", func() {
				reader.DisableEcho()

				output := SimulatePipes(reader, "Lots of mutable state!", func() {
					reader.Readline("Developer, what is best in life?")
				})

				Expect(output).To(Equal([]string{"Developer, what is best in life?"}))
			})
		})

		Describe("signals", func() {

		})
	})
})
