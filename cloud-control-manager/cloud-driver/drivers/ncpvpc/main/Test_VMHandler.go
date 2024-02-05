// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Tester Example.
//
// by ETRI, 2020.12.

package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	idrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"
	cblog "github.com/cloud-barista/cb-log"

	// ncpvpcdrv "github.com/cloud-barista/ncpvpc/ncpvpc"  // For local test
	ncpvpcdrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/ncpvpc"
)

var cblogger *logrus.Logger

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("NCP VPC Resource Test")
	cblog.SetLevel("info")
}

func testErr() error {

	return errors.New("")
	// return ncloud.New("504", "찾을 수 없음", nil)
}

// Test VM Lifecycle Management (Create/Suspend/Resume/Reboot/Terminate)
func handleVM() {
	cblogger.Debug("Start VMHandler Resource Test")

	ResourceHandler, err := getResourceHandler("VM")
	if err != nil {
		panic(err)
	}

	vmHandler := ResourceHandler.(irs.VMHandler)

	for {
		fmt.Println("\n============================================================================================")
		fmt.Println("[ VM Management Test ]")
		fmt.Println("1. Start(Create) VM")
		fmt.Println("2. Get VM Info")
		fmt.Println("3. Suspend VM")
		fmt.Println("4. Resume VM")
		fmt.Println("5. Reboot VM")

		fmt.Println("6. Terminate VM")
		fmt.Println("7. Get VMStatus")
		fmt.Println("8. List VMStatus")
		fmt.Println("9. List VM")
		fmt.Println("0. Exit")
		fmt.Println("\n   Select a number above!! : ")
		fmt.Println("============================================================================================")

		//config := readConfigFile()
		VmID := irs.IID{SystemId: "22356052"}

		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)

		if err != nil {
			panic(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 0:
				return

			case 1:
				vmReqInfo := irs.VMReqInfo{
					// # NCP에서는 VM instance 이름에 대문자 허용 안되므로, VMHandler 내부에서 소문자로 변환되어 반영됨.
					IId: irs.IID{NameId: "My-VM-01"},
					
					ImageType: "PublicImage", // "", "default", "PublicImage" or "MyImage"
					// ImageType: "MyImage", // "", "default", "PublicImage" or "MyImage"

					// # For Public Image Test
					// # (Note) NCP VPC infra service와 Classic 2세대 service는 ImageID, VMSpecName 체계가 다름.
					ImageIID:   irs.IID{NameId: "ubuntu-18.04", SystemId: "SW.VSVR.OS.LNX64.UBNTU.SVR1804.B050"},
					// ImageIID:   irs.IID{NameId: "Windows-Server-2016(64bit)", SystemId: "SW.VSVR.OS.WND64.WND.SVR2016EN.B100"},
					// ImageIID:   irs.IID{NameId: "CentOS 7.8 (64-bit)", SystemId: "SW.VSVR.OS.LNX64.CNTOS.0708.B050"}, 
					// ImageIID:   irs.IID{NameId: "Rocky Linux 8.6", SystemId: "SW.VSVR.OS.LNX64.ROCKY.0806.B050"}, 

					// # For MyImage Test
					// ImageIID:   irs.IID{NameId: "ubuntu-18.04", SystemId: "18306970"}, 
					// ImageIID:   irs.IID{NameId: "ubuntu-18.04", SystemId: "13233382"},  // In case of "MyImage"
					// ImageIID:   irs.IID{NameId: "ncpvpc-winimage-02", SystemId: "14917995"}, // In case of "MyImage"

					// ### Caution!! ### : NCP Classic 2세대 infra가 아닌 NCP VPC infra service에서는 VPC, subnet 지정 필수!!
					VpcIID:    irs.IID{SystemId: "41565"}, // ncp-vpc-01
					SubnetIID: irs.IID{SystemId: "92493"}, // ncp-subnet-01
					// VpcIID:    irs.IID{SystemId: "1363"},
					// SubnetIID: irs.IID{SystemId: "3325"},

					// VMSpecName: "SVR.VSVR.HICPU.C004.M008.NET.SSD.B050.G002", // For Image : "SW.VSVR.OS.LNX64.UBNTU.SVR1804.B050"
					// VMSpecName: "SVR.VSVR.HICPU.C002.M004.NET.SSD.B100.G002", // For Image : "SW.VSVR.OS.WND64.WND.SVR2016EN.B100"
					// VMSpecName: "SVR.VSVR.HICPU.C004.M008.NET.SSD.B050.G002", // For Image : "SW.VSVR.OS.LNX64.CNTOS.0708.B050"
					VMSpecName: "SVR.VSVR.HICPU.C004.M008.NET.SSD.B050.G002", // For Image : "SW.VSVR.OS.LNX64.ROCKY.0806.B050"					

					KeyPairIID: irs.IID{SystemId: "NCP-keypair-05"}, // Caution : Not NameId!!
					// KeyPairIID: irs.IID{SystemId: "ns01-ncpv-cij7lb1jcupork04fnr0"}, // Caution : Not NameId!!

					// ### Caution!! ### : AccessControlGroup은 NCP console상의 VPC 메뉴의 'Network ACL'이 아닌 Server 메뉴의 'ACG'에 해당됨.
					// SecurityGroupIIDs: []irs.IID{{SystemId: "44518"}}, // ncp-sg-02
					SecurityGroupIIDs: []irs.IID{{SystemId: "114954"}}, // ncp-vpc-01-default-acg
					// SecurityGroupIIDs: []irs.IID{{SystemId: "3486"}},

					VMUserPasswd: "cdcdcd353535**", 
				}

				vmInfo, err := vmHandler.StartVM(vmReqInfo)
				if err != nil {
					//panic(err)
					cblogger.Error(err)
				} else {
					cblogger.Info("VM 생성 완료!!", vmInfo)
					spew.Dump(vmInfo)
				}
				//cblogger.Info(vm)

				cblogger.Info("\nCreateVM Test Finished")

			case 2:
				vmInfo, err := vmHandler.GetVM(VmID)
				if err != nil {
					cblogger.Errorf("Failed Get the VM info.: [%s]", VmID)
					cblogger.Error(err)
				} else {
					cblogger.Infof("VM info. of [%s]", VmID)
					cblogger.Info(vmInfo)
					spew.Dump(vmInfo)
				}

				cblogger.Info("\nGetVM Test Finished")

			case 3:
				cblogger.Info("Start Suspend VM ...")
				result, err := vmHandler.SuspendVM(VmID)
				if err != nil {
					cblogger.Errorf("[%s] VM Suspend 실패 - [%s]", VmID, result)
					cblogger.Error(err)
				} else {
					cblogger.Infof("[%s] VM Suspend 실행 성공 - [%s]", VmID, result)
				}

				cblogger.Info("\nSuspendVM Test Finished")

			case 4:
				cblogger.Info("Start Resume  VM ...")
				result, err := vmHandler.ResumeVM(VmID)
				if err != nil {
					cblogger.Errorf("[%s] VM Resume 실패 - [%s]", VmID, result)
					cblogger.Error(err)
				} else {
					cblogger.Infof("[%s] VM Resume 실행 성공 - [%s]", VmID, result)
				}

				cblogger.Info("\nResumeVM Test Finished")

			case 5:
				cblogger.Info("Start Reboot  VM ...")
				result, err := vmHandler.RebootVM(VmID)
				if err != nil {
					cblogger.Errorf("[%s] VM Reboot 실패 - [%s]", VmID, result)
					cblogger.Error(err)
				} else {
					cblogger.Infof("[%s] VM Reboot 실행 성공 - [%s]", VmID, result)
				}

				cblogger.Info("\nRebootVM Test Finished")

			case 6:
				cblogger.Info("Start Terminate  VM ...")
				result, err := vmHandler.TerminateVM(VmID)
				if err != nil {
					cblogger.Errorf("[%s] VM Terminate 실패 - [%s]", VmID, result)
					cblogger.Error(err)
				} else {
					cblogger.Infof("[%s] VM Terminate 실행 성공 - [%s]", VmID, result)
				}

				cblogger.Info("\nTerminateVM Test Finished")

			case 7:
				cblogger.Info("Start Get VM Status...")
				vmStatus, err := vmHandler.GetVMStatus(VmID)
				if err != nil {
					cblogger.Errorf("[%s] Get VM Status 실패", VmID)
					cblogger.Error(err)
				} else {
					cblogger.Infof("[%s] Get VM Status 실행 성공 : [%s]", VmID, vmStatus)
				}

				cblogger.Info("\nGet VMStatus Test Finished")

			case 8:
				cblogger.Info("Start ListVMStatus ...")
				vmStatusInfos, err := vmHandler.ListVMStatus()
				if err != nil {
					cblogger.Error("ListVMStatus 실패")
					cblogger.Error(err)
				} else {
					cblogger.Info("ListVMStatus 실행 성공")
					//cblogger.Info(vmStatusInfos)
					spew.Dump(vmStatusInfos)
				}

				cblogger.Info("\nListVM Status Test Finished")

			case 9:
				cblogger.Info("Start ListVM ...")
				vmList, err := vmHandler.ListVM()
				if err != nil {
					cblogger.Error("ListVM 실패")
					cblogger.Error(err)
				} else {
					cblogger.Info("ListVM 실행 성공")
					cblogger.Info("=========== VM 목록 ================")
					// cblogger.Info(vmList)
					spew.Dump(vmList)
					cblogger.Infof("=========== VM 목록 수 : [%d] ================", len(vmList))
					if len(vmList) > 0 {
						VmID = vmList[0].IId
					}
				}

				cblogger.Info("\nListVM Test Finished")

			}
		}
	}
}

