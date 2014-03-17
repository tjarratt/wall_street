package wall_street_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "wall_street/testhelpers"
	. "wall_street"
)

var _ = Describe("Wall Street", func() {
	var (
		reader *ReadlineReader
	)

	BeforeEach(func() {
		reader = NewReadline()
	})

	Describe("Readline", func() {
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

				Expect(out).To(Equal([]string{"Tonight on The Outer Limits"}))
			})

			It("can be suppressed", func() {
				reader.DisableEcho()

				output := SimulatePipes(reader, "Where No Man Has Gone Before", func() {
					reader.Readline("Tonight on The Outer Limits")
				})

				Expect(output).To(Equal([]string{}))
			})
		})
	})
})
