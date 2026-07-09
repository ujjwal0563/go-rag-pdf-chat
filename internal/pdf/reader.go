package pdf

import (
	"os"
	"strings"

	pdf "github.com/ledongthuc/pdf"
)

func ReadPDF(filePath string) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	reader, err := pdf.NewReader(file, info.Size())
	if err != nil {
		return "", err
	}

	var text strings.Builder

	numPages := reader.NumPage()

	for i := 1; i <= numPages; i++ {

		page := reader.Page(i)

		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		text.WriteString(content)
		text.WriteString("\n")
	}

	return text.String(), nil
}