func main() {
	cblogger.Info("NCP VPC Resource Test")

	handleVM()
}

// handlerType : resources폴더의 xxxHandler.go에서 Handler이전까지의 문자열
// (예) ImageHandler.go -> "Image"
func getResourceHandler(handlerType string) (interface{}, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(ncpvpcdrv.NcpVpcDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			ClientId:     config.Ncp.NcpAccessKeyID,
			ClientSecret: config.Ncp.NcpSecretKey,
		},
		RegionInfo: idrv.RegionInfo{
			Region: config.Ncp.Region,
			Zone:   config.Ncp.Zone,
		},
	}

	// NOTE Just for test
	//cblogger.Info(config.Ncp.NcpAccessKeyID)
	//cblogger.Info(config.Ncp.NcpSecretKey)

	cblogger.Info(connectionInfo.RegionInfo.Zone)

	cloudConnection, errCon := cloudDriver.ConnectCloud(connectionInfo)
	if errCon != nil {
		return nil, errCon
	}

	var resourceHandler interface{}
	var err error

	switch handlerType {
	case "Image":
		resourceHandler, err = cloudConnection.CreateImageHandler()
	case "Security":
		resourceHandler, err = cloudConnection.CreateSecurityHandler()
	case "VNetwork":
		resourceHandler, err = cloudConnection.CreateVPCHandler()
	case "VM":
		resourceHandler, err = cloudConnection.CreateVMHandler()
	case "VMSpec":
		resourceHandler, err = cloudConnection.CreateVMSpecHandler()
	}

	if err != nil {
		return nil, err
	}
	return resourceHandler, nil
}

