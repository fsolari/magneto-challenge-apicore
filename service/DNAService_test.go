package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestIsMutantMustReturnTrueIfDNAMutant(t *testing.T) {
	var matrix = [][]rune{{65, 84, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 65, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	ok := isMutant(matrix, len(matrix))

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", ok))

}

func TestIsMutantMustReturnFalseIfDNAHuman(t *testing.T) {

	var matrix = [][]rune{{84, 65, 71, 67, 71, 65}, {67, 65, 84, 84, 71, 67}, {84, 84, 65, 84, 65, 84}, {65, 71, 65, 67, 71, 84}, {71, 67, 67, 65, 84, 65}, {84, 67, 65, 67, 84, 71}}

	ok := isMutant(matrix, len(matrix))

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestLoopHorizontallyMustSendToChannelSequenceBeforeDone(t *testing.T) {

	var matrix = [][]rune{{65, 65, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 67, 71, 71}, {67, 67, 67, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopHorizontally(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, true, ok, fmt.Sprintf("ok must be true %v", ok))

}

func TestLoopHorizontallyIfNoSequenceMustSendToDoneChannel(t *testing.T) {

	var matrix = [][]rune{{84, 84, 71, 67, 71, 65}, {67, 65, 71, 84, 71, 67}, {84, 84, 65, 84, 71, 84}, {65, 71, 65, 65, 71, 71}, {67, 67, 71, 67, 84, 65}, {84, 67, 65, 67, 84, 71}}

	cols := len(matrix)

	sequence := make(chan bool)
	done := make(chan bool)

	go loopHorizontally(matrix, cols, sequence, done)

	ok := readChannels(sequence, done)

	assert.Equal(t, false, ok, fmt.Sprintf("ok must be false %v", ok))
}

func TestLoopVerticallyMustSendToBothChannelsOnSequence(t *testing.T) {

}

func TestLoopVerticallyMustSendToDoneChannelAlways(t *testing.T) {

}

func TestLoopDiagonallyMustSendToBothChannelsOnSequence(t *testing.T) {

}

func TestLoopDiagonallyMustSendToDoneChannelAlways(t *testing.T) {

}

func TestCheckConditionIfMeetConditionReturnTrue(t *testing.T) {}

func TestCheckConditionIfFailsConditionReturnFalse(t *testing.T) {}

func TestOnSequenceSendTrueMustSendToChannelIfIndexEqual(t *testing.T) {}

func TestChannelReaderMustReturnTrueIfChannelAIsDoneFirst(t *testing.T) {}

func TestChannelReaderMustReturnFalseIfChannelBIsDone(t *testing.T) {}

func readChannels(seq chan bool, done chan bool) bool {
	select {
	case <-seq:
		return true
	case <-done:
		return false
	}
}