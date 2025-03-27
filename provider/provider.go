package provider

import (
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/marmotdata/pulumi-marmot/provider/internal/client/client"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsGen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "marmot"

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Asset](),
			infer.Resource[Lineage](),
		},
		Config: infer.Config[Config](),
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Metadata: schema.Metadata{
			DisplayName: "marmot",
			Description: "Provider for Marmot",
			LanguageMap: map[string]any{
				"go": gen.GoPackageInfo{
					GenerateExtraInputTypes:        true,
					GenerateResourceContainerTypes: true,
					ImportBasePath:                 "github.com/marmotdata/pulumi-marmot/sdk/go/marmot",
				},
				"nodejs": nodejsGen.NodePackageInfo{
					PackageName: "@marmotdata/pulumi",
				},
			},
		},
	})
}

type Config struct {
	Host   string `pulumi:"host"`
	APIKey string `pulumi:"apiKey"`
}

func (c *Config) GetClient() (*client.Marmot, error) {
	// Parse host to determine scheme
	host := c.Host
	scheme := "https"

	if strings.HasPrefix(host, "http://") {
		scheme = "http"
		host = strings.TrimPrefix(host, "http://")
	} else if strings.HasPrefix(host, "https://") {
		host = strings.TrimPrefix(host, "https://")
	}

	transport := httptransport.New(host, "/api/v1", []string{scheme})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("X-API-Key", "header", c.APIKey)
	return client.New(transport, nil), nil
}
