// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/OS-M/aws-sdk-go/aws"
	"github.com/OS-M/aws-sdk-go/aws/awserr"
	"github.com/OS-M/aws-sdk-go/aws/session"
	"github.com/OS-M/aws-sdk-go/service/accessanalyzer"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// Passing check. Restrictive identity policy.
//

func ExampleAccessAnalyzer_CheckAccessNotGranted_shared00() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.CheckAccessNotGrantedInput{
		Access: []*accessanalyzer.Access{
			{
				Actions: []*string{
					aws.String("s3:PutObject"),
				},
			},
		},
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Id\":\"123\",\"Statement\":[{\"Sid\":\"AllowJohnDoe\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::123456789012:user/JohnDoe\"},\"Action\":\"s3:GetObject\",\"Resource\":\"*\"}]}"),
		PolicyType:     aws.String("RESOURCE_POLICY"),
	}

	result, err := svc.CheckAccessNotGranted(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeInvalidParameterException:
				fmt.Println(accessanalyzer.ErrCodeInvalidParameterException, aerr.Error())
			case accessanalyzer.ErrCodeUnprocessableEntityException:
				fmt.Println(accessanalyzer.ErrCodeUnprocessableEntityException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Passing check. Restrictive S3 Bucket resource policy.
//

func ExampleAccessAnalyzer_CheckAccessNotGranted_shared01() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.CheckAccessNotGrantedInput{
		Access: []*accessanalyzer.Access{
			{
				Resources: []*string{
					aws.String("arn:aws:s3:::sensitive-bucket/*"),
				},
			},
		},
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Id\":\"123\",\"Statement\":[{\"Sid\":\"AllowJohnDoe\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::123456789012:user/JohnDoe\"},\"Action\":\"s3:PutObject\",\"Resource\":\"arn:aws:s3:::non-sensitive-bucket/*\"}]}"),
		PolicyType:     aws.String("RESOURCE_POLICY"),
	}

	result, err := svc.CheckAccessNotGranted(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeInvalidParameterException:
				fmt.Println(accessanalyzer.ErrCodeInvalidParameterException, aerr.Error())
			case accessanalyzer.ErrCodeUnprocessableEntityException:
				fmt.Println(accessanalyzer.ErrCodeUnprocessableEntityException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Failing check. Permissive S3 Bucket resource policy.
//

func ExampleAccessAnalyzer_CheckAccessNotGranted_shared02() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.CheckAccessNotGrantedInput{
		Access: []*accessanalyzer.Access{
			{
				Resources: []*string{
					aws.String("arn:aws:s3:::my-bucket/*"),
				},
			},
		},
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Id\":\"123\",\"Statement\":[{\"Sid\":\"AllowJohnDoe\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::123456789012:user/JohnDoe\"},\"Action\":\"s3:PutObject\",\"Resource\":\"arn:aws:s3:::my-bucket/*\"}]}"),
		PolicyType:     aws.String("RESOURCE_POLICY"),
	}

	result, err := svc.CheckAccessNotGranted(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeInvalidParameterException:
				fmt.Println(accessanalyzer.ErrCodeInvalidParameterException, aerr.Error())
			case accessanalyzer.ErrCodeUnprocessableEntityException:
				fmt.Println(accessanalyzer.ErrCodeUnprocessableEntityException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Passing check. S3 Bucket policy without public access.
//

func ExampleAccessAnalyzer_CheckNoPublicAccess_shared00() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.CheckNoPublicAccessInput{
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Bob\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::111122223333:user/JohnDoe\"},\"Action\":[\"s3:GetObject\"]}]}"),
		ResourceType:   aws.String("AWS::S3::Bucket"),
	}

	result, err := svc.CheckNoPublicAccess(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeInvalidParameterException:
				fmt.Println(accessanalyzer.ErrCodeInvalidParameterException, aerr.Error())
			case accessanalyzer.ErrCodeUnprocessableEntityException:
				fmt.Println(accessanalyzer.ErrCodeUnprocessableEntityException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Failing check. S3 Bucket policy with public access.
//

func ExampleAccessAnalyzer_CheckNoPublicAccess_shared01() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.CheckNoPublicAccessInput{
		PolicyDocument: aws.String("{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"Bob\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":[\"s3:GetObject\"]}]}"),
		ResourceType:   aws.String("AWS::S3::Bucket"),
	}

	result, err := svc.CheckNoPublicAccess(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeInvalidParameterException:
				fmt.Println(accessanalyzer.ErrCodeInvalidParameterException, aerr.Error())
			case accessanalyzer.ErrCodeUnprocessableEntityException:
				fmt.Println(accessanalyzer.ErrCodeUnprocessableEntityException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Successfully started generating finding recommendation
//

func ExampleAccessAnalyzer_GenerateFindingRecommendation_shared00() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GenerateFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("finding-id"),
	}

	result, err := svc.GenerateFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Failed field validation for id value
//

func ExampleAccessAnalyzer_GenerateFindingRecommendation_shared01() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GenerateFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("!"),
	}

	result, err := svc.GenerateFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Successfully fetched finding recommendation
//

func ExampleAccessAnalyzer_GetFindingRecommendation_shared00() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GetFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("finding-id"),
		MaxResults:  aws.Int64(3),
		NextToken:   aws.String("token"),
	}

	result, err := svc.GetFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeResourceNotFoundException:
				fmt.Println(accessanalyzer.ErrCodeResourceNotFoundException, aerr.Error())
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// In progress finding recommendation
//

func ExampleAccessAnalyzer_GetFindingRecommendation_shared01() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GetFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("finding-id"),
		MaxResults:  aws.Int64(3),
	}

	result, err := svc.GetFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeResourceNotFoundException:
				fmt.Println(accessanalyzer.ErrCodeResourceNotFoundException, aerr.Error())
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Failed finding recommendation
//

func ExampleAccessAnalyzer_GetFindingRecommendation_shared02() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GetFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("finding-id"),
		MaxResults:  aws.Int64(3),
	}

	result, err := svc.GetFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeResourceNotFoundException:
				fmt.Println(accessanalyzer.ErrCodeResourceNotFoundException, aerr.Error())
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Failed field validation for id value
//

func ExampleAccessAnalyzer_GetFindingRecommendation_shared03() {
	svc := accessanalyzer.New(session.New())
	input := &accessanalyzer.GetFindingRecommendationInput{
		AnalyzerArn: aws.String("arn:aws:access-analyzer:us-east-1:111122223333:analyzer/a"),
		Id:          aws.String("!"),
	}

	result, err := svc.GetFindingRecommendation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case accessanalyzer.ErrCodeResourceNotFoundException:
				fmt.Println(accessanalyzer.ErrCodeResourceNotFoundException, aerr.Error())
			case accessanalyzer.ErrCodeValidationException:
				fmt.Println(accessanalyzer.ErrCodeValidationException, aerr.Error())
			case accessanalyzer.ErrCodeInternalServerException:
				fmt.Println(accessanalyzer.ErrCodeInternalServerException, aerr.Error())
			case accessanalyzer.ErrCodeThrottlingException:
				fmt.Println(accessanalyzer.ErrCodeThrottlingException, aerr.Error())
			case accessanalyzer.ErrCodeAccessDeniedException:
				fmt.Println(accessanalyzer.ErrCodeAccessDeniedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
