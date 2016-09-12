package gpg

import (
	"errors"
	"github.com/queeno/aptlify/exec"
	"strings"
)

var execExec = exec.Exec

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
	fingerprintArray, _ := execExec("apt-key finger --with-colons")
	extractedFingerprints, _ := extractFingerprints(fingerprintArray)
	for _, extractedFingerprint := range extractedFingerprints {
		if keyFingerprint == extractedFingerprint {
			return true
		}
	}
	return false
}
