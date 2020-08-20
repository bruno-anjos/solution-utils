package generic_utils

const (
	ServiceEnvVarName  = "SERVICE_ID"
	InstanceEnvVarName = "INSTANCE_ID"
)

const (
	TCP string = "tcp"
	UDP string = "udp"
)

type Node struct {
	Id   string
	Addr string
}
