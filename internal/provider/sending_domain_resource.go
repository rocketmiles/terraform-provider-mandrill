package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rocketmiles/terraform-provider-mandrill/internal/provider/mandrill"
)

const (
	mandrillVerifyTxtRecordPrefix = "mandrill_verify."
)

type sendingDomainResourceType struct{}

func (t sendingDomainResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Sending domain",

		Attributes: map[string]tfsdk.Attribute{
			"domain_name": {
				MarkdownDescription: "Sending domain name",
				Required:            true,
				Type:                types.StringType,
			},
			"id": {
				Computed:            true,
				MarkdownDescription: "Sending domain name as ID",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"verify_txt_record": {
				Computed:            true,
				MarkdownDescription: "The full verify TXT record value to be set on domain zone",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"spf_valid": {
				Computed:            true,
				MarkdownDescription: "Set to true if SPF DNS configuration is valid",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.BoolType,
			},
			"dkim_valid": {
				Computed:            true,
				MarkdownDescription: "Set to true if DKIM DNS configuration is valid",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.BoolType,
			},
		},
	}, nil
}

func (t sendingDomainResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return sendingDomainResource{
		provider: provider,
	}, diags
}

type sendingDomainResourceData struct {
	DomainName      types.String `tfsdk:"domain_name"`
	Id              types.String `tfsdk:"id"`
	VerifyTxtRecord types.String `tfsdk:"verify_txt_record"`
	SpfValid        types.Bool   `tfsdk:"spf_valid"`
	DkimValid       types.Bool   `tfsdk:"dkim_valid"`
}

type sendingDomainResource struct {
	provider provider
}

func checkDomain(data *sendingDomainResourceData, r sendingDomainResource, diagnostics *diag.Diagnostics) {
	request := mandrill.SendersCheckDomainRequest{Domain: data.DomainName.Value}
	response, err := r.provider.client.SendersCheckDomain(request)

	if err != nil {
		diagnostics.AddError("Error sending request", err.Error())
	}

	data.Id = types.String{Value: response.Domain}
	data.DomainName = types.String{Value: response.Domain}
	data.VerifyTxtRecord = types.String{Value: mandrillVerifyTxtRecordPrefix + response.VerifyTxtKey}
	data.SpfValid = types.Bool{Value: response.Spf.Valid}
	data.DkimValid = types.Bool{Value: response.Dkim.Valid}
}

func (r sendingDomainResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data sendingDomainResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	checkDomain(&data, r, &resp.Diagnostics)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r sendingDomainResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data sendingDomainResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	checkDomain(&data, r, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r sendingDomainResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data sendingDomainResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	checkDomain(&data, r, &resp.Diagnostics)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r sendingDomainResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data sendingDomainResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.DeleteExample(...)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }

	resp.State.RemoveResource(ctx)
}

func (r sendingDomainResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("domain_name"), req, resp)
}
