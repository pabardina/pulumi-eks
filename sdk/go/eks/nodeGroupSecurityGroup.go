// *** WARNING: this file was generated by pulumi-gen-eks. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NodeGroupSecurityGroup is a component that wraps creating a security group for node groups with the default ingress & egress rules required to connect and work with the EKS cluster security group.
type NodeGroupSecurityGroup struct {
	pulumi.ResourceState

	// The security group for node groups with the default ingress & egress rules required to connect and work with the EKS cluster security group.
	SecurityGroup ec2.SecurityGroupOutput `pulumi:"securityGroup"`
	// The EKS cluster ingress rule.
	SecurityGroupRule ec2.SecurityGroupRuleOutput `pulumi:"securityGroupRule"`
}

// NewNodeGroupSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroupSecurityGroup(ctx *pulumi.Context,
	name string, args *NodeGroupSecurityGroupArgs, opts ...pulumi.ResourceOption) (*NodeGroupSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterSecurityGroup == nil {
		return nil, errors.New("invalid value for required argument 'ClusterSecurityGroup'")
	}
	if args.EksCluster == nil {
		return nil, errors.New("invalid value for required argument 'EksCluster'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource NodeGroupSecurityGroup
	err := ctx.RegisterRemoteComponentResource("eks:index:NodeGroupSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type nodeGroupSecurityGroupArgs struct {
	// The security group associated with the EKS cluster.
	ClusterSecurityGroup *ec2.SecurityGroup `pulumi:"clusterSecurityGroup"`
	// The EKS cluster associated with the worker node group
	EksCluster *eks.Cluster `pulumi:"eksCluster"`
	// Key-value mapping of tags to apply to this security group.
	Tags map[string]string `pulumi:"tags"`
	// The VPC in which to create the worker node group.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a NodeGroupSecurityGroup resource.
type NodeGroupSecurityGroupArgs struct {
	// The security group associated with the EKS cluster.
	ClusterSecurityGroup ec2.SecurityGroupInput
	// The EKS cluster associated with the worker node group
	EksCluster eks.ClusterInput
	// Key-value mapping of tags to apply to this security group.
	Tags pulumi.StringMapInput
	// The VPC in which to create the worker node group.
	VpcId pulumi.StringInput
}

func (NodeGroupSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupSecurityGroupArgs)(nil)).Elem()
}

type NodeGroupSecurityGroupInput interface {
	pulumi.Input

	ToNodeGroupSecurityGroupOutput() NodeGroupSecurityGroupOutput
	ToNodeGroupSecurityGroupOutputWithContext(ctx context.Context) NodeGroupSecurityGroupOutput
}

func (*NodeGroupSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupSecurityGroup)(nil))
}

func (i *NodeGroupSecurityGroup) ToNodeGroupSecurityGroupOutput() NodeGroupSecurityGroupOutput {
	return i.ToNodeGroupSecurityGroupOutputWithContext(context.Background())
}

func (i *NodeGroupSecurityGroup) ToNodeGroupSecurityGroupOutputWithContext(ctx context.Context) NodeGroupSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupSecurityGroupOutput)
}

func (i *NodeGroupSecurityGroup) ToNodeGroupSecurityGroupPtrOutput() NodeGroupSecurityGroupPtrOutput {
	return i.ToNodeGroupSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *NodeGroupSecurityGroup) ToNodeGroupSecurityGroupPtrOutputWithContext(ctx context.Context) NodeGroupSecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupSecurityGroupPtrOutput)
}

type NodeGroupSecurityGroupPtrInput interface {
	pulumi.Input

	ToNodeGroupSecurityGroupPtrOutput() NodeGroupSecurityGroupPtrOutput
	ToNodeGroupSecurityGroupPtrOutputWithContext(ctx context.Context) NodeGroupSecurityGroupPtrOutput
}

type nodeGroupSecurityGroupPtrType NodeGroupSecurityGroupArgs

func (*nodeGroupSecurityGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupSecurityGroup)(nil))
}

func (i *nodeGroupSecurityGroupPtrType) ToNodeGroupSecurityGroupPtrOutput() NodeGroupSecurityGroupPtrOutput {
	return i.ToNodeGroupSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *nodeGroupSecurityGroupPtrType) ToNodeGroupSecurityGroupPtrOutputWithContext(ctx context.Context) NodeGroupSecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupSecurityGroupPtrOutput)
}

// NodeGroupSecurityGroupArrayInput is an input type that accepts NodeGroupSecurityGroupArray and NodeGroupSecurityGroupArrayOutput values.
// You can construct a concrete instance of `NodeGroupSecurityGroupArrayInput` via:
//
//          NodeGroupSecurityGroupArray{ NodeGroupSecurityGroupArgs{...} }
type NodeGroupSecurityGroupArrayInput interface {
	pulumi.Input

	ToNodeGroupSecurityGroupArrayOutput() NodeGroupSecurityGroupArrayOutput
	ToNodeGroupSecurityGroupArrayOutputWithContext(context.Context) NodeGroupSecurityGroupArrayOutput
}

type NodeGroupSecurityGroupArray []NodeGroupSecurityGroupInput

func (NodeGroupSecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NodeGroupSecurityGroup)(nil))
}

