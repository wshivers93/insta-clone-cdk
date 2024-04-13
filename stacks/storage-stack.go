package storageStack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func StorageStack(scope constructs.Construct, id string, props awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props)

	awss3.NewBucket(stack, jsii.String("ImageBucket"), &awss3.BucketProps{
		Encryption: awss3.BucketEncryption_KMS_MANAGED,
	})

	awsdynamodb.NewTable(stack, jsii.String("UserTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		Encryption: awsdynamodb.TableEncryption_AWS_MANAGED,
		TableName: jsii.String("Users"),
	})

	awsdynamodb.NewTable(stack, jsii.String("PostsTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("userId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("postId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("Posts"),
	})

	return stack
}
