package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	tags "github.com/ignacio-magno/utils/structTags"
	"strings"
)

// GenerateProjectionExpressionExclude  filters is the name of field in struct, and exclude the fields in the filter
func GenerateProjectionExpressionExclude(structure interface{}, filters []string) *string {
	fields := tags.NewTags(structure, "dynamodbav").GetMapFieldsAndTagCleans()

	containValue := func(value string) bool {
		for _, v := range filters {
			if v == value {
				return true
			}
		}
		return false
	}

	var projection []string
	for i := range fields {
		if !containValue(i) {
			projection = append(projection, fields[i])
		}
	}

	return aws.String(strings.Join(projection, ","))
}
