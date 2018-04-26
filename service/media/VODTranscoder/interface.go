package VODTranscoder

import (
	"net/http"
	m 	"github.com/rightly/ncloud-sdk-go/ncloud"
)

//to use Testing
type VODTranscoderAPI 	interface {
	// Job Operation
	createJob(credentials m.Credentials) *http.Request
	cancelJob(credentials m.Credentials, jobId string) *http.Request
	lookupJobList(credentials m.Credentials) *http.Request
	lookupJobInfo(credentials m.Credentials, jobId string) *http.Request

	// Preset Operation
	lookupPresetList(credentials m.Credentials) *http.Request
	lookupPresetInfo(credentials m.Credentials, presetId string) *http.Request
}