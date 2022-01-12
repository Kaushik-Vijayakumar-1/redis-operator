package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailover represents a Redis failover
// +kubebuilder:printcolumn:name="NAME",type="string",JSONPath=".metadata.name"
// +kubebuilder:printcolumn:name="REDIS",type="integer",JSONPath=".spec.redis.replicas"
// +kubebuilder:printcolumn:name="SENTINELS",type="integer",JSONPath=".spec.sentinel.replicas"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:singular=redisfailover,path=redisfailovers,shortName=rf,scope=Namespaced
type RedisFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFailoverSpec `json:"spec"`
}

// RedisFailoverSpec represents a Redis failover spec
type RedisFailoverSpec struct {
	Redis          RedisSettings      `json:"redis,omitempty"`
	Sentinel       SentinelSettings   `json:"sentinel,omitempty"`
	Auth           AuthSettings       `json:"auth,omitempty"`
	LabelWhitelist []string           `json:"labelWhitelist,omitempty"`
	BootstrapNode  *BootstrapSettings `json:"bootstrapNode,omitempty"`
}

// RedisCommandRename defines the specification of a "rename-command" configuration option
type RedisCommandRename struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// RedisSettings defines the specification of the redis cluster
type RedisSettings struct {
	Image                         string                        `json:"image,omitempty"`
	ImagePullPolicy               corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	Replicas                      int32                         `json:"replicas,omitempty"`
	Resources                     corev1.ResourceRequirements   `json:"resources,omitempty"`
	CustomConfig                  []string                      `json:"customConfig,omitempty"`
	CustomCommandRenames          []RedisCommandRename          `json:"customCommandRenames,omitempty"`
	Command                       []string                      `json:"command,omitempty"`
	ShutdownConfigMap             string                        `json:"shutdownConfigMap,omitempty"`
	Storage                       RedisStorage                  `json:"storage,omitempty"`
	Exporter                      RedisExporter                 `json:"exporter,omitempty"`
	Affinity                      *corev1.Affinity              `json:"affinity,omitempty"`
	SecurityContext               *corev1.PodSecurityContext    `json:"securityContext,omitempty"`
	ImagePullSecrets              []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Tolerations                   []corev1.Toleration           `json:"tolerations,omitempty"`
	NodeSelector                  map[string]string             `json:"nodeSelector,omitempty"`
	PodAnnotations                map[string]string             `json:"podAnnotations,omitempty"`
	ServiceAnnotations            map[string]string             `json:"serviceAnnotations,omitempty"`
	HostNetwork                   bool                          `json:"hostNetwork,omitempty"`
	DNSPolicy                     corev1.DNSPolicy              `json:"dnsPolicy,omitempty"`
	PriorityClassName             string                        `json:"priorityClassName,omitempty"`
	ServiceAccountName            string                        `json:"serviceAccountName,omitempty"`
	TerminationGracePeriodSeconds int64                         `json:"terminationGracePeriod,omitempty"`
}

// SentinelSettings defines the specification of the sentinel cluster
type SentinelSettings struct {
	Image              string                        `json:"image,omitempty"`
	ImagePullPolicy    corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	Replicas           int32                         `json:"replicas,omitempty"`
	Resources          corev1.ResourceRequirements   `json:"resources,omitempty"`
	CustomConfig       []string                      `json:"customConfig,omitempty"`
	Command            []string                      `json:"command,omitempty"`
	Affinity           *corev1.Affinity              `json:"affinity,omitempty"`
	SecurityContext    *corev1.PodSecurityContext    `json:"securityContext,omitempty"`
	ImagePullSecrets   []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Tolerations        []corev1.Toleration           `json:"tolerations,omitempty"`
	NodeSelector       map[string]string             `json:"nodeSelector,omitempty"`
	PodAnnotations     map[string]string             `json:"podAnnotations,omitempty"`
	ServiceAnnotations map[string]string             `json:"serviceAnnotations,omitempty"`
	Exporter           SentinelExporter              `json:"exporter,omitempty"`
	HostNetwork        bool                          `json:"hostNetwork,omitempty"`
	DNSPolicy          corev1.DNSPolicy              `json:"dnsPolicy,omitempty"`
	PriorityClassName  string                        `json:"priorityClassName,omitempty"`
	ServiceAccountName string                        `json:"serviceAccountName,omitempty"`
}

// AuthSettings contains settings about auth
type AuthSettings struct {
	SecretPath string `json:"secretPath,omitempty"`
}

// BootstrapSettings contains settings about a potential bootstrap node
type BootstrapSettings struct {
	Host           string `json:"host,omitempty"`
	Port           string `json:"port,omitempty"`
	AllowSentinels bool   `json:"allowSentinels,omitempty"`
}

// RedisExporter defines the specification for the redis exporter
type RedisExporter struct {
	Enabled         bool              `json:"enabled,omitempty"`
	Image           string            `json:"image,omitempty"`
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Env             []corev1.EnvVar   `json:"env,omitempty"`
}

// SentinelExporter defines the specification for the sentinel exporter
type SentinelExporter struct {
	Enabled         bool              `json:"enabled,omitempty"`
	Image           string            `json:"image,omitempty"`
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Env             []corev1.EnvVar   `json:"env,omitempty"`
}

// RedisStorage defines the structure used to store the Redis Data
type RedisStorage struct {
	KeepAfterDeletion     bool                          `json:"keepAfterDeletion,omitempty"`
	EmptyDir              *corev1.EmptyDirVolumeSource  `json:"emptyDir,omitempty"`
	PersistentVolumeClaim *corev1.PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailoverList represents a Redis failover list
type RedisFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RedisFailover `json:"items"`
}
