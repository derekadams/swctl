/**
 * Copyright © 2014-2020 The SiteWhere Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package alpha3 defines SiteWhere Structures
package alpha3

// SiteWhereInstanceInfrastructureGRPCConfiguration SiteWhere Instance Infrastructure gRPC configurations
type SiteWhereInstanceInfrastructureGRPCConfiguration struct {
	BackoffMultiplier     float64 `json:"backoffMultiplier"`
	InitialBackoffSeconds int64   `json:"initialBackoffSeconds"`
	MaxBackoffSeconds     int64   `json:"maxBackoffSeconds"`
	MaxRetryCount         int64   `json:"maxRetryCount"`
	ResolveFQDN           bool    `json:"resolveFQDN"`
}

// SiteWhereInstanceInfrastructureKafkaConfiguration SiteWhere Instance Infrastrucre Kafka configurations
type SiteWhereInstanceInfrastructureKafkaConfiguration struct {
	Hostname                      string `json:"hostname"`
	Port                          int64  `json:"port"`
	DefaultTopicPartitions        int64  `json:"defaultTopicPartitions"`
	DefaultTopicReplicationFactor int64  `json:"defaultTopicReplicationFactor"`
}

// SiteWhereInstanceInfrastructureMetricsConfiguration SiteWhere Instance Infrastrucre Metrics configurations
type SiteWhereInstanceInfrastructureMetricsConfiguration struct {
	Enabled  bool  `json:"enabled"`
	HTTPPort int64 `json:"httpPort"`
}

// SiteWhereInstanceInfrastructureRedisConfiguration SiteWhere Instance Infrastrucre Redis configurations
type SiteWhereInstanceInfrastructureRedisConfiguration struct {
	Hostname        string `json:"hostname"`
	Port            int64  `json:"port"`
	NodeCount       int64  `json:"nodeCount"`
	MasterGroupName string `json:"masterGroupName"`
}

// SiteWhereInstanceInfrastructureConfiguration SiteWhere Instance Infrastructure configurations
type SiteWhereInstanceInfrastructureConfiguration struct {
	Namespace string                                               `json:"namespace"`
	GRPC      *SiteWhereInstanceInfrastructureGRPCConfiguration    `json:"grpc"`
	Kafka     *SiteWhereInstanceInfrastructureKafkaConfiguration   `json:"kafka"`
	Metrics   *SiteWhereInstanceInfrastructureMetricsConfiguration `json:"metrics"`
	Redis     *SiteWhereInstanceInfrastructureRedisConfiguration   `json:"redis"`
}

// SiteWhereInstancePersistenceCassandraConfiguration SiteWhere Instance Persistence Cassandra configurations
type SiteWhereInstancePersistenceCassandraConfiguration struct {
	ContactPoints string `json:"contactPoints"`
	Keyspace      string `json:"keyspace"`
}

// SiteWhereInstancePersistenceInfluxDBConfiguration SiteWhere Instance Persistence InfuxDB configurations
type SiteWhereInstancePersistenceInfluxDBConfiguration struct {
	Hostname     string `json:"hostname"`
	Port         int64  `json:"port"`
	DatabaseName string `json:"databaseName"`
}

// SiteWhereInstancePersistenceRDBConfiguration SiteWhere Instance Persistence Relational Database configurations
type SiteWhereInstancePersistenceRDBConfiguration struct {
}

// SiteWhereInstancePersistenceConfiguration SiteWhere Instance Persistence configurations
type SiteWhereInstancePersistenceConfiguration struct {
	CassandraConfigurations map[string]SiteWhereInstancePersistenceCassandraConfiguration `json:"cassandraConfigurations"`
	InfluxDBConfigurations  map[string]SiteWhereInstancePersistenceInfluxDBConfiguration  `json:"influxDbConfigurations"`
	RDBConfigurations       map[string]SiteWhereInstancePersistenceRDBConfiguration       `json:"rdbConfigurations"`
}

// SiteWhereInstanceConfiguration SiteWhere Instance configurations
type SiteWhereInstanceConfiguration struct {
	Infrastructure *SiteWhereInstanceInfrastructureConfiguration `json:"infrastructure"`
	Persistence    *SiteWhereInstancePersistenceConfiguration    `json:"persistenceConfigurations"`
}

// SiteWhereInstanceStatus SiteWhere Instance Tenant Management and User Management Status
type SiteWhereInstanceStatus struct {
	TenantManagementStatus string `json:"tenantManagementStatus"`
	UserManagementStatus   string `json:"userManagementStatus"`
}

// SiteWhereMicroserviceStatus SiteWhere Instance Microservice Status
type SiteWhereMicroserviceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// SiteWhereProfile profiles to use the application
type SiteWhereProfile string

const (
	// Default profile, use all microservices
	Default SiteWhereProfile = "Default"
	// Minimal profile, use a reduce set of microservices
	Minimal SiteWhereProfile = "Minimal"
)

// SiteWhereInstance represents an Instacen in SiteWhere
type SiteWhereInstance struct {
	Name                  string                          `json:"name"`
	Namespace             string                          `json:"namespace"`
	Tag                   string                          `json:"tag"`
	Debug                 bool                            `json:"debug"`
	ConfigurationTemplate string                          `json:"configurationTemaplate"`
	DatasetTemplate       string                          `json:"datasetTemplate"`
	Configuration         *SiteWhereInstanceConfiguration `json:"configuration"`
	Status                *SiteWhereInstanceStatus        `json:"status"`
	Microservices         []SiteWhereMicroserviceStatus   `json:"microservices"`
	Profile               SiteWhereProfile                `json:"profile"`
}

// SiteWhereMicroservice defines a microservice
type SiteWhereMicroservice struct {
	ID          string
	Name        string
	Description string
	Icon        string
	PortOffset  int32
	Logger      string
	Profile     SiteWhereProfile
}

var microservices = []SiteWhereMicroservice{
	{
		ID:          "asset-management",
		Name:        "Asset Management",
		Description: "Provides APIs for managing assets associated with device assignments",
		Icon:        "devices_other",
		PortOffset:  6,
		Logger:      "com.sitewhere.asset",
		Profile:     Default},
	{
		ID:          "batch-operations",
		Name:        "Batch Operations",
		Description: "Handles processing of operations which affect a large number of devices",
		Icon:        "view_module",
		Logger:      "com.sitewhere.batch",
		PortOffset:  11,
		Profile:     Minimal},
	{
		ID:          "command-delivery",
		Name:        "Command Delivery",
		Description: "Manages delivery of commands in various formats based on invocation events",
		Icon:        "call_made",
		Logger:      "com.sitewhere.commands",
		PortOffset:  12,
		Profile:     Minimal},
	{
		ID:          "device-management",
		Name:        "Device Management",
		Description: "Provides APIs for managing the device object model",
		Icon:        "developer_board",
		Logger:      "com.sitewhere.device",
		PortOffset:  4,
		Profile:     Default},
	{
		ID:          "device-registration",
		Name:        "Device Registration",
		Description: "Handles registration of new devices with the system",
		Icon:        "add_box",
		Logger:      "com.sitewhere.registration",
		PortOffset:  13,
		Profile:     Minimal},
	{
		ID:          "device-state",
		Name:        "Device State",
		Description: "Provides device state management features such as device shadows",
		Icon:        "warning",
		Logger:      "com.sitewhere.devicestate",
		PortOffset:  14,
		Profile:     Minimal},
	{
		ID:          "event-management",
		Name:        "Event Management",
		Description: "Provides APIs for persisting and accessing events generated by devices",
		Icon:        "dynamic_feed",
		Logger:      "com.sitewhere.event",
		PortOffset:  5,
		Profile:     Default},
	{
		ID:          "event-sources",
		Name:        "Event Sources",
		Description: "Handles inbound device data from various sources, protocols, and formats",
		Icon:        "forward",
		Logger:      "com.sitewhere.sources",
		PortOffset:  8,
		Profile:     Default},
	{
		ID:          "inbound-processing",
		Name:        "Inbound Processing",
		Description: "Common processing logic applied to enrich and direct inbound events",
		Icon:        "input",
		PortOffset:  7,
		Logger:      "com.sitewhere.inbound",
		Profile:     Default},
	{
		ID:          "instance-management",
		Name:        "Instance Management",
		Description: "Handles APIs for managing global aspects of an instance",
		Icon:        "language",
		PortOffset:  1,
		Logger:      "com.sitewhere.instance",
		Profile:     Default},
	{
		ID:          "label-generation",
		Name:        "Label Generation",
		Description: "Supports generating labels such as bar codes and QR codes for devices",
		Icon:        "label",
		PortOffset:  9,
		Logger:      "com.sitewhere.labels",
		Profile:     Minimal},
	{
		ID:          "outbound-connectors",
		Name:        "Outbound Connectors",
		Description: "Allows event streams to be delivered to external systems for additional processing",
		Icon:        "label",
		PortOffset:  16,
		Logger:      "com.sitewhere.connectors",
		Profile:     Default},
	{
		ID:          "schedule-management",
		Name:        "Schedule Management",
		Description: "Supports scheduling of various system operations",
		Icon:        "label",
		PortOffset:  18,
		Logger:      "com.sitewhere.schedule",
		Profile:     Minimal},
}

// GetSiteWhereMicroservicesList Returns the list of SiteWhere Microservices Names
func GetSiteWhereMicroservicesList() []SiteWhereMicroservice {
	return microservices
}
