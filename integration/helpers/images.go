/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/base32"
	"fmt"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/osdi23p228/fabric/common/util"
)

func AssertImagesExist(imageNames ...string) {
	dockerClient, err := docker.NewClientFromEnv()
	Expect(err).NotTo(HaveOccurred())

	for _, imageName := range imageNames {
		images, err := dockerClient.ListImages(docker.ListImagesOptions{
			Filters: map[string][]string{"reference": {imageName}},
		})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())

		if len(images) != 1 {
			Fail(fmt.Sprintf("missing required image: %s", imageName), 1)
		}
	}
}

// UniqueName generates base-32 enocded UUIDs for container names.
func UniqueName() string {
	name := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(util.GenerateBytesUUID())
	return strings.ToLower(name)
}
