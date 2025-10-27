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

output "security_group_id" {
  description = "ID of the test security group"
  value       = aws_security_group.test.id
}

output "egress_rule_id" {
  description = "ID of the egress rule"
  value       = module.egress_ssh.id
}

output "egress_rule_arn" {
  description = "ARN of the egress rule"
  value       = module.egress_ssh.arn
}

output "security_group_rule_id" {
  description = "AWS-assigned security group rule ID"
  value       = module.egress_ssh.security_group_rule_id
}

output "effective_source" {
  description = "Effective destination for the egress rule"
  value       = module.egress_ssh.egress_rule_effective_source
}
