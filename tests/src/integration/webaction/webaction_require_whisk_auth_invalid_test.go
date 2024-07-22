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

const (
	ANNOTATION_ERROR = "ERROR_ACTION_ANNOTATION"
	MSG_ASSERT_EXPECTED_ERROR = "Expected error [%s], but got:\n [%v]\n"
)

var (
	manifestStringEmpty = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/webaction/manifest_require_whisk_auth_invalid_str_empty.yaml"
	manifestStringNil = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/webaction/manifest_require_whisk_auth_invalid_str_nil.yaml"
	manifestIntTooBig = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/webaction/manifest_require_whisk_auth_invalid_int_big.yaml"
	manifestIntNegative = os.Getenv("GOPATH") + "/src/github.com/sciabarracom/openwhisk-wskdeploy/tests/src/integration/webaction/manifest_require_whisk_auth_invalid_int_neg.yaml"
)

func TestRequireWhiskAuthAnnotationInvalid(t *testing.T) {
	wskdeploy := common.NewWskdeploy()

	_, err1 := wskdeploy.DeployManifestPathOnly(manifestStringEmpty)
	assert.Contains(t, err1.Error(), ANNOTATION_ERROR)

	_, err2 := wskdeploy.DeployManifestPathOnly(manifestStringNil)
	assert.Contains(t, err2.Error(), ANNOTATION_ERROR)

	_, err3 := wskdeploy.DeployManifestPathOnly(manifestIntTooBig)
	assert.Contains(t, err3.Error(), ANNOTATION_ERROR)

	_, err4 := wskdeploy.DeployManifestPathOnly(manifestIntNegative)
	assert.Contains(t, err4.Error(), ANNOTATION_ERROR)
}
