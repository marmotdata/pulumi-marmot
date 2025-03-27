package provider

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/marmotdata/pulumi-marmot/provider/internal/client/client/lineage"
	"github.com/marmotdata/pulumi-marmot/provider/internal/client/models"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Lineage struct{}

type LineageArgs struct {
	Source string `pulumi:"source"`
	Target string `pulumi:"target"`
}

type LineageState struct {
	Source     string  `pulumi:"source"`
	Target     string  `pulumi:"target"`
	ResourceID string  `pulumi:"resourceId"`
	Type       *string `pulumi:"type,optional"`
}

func (Lineage) Create(ctx context.Context, name string, input LineageArgs, preview bool) (string, LineageState, error) {
	state := LineageState{
		Source: input.Source,
		Target: input.Target,
	}

	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return "", state, err
	}

	params := lineage.NewPostLineageDirectParams().WithEdge(&models.LineageLineageEdge{
		Source: input.Source,
		Target: input.Target,
	})

	result, err := client.Lineage.PostLineageDirect(params)
	if err != nil {
		return "", state, err
	}

	state.ResourceID = result.Payload.ID
	state.Type = &result.Payload.Type

	return result.Payload.ID, state, nil
}

func (Lineage) Read(ctx context.Context, id string, inputs LineageArgs, state LineageState) (string, LineageArgs, LineageState, error) {
	// Compare the actual string values, not output types
	if state.Source == inputs.Source && state.Target == inputs.Target {
		return id, inputs, state, nil
	}

	// Only create a new state if the values actually differ
	if state.Source != inputs.Source || state.Target != inputs.Target {
		return id, inputs, LineageState{
			Source:     inputs.Source,
			Target:     inputs.Target,
			ResourceID: state.ResourceID,
			Type:       state.Type,
		}, nil
	}

	return id, inputs, state, nil
}

func (Lineage) Delete(ctx context.Context, id string, state LineageState) error {
	config := infer.GetConfig[Config](ctx)
	client, err := config.GetClient()
	if err != nil {
		return err
	}

	params := lineage.NewDeleteLineageDirectIDParams().WithID(strfmt.UUID(id))
	_, err = client.Lineage.DeleteLineageDirectID(params)
	return err
}
