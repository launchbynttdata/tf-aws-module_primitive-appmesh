// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "name" {
  description = "Name to use for the service mesh. Must be between 1 and 255 characters in length"
  type        = string
}

variable "spec_egress_filter_type" {
  description = "Egress filter type. By default, the type is DROP_ALL. Valid values are ALLOW_ALL and DROP_ALL"
  type        = string
  default     = "DROP_ALL"
}

variable "tags" {
  description = "A map of custom tags to be attached to this resource"
  type        = map(string)
  default     = {}
}