func (i NodeGroupSecurityGroupArray) ToNodeGroupSecurityGroupArrayOutput() NodeGroupSecurityGroupArrayOutput {
	return i.ToNodeGroupSecurityGroupArrayOutputWithContext(context.Background())
}

func (i NodeGroupSecurityGroupArray) ToNodeGroupSecurityGroupArrayOutputWithContext(ctx context.Context) NodeGroupSecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupSecurityGroupArrayOutput)
}

// NodeGroupSecurityGroupMapInput is an input type that accepts NodeGroupSecurityGroupMap and NodeGroupSecurityGroupMapOutput values.
// You can construct a concrete instance of `NodeGroupSecurityGroupMapInput` via:
//
//          NodeGroupSecurityGroupMap{ "key": NodeGroupSecurityGroupArgs{...} }
type NodeGroupSecurityGroupMapInput interface {
	pulumi.Input

	ToNodeGroupSecurityGroupMapOutput() NodeGroupSecurityGroupMapOutput
	ToNodeGroupSecurityGroupMapOutputWithContext(context.Context) NodeGroupSecurityGroupMapOutput
}

type NodeGroupSecurityGroupMap map[string]NodeGroupSecurityGroupInput

func (NodeGroupSecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NodeGroupSecurityGroup)(nil))
}

func (i NodeGroupSecurityGroupMap) ToNodeGroupSecurityGroupMapOutput() NodeGroupSecurityGroupMapOutput {
	return i.ToNodeGroupSecurityGroupMapOutputWithContext(context.Background())
}

func (i NodeGroupSecurityGroupMap) ToNodeGroupSecurityGroupMapOutputWithContext(ctx context.Context) NodeGroupSecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupSecurityGroupMapOutput)
}

type NodeGroupSecurityGroupOutput struct {
	*pulumi.OutputState
}

func (NodeGroupSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupSecurityGroup)(nil))
}

func (o NodeGroupSecurityGroupOutput) ToNodeGroupSecurityGroupOutput() NodeGroupSecurityGroupOutput {
	return o
}

func (o NodeGroupSecurityGroupOutput) ToNodeGroupSecurityGroupOutputWithContext(ctx context.Context) NodeGroupSecurityGroupOutput {
	return o
}

func (o NodeGroupSecurityGroupOutput) ToNodeGroupSecurityGroupPtrOutput() NodeGroupSecurityGroupPtrOutput {
	return o.ToNodeGroupSecurityGroupPtrOutputWithContext(context.Background())
}

func (o NodeGroupSecurityGroupOutput) ToNodeGroupSecurityGroupPtrOutputWithContext(ctx context.Context) NodeGroupSecurityGroupPtrOutput {
	return o.ApplyT(func(v NodeGroupSecurityGroup) *NodeGroupSecurityGroup {
		return &v
	}).(NodeGroupSecurityGroupPtrOutput)
}

type NodeGroupSecurityGroupPtrOutput struct {
	*pulumi.OutputState
}

func (NodeGroupSecurityGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupSecurityGroup)(nil))
}

func (o NodeGroupSecurityGroupPtrOutput) ToNodeGroupSecurityGroupPtrOutput() NodeGroupSecurityGroupPtrOutput {
	return o
}

func (o NodeGroupSecurityGroupPtrOutput) ToNodeGroupSecurityGroupPtrOutputWithContext(ctx context.Context) NodeGroupSecurityGroupPtrOutput {
	return o
}

type NodeGroupSecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (NodeGroupSecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeGroupSecurityGroup)(nil))
}

func (o NodeGroupSecurityGroupArrayOutput) ToNodeGroupSecurityGroupArrayOutput() NodeGroupSecurityGroupArrayOutput {
	return o
}

func (o NodeGroupSecurityGroupArrayOutput) ToNodeGroupSecurityGroupArrayOutputWithContext(ctx context.Context) NodeGroupSecurityGroupArrayOutput {
	return o
}

func (o NodeGroupSecurityGroupArrayOutput) Index(i pulumi.IntInput) NodeGroupSecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeGroupSecurityGroup {
		return vs[0].([]NodeGroupSecurityGroup)[vs[1].(int)]
	}).(NodeGroupSecurityGroupOutput)
}

type NodeGroupSecurityGroupMapOutput struct{ *pulumi.OutputState }

func (NodeGroupSecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NodeGroupSecurityGroup)(nil))
}

func (o NodeGroupSecurityGroupMapOutput) ToNodeGroupSecurityGroupMapOutput() NodeGroupSecurityGroupMapOutput {
	return o
}

func (o NodeGroupSecurityGroupMapOutput) ToNodeGroupSecurityGroupMapOutputWithContext(ctx context.Context) NodeGroupSecurityGroupMapOutput {
	return o
}

func (o NodeGroupSecurityGroupMapOutput) MapIndex(k pulumi.StringInput) NodeGroupSecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NodeGroupSecurityGroup {
		return vs[0].(map[string]NodeGroupSecurityGroup)[vs[1].(string)]
	}).(NodeGroupSecurityGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeGroupSecurityGroupOutput{})
	pulumi.RegisterOutputType(NodeGroupSecurityGroupPtrOutput{})
	pulumi.RegisterOutputType(NodeGroupSecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(NodeGroupSecurityGroupMapOutput{})
}