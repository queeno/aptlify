package gpg

import (
	"fmt"
	ctx "github.com/queeno/aptlify/context"
	"github.com/queeno/aptlify/utils"
	"strings"
)

func extractFingerprints(output []string) []string {
	var s []string
	for _, line := range output {
		if strings.HasPrefix(line, "fpr") {
			s = append(s, line[36:52])
		}
	}
	return s
}

func keyLoaded(keyFingerprint string) bool {

	ctx.Logging.Info.Println(fmt.Sprintf("Can we find %s", keyFingerprint))
	fingerprintArray, _ := utils.Exec("apt-key finger --with-colons")
	for _, extractedFingerprint := range extractFingerprints(fingerprintArray) {
		if keyFingerprint == extractedFingerprint {
			return true
		}
	}
	return false
}
