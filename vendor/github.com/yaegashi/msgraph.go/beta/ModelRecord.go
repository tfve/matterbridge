// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RecordOperation undocumented
type RecordOperation struct {
	// CommsOperation is the base model of RecordOperation
	CommsOperation
	// RecordingLocation undocumented
	RecordingLocation *string `json:"recordingLocation,omitempty"`
	// RecordingAccessToken undocumented
	RecordingAccessToken *string `json:"recordingAccessToken,omitempty"`
	// CompletionReason undocumented
	CompletionReason *RecordCompletionReason `json:"completionReason,omitempty"`
}
