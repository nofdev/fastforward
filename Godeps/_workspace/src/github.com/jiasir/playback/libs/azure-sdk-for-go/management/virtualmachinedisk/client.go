// Package virtualmachinedisk provides a client for Virtual Machine Disks.
package virtualmachinedisk

import (
	"encoding/xml"
	"fmt"
	"github.com/nofdev/fastforward/Godeps/_workspace/src/github.com/jiasir/playback/libs/azure-sdk-for-go/management"
)

const (
	addDataDiskURL    = "services/hostedservices/%s/deployments/%s/roles/%s/DataDisks"
	addDiskURL        = "services/disks"
	deleteDataDiskURL = "services/hostedservices/%s/deployments/%s/roles/%s/DataDisks/%d"
	deleteDiskURL     = "services/disks/%s"
	getDataDiskURL    = "services/hostedservices/%s/deployments/%s/roles/%s/DataDisks/%d"
	getDiskURL        = "services/disks/%s"
	listDisksURL      = "services/disks"
	updateDataDiskURL = "services/hostedservices/%s/deployments/%s/roles/%s/DataDisks/%d"
	updateDiskURL     = "services/disks/%s"

	errParamNotSpecified = "Parameter %s is not specified."
)

//NewClient is used to instantiate a new DiskClient from an Azure client
func NewClient(client management.Client) DiskClient {
	return DiskClient{client: client}
}

// AddDataDisk adds a data disk to a Virtual Machine
//
// https://msdn.microsoft.com/en-us/library/azure/jj157199.aspx
func (c DiskClient) AddDataDisk(
	service string,
	deployment string,
	role string,
	params CreateDataDiskParameters) (management.OperationID, error) {
	if service == "" {
		return "", fmt.Errorf(errParamNotSpecified, "service")
	}
	if deployment == "" {
		return "", fmt.Errorf(errParamNotSpecified, "deployment")
	}
	if role == "" {
		return "", fmt.Errorf(errParamNotSpecified, "role")
	}

	requestURL := fmt.Sprintf(addDataDiskURL, service, deployment, role)

	req, err := xml.Marshal(params)
	if err != nil {
		return "", err
	}

	return c.client.SendAzurePostRequest(requestURL, req)
}

// AddDisk adds an operating system disk or data disk to the user image repository
//
// https://msdn.microsoft.com/en-us/library/azure/jj157178.aspx
func (c DiskClient) AddDisk(params CreateDiskParameters) (management.OperationID, error) {
	req, err := xml.Marshal(params)
	if err != nil {
		return "", err
	}

	return c.client.SendAzurePostRequest(addDiskURL, req)
}

// DeleteDataDisk removes the specified data disk from a Virtual Machine
//
// https://msdn.microsoft.com/en-us/library/azure/jj157179.aspx
func (c DiskClient) DeleteDataDisk(
	service string,
	deployment string,
	role string,
	lun int,
	deleteVHD bool) (management.OperationID, error) {
	if service == "" {
		return "", fmt.Errorf(errParamNotSpecified, "service")
	}
	if deployment == "" {
		return "", fmt.Errorf(errParamNotSpecified, "deployment")
	}
	if role == "" {
		return "", fmt.Errorf(errParamNotSpecified, "role")
	}

	requestURL := fmt.Sprintf(deleteDataDiskURL, service, deployment, role, lun)
	if deleteVHD {
		requestURL += "?comp=media"
	}

	return c.client.SendAzureDeleteRequest(requestURL)
}

// DeleteDisk deletes the specified data or operating system disk from the image
// repository that is associated with the specified subscription
//
// https://msdn.microsoft.com/en-us/library/azure/jj157200.aspx
func (c DiskClient) DeleteDisk(name string, deleteVHD bool) (management.OperationID, error) {
	if name == "" {
		return "", fmt.Errorf(errParamNotSpecified, "name")
	}

	requestURL := fmt.Sprintf(deleteDiskURL, name)
	if deleteVHD {
		requestURL += "?comp=media"
	}

	return c.client.SendAzureDeleteRequest(requestURL)
}

// GetDataDisk retrieves the specified data disk from a Virtual Machine
//
// https://msdn.microsoft.com/en-us/library/azure/jj157180.aspx
func (c DiskClient) GetDataDisk(
	service string,
	deployment string,
	role string,
	lun int) (DataDiskResponse, error) {
	var response DataDiskResponse
	if service == "" {
		return response, fmt.Errorf(errParamNotSpecified, "service")
	}
	if deployment == "" {
		return response, fmt.Errorf(errParamNotSpecified, "deployment")
	}
	if role == "" {
		return response, fmt.Errorf(errParamNotSpecified, "role")
	}

	requestURL := fmt.Sprintf(getDataDiskURL, service, deployment, role, lun)

	data, err := c.client.SendAzureGetRequest(requestURL)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal(data, &response)
	return response, err
}

// GetDisk retrieves information about the specified disk
//
// https://msdn.microsoft.com/en-us/library/azure/dn775053.aspx
func (c DiskClient) GetDisk(name string) (DiskResponse, error) {
	var response DiskResponse
	if name == "" {
		return response, fmt.Errorf(errParamNotSpecified, "name")
	}

	requestURL := fmt.Sprintf(getDiskURL, name)

	data, err := c.client.SendAzureGetRequest(requestURL)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal(data, &response)
	return response, err
}

// ListDisks retrieves a list of the disks in the image repository that is associated
// with the specified subscription
//
// https://msdn.microsoft.com/en-us/library/azure/jj157176.aspx
func (c DiskClient) ListDisks() (ListDiskResponse, error) {
	var response ListDiskResponse

	data, err := c.client.SendAzureGetRequest(listDisksURL)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal(data, &response)
	return response, err
}

// UpdateDataDisk updates the configuration of the specified data disk that is
// attached to the specified Virtual Machine
//
// https://msdn.microsoft.com/en-us/library/azure/jj157190.aspx
func (c DiskClient) UpdateDataDisk(
	service string,
	deployment string,
	role string,
	lun int,
	params UpdateDataDiskParameters) (management.OperationID, error) {
	if service == "" {
		return "", fmt.Errorf(errParamNotSpecified, "service")
	}
	if deployment == "" {
		return "", fmt.Errorf(errParamNotSpecified, "deployment")
	}
	if role == "" {
		return "", fmt.Errorf(errParamNotSpecified, "role")
	}

	requestURL := fmt.Sprintf(updateDataDiskURL, service, deployment, role, lun)

	req, err := xml.Marshal(params)
	if err != nil {
		return "", err
	}

	return c.client.SendAzurePutRequest(requestURL, "", req)
}

// UpdateDisk updates the label of an existing disk in the image repository that is
// associated with the specified subscription
//
// https://msdn.microsoft.com/en-us/library/azure/jj157205.aspx
func (c DiskClient) UpdateDisk(
	name string,
	params UpdateDiskParameters) (management.OperationID, error) {
	if name == "" {
		return "", fmt.Errorf(errParamNotSpecified, "name")
	}

	requestURL := fmt.Sprintf(updateDiskURL, name)

	req, err := xml.Marshal(params)
	if err != nil {
		return "", err
	}

	return c.client.SendAzurePutRequest(requestURL, "", req)
}
