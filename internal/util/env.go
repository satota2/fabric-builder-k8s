// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"os"
)

func GetOptionalEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func GetRequiredEnv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}

	return "", fmt.Errorf("environment variable not set: %s", key)
}