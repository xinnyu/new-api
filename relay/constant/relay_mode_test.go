package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTextRelayMode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mode     int
		expected bool
	}{
		{name: "chat completions", mode: RelayModeChatCompletions, expected: true},
		{name: "legacy completions", mode: RelayModeCompletions, expected: true},
		{name: "responses", mode: RelayModeResponses, expected: true},
		{name: "responses compact", mode: RelayModeResponsesCompact, expected: true},
		{name: "unknown", mode: RelayModeUnknown, expected: false},
		{name: "embeddings excluded", mode: RelayModeEmbeddings, expected: false},
		{name: "moderations excluded", mode: RelayModeModerations, expected: false},
		{name: "rerank excluded", mode: RelayModeRerank, expected: false},
		{name: "images generations excluded", mode: RelayModeImagesGenerations, expected: false},
		{name: "images edits excluded", mode: RelayModeImagesEdits, expected: false},
		{name: "audio speech excluded", mode: RelayModeAudioSpeech, expected: false},
		{name: "audio transcription excluded", mode: RelayModeAudioTranscription, expected: false},
		{name: "suno submit excluded", mode: RelayModeSunoSubmit, expected: false},
		{name: "video submit excluded", mode: RelayModeVideoSubmit, expected: false},
		{name: "midjourney imagine excluded", mode: RelayModeMidjourneyImagine, expected: false},
		{name: "realtime excluded", mode: RelayModeRealtime, expected: false},
		{name: "gemini excluded (coarse bucket, includes embedContent)", mode: RelayModeGemini, expected: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, IsTextRelayMode(tc.mode))
		})
	}
}
