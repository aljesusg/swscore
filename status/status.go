// status is a simple package for offering up various status information from Kiali.
package status

const (
	name           = "Kiali"
	ConsoleVersion = name + " console version"
	CoreVersion    = name + " core version"
	CoreCommitHash = name + " core commit hash"
	State          = name + " state"
	StateRunning   = "running"
)

// HTTP status code 200 and user model in data
// swagger:response statusInfo
type swaggStatusInfoResp struct {
	// in:body
	Body struct {
		// HTTP status code 200
		Code int `json:"code"`
		// StatusInfo model
		Data StatusInfo `json:"data"`
	}
}

// StatusInfo statusInfo
//
// This is used for returning a response of Kiali Status
//
// swagger:model StatusInfo
type StatusInfo struct {
	// The state ok Kiali
	// A hash of key,values with versions of Kiali and state
	//
	// required: true
	Status map[string]string `json:"status"`
	// An array of external services installed
	//
	// required: true
	ExternalServices []ExternalServiceInfo `json:"externalServices"`
	// An array of warningMessages
	WarningMessages []string `json:"warningMessages"`
}

var info StatusInfo

// Status response model
//
// This is used for returning a response of Kiali Status
//
// swagger:model externalServiceInfo
type ExternalServiceInfo struct {
	// The name of the service
	//
	// required: true
	Name string `json:"name"`

	// The installed version of the service
	//
	// required: true
	Version string `json:"version"`
}

func init() {
	info = StatusInfo{Status: make(map[string]string)}
	info.Status[State] = StateRunning
}

// Put adds or replaces status info for the provided name. Any previous setting is returned.
func Put(name, value string) (previous string, hasPrevious bool) {
	previous, hasPrevious = info.Status[name]
	info.Status[name] = value
	return previous, hasPrevious
}

// Get returns a copy of the current status info.
func Get() (status StatusInfo) {
	info.ExternalServices = []ExternalServiceInfo{}
	info.WarningMessages = []string{}
	getVersions()
	return info
}
