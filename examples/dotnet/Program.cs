using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Marmot = Pulumi.Marmot;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var marmotHost = config.Require("marmot:host");
    var marmotApiKey = config.Require("marmot:apiKey");
    var kafkaAsset = new Marmot.Asset("kafkaAsset", new()
    {
        Name = "customer-events-stream",
        Type = "Topic",
        Description = "Kafka stream for customer events",
        Services = new[]
        {
            "Kafka",
        },
        Tags = new[]
        {
            "events",
            "streaming",
            "real-time",
            "customer-data",
        },
        Metadata = 
        {
            { "owner", "platform-team" },
            { "partitions", "24" },
            { "replication_factor", "3" },
            { "retention_ms", "604800000" },
            { "security_protocol", "SASL_SSL" },
            { "compression_type", "lz4" },
            { "max_message_size", "1048576" },
            { "group_id", "customer-events-consumers" },
        },
        Schema = 
        {
            { "type", "avro" },
            { "doc", "Schema for customer events" },
            { "name", "CustomerEvent" },
            { "namespace", "com.example.events" },
            { "fields", new[]
            {
                
                {
                    { "name", "event_id" },
                    { "type", "string" },
                    { "doc", "Unique identifier for the event" },
                },
                
                {
                    { "name", "customer_id" },
                    { "type", "string" },
                    { "doc", "Customer identifier" },
                },
                
                {
                    { "name", "event_type" },
                    { "type", 
                    {
                        { "type", "enum" },
                        { "name", "EventType" },
                        { "symbols", new[]
                        {
                            "SIGN_UP",
                            "LOGIN",
                            "PURCHASE",
                            "ACCOUNT_UPDATE",
                            "LOGOUT",
                        } },
                    } },
                },
                
                {
                    { "name", "timestamp" },
                    { "type", "long" },
                    { "doc", "Event timestamp in milliseconds since epoch" },
                },
                
                {
                    { "name", "payload" },
                    { "type", 
                    {
                        { "type", "record" },
                        { "name", "EventPayload" },
                        { "fields", new[]
                        {
                            
                            {
                                { "name", "session_id" },
                                { "type", new[]
                                {
                                    "null",
                                    "string",
                                } },
                            },
                            
                            {
                                { "name", "device_info" },
                                { "type", 
                                {
                                    { "type", "record" },
                                    { "name", "DeviceInfo" },
                                    { "fields", new[]
                                    {
                                        
                                        {
                                            { "name", "type" },
                                            { "type", new[]
                                            {
                                                "null",
                                                "string",
                                            } },
                                        },
                                        
                                        {
                                            { "name", "os" },
                                            { "type", new[]
                                            {
                                                "null",
                                                "string",
                                            } },
                                        },
                                        
                                        {
                                            { "name", "browser" },
                                            { "type", new[]
                                            {
                                                "null",
                                                "string",
                                            } },
                                        },
                                    } },
                                } },
                            },
                            
                            {
                                { "name", "properties" },
                                { "type", 
                                {
                                    { "type", "map" },
                                    { "values", "string" },
                                } },
                            },
                        } },
                    } },
                },
            } },
        },
        ExternalLinks = new[]
        {
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Kafka UI",
                Url = "http://kafka-ui.example.com",
            },
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Monitoring",
                Url = "http://grafana.example.com/kafka-dashboard",
            },
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Schema Registry",
                Url = "http://schema-registry.example.com/subjects/customer-events/versions/latest",
                Icon = "database-schema",
            },
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Documentation",
                Url = "http://docs.example.com/kafka/customer-events",
                Icon = "book",
            },
        },
        Environments = 
        {
            { "dev", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Development",
                Path = "dev-customer-events",
                Metadata = 
                {
                    { "retention_ms", "86400000" },
                    { "auto_create_topics", "true" },
                    { "cleanup_policy", "delete" },
                    { "min_insync_replicas", "1" },
                    { "max_message_bytes", "1048576" },
                },
            } },
            { "test", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Testing",
                Path = "test-customer-events",
                Metadata = 
                {
                    { "retention_ms", "259200000" },
                    { "auto_create_topics", "true" },
                    { "cleanup_policy", "delete" },
                    { "min_insync_replicas", "2" },
                    { "max_message_bytes", "1048576" },
                },
            } },
            { "prod", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Production",
                Path = "prod-customer-events",
                Metadata = 
                {
                    { "retention_ms", "604800000" },
                    { "auto_create_topics", "false" },
                    { "cleanup_policy", "compact,delete" },
                    { "min_insync_replicas", "2" },
                    { "max_message_bytes", "1048576" },
                    { "monitoring_enabled", "true" },
                },
            } },
        },
    });

    var postgresAsset = new Marmot.Asset("postgresAsset", new()
    {
        Name = "customer-data-warehouse",
        Type = "Database",
        Description = "PostgreSQL database for customer data",
        Services = new[]
        {
            "PostgreSQL",
        },
        Tags = new[]
        {
            "database",
            "warehouse",
            "structured-data",
        },
        Metadata = 
        {
            { "owner", "data-team" },
            { "version", "14.5" },
            { "size", "medium" },
        },
        Schema = 
        {
            { "tables", new[]
            {
                
                {
                    { "name", "customers" },
                    { "description", "Customer records" },
                    { "columns", new[]
                    {
                        
                        {
                            { "name", "id" },
                            { "type", "uuid" },
                        },
                        
                        {
                            { "name", "email" },
                            { "type", "varchar(255)" },
                        },
                        
                        {
                            { "name", "created_at" },
                            { "type", "timestamp" },
                        },
                    } },
                },
                
                {
                    { "name", "orders" },
                    { "description", "Customer orders" },
                    { "columns", new[]
                    {
                        
                        {
                            { "name", "id" },
                            { "type", "uuid" },
                        },
                        
                        {
                            { "name", "customer_id" },
                            { "type", "uuid" },
                        },
                        
                        {
                            { "name", "total" },
                            { "type", "decimal(10,2)" },
                        },
                    } },
                },
            } },
        },
        ExternalLinks = new[]
        {
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Database Admin",
                Url = "http://pgadmin.example.com",
            },
        },
        Environments = 
        {
            { "dev", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Development",
                Path = "dev-db/customer_data",
                Metadata = 
                {
                    { "backup_frequency", "daily" },
                },
            } },
            { "prod", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Production",
                Path = "prod-db/customer_data",
                Metadata = 
                {
                    { "backup_frequency", "hourly" },
                    { "ha_enabled", "true" },
                },
            } },
        },
    });

    var s3Asset = new Marmot.Asset("S3Asset", new()
    {
        Name = "customer-data-lake",
        Type = "Bucket",
        Description = "S3 bucket for customer data lake storage",
        Services = new[]
        {
            "S3",
        },
        Tags = new[]
        {
            "storage",
            "data-lake",
            "raw-data",
        },
        Metadata = 
        {
            { "owner", "data-platform-team" },
            { "region", "us-west-2" },
            { "lifecycle_policy", "glacier-after-90-days" },
        },
        ExternalLinks = new[]
        {
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "AWS Console",
                Url = "https://console.aws.amazon.com/S3/buckets/customer-data-lake",
            },
        },
        Environments = 
        {
            { "dev", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Development",
                Path = "dev-customer-data-lake",
                Metadata = 
                {
                    { "versioning", "enabled" },
                },
            } },
            { "prod", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Production",
                Path = "prod-customer-data-lake",
                Metadata = 
                {
                    { "versioning", "enabled" },
                    { "encryption", "AES-256" },
                },
            } },
        },
    });

    var serviceAsset = new Marmot.Asset("serviceAsset", new()
    {
        Name = "order-processing-service",
        Type = "Service",
        Description = "Microservice for processing customer orders",
        Services = new[]
        {
            "Kubernetes",
        },
        Tags = new[]
        {
            "microservice",
            "orders",
            "processing",
        },
        Metadata = 
        {
            { "owner", "order-team" },
            { "language", "golang" },
            { "version", "1.2.3" },
            { "repository", "github.com/example/order-processor" },
        },
        ExternalLinks = new[]
        {
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "API Documentation",
                Url = "http://docs.example.com/order-api",
            },
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Monitoring Dashboard",
                Url = "http://grafana.example.com/order-service",
            },
            new Marmot.Inputs.ExternalLinkArgs
            {
                Name = "Repository",
                Url = "https://github.com/example/order-processor",
                Icon = "github",
            },
        },
        Environments = 
        {
            { "dev", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Development",
                Path = "dev/order-processor",
                Metadata = 
                {
                    { "replicas", "1" },
                },
            } },
            { "prod", new Marmot.Inputs.AssetEnvironmentArgs
            {
                Name = "Production",
                Path = "prod/order-processor",
                Metadata = 
                {
                    { "replicas", "3" },
                    { "autoscaling", "enabled" },
                },
            } },
        },
    });

    var kafkaToServiceLineage = new Marmot.Lineage("kafkaToServiceLineage", new()
    {
        Source = kafkaAsset.Mrn,
        Target = serviceAsset.Mrn,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            kafkaAsset,
            serviceAsset,
        },
        IgnoreChanges =
        {
            "resourceId",
        },
    });

    var serviceToPostgresLineage = new Marmot.Lineage("serviceToPostgresLineage", new()
    {
        Source = serviceAsset.Mrn,
        Target = postgresAsset.Mrn,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            serviceAsset,
            postgresAsset,
        },
        IgnoreChanges =
        {
            "resourceId",
        },
    });

    var serviceToS3Lineage = new Marmot.Lineage("serviceToS3Lineage", new()
    {
        Source = serviceAsset.Mrn,
        Target = s3Asset.Mrn,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            serviceAsset,
            s3Asset,
        },
        IgnoreChanges =
        {
            "resourceId",
        },
    });

    return new Dictionary<string, object?>
    {
        ["kafkaAssetId"] = 
        {
            { "value", kafkaAsset.ResourceId },
        },
        ["kafkaAssetMrn"] = 
        {
            { "value", kafkaAsset.Mrn },
        },
        ["postgresAssetId"] = 
        {
            { "value", postgresAsset.ResourceId },
        },
        ["postgresAssetMrn"] = 
        {
            { "value", postgresAsset.Mrn },
        },
        ["S3AssetId"] = 
        {
            { "value", s3Asset.ResourceId },
        },
        ["S3AssetMrn"] = 
        {
            { "value", s3Asset.Mrn },
        },
        ["serviceAssetId"] = 
        {
            { "value", serviceAsset.ResourceId },
        },
        ["serviceAssetMrn"] = 
        {
            { "value", serviceAsset.Mrn },
        },
        ["kafkaToServiceLineageId"] = 
        {
            { "value", kafkaToServiceLineage.ResourceId },
        },
        ["serviceToPostgresLineageId"] = 
        {
            { "value", serviceToPostgresLineage.ResourceId },
        },
    };
});

