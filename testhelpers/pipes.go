package testhelpers

import(
	. "github.com/onsi/gomega"
	"bytes"
  "io"
	"os"
  "strings"
  "github.com/tjarratt/wall_street"
)

func SimulatePipes(reader *wall_street.ReadlineReader, input string, block func()) []string {
	in, out, err := os.Pipe()
	Expect(err).NotTo(HaveOccurred())
	reader.SetReadPipe(in)

	go func() {
		defer out.Close()
		out.Write([]byte(input + "\n"))
	}()

	return CaptureSTDOUT(reader, func() { block() })
}

func CaptureSTDOUT(reader *wall_street.ReadlineReader, block func()) []string {
	in, out, err := os.Pipe()
	Expect(err).ToNot(HaveOccurred())

	reader.SetWritePipe(out)

	block()
	out.Close()

	var buf bytes.Buffer
	io.Copy(&buf, in)
	if len(buf.String()) == 0 {
		return []string{}
	}

	return strings.Split(buf.String(), "\n")
}
