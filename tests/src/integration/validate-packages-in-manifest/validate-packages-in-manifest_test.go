// +build integration

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tests

import (
	"github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/common"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPackagesInManifest(t *testing.T) {
	wskdeploy := common.NewWskdeploy()
	_, err := wskdeploy.Deploy(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")
	_, err = wskdeploy.Undeploy(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to undeploy based on the manifest and deployment files.")
}

var (
	manifestPath   = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/validate-packages-in-manifest/manifest.yaml"
	deploymentPath = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/validate-packages-in-manifest/deployment.yaml"
)
