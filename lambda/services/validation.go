package services

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/ajaxe/route53updater/cli/shared"
)

const nonceSkew = 45

func ValidateRequest(payload *shared.Payload, preSharedKey string) error {
	now := time.Now().Unix()
	nonce, err := strconv.ParseInt(payload.Nonce, 10, 64)
	if err != nil {
		return err
	}
	if now-nonce > nonceSkew {
		return fmt.Errorf("invalid nonce")
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s%s", payload.Nonce, preSharedKey)))
	if fmt.Sprintf("%x", sum) != payload.HashKey {
		return fmt.Errorf("invalid hash key")
	}
	return nil
}
