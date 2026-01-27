package environment

import (
	"context"
	"fmt"

	"github.com/Altiorem-Labs/terraform-provider-coolify/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &EnvironmentResource{}
var _ resource.ResourceWithConfigure = &EnvironmentResource{}

type EnvironmentResource struct {
	client *client.CoolifyClient
}

type EnvironmentResourceModel struct {
	ProjectID types.String `tfsdk:"project_id"`
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
}

func NewEnvironmentResource() resource.Resource {
	return &EnvironmentResource{}
}

func (r *EnvironmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments"
}

func (r *EnvironmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a Coolify environment.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the project to create the environment in.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier (UUID) for the environment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the environment.",
			},
		},
	}
}

func (r *EnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.CoolifyClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.CoolifyClient, got: %T", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *EnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EnvironmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environment, err := r.client.CreateEnvironment(plan.ProjectID.ValueString(), plan.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating environment", err.Error())
		return
	}

	plan.ID = types.StringValue(environment.UUID)
	// Keep the original planned name since API might return empty name
	if environment.Name != "" {
		plan.Name = types.StringValue(environment.Name)
	}
	plan.ProjectID = types.StringValue(plan.ProjectID.ValueString())

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environment, err := r.client.GetEnvironment(state.ProjectID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading environment", err.Error())
		return
	}

	// If environment not found, mark as removed
	if environment == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	// Update state with current values
	state.ID = types.StringValue(environment.UUID)
	if environment.Name != "" {
		state.Name = types.StringValue(environment.Name)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *EnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Coolify API doesn't support updating environments
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"The Coolify API does not support updating environments. Please recreate the resource with the new configuration.",
	)
}

func (r *EnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteEnvironment(state.ProjectID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting environment", err.Error())
		return
	}
}
