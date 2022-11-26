#!/usr/bin/env bash
#
# Adapted from https://github.com/terraform-providers/terraform-provider-kubernetes
#
# Copyright 2022 Robert Patrick <rhpatrick@gmail.com>
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
#
# Update gofmt
echo "==> Reformatting code to comply with gofmt requirements..."
gofmt_files=$(gofmt -l -w `find . -name '*.go' | grep -v vendor`)
if [[ -n "${gofmt_files}" ]]; then
    echo 'gofmt reformatted the following files:'
    echo "${gofmt_files}"
else
    echo "No files required reformatting"
fi
