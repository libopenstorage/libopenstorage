package osdconfig

import "time"

// NodesConfig contains all of node level data in the form of a map with node ID's as keys
type NodesConfig struct {
	NodeConf map[string]*NodeConfig `json:"node_conf,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
}

// NodeConfig is a node level config data
type NodeConfig struct {
	NodeId    string         `json:"node_id,omitempty" enable:"true" hidden:"false" usage:"ID for the node"`
	NodeIndex int            `json:"node_index,omitempty" enable:"true" hidden:"false" usage:"Index for the node"`
	Network   *NetworkConfig `json:"network,omitempty" enable:"true" hidden:"false" usage:"Network configuration" description:"Configure network values for a node"`
	Storage   *StorageConfig `json:"storage,omitempty" enable:"true" hidden:"false" usage:"Storage configuration" description:"Configure storage values for a node"`
	Private   interface{}    `json:"generic,omitempty" enable:"false" hidden:"false" usage:"Private node data"`
}

func (conf *NodeConfig) Init() *NodeConfig {
	conf.Network = new(NetworkConfig).Init()
	conf.Storage = new(StorageConfig).Init()
	return conf
}

// KvdbConfig stores parameters defining kvdb configuration
type KvdbConfig struct {
	Name           string   `json:"name,omitempty" enable:"true" hidden:"false" usage:"Name for kvdb"`
	Username       string   `json:"username,omitempty" enable:"true" hidden:"false" usage:"Username for kvdb"`
	Password       string   `json:"password,omitempty" enable:"true" hidden:"false" usage:"Passwd for kvdb"`
	CaFile         string   `json:"ca_file,omitempty" enable:"true" hidden:"false" usage:"CA file for kvdb"`
	CertFile       string   `json:"cert_file,omitempty" enable:"true" hidden:"false" usage:"Cert file for kvdb"`
	CertKeyFile    string   `json:"cert_key_file,omitempty" enable:"true" hidden:"false" usage:"Cert key file for kvdb"`
	TrustedCaFile  string   `json:"trusted_ca_file,omitempty" enable:"true" hidden:"false" usage:"Trusted CA file for kvdb"`
	ClientCertAuth string   `json:"client_cert_auth,omitempty" enable:"true" hidden:"false" usage:"Client cert auth"`
	AclToken       string   `json:"acl_token,omitempty" enable:"true" hidden:"false" usage:"ACL token"`
	Discovery      []string `json:"discovery,omitempty" enable:"true" hidden:"false" usage:"kvdb addresses"`
}

func (conf *KvdbConfig) Init() *KvdbConfig {
	conf.Discovery = make([]string, 0, 0)
	return conf
}

// ClusterConfig is a cluster level config parameter struct
// swagger:model
type ClusterConfig struct {
	Description    string         `json:"description,omitempty" enable:"true" hidden:"false" usage:"Cluster description"`
	Mode           string         `json:"mode,omitempty" enable:"true" hidden:"false" usage:"Mode for cluster"`
	Version        string         `json:"version,omitempty" enable:"true" hidden:"false" usage:"Version info for cluster"`
	Created        time.Time      `json:"created,omitempty" enable:"true" hidden:"false" usage:"Creation info for cluster"`
	ClusterId      string         `json:"cluster_id,omitempty" enable:"true" hidden:"false" usage:"Cluster ID info"`
	ClusterUUId    string         `json:"cluster_uuid,omitempty" enable:"true" hidden:"false" usage:"Cluster UUID info"`
	NodeId         []string       `json:"node_id,omitempty" enable:"true" hidden:"false" usage:"Node ID info"`
	LoggingUrl     string         `json:"logging_url,omitempty" enable:"true" hidden:"false" usage:"URL for logging"`
	AlertingUrl    string         `json:"alerting_url,omitempty" enable:"true" hidden:"false" usage:"URL for altering"`
	Scheduler      string         `json:"scheduler,omitempty" enable:"true" hidden:"false" usage:"Cluster scheduler"`
	Multicontainer bool           `json:"multicontainer,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Nolh           bool           `json:"nolh,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Callhome       bool           `json:"callhome,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Bootstrap      bool           `json:"bootstrap,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	TunnelEndPoint string         `json:"tunnel_end_point,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	TunnelCerts    []string       `json:"tunnel_certs,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Driver         string         `json:"driver,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	DebugLevel     string         `json:"debug_level,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Domain         string         `json:"domain,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	LicServerUrl   string         `json:"lic_server_url,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	BootstrapType  string         `json:"bootstrap_type,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	APIRootCA      string         `json:"api_root_ca,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	APICert        string         `json:"api_cert,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	APIKey         string         `json:"api_key,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	CSIEndpoint    string         `json:"csi_endpoint,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
	Secrets        *SecretsConfig `json:"secrets,omitempty" enable:"true" hidden:"false" usage:"usage to be added" description:"description to be added"`
	Kvdb           *KvdbConfig    `json:"kvdb,omitempty" enable:"false" hidden:"false" usage:"usage to be added" description:"description to be added"`
	Private        interface{}    `json:"generic,omitempty" enable:"true" hidden:"false" usage:"usage to be added"`
}

func (conf *ClusterConfig) Init() *ClusterConfig {
	conf.TunnelCerts = make([]string, 0, 0)
	conf.NodeId = make([]string, 0, 0)
	conf.Secrets = new(SecretsConfig).Init()
	conf.Kvdb = new(KvdbConfig).Init()
	return conf
}

// NetworkConfig is a network configuration parameters struct
type NetworkConfig struct {
	MgtIface    string `json:"mgt_iface,omitempty" enable:"true" hidden:"false" usage:"Management interface"`
	DataIface   string `json:"data_iface,omitempty" enable:"true" hidden:"false" usage:"Data interface"`
	MgtIfaceIP  string `json:"mgt_iface_ip,omitempty" enable:"true" hidden:"false" usage:"Management interface IP"`
	DataIfaceIP string `json:"data_iface_ip,omitempty" enable:"true" hidden:"false" usage:"Data interface IP"`
}

func (conf *NetworkConfig) Init() *NetworkConfig {
	return conf
}

// SecretsConfig is a secrets configuration parameters struct
type SecretsConfig struct {
	SecretType       string       `json:"secret_type,omitempty" enable:"true" hidden:"false" usage:"Secret type"`
	ClusterSecretKey string       `json:"cluster_secret_key,omitempty" enable:"true" hidden:"false" usage:"Secret key"`
	Vault            *VaultConfig `json:"vault,omitempty" enable:"true" hidden:"false" usage:"Vault configuration"`
	Aws              *AWSConfig   `json:"aws,omitempty" enable:"true" hidden:"false" usage:"AWS configuration"`
}

func (conf *SecretsConfig) Init() *SecretsConfig {
	conf.Vault = new(VaultConfig).Init()
	conf.Aws = new(AWSConfig).Init()
	return conf
}

// VaultConfig is a vault configuration parameters struct
type VaultConfig struct {
	Token         string `json:"token,omitempty" enable:"true" hidden:"false" usage:"Vault token"`
	Addr          string `json:"addr,omitempty" enable:"true" hidden:"false" usage:"Vault address"`
	Cacert        string `json:"cacert,omitempty" enable:"true" hidden:"false" usage:"Vault CA certificate"`
	Capath        string `json:"capath,omitempty" enable:"true" hidden:"false" usage:"Vault CA path"`
	ClientCert    string `json:"client_cert,omitempty" enable:"true" hidden:"false" usage:"Vault client certificate"`
	ClientKey     string `json:"client_key,omitempty" enable:"true" hidden:"false" usage:"Vault client key"`
	SkipVerify    string `json:"skip_verify,omitempty" enable:"true" hidden:"false" usage:"Vault skip verification"`
	TlsServerName string `json:"tls_server_name,omitempty" enable:"true" hidden:"false" usage:"Vault TLS server name"`
	BasePath      string `json:"base_path,omitempty" enable:"true" hidden:"false" usage:"Vault base path"`
}

func (conf *VaultConfig) Init() *VaultConfig {
	return conf
}

// AWS configuration parameters struct
type AWSConfig struct {
	AwsAccessKeyId     string `json:"aws_access_key_id,omitempty" enable:"true" hidden:"false" usage:"AWS access key ID"`
	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty" enable:"true" hidden:"false" usage:"AWS secret access key"`
	AwsSecretTokenKey  string `json:"aws_secret_token_key,omitempty" enable:"true" hidden:"false" usage:"AWS secret token key"`
	AwsCmk             string `json:"aws_cmk,omitempty" enable:"true" hidden:"false" usage:"AWS CMK"`
	AwsRegion          string `json:"aws_region,omitempty" enable:"true" hidden:"false" usage:"AWS region"`
}

