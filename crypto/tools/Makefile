# Copyright ArxanFintech Technology Ltd. 2018 All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#		 http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# -------------------------------------------------------------
# This makefile defines the following targets
#
#   - all (default) - builds all targets and runs all tests/checks
#   - checks - runs all tests/checks
#   - unit-test - runs the go-test based unit tests
#   - gotools - installs go tools like golint
#   - linter - runs all code checks

PROJECT_NAME=arxanchain/sdk-go-common/crypto/tools
PKGNAME = github.com/$(PROJECT_NAME)

CGO_FLAGS = CGO_CFLAGS=" " CGO_LDFLAGS="-lstdc++ -lm -lz -lbz2 -lsnappy"
EXT_LDFLAGS= --ldflags '-extldflags "-lstdc++ -lm -static"'

all: srvc

srvc: crypto-util sign-util

.PHONY: crypto-util
crypto-util: build/bin/crypto-util

.PHONY: sign-util
sign-util: build/bin/sign-util

build/bin:
	mkdir -p $@

build/bin/%:
	@mkdir -p $(@D)
	@echo "$@"
	$(CGO_FLAGS) GOBIN=$(abspath $(@D)) go install $(EXT_LDFLAGS) $(PKGNAME)/$(@F)
	@echo "Binary available as $@"
	@touch $@

.PHONY: clean
	-@rm -rf build ||:

