package wall_street_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/wall_street"
	. "github.com/tjarratt/wall_street/testhelpers"
	"github.com/tjarratt/wall_street/testhelpers/fakes"
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

		Describe("masking user input", func() {
			var (
				stdout    []string
				theAnswer string
				fakeEcho  *fakes.Echoer
			)

			BeforeEach(func() {
				reader.MaskUserInput = true
				fakeEcho = &fakes.Echoer{}
				reader.Echo = fakeEcho

				stdout = SimulatePipes(reader, "terrible secrets", func() {
					theAnswer = reader.Readline("Tell me a saucy secret")
				})
			})

			It("changes characters passed to stdout to asterisks", func() {
				Expect(stdout).To(Equal([]string{"****************"}))
			})

			It("returns user input, unchanged", func() {
				Expect(theAnswer).To(Equal("terrible secrets"))
			})

			It("disables the echo in the tty", func() {
				Expect(fakeEcho.DisableCallCount()).To(Equal(1))
			})

			It("enables the echo in the tty when it's done", func() {
				Expect(fakeEcho.EnableCallCount()).To(Equal(1))
				Expect(fakeEcho.CurrentState).To(Equal("enabled"))
			})
		})

		Describe("control characters", func() {
			It("omits them from the output", func() {
				userInput := fmt.Sprintf("%s%s%s%s%s%s%s%sselect start", Up, Up, Down, Down, Left, Right, Left, Right)
				stdout := SimulatePipes(reader, userInput, func() {
					Expect(reader.Readline("Konami code")).To(Equal("select start"))
				})

				Expect(stdout).To(Equal([]string{"select start"}))
			})
		})
	})
})
