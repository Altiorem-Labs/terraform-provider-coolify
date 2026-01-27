package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/Altiorem-Labs/terraform-provider-coolify/internal/client"
	"github.com/Altiorem-Labs/terraform-provider-coolify/internal/resources/environment"
	"github.com/Altiorem-Labs/terraform-provider-coolify/internal/resources/project"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CoolifyProvider{}

type CoolifyProvider struct {
	version string
}

// DataSources implements [provider.Provider].
func (p *CoolifyProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Metadata implements [provider.Provider].
func (p *CoolifyProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "coolify"
	resp.Version = p.version
}

// Resources implements [provider.Provider].
func (p *CoolifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		project.NewProjectResource,
		environment.NewEnvironmentResource,
	}
}

// Schema implements [provider.Provider].
func (p *CoolifyProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with your self-hosted instance of Coolify.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The base URL of your Coolify instance (e.g., https://app.coolify.io). Also configurable via the `COOLIFY_ENDPOINT` environment variable.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Your Coolify API Token. Also configurable via the `COOLIFY_TOKEN` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

type CoolifyProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *CoolifyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config CoolifyProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.IsUnknown() || config.Token.IsUnknown() {
		return
	}

	endpoint := os.Getenv("COOLIFY_ENDPOINT")
	token := os.Getenv("COOLIFY_TOKEN")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Coolify endpoint is missing",
			"The provider cannot function without the API URL. Configure it in the provider block or use the COOLIFY_ENDPOINT environment variable.",
		)
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Coolify token is missing",
			"The provider requires authentication. Configure the token in the provider block or use the COOLIFY_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client := &client.CoolifyClient{
		Endpoint: endpoint,
		Token:    token,
		Client:   &http.Client{},
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CoolifyProvider{
			version: version,
		}
	}
}