// Region : 사용할 리전명 (ex) ap-northeast-2
// ImageID : VM 생성에 사용할 AMI ID (ex) ami-047f7b46bd6dd5d84
// BaseName : 다중 VM 생성 시 사용할 Prefix이름 ("BaseName" + "_" + "숫자" 형식으로 VM을 생성 함.) (ex) mcloud-barista
// VmID : 라이프 사이트클을 테스트할 EC2 인스턴스ID
// InstanceType : VM 생성시 사용할 인스턴스 타입 (ex) t2.micro
// KeyName : VM 생성시 사용할 키페어 이름 (ex) mcloud-barista-keypair
// MinCount :
// MaxCount :
// SubnetId : VM이 생성될 VPC의 SubnetId (ex) subnet-cf9ccf83
// SecurityGroupID : 생성할 VM에 적용할 보안그룹 ID (ex) sg-0df1c209ea1915e4b
type Config struct {
	Ncp struct {
		NcpAccessKeyID string `yaml:"ncp_access_key_id"`
		NcpSecretKey   string `yaml:"ncp_secret_key"`
		Region         string `yaml:"region"`
		Zone           string `yaml:"zone"`

		ImageID string `yaml:"image_id"`

		VmID         string `yaml:"ncp_instance_id"`
		BaseName     string `yaml:"base_name"`
		InstanceType string `yaml:"instance_type"`
		KeyName      string `yaml:"key_name"`
		MinCount     int64  `yaml:"min_count"`
		MaxCount     int64  `yaml:"max_count"`

		SubnetID        string `yaml:"subnet_id"`
		SecurityGroupID string `yaml:"security_group_id"`

		PublicIP string `yaml:"public_ip"`
	} `yaml:"ncpvpc"`
}

func readConfigFile() Config {
	// # Set Environment Value of Project Root Path
	// goPath := os.Getenv("GOPATH")
	// rootPath := goPath + "/src/github.com/cloud-barista/ncp/ncp/main"
	// cblogger.Debugf("Test Config file : [%]", rootPath+"/config/config.yaml")
	rootPath 	:= os.Getenv("CBSPIDER_ROOT")
	configPath 	:= rootPath + "/cloud-control-manager/cloud-driver/drivers/ncpvpc/main/config/config.yaml"
	cblogger.Debugf("Test Config file : [%s]", configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	cblogger.Info("ConfigFile Loaded ...")

	// Just for test
	cblogger.Debug(config.Ncp.NcpAccessKeyID, " ", config.Ncp.Region)

	return config
}
