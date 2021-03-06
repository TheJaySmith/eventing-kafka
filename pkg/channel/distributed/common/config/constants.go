/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

// Package Constants
const (
	// The name of the configmap used to hold eventing-kafka settings
	SettingsConfigMapName = "config-eventing-kafka"
	// The name of the keys in the Data section of the eventing-kafka configmap that holds Sarama and Eventing-Kafka configuration YAML
	SaramaSettingsConfigKey        = "sarama"
	EventingKafkaSettingsConfigKey = "eventing-kafka"
)
