package detector

import (
	"fmt"

	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/schema"
)

type AwsRouteInvalidNetworkInterfaceDetector struct {
	*Detector
	networkInterfaces map[string]bool
}

func (d *Detector) CreateAwsRouteInvalidNetworkInterfaceDetector() *AwsRouteInvalidNetworkInterfaceDetector {
	nd := &AwsRouteInvalidNetworkInterfaceDetector{
		Detector:          d,
		networkInterfaces: map[string]bool{},
	}
	nd.Name = "aws_route_invalid_network_interface"
	nd.IssueType = issue.ERROR
	nd.TargetType = "resource"
	nd.Target = "aws_route"
	nd.DeepCheck = true
	nd.Enabled = true
	return nd
}

func (d *AwsRouteInvalidNetworkInterfaceDetector) PreProcess() {
	resp, err := d.AwsClient.DescribeNetworkInterfaces()
	if err != nil {
		d.Logger.Error(err)
		d.Error = true
		return
	}

	for _, networkInterface := range resp.NetworkInterfaces {
		d.networkInterfaces[*networkInterface.NetworkInterfaceId] = true
	}
}

func (d *AwsRouteInvalidNetworkInterfaceDetector) Detect(resource *schema.Resource, issues *[]*issue.Issue) {
	networkInterfaceToken, ok := resource.GetToken("network_interface_id")
	if !ok {
		return
	}
	networkInterface, err := d.evalToString(networkInterfaceToken.Text)
	if err != nil {
		d.Logger.Error(err)
		return
	}

	if !d.networkInterfaces[networkInterface] {
		issue := &issue.Issue{
			Detector: d.Name,
			Type:     d.IssueType,
			Message:  fmt.Sprintf("\"%s\" is invalid network interface ID.", networkInterface),
			Line:     networkInterfaceToken.Pos.Line,
			File:     networkInterfaceToken.Pos.Filename,
		}
		*issues = append(*issues, issue)
	}
}
