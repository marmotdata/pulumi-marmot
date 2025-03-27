package main

import (
	"github.com/marmotdata/pulumi-marmot/sdk/go/marmot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		marmotHost := cfg.Require("marmot:host")
		marmotApiKey := cfg.Require("marmot:apiKey")
		kafkaAsset, err := marmot.NewAsset(ctx, "kafkaAsset", &marmot.AssetArgs{
			Name:        pulumi.String("customer-events-stream"),
			Type:        pulumi.String("Topic"),
			Description: pulumi.String("Kafka stream for customer events"),
			Services: pulumi.StringArray{
				pulumi.String("Kafka"),
			},
			Tags: pulumi.StringArray{
				pulumi.String("events"),
				pulumi.String("streaming"),
				pulumi.String("real-time"),
				pulumi.String("customer-data"),
			},
			Metadata: pulumi.Map{
				"owner":              pulumi.Any("platform-team"),
				"partitions":         pulumi.Any("24"),
				"replication_factor": pulumi.Any("3"),
				"retention_ms":       pulumi.Any("604800000"),
				"security_protocol":  pulumi.Any("SASL_SSL"),
				"compression_type":   pulumi.Any("lz4"),
				"max_message_size":   pulumi.Any("1048576"),
				"group_id":           pulumi.Any("customer-events-consumers"),
			},
			Schema: pulumi.Map{
				"type":      pulumi.Any("avro"),
				"doc":       pulumi.Any("Schema for customer events"),
				"name":      pulumi.Any("CustomerEvent"),
				"namespace": pulumi.Any("com.example.events"),
				"fields": pulumi.Any{
					map[string]interface{}{
						"name": "event_id",
						"type": "string",
						"doc":  "Unique identifier for the event",
					},
					map[string]interface{}{
						"name": "customer_id",
						"type": "string",
						"doc":  "Customer identifier",
					},
					map[string]interface{}{
						"name": "event_type",
						"type": map[string]interface{}{
							"type": "enum",
							"name": "EventType",
							"symbols": []string{
								"SIGN_UP",
								"LOGIN",
								"PURCHASE",
								"ACCOUNT_UPDATE",
								"LOGOUT",
							},
						},
					},
					map[string]interface{}{
						"name": "timestamp",
						"type": "long",
						"doc":  "Event timestamp in milliseconds since epoch",
					},
					map[string]interface{}{
						"name": "payload",
						"type": map[string]interface{}{
							"type": "record",
							"name": "EventPayload",
							"fields": []interface{}{
								map[string]interface{}{
									"name": "session_id",
									"type": []string{
										"null",
										"string",
									},
								},
								map[string]interface{}{
									"name": "device_info",
									"type": map[string]interface{}{
										"type": "record",
										"name": "DeviceInfo",
										"fields": []map[string]interface{}{
											map[string]interface{}{
												"name": "type",
												"type": []string{
													"null",
													"string",
												},
											},
											map[string]interface{}{
												"name": "os",
												"type": []string{
													"null",
													"string",
												},
											},
											map[string]interface{}{
												"name": "browser",
												"type": []string{
													"null",
													"string",
												},
											},
										},
									},
								},
								map[string]interface{}{
									"name": "properties",
									"type": map[string]interface{}{
										"type":   "map",
										"values": "string",
									},
								},
							},
						},
					},
				},
			},
			ExternalLinks: marmot.ExternalLinkArray{
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Kafka UI"),
					Url:  pulumi.String("http://kafka-ui.example.com"),
				},
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Monitoring"),
					Url:  pulumi.String("http://grafana.example.com/kafka-dashboard"),
				},
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Schema Registry"),
					Url:  pulumi.String("http://schema-registry.example.com/subjects/customer-events/versions/latest"),
					Icon: pulumi.String("database-schema"),
				},
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Documentation"),
					Url:  pulumi.String("http://docs.example.com/kafka/customer-events"),
					Icon: pulumi.String("book"),
				},
			},
			Environments: marmot.AssetEnvironmentMap{
				"dev": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Development"),
					Path: pulumi.String("dev-customer-events"),
					Metadata: pulumi.Map{
						"retention_ms":        pulumi.Any("86400000"),
						"auto_create_topics":  pulumi.Any("true"),
						"cleanup_policy":      pulumi.Any("delete"),
						"min_insync_replicas": pulumi.Any("1"),
						"max_message_bytes":   pulumi.Any("1048576"),
					},
				},
				"test": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Testing"),
					Path: pulumi.String("test-customer-events"),
					Metadata: pulumi.Map{
						"retention_ms":        pulumi.Any("259200000"),
						"auto_create_topics":  pulumi.Any("true"),
						"cleanup_policy":      pulumi.Any("delete"),
						"min_insync_replicas": pulumi.Any("2"),
						"max_message_bytes":   pulumi.Any("1048576"),
					},
				},
				"prod": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Production"),
					Path: pulumi.String("prod-customer-events"),
					Metadata: pulumi.Map{
						"retention_ms":        pulumi.Any("604800000"),
						"auto_create_topics":  pulumi.Any("false"),
						"cleanup_policy":      pulumi.Any("compact,delete"),
						"min_insync_replicas": pulumi.Any("2"),
						"max_message_bytes":   pulumi.Any("1048576"),
						"monitoring_enabled":  pulumi.Any("true"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		postgresAsset, err := marmot.NewAsset(ctx, "postgresAsset", &marmot.AssetArgs{
			Name:        pulumi.String("customer-data-warehouse"),
			Type:        pulumi.String("Database"),
			Description: pulumi.String("PostgreSQL database for customer data"),
			Services: pulumi.StringArray{
				pulumi.String("PostgreSQL"),
			},
			Tags: pulumi.StringArray{
				pulumi.String("database"),
				pulumi.String("warehouse"),
				pulumi.String("structured-data"),
			},
			Metadata: pulumi.Map{
				"owner":   pulumi.Any("data-team"),
				"version": pulumi.Any("14.5"),
				"size":    pulumi.Any("medium"),
			},
			Schema: pulumi.Map{
				"tables": pulumi.Any{
					map[string]interface{}{
						"name":        "customers",
						"description": "Customer records",
						"columns": []map[string]interface{}{
							map[string]interface{}{
								"name": "id",
								"type": "uuid",
							},
							map[string]interface{}{
								"name": "email",
								"type": "varchar(255)",
							},
							map[string]interface{}{
								"name": "created_at",
								"type": "timestamp",
							},
						},
					},
					map[string]interface{}{
						"name":        "orders",
						"description": "Customer orders",
						"columns": []map[string]interface{}{
							map[string]interface{}{
								"name": "id",
								"type": "uuid",
							},
							map[string]interface{}{
								"name": "customer_id",
								"type": "uuid",
							},
							map[string]interface{}{
								"name": "total",
								"type": "decimal(10,2)",
							},
						},
					},
				},
			},
			ExternalLinks: marmot.ExternalLinkArray{
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Database Admin"),
					Url:  pulumi.String("http://pgadmin.example.com"),
				},
			},
			Environments: marmot.AssetEnvironmentMap{
				"dev": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Development"),
					Path: pulumi.String("dev-db/customer_data"),
					Metadata: pulumi.Map{
						"backup_frequency": pulumi.Any("daily"),
					},
				},
				"prod": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Production"),
					Path: pulumi.String("prod-db/customer_data"),
					Metadata: pulumi.Map{
						"backup_frequency": pulumi.Any("hourly"),
						"ha_enabled":       pulumi.Any("true"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		s3Asset, err := marmot.NewAsset(ctx, "S3Asset", &marmot.AssetArgs{
			Name:        pulumi.String("customer-data-lake"),
			Type:        pulumi.String("Bucket"),
			Description: pulumi.String("S3 bucket for customer data lake storage"),
			Services: pulumi.StringArray{
				pulumi.String("S3"),
			},
			Tags: pulumi.StringArray{
				pulumi.String("storage"),
				pulumi.String("data-lake"),
				pulumi.String("raw-data"),
			},
			Metadata: pulumi.Map{
				"owner":            pulumi.Any("data-platform-team"),
				"region":           pulumi.Any("us-west-2"),
				"lifecycle_policy": pulumi.Any("glacier-after-90-days"),
			},
			ExternalLinks: marmot.ExternalLinkArray{
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("AWS Console"),
					Url:  pulumi.String("https://console.aws.amazon.com/S3/buckets/customer-data-lake"),
				},
			},
			Environments: marmot.AssetEnvironmentMap{
				"dev": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Development"),
					Path: pulumi.String("dev-customer-data-lake"),
					Metadata: pulumi.Map{
						"versioning": pulumi.Any("enabled"),
					},
				},
				"prod": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Production"),
					Path: pulumi.String("prod-customer-data-lake"),
					Metadata: pulumi.Map{
						"versioning": pulumi.Any("enabled"),
						"encryption": pulumi.Any("AES-256"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		serviceAsset, err := marmot.NewAsset(ctx, "serviceAsset", &marmot.AssetArgs{
			Name:        pulumi.String("order-processing-service"),
			Type:        pulumi.String("Service"),
			Description: pulumi.String("Microservice for processing customer orders"),
			Services: pulumi.StringArray{
				pulumi.String("Kubernetes"),
			},
			Tags: pulumi.StringArray{
				pulumi.String("microservice"),
				pulumi.String("orders"),
				pulumi.String("processing"),
			},
			Metadata: pulumi.Map{
				"owner":      pulumi.Any("order-team"),
				"language":   pulumi.Any("golang"),
				"version":    pulumi.Any("1.2.3"),
				"repository": pulumi.Any("github.com/example/order-processor"),
			},
			ExternalLinks: marmot.ExternalLinkArray{
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("API Documentation"),
					Url:  pulumi.String("http://docs.example.com/order-api"),
				},
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Monitoring Dashboard"),
					Url:  pulumi.String("http://grafana.example.com/order-service"),
				},
				&marmot.ExternalLinkArgs{
					Name: pulumi.String("Repository"),
					Url:  pulumi.String("https://github.com/example/order-processor"),
					Icon: pulumi.String("github"),
				},
			},
			Environments: marmot.AssetEnvironmentMap{
				"dev": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Development"),
					Path: pulumi.String("dev/order-processor"),
					Metadata: pulumi.Map{
						"replicas": pulumi.Any("1"),
					},
				},
				"prod": &marmot.AssetEnvironmentArgs{
					Name: pulumi.String("Production"),
					Path: pulumi.String("prod/order-processor"),
					Metadata: pulumi.Map{
						"replicas":    pulumi.Any("3"),
						"autoscaling": pulumi.Any("enabled"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		kafkaToServiceLineage, err := marmot.NewLineage(ctx, "kafkaToServiceLineage", &marmot.LineageArgs{
			Source: kafkaAsset.Mrn,
			Target: serviceAsset.Mrn,
		}, pulumi.DependsOn([]pulumi.Resource{
			kafkaAsset,
			serviceAsset,
		}), pulumi.IgnoreChanges([]string{
			resourceId,
		}))
		if err != nil {
			return err
		}
		serviceToPostgresLineage, err := marmot.NewLineage(ctx, "serviceToPostgresLineage", &marmot.LineageArgs{
			Source: serviceAsset.Mrn,
			Target: postgresAsset.Mrn,
		}, pulumi.DependsOn([]pulumi.Resource{
			serviceAsset,
			postgresAsset,
		}), pulumi.IgnoreChanges([]string{
			resourceId,
		}))
		if err != nil {
			return err
		}
		_, err = marmot.NewLineage(ctx, "serviceToS3Lineage", &marmot.LineageArgs{
			Source: serviceAsset.Mrn,
			Target: s3Asset.Mrn,
		}, pulumi.DependsOn([]pulumi.Resource{
			serviceAsset,
			s3Asset,
		}), pulumi.IgnoreChanges([]string{
			resourceId,
		}))
		if err != nil {
			return err
		}
		ctx.Export("kafkaAssetId", pulumi.StringMap{
			"value": kafkaAsset.ResourceId,
		})
		ctx.Export("kafkaAssetMrn", pulumi.StringMap{
			"value": kafkaAsset.Mrn,
		})
		ctx.Export("postgresAssetId", pulumi.StringMap{
			"value": postgresAsset.ResourceId,
		})
		ctx.Export("postgresAssetMrn", pulumi.StringMap{
			"value": postgresAsset.Mrn,
		})
		ctx.Export("S3AssetId", pulumi.StringMap{
			"value": s3Asset.ResourceId,
		})
		ctx.Export("S3AssetMrn", pulumi.StringMap{
			"value": s3Asset.Mrn,
		})
		ctx.Export("serviceAssetId", pulumi.StringMap{
			"value": serviceAsset.ResourceId,
		})
		ctx.Export("serviceAssetMrn", pulumi.StringMap{
			"value": serviceAsset.Mrn,
		})
		ctx.Export("kafkaToServiceLineageId", pulumi.StringMap{
			"value": kafkaToServiceLineage.ResourceId,
		})
		ctx.Export("serviceToPostgresLineageId", pulumi.StringMap{
			"value": serviceToPostgresLineage.ResourceId,
		})
		return nil
	})
}
