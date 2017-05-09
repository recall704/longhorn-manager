package types

type VolumeState string

const (
	VolumeStateCreated  = VolumeState("created")
	VolumeStateDetached = VolumeState("detached")
	VolumeStateFault    = VolumeState("fault")
	VolumeStateHealthy  = VolumeState("healthy")
	VolumeStateDegraded = VolumeState("degraded")
)

type KVMetadata struct {
	KVIndex uint64 `json:"-"`
}

type VolumeInfo struct {
	// Attributes
	Name                string
	Size                int64 `json:",string"`
	BaseImage           string
	FromBackup          string
	NumberOfReplicas    int
	StaleReplicaTimeout int

	// Running state
	Created      string
	TargetHostID string
	HostID       string
	State        VolumeState
	DesireState  VolumeState
	Endpoint     string

	KVMetadata
}

type InstanceInfo struct {
	ID         string
	Type       InstanceType
	Name       string
	HostID     string
	Address    string
	Running    bool
	VolumeName string

	KVMetadata
}

type ControllerInfo struct {
	InstanceInfo
}

type ReplicaInfo struct {
	InstanceInfo

	Mode         ReplicaMode
	BadTimestamp string
}

type HostInfo struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address"`

	KVMetadata
}

type SettingsInfo struct {
	BackupTarget string `json:"backupTarget"`
	EngineImage  string `json:"engineImage"`

	KVMetadata
}