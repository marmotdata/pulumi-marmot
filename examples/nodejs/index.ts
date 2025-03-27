import * as pulumi from "@pulumi/pulumi";
import * as marmot from "@marmotdata/pulumi";

const config = new pulumi.Config();
const marmotHost = config.require("marmot:host");
const marmotApiKey = config.require("marmot:apiKey");
const kafkaAsset = new marmot.Asset("kafkaAsset", {
    name: "customer-events-stream",
    type: "Topic",
    description: "Kafka stream for customer events",
    services: ["Kafka"],
    tags: [
        "events",
        "streaming",
        "real-time",
        "customer-data",
    ],
    metadata: {
        owner: "platform-team",
        partitions: "24",
        replication_factor: "3",
        retention_ms: "604800000",
        security_protocol: "SASL_SSL",
        compression_type: "lz4",
        max_message_size: "1048576",
        group_id: "customer-events-consumers",
    },
    schema: {
        type: "avro",
        doc: "Schema for customer events",
        name: "CustomerEvent",
        namespace: "com.example.events",
        fields: [
            {
                name: "event_id",
                type: "string",
                doc: "Unique identifier for the event",
            },
            {
                name: "customer_id",
                type: "string",
                doc: "Customer identifier",
            },
            {
                name: "event_type",
                type: {
                    type: "enum",
                    name: "EventType",
                    symbols: [
                        "SIGN_UP",
                        "LOGIN",
                        "PURCHASE",
                        "ACCOUNT_UPDATE",
                        "LOGOUT",
                    ],
                },
            },
            {
                name: "timestamp",
                type: "long",
                doc: "Event timestamp in milliseconds since epoch",
            },
            {
                name: "payload",
                type: {
                    type: "record",
                    name: "EventPayload",
                    fields: [
                        {
                            name: "session_id",
                            type: [
                                "null",
                                "string",
                            ],
                        },
                        {
                            name: "device_info",
                            type: {
                                type: "record",
                                name: "DeviceInfo",
                                fields: [
                                    {
                                        name: "type",
                                        type: [
                                            "null",
                                            "string",
                                        ],
                                    },
                                    {
                                        name: "os",
                                        type: [
                                            "null",
                                            "string",
                                        ],
                                    },
                                    {
                                        name: "browser",
                                        type: [
                                            "null",
                                            "string",
                                        ],
                                    },
                                ],
                            },
                        },
                        {
                            name: "properties",
                            type: {
                                type: "map",
                                values: "string",
                            },
                        },
                    ],
                },
            },
        ],
    },
    externalLinks: [
        {
            name: "Kafka UI",
            url: "http://kafka-ui.example.com",
        },
        {
            name: "Monitoring",
            url: "http://grafana.example.com/kafka-dashboard",
        },
        {
            name: "Schema Registry",
            url: "http://schema-registry.example.com/subjects/customer-events/versions/latest",
            icon: "database-schema",
        },
        {
            name: "Documentation",
            url: "http://docs.example.com/kafka/customer-events",
            icon: "book",
        },
    ],
    environments: {
        dev: {
            name: "Development",
            path: "dev-customer-events",
            metadata: {
                retention_ms: "86400000",
                auto_create_topics: "true",
                cleanup_policy: "delete",
                min_insync_replicas: "1",
                max_message_bytes: "1048576",
            },
        },
        test: {
            name: "Testing",
            path: "test-customer-events",
            metadata: {
                retention_ms: "259200000",
                auto_create_topics: "true",
                cleanup_policy: "delete",
                min_insync_replicas: "2",
                max_message_bytes: "1048576",
            },
        },
        prod: {
            name: "Production",
            path: "prod-customer-events",
            metadata: {
                retention_ms: "604800000",
                auto_create_topics: "false",
                cleanup_policy: "compact,delete",
                min_insync_replicas: "2",
                max_message_bytes: "1048576",
                monitoring_enabled: "true",
            },
        },
    },
});
const postgresAsset = new marmot.Asset("postgresAsset", {
    name: "customer-data-warehouse",
    type: "Database",
    description: "PostgreSQL database for customer data",
    services: ["PostgreSQL"],
    tags: [
        "database",
        "warehouse",
        "structured-data",
    ],
    metadata: {
        owner: "data-team",
        version: "14.5",
        size: "medium",
    },
    schema: {
        tables: [
            {
                name: "customers",
                description: "Customer records",
                columns: [
                    {
                        name: "id",
                        type: "uuid",
                    },
                    {
                        name: "email",
                        type: "varchar(255)",
                    },
                    {
                        name: "created_at",
                        type: "timestamp",
                    },
                ],
            },
            {
                name: "orders",
                description: "Customer orders",
                columns: [
                    {
                        name: "id",
                        type: "uuid",
                    },
                    {
                        name: "customer_id",
                        type: "uuid",
                    },
                    {
                        name: "total",
                        type: "decimal(10,2)",
                    },
                ],
            },
        ],
    },
    externalLinks: [{
        name: "Database Admin",
        url: "http://pgadmin.example.com",
    }],
    environments: {
        dev: {
            name: "Development",
            path: "dev-db/customer_data",
            metadata: {
                backup_frequency: "daily",
            },
        },
        prod: {
            name: "Production",
            path: "prod-db/customer_data",
            metadata: {
                backup_frequency: "hourly",
                ha_enabled: "true",
            },
        },
    },
});
const s3Asset = new marmot.Asset("S3Asset", {
    name: "customer-data-lake",
    type: "Bucket",
    description: "S3 bucket for customer data lake storage",
    services: ["S3"],
    tags: [
        "storage",
        "data-lake",
        "raw-data",
    ],
    metadata: {
        owner: "data-platform-team",
        region: "us-west-2",
        lifecycle_policy: "glacier-after-90-days",
    },
    externalLinks: [{
        name: "AWS Console",
        url: "https://console.aws.amazon.com/S3/buckets/customer-data-lake",
    }],
    environments: {
        dev: {
            name: "Development",
            path: "dev-customer-data-lake",
            metadata: {
                versioning: "enabled",
            },
        },
        prod: {
            name: "Production",
            path: "prod-customer-data-lake",
            metadata: {
                versioning: "enabled",
                encryption: "AES-256",
            },
        },
    },
});
const serviceAsset = new marmot.Asset("serviceAsset", {
    name: "order-processing-service",
    type: "Service",
    description: "Microservice for processing customer orders",
    services: ["Kubernetes"],
    tags: [
        "microservice",
        "orders",
        "processing",
    ],
    metadata: {
        owner: "order-team",
        language: "golang",
        version: "1.2.3",
        repository: "github.com/example/order-processor",
    },
    externalLinks: [
        {
            name: "API Documentation",
            url: "http://docs.example.com/order-api",
        },
        {
            name: "Monitoring Dashboard",
            url: "http://grafana.example.com/order-service",
        },
        {
            name: "Repository",
            url: "https://github.com/example/order-processor",
            icon: "github",
        },
    ],
    environments: {
        dev: {
            name: "Development",
            path: "dev/order-processor",
            metadata: {
                replicas: "1",
            },
        },
        prod: {
            name: "Production",
            path: "prod/order-processor",
            metadata: {
                replicas: "3",
                autoscaling: "enabled",
            },
        },
    },
});
const kafkaToServiceLineage = new marmot.Lineage("kafkaToServiceLineage", {
    source: kafkaAsset.mrn,
    target: serviceAsset.mrn,
}, {
    dependsOn: [
        kafkaAsset,
        serviceAsset,
    ],
    ignoreChanges: [resourceId],
});
const serviceToPostgresLineage = new marmot.Lineage("serviceToPostgresLineage", {
    source: serviceAsset.mrn,
    target: postgresAsset.mrn,
}, {
    dependsOn: [
        serviceAsset,
        postgresAsset,
    ],
    ignoreChanges: [resourceId],
});
const serviceToS3Lineage = new marmot.Lineage("serviceToS3Lineage", {
    source: serviceAsset.mrn,
    target: s3Asset.mrn,
}, {
    dependsOn: [
        serviceAsset,
        s3Asset,
    ],
    ignoreChanges: [resourceId],
});
export const kafkaAssetId = {
    value: kafkaAsset.resourceId,
};
export const kafkaAssetMrn = {
    value: kafkaAsset.mrn,
};
export const postgresAssetId = {
    value: postgresAsset.resourceId,
};
export const postgresAssetMrn = {
    value: postgresAsset.mrn,
};
export const s3AssetId = {
    value: s3Asset.resourceId,
};
export const s3AssetMrn = {
    value: s3Asset.mrn,
};
export const serviceAssetId = {
    value: serviceAsset.resourceId,
};
export const serviceAssetMrn = {
    value: serviceAsset.mrn,
};
export const kafkaToServiceLineageId = {
    value: kafkaToServiceLineage.resourceId,
};
export const serviceToPostgresLineageId = {
    value: serviceToPostgresLineage.resourceId,
};