func (conf *AWSConfig) Init() *AWSConfig {
	return conf
}

// StorageConfig is a storage configuration parameters struct
type StorageConfig struct {
	DevicesMd        []string    `json:"devices_md,omitempty" enable:"true" hidden:"false" usage:"Devices MD"`
	MaxCount         uint32      `json:"max_count,omitempty" enable:"true" hidden:"false" usage:"Maximum count"`
	MaxDriveSetCount uint32      `json:"max_drive_set_count,omitempty" enable:"true" hidden:"false" usage:"Max drive set count"`
	Devices          []string    `json:"devices,omitempty" enable:"true" hidden:"false" usage:"Devices list"`
	JournalDev       string      `json:"journal_dev,omitempty" enable:"true" hidden:"false" usage:"Journal device"`
	Driver           string      `json:"driver,omitempty" enable:"true" hidden:"false" usage:"Driver"`
	DebugLevel       string      `json:"debug_level,omitempty" enable:"true" hidden:"false" usage:"Debug level"`
	RaidLevel        string      `json:"raid_level,omitempty" enable:"true" hidden:"false" usage:"RAID level info"`
	RaidLevelMd      string      `json:"raid_level_md,omitempty" enable:"true" hidden:"false" usage:"RAID level MD"`
	AsyncIo          bool        `json:"async_io,omitempty" enable:"true" hidden:"false" usage:"Async I/O"`
	NumThreads       uint32      `json:"num_threads,omitempty" enable:"true" hidden:"false" usage:"Number of threads"`
	MaxNumThreads    uint32      `json:"max_num_threads,omitempty" enable:"true" hidden:"false" usage:"Max number of threads"`
	Private          interface{} `json:"private,omitempty" enable:"true" hidden:"false" usage:"Runtime options"`
}

func (conf *StorageConfig) Init() *StorageConfig {
	conf.Devices = make([]string, 0, 0)
	conf.DevicesMd = make([]string, 0, 0)
	return conf
}
