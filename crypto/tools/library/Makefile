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

EXT_LDFLAGS='-buildmode=c-shared'

all: utils

.PHONY: utils
utils: build/utils.so

build:
	@mkdir -p $@

build/%:
	@mkdir -p $(@D)
	@echo "$@"
	go build $(EXT_LDFLAGS) -o $@ "./utils.go"
	@echo "Libraries available as $@"
	@touch $@

.PHONY: clean
	-@rm -rf build ||:

