#
# Adapted from https://github.com/terraform-providers/terraform-provider-kubernetes
#
# Copyright 2018 Robert Patrick <rhpatrick@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
PKG_NAME=terraform-provider-environment

test: fmtcheck
	go test -v -coverprofile=coverage.out './...'

fmtcheck:
	@chmod +x ./scripts/gofmtcheck.sh
	@sh -c ./scripts/gofmtcheck.sh

errcheck:
	@chmod +x ./scripts/errcheck.sh
	@sh -c ./scripts/errcheck.sh

compile:
	@go build -o build/$(PKG_NAME)

cp:
	@mkdir -p ~/.terraform.d/plugins/terraform.com/provider/environment/1.0.0/$(go env GOHOSTOS)_$(go env GOHOSTARCH)
	@cp build/$(PKG_NAME) ~/.terraform.d/plugins/terraform.com/provider/environment/1.0.0/$(go env GOHOSTOS)_$(go env GOHOSTARCH)
