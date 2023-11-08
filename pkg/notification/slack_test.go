package notification

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestSlackSplitLargerMessage(t *testing.T) {
	parts := SplitCodeBlocks(":bangbang: drift detected\n\n ```\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\n\n\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\n\n```")
	assert.Equal(t, 2, len(parts))
	assert.Equal(t, 2, strings.Count(parts[0], "```"))
	assert.Equal(t, 2, strings.Count(parts[1], "```"))
}

func TestSlackSmallerMessageNotSplit(t *testing.T) {
	msg := ":bangbang: drift detected\n\n ```\nhere it is\nhere it is```"
	parts := SplitCodeBlocks(msg)
	assert.Equal(t, 1, len(parts))
	// TODO: Fix the func then update test to remove the first newline char
	assert.Equal(t, "\n"+msg, parts[0])
}

func TestSendSlackMessageThatIsLargerThan2Parts(t *testing.T) {
	url := os.Getenv("TEST_SLACK_NOTIFICATION_URL")
	if url == "" {
		t.Skip("Skipping slack message test: $TEST_SLACK_NOTIFICATION_URL not set")
	}
	msg := ":bangbang: drift detected\n\n ```\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\n\n\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\nhere it is\n\n```"
	notification := SlackNotification{Url: url}
	err := notification.Send(msg)
	assert.Equal(t, nil, err)
}
