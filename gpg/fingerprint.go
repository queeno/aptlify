package gpg

import (
	"errors"
	"fmt"
	"github.com/queeno/aptlify/utils"
	"strings"
)

func extractFingerprints(output []string) ([]string, error) {
	var s []string
	for _, line := range output {
		if strings.HasPrefix(line, "fpr") {
			if len(line) < 52 {
				return nil, errors.New("malformed gpg-fingerprint returned by apt-key")
			}
			s = append(s, line[36:52])
		}
	}
	return s, nil
}

func keyLoaded(keyFingerprint string) bool {

	//ctx.Logging.Info.Println(fmt.Sprintf("Can we find %s", keyFingerprint))
	fingerprintArray, _ := utils.Exec("apt-key finger --with-colons")
	extractedFingerprints, _ := extractFingerprints(fingerprintArray)
	fmt.Println(extractedFingerprints)
	for _, extractedFingerprint := range extractedFingerprints {
		if keyFingerprint == extractedFingerprint {
			return true
		}
	}
	return false
}
