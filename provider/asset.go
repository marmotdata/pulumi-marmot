package provider

import (
	"context"
	"fmt"
	"reflect"

	"github.com/marmotdata/pulumi-marmot/provider/internal/client/client/assets"
	"github.com/marmotdata/pulumi-marmot/provider/internal/client/models"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Asset struct{}

type AssetSource struct {
	Name       string                 `pulumi:"name"`
	Priority   *int64                 `pulumi:"priority,optional"`
	Properties map[string]interface{} `pulumi:"properties,optional"`
}

type AssetEnvironment struct {
	Name     string                 `pulumi:"name"`
	Path     string                 `pulumi:"path"`
	Metadata map[string]interface{} `pulumi:"metadata,optional"`
}

type ExternalLink struct {
	Icon *string `pulumi:"icon,optional"`
	Name string  `pulumi:"name"`
	URL  string  `pulumi:"url"`
}

type AssetArgs struct {
	Name          string                      `pulumi:"name"`
	Type          string                      `pulumi:"type"`
	Description   string                      `pulumi:"description,optional"`
	Providers     []string                    `pulumi:"services"`
	Tags          []string                    `pulumi:"tags,optional"`
	Metadata      map[string]interface{}      `pulumi:"metadata,optional"`
	Schema        map[string]interface{}      `pulumi:"schema,optional"`
	ExternalLinks []ExternalLink              `pulumi:"externalLinks,optional"`
	Sources       []AssetSource               `pulumi:"sources,optional"`
	Environments  map[string]AssetEnvironment `pulumi:"environments,optional"`
}

type AssetState struct {
	AssetArgs
	ResourceID string  `pulumi:"resourceId"`
	CreatedAt  string  `pulumi:"createdAt"`
	CreatedBy  string  `pulumi:"createdBy"`
	UpdatedAt  string  `pulumi:"updatedAt"`
	LastSyncAt *string `pulumi:"lastSyncAt,optional"`
	MRN        string  `pulumi:"mrn"`
}

func normalizeMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			result[k] = normalizeMap(val)
		case []interface{}:
			result[k] = normalizeArray(val)
		case float64:
			if val == float64(int64(val)) {
				result[k] = int64(val)
			} else {
				result[k] = val
			}
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func normalizeArray(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		switch val := v.(type) {
		case map[string]interface{}:
			result[i] = normalizeMap(val)
		case []interface{}:
			result[i] = normalizeArray(val)
		default:
			result[i] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func (Asset) Diff(ctx context.Context, id string, olds AssetState, news AssetArgs) (p.DiffResponse, error) {
	detailedDiff := map[string]p.PropertyDiff{}
	hasChanges := false

	// Basic field comparisons
	if olds.Name != news.Name {
		detailedDiff["name"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}
	if olds.Type != news.Type {
		detailedDiff["type"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}
	if olds.Description != news.Description {
		detailedDiff["description"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}
	if !reflect.DeepEqual(olds.Providers, news.Providers) {
		detailedDiff["services"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}
	if !reflect.DeepEqual(olds.Tags, news.Tags) {
		detailedDiff["tags"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}
	if !reflect.DeepEqual(olds.ExternalLinks, news.ExternalLinks) {
		detailedDiff["externalLinks"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}

	// Compare metadata with normalization
	oldMeta := normalizeMapValues(olds.Metadata)
	newMeta := normalizeMapValues(news.Metadata)
	if !reflect.DeepEqual(oldMeta, newMeta) {
		for k := range oldMeta {
			if _, exists := newMeta[k]; !exists {
				detailedDiff[fmt.Sprintf("metadata[%q]", k)] = p.PropertyDiff{Kind: p.Delete}
				hasChanges = true
			}
		}
		for k := range newMeta {
			if _, exists := oldMeta[k]; !exists || !reflect.DeepEqual(oldMeta[k], newMeta[k]) {
				detailedDiff[fmt.Sprintf("metadata[%q]", k)] = p.PropertyDiff{Kind: p.Update}
				hasChanges = true
			}
		}
	}

	// Compare schema
	if !reflect.DeepEqual(normalizeMapValues(olds.Schema), normalizeMapValues(news.Schema)) {
		detailedDiff["schema"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}

	// Compare sources and environments
	if !reflect.DeepEqual(normalizeSources(olds.Sources), normalizeSources(news.Sources)) {
		detailedDiff["sources"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}

	if !reflect.DeepEqual(normalizeEnvironments(olds.Environments), normalizeEnvironments(news.Environments)) &&
		!(len(olds.Environments) == 0 && len(news.Environments) == 0) {
		detailedDiff["environments"] = p.PropertyDiff{Kind: p.Update}
		hasChanges = true
	}

	return p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          hasChanges,
		DetailedDiff:        detailedDiff,
	}, nil
}

func normalizeSources(sources []AssetSource) []map[string]interface{} {
	if sources == nil {
		return nil
	}
	result := make([]map[string]interface{}, len(sources))
	for i, s := range sources {
		m := map[string]interface{}{
			"name": s.Name,
		}
		if s.Properties != nil {
			m["properties"] = normalizeMapValues(s.Properties)
		}
		if s.Priority != nil {
			m["priority"] = fmt.Sprintf("%v", *s.Priority)
		}
		result[i] = m
	}
	return result
}

func normalizeEnvironments(envs map[string]AssetEnvironment) map[string]interface{} {
	if envs == nil || len(envs) == 0 {
		return nil
	}

	result := make(map[string]interface{})
	for k, v := range envs {
		m := map[string]interface{}{
			"name": v.Name,
			"path": v.Path,
		}
		if v.Metadata != nil {
			m["metadata"] = normalizeMapValues(v.Metadata)
		}
		result[k] = m
	}
	return result
}

func normalizeMapValues(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			result[k] = normalizeMapValues(val)
		case []interface{}:
			newArr := make([]interface{}, len(val))
			for i, item := range val {
				if mapItem, ok := item.(map[string]interface{}); ok {
					newArr[i] = normalizeMapValues(mapItem)
				} else {
					newArr[i] = fmt.Sprintf("%v", item)
				}
			}
			result[k] = newArr
		default:
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func sourcesEqual(a, b []AssetSource) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(normalizeMap(a[i].Properties), normalizeMap(b[i].Properties)) {
			return false
		}
	}
	return true
}

func environmentsEqual(a, b map[string]AssetEnvironment) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || !reflect.DeepEqual(normalizeMap(v.Metadata), normalizeMap(bv.Metadata)) {
			return false
		}
	}
	return true
}

func (Asset) Create(ctx context.Context, name string, input AssetArgs, preview bool) (string, AssetState, error) {
	state := AssetState{AssetArgs: input}
	if preview {
		return name, state, nil
	}

	asset, err := parseToAsset(input)
	if err != nil {
		return "", state, err
	}

	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return "", state, err
	}

	params := assets.NewPostAssetsParams().WithAsset(asset)
	result, err := client.Assets.PostAssets(params)
	if err != nil {
		return "", state, err
	}

	state = parseToAssetState(result.Payload)
	return result.Payload.ID, state, nil
}

func (Asset) Read(ctx context.Context, id string, inputs AssetArgs, state AssetState) (string, AssetArgs, AssetState, error) {
	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return "", inputs, state, err
	}

	params := assets.NewGetAssetsIDParams().WithID(id)
	result, err := client.Assets.GetAssetsID(params)
	if err != nil {
		return "", inputs, state, err
	}

	newState := parseToAssetState(result.Payload)
	return id, inputs, newState, nil
}

func (Asset) Update(ctx context.Context, id string, olds AssetState, news AssetArgs, preview bool) (AssetState, error) {
	if preview {
		return AssetState{
			AssetArgs:  news,
			ResourceID: olds.ResourceID,
			CreatedAt:  olds.CreatedAt,
			CreatedBy:  olds.CreatedBy,
			UpdatedAt:  olds.UpdatedAt,
			LastSyncAt: olds.LastSyncAt,
			MRN:        olds.MRN,
		}, nil
	}

	if len(news.Providers) == 0 {
		news.Providers = olds.Providers
	}

	asset, err := parseToAsset(news)
	if err != nil {
		return AssetState{}, err
	}

	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return AssetState{}, err
	}

	params := assets.NewPutAssetsIDParams().WithID(id).WithAsset(&models.AssetsUpdateRequest{
		Name:          news.Name,
		Type:          news.Type,
		Description:   news.Description,
		Providers:     news.Providers,
		Tags:          news.Tags,
		Metadata:      news.Metadata,
		Schema:        news.Schema,
		ExternalLinks: convertExternalLinks(news.ExternalLinks),
		Sources:       asset.Sources,
		Environments:  asset.Environments,
	})

	result, err := client.Assets.PutAssetsID(params)
	if err != nil {
		return AssetState{}, err
	}

	return parseToAssetState(result.Payload), nil
}

func (Asset) Delete(ctx context.Context, id string, state AssetState) error {
	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return err
	}

	params := assets.NewDeleteAssetsIDParams().WithID(id)
	_, err = client.Assets.DeleteAssetsID(params)
	return err
}

func parseToAsset(input AssetArgs) (*models.AssetsCreateRequest, error) {
	sources := make([]*models.AssetAssetSource, 0)
	for _, source := range input.Sources {
		sources = append(sources, &models.AssetAssetSource{
			Name:       source.Name,
			Priority:   int64Value(source.Priority),
			Properties: normalizeMap(source.Properties),
		})
	}

	environments := make(map[string]models.AssetEnvironment)
	for k, v := range input.Environments {
		environments[k] = models.AssetEnvironment{
			Name:     v.Name,
			Path:     v.Path,
			Metadata: normalizeMap(v.Metadata),
		}
	}

	normalizedMetadata := make(map[string]interface{})
	for k, v := range input.Metadata {
		switch val := v.(type) {
		case float64:
			if val == float64(int64(val)) {
				normalizedMetadata[k] = int64(val)
			} else {
				normalizedMetadata[k] = val
			}
		default:
			normalizedMetadata[k] = v
		}
	}

	return &models.AssetsCreateRequest{
		Name:          &input.Name,
		Type:          &input.Type,
		Description:   input.Description,
		Providers:     input.Providers,
		Tags:          input.Tags,
		Metadata:      normalizedMetadata,
		Schema:        normalizeMap(input.Schema),
		ExternalLinks: convertExternalLinks(input.ExternalLinks),
		Sources:       sources,
		Environments:  environments,
	}, nil
}

func parseToAssetState(asset *models.AssetAsset) AssetState {
	sources := make([]AssetSource, 0)
	if asset.Sources != nil {
		for _, source := range asset.Sources {
			var props map[string]interface{}
			if p, ok := source.Properties.(map[string]interface{}); ok {
				props = normalizeMap(p)
			}
			sources = append(sources, AssetSource{
				Name:       source.Name,
				Priority:   &source.Priority,
				Properties: props,
			})
		}
	}

	environments := make(map[string]AssetEnvironment)
	if asset.Environments != nil {
		for k, v := range asset.Environments {
			var meta map[string]interface{}
			if m, ok := v.Metadata.(map[string]interface{}); ok {
				meta = normalizeMap(m)
			}
			environments[k] = AssetEnvironment{
				Name:     v.Name,
				Path:     v.Path,
				Metadata: meta,
			}
		}
	}

	var metadata, schema map[string]interface{}
	if m, ok := asset.Metadata.(map[string]interface{}); ok {
		metadata = normalizeMap(m)
	}
	if s, ok := asset.Schema.(map[string]interface{}); ok {
		schema = normalizeMap(s)
	}

	return AssetState{
		AssetArgs: AssetArgs{
			Name:          asset.Name,
			Type:          asset.Type,
			Description:   asset.Description,
			Providers:     asset.Providers,
			Tags:          asset.Tags,
			Metadata:      metadata,
			Schema:        schema,
			ExternalLinks: convertModelExternalLinks(asset.ExternalLinks),
			Sources:       sources,
			Environments:  environments,
		},
		ResourceID: asset.ID,
		MRN:        asset.Mrn,
		CreatedAt:  asset.CreatedAt,
		CreatedBy:  asset.CreatedBy,
		UpdatedAt:  asset.UpdatedAt,
		LastSyncAt: &asset.LastSyncAt,
	}
}

func convertExternalLinks(links []ExternalLink) []*models.AssetExternalLink {
	if len(links) == 0 {
		return nil
	}
	result := make([]*models.AssetExternalLink, len(links))
	for i, link := range links {
		result[i] = &models.AssetExternalLink{
			Icon: stringValue(link.Icon),
			Name: link.Name,
			URL:  link.URL,
		}
	}
	return result
}

func convertModelExternalLinks(links []*models.AssetExternalLink) []ExternalLink {
	if len(links) == 0 {
		return nil
	}
	result := make([]ExternalLink, len(links))
	for i, link := range links {
		result[i] = ExternalLink{
			Name: link.Name,
			URL:  link.URL,
		}
		if link.Icon != "" {
			icon := link.Icon
			result[i].Icon = &icon
		}
	}
	return result
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
