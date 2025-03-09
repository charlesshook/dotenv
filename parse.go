package dotenv

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func parsefile(filename string) (map[string]string, error) {
	environmentVariables := make(map[string]string)

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 || strings.HasPrefix(text, "#") {
			continue
		}

		text = removeInlineComment(text)
		text = strings.TrimSpace(text)

		textSplits := strings.SplitN(text, "=", 2)

		if len(textSplits) != 2 {
			continue
		}

		textSplits[1] = removeQuotes(textSplits[1])

		environmentVariables[textSplits[0]] = textSplits[1]
	}

	return environmentVariables, nil
}

func removeInlineComment(value string) string {
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) ||
		(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
		return value
	}

	return regexp.MustCompile(`\s+#.*$`).ReplaceAllString(value, "")
}

func removeQuotes(value string) string {
	if len(value) > 2 && (value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'') {
		value = value[1 : len(value)-1]
	}

	return value
}
