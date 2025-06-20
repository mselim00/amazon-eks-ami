package ec2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceRetryable(t *testing.T) {
	testCases := []struct {
		input          error
		expectedOutput bool
	}{
		{
			input:          fmt.Errorf("An error occurred (InvalidInstanceID.NotFound) when calling the DescribeInstances operation: Invalid id: \"i-0123456789abcdefg\""),
			expectedOutput: true,
		},
		{
			input:          fmt.Errorf("EC2: DescribeInstances, https response error StatusCode: 0, RequestID: , request send failed, Post \"https://ec2.us-east-1.amazonaws.com/\": dial tcp 52.46.142.78:443: i/o timeout"),
			expectedOutput: true,
		},
		{
			input:          fmt.Errorf("fake error"),
			expectedOutput: false,
		},
	}
	for _, testCase := range testCases {
		retryable, _ := instanceRetryable(testCase.input)
		assert.Equal(t, retryable, testCase.expectedOutput)
	}
}
