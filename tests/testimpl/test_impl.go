package testimpl

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	testTypes "github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	failedToGetSecurityGroupMsg = "Failed to get security group"
	failedToFindEgressRuleMsg   = "Failed to find egress rule"
)

func TestComposableComplete(t *testing.T, ctx testTypes.TestContext) {
	ec2Client := GetAWSEC2Client(t)

	egressRuleId := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "egress_rule_id")
	securityGroupId := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "security_group_id")
	effectiveSource := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "effective_source")

	t.Run("TestSecurityGroupEgressRuleExists", func(t *testing.T) {
		testSecurityGroupEgressRuleExists(t, ec2Client, securityGroupId, egressRuleId)
	})

	t.Run("TestSecurityGroupEgressRuleProperties", func(t *testing.T) {
		testSecurityGroupEgressRuleProperties(t, ec2Client, securityGroupId, egressRuleId)
	})

	t.Run("TestEffectiveSource", func(t *testing.T) {
		testEffectiveSource(t, effectiveSource)
	})
}

func testSecurityGroupEgressRuleExists(t *testing.T, ec2Client *ec2.Client, securityGroupId, egressRuleId string) {
	// Get security group
	sg, err := ec2Client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{securityGroupId},
	})
	require.NoError(t, err, failedToGetSecurityGroupMsg)
	require.NotEmpty(t, sg.SecurityGroups, "Security group should exist")

	// Verify security group has at least one egress rule
	assert.NotEmpty(t, sg.SecurityGroups[0].IpPermissionsEgress, "Security group should have at least one egress rule")
}

func testSecurityGroupEgressRuleProperties(t *testing.T, ec2Client *ec2.Client, securityGroupId, egressRuleId string) {
	// Get security group rules
	rules, err := ec2Client.DescribeSecurityGroupRules(context.TODO(), &ec2.DescribeSecurityGroupRulesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("group-id"),
				Values: []string{securityGroupId},
			},
		},
	})
	require.NoError(t, err, "Failed to describe security group rules")

	// Find the specific egress rule by ID
	var egressRule *types.SecurityGroupRule
	for i := range rules.SecurityGroupRules {
		rule := &rules.SecurityGroupRules[i]
		if rule.SecurityGroupRuleId != nil && *rule.SecurityGroupRuleId == egressRuleId {
			egressRule = rule
			break
		}
	}

	require.NotNil(t, egressRule, "Egress rule should be found")

	// Verify basic rule properties (protocol-agnostic)
	assert.True(t, *egressRule.IsEgress, "Rule should be an egress rule")
	assert.NotNil(t, egressRule.IpProtocol, "Rule should have a protocol specified")

	// Verify at least one destination is set
	hasDestination := egressRule.CidrIpv4 != nil ||
		egressRule.CidrIpv6 != nil ||
		egressRule.PrefixListId != nil ||
		egressRule.ReferencedGroupInfo != nil
	assert.True(t, hasDestination, "Rule should have at least one destination (CIDR, prefix list, or SG)")
}

func testEffectiveSource(t *testing.T, effectiveSource string) {
	assert.NotEmpty(t, effectiveSource, "Effective source should not be empty")
	// Verify it has one of the expected prefixes
	hasValidPrefix := strings.Contains(effectiveSource, "cidr_ipv4:") ||
		strings.Contains(effectiveSource, "cidr_ipv6:") ||
		strings.Contains(effectiveSource, "prefix_list:") ||
		strings.Contains(effectiveSource, "security_group:")
	assert.True(t, hasValidPrefix, "Effective source should have a valid prefix (cidr_ipv4, cidr_ipv6, prefix_list, or security_group)")
}

func GetAWSEC2Client(t *testing.T) *ec2.Client {
	awsEC2Client := ec2.NewFromConfig(GetAWSConfig(t))
	return awsEC2Client
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
