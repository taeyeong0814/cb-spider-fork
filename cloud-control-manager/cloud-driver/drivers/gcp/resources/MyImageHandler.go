package resources

import (
	"context"
	"strings"
	"time"

	idrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"
	compute "google.golang.org/api/compute/v1"
)

type GCPMyImageHandler struct {
	Region     idrv.RegionInfo
	Ctx        context.Context
	Client     *compute.Service
	Credential idrv.CredentialInfo
}

const (
	GCPMyImageReady   string = "READY"
	GCPMyImageInvalid string = "INVALID"
)

func (MyImageHandler *GCPMyImageHandler) SnapshotVM(snapshotReqInfo irs.MyImageInfo) (irs.MyImageInfo, error) {
	projectID := MyImageHandler.Credential.ProjectID
	myImageName := snapshotReqInfo.IId.NameId

	machineImage := &compute.MachineImage{
		SourceInstance: snapshotReqInfo.SourceVM.SystemId,
	}

	op, err := MyImageHandler.Client.MachineImages.Insert(projectID, machineImage).Do()
	if err != nil {
		cblogger.Error(err)
		return irs.MyImageInfo{}, err
	}

	WaitUntilComplete(MyImageHandler.Client, projectID, "", op.Name, true)

	myImageInfo, errMyImage := MyImageHandler.GetMyImage(irs.IID{NameId: myImageName, SystemId: myImageName})
	if errMyImage != nil {
		cblogger.Error(errMyImage)
		return irs.MyImageInfo{}, errMyImage
	}

	return myImageInfo, nil
}

func (MyImageHandler *GCPMyImageHandler) ListMyImage() ([]*irs.MyImageInfo, error) {
	myImageInfoList := []*irs.MyImageInfo{}

	projectID := MyImageHandler.Credential.ProjectID

	myImageList, err := MyImageHandler.Client.MachineImages.List(projectID).Do()
	if err != nil {
		cblogger.Error(err)
		return nil, err
	}

	for _, myImage := range myImageList.Items {
		myImageInfo, err := convertMyImageInfo(myImage)
		if err != nil {
			cblogger.Error(err)
			return nil, err
		}
		myImageInfoList = append(myImageInfoList, &myImageInfo)
	}

	return myImageInfoList, nil
}

func (MyImageHandler *GCPMyImageHandler) GetMyImage(myImageIID irs.IID) (irs.MyImageInfo, error) {
	projectID := MyImageHandler.Credential.ProjectID

	myImageResp, err := MyImageHandler.Client.MachineImages.Get(projectID, myImageIID.SystemId).Do()
	if err != nil {
		cblogger.Error(err)
		return irs.MyImageInfo{}, err
	}

	myImageInfo, errMyImage := convertMyImageInfo(myImageResp)
	if errMyImage != nil {
		cblogger.Error(errMyImage)
		return irs.MyImageInfo{}, errMyImage
	}

	return myImageInfo, nil
}

func (MyImageHandler *GCPMyImageHandler) DeleteMyImage(myImageIID irs.IID) (bool, error) {
	projectID := MyImageHandler.Credential.ProjectID
	myImage := myImageIID.SystemId

	op, err := MyImageHandler.Client.MachineImages.Delete(projectID, myImage).Do()
	if err != nil {
		cblogger.Error(err)
		return false, err
	}

	WaitUntilComplete(MyImageHandler.Client, projectID, "", op.Name, true)

	return true, nil
}

func convertMyImageInfo(myImageResp *compute.MachineImage) (irs.MyImageInfo, error) {
	myImageInfo := irs.MyImageInfo{}

	myImageInfo.IId = irs.IID{NameId: myImageResp.Name, SystemId: myImageResp.Name}
	arrSourceVM := strings.Split(myImageResp.SourceInstance, "/")
	sourceVM := arrSourceVM[len(arrSourceVM)-1]
	myImageInfo.SourceVM = irs.IID{SystemId: sourceVM}

	myImageStatus, err := convertMyImageStatus(myImageResp.Status)
	if err != nil {
		cblogger.Error(err)
		return irs.MyImageInfo{}, err
	}

	myImageInfo.Status = myImageStatus

	myImageInfo.CreatedTime, _ = time.Parse(time.RFC3339, myImageResp.CreationTimestamp)

	return myImageInfo, nil
}

func convertMyImageStatus(status string) (irs.MyImageStatus, error) {
	var returnStatus irs.MyImageStatus

	if status == GCPMyImageReady {
		returnStatus = irs.MyImageAvailable
	} else if status == GCPMyImageInvalid {
		returnStatus = irs.MyImageUnavailable
	}
	return returnStatus, nil

}
