package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	// "github.com/liatrio/engineering-bootcamp/examples/terraform-provider/devops-resources"
)

var (
	_ datasource.DataSource              = &engineersDataSource{}
	_ datasource.DataSourceWithConfigure = &engineersDataSource{}
)

func NewEngineersDataSource() datasource.DataSource {
	return &engineersDataSource{}
}

type engineersDataSource struct {
	client *Client
}

type EngineersDataSourceModel struct {
	Engineers []engineerModel `tfsdk:"engineers"`
}

// In the client.go file
// type Engineer struct {
// 	ID    string `tfsdk:"id"`
// 	Name  string `tfsdk:"name"`
// 	Email string `tfsdk:"email"`
// }

type engineerModel struct {
	ID    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Email types.String `tfsdk:"email"`
}

func (d *engineersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_engineers"
}

func (d *engineersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T", req.ProviderData))
		return
	}
	d.client = client
}

func (d *engineersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"engineers": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The unique identifier for the engineer.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the engineer.",
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "The email address of the engineer.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *engineersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EngineersDataSourceModel

	engineers, err := d.client.GetEngineers()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HashiCups Engineers",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, engineer := range engineers {
		engineerState := engineerModel{
			ID:    types.StringValue(engineer.ID),
			Name:  types.StringValue(engineer.Name),
			Email: types.StringValue(engineer.Email),
		}

		state.Engineers = append(state.Engineers, engineerState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
