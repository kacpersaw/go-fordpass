package fordpass

import (
	"fmt"
	"net/http"
)

type Status struct {
	VehicleStatus VehicleStatus `json:"vehiclestatus"`
	Version       string        `json:"version"`
	Status        int           `json:"status"`
}
type LockStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type Alarm struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type PrmtAlarmEvent struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type Odometer struct {
	Value     float64 `json:"value"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}
type Fuel struct {
	FuelLevel       float64 `json:"fuelLevel"`
	DistanceToEmpty float64 `json:"distanceToEmpty"`
	Status          string  `json:"status"`
	Timestamp       string  `json:"timestamp"`
}
type Gps struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	GpsState  string `json:"gpsState"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RemoteStart struct {
	RemoteStartDuration int    `json:"remoteStartDuration"`
	RemoteStartTime     int    `json:"remoteStartTime"`
	Status              string `json:"status"`
	Timestamp           string `json:"timestamp"`
}
type RemoteStartStatus struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type BatteryHealth struct {
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}
type BatteryStatusActual struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type Battery struct {
	BatteryHealth       BatteryHealth       `json:"batteryHealth"`
	BatteryStatusActual BatteryStatusActual `json:"batteryStatusActual"`
}
type Oil struct {
	OilLife       string `json:"oilLife"`
	OilLifeActual int    `json:"oilLifeActual"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}
type TirePressure struct {
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}
type TirePressureByLocation struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type TirePressureSystemStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type DualRearWheel struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type LeftFrontTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type LeftFrontTirePressure struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RightFrontTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RightFrontTirePressure struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type OuterLeftRearTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type OuterLeftRearTirePressure struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type OuterRightRearTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type OuterRightRearTirePressure struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type InnerLeftRearTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type InnerRightRearTireStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RecommendedFrontTirePressure struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RecommendedRearTirePressure struct {
	Value     int    `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type TPMS struct {
	TirePressureByLocation       TirePressureByLocation       `json:"tirePressureByLocation"`
	TirePressureSystemStatus     TirePressureSystemStatus     `json:"tirePressureSystemStatus"`
	DualRearWheel                DualRearWheel                `json:"dualRearWheel"`
	LeftFrontTireStatus          LeftFrontTireStatus          `json:"leftFrontTireStatus"`
	LeftFrontTirePressure        LeftFrontTirePressure        `json:"leftFrontTirePressure"`
	RightFrontTireStatus         RightFrontTireStatus         `json:"rightFrontTireStatus"`
	RightFrontTirePressure       RightFrontTirePressure       `json:"rightFrontTirePressure"`
	OuterLeftRearTireStatus      OuterLeftRearTireStatus      `json:"outerLeftRearTireStatus"`
	OuterLeftRearTirePressure    OuterLeftRearTirePressure    `json:"outerLeftRearTirePressure"`
	OuterRightRearTireStatus     OuterRightRearTireStatus     `json:"outerRightRearTireStatus"`
	OuterRightRearTirePressure   OuterRightRearTirePressure   `json:"outerRightRearTirePressure"`
	InnerLeftRearTireStatus      InnerLeftRearTireStatus      `json:"innerLeftRearTireStatus"`
	InnerLeftRearTirePressure    interface{}                  `json:"innerLeftRearTirePressure"`
	InnerRightRearTireStatus     InnerRightRearTireStatus     `json:"innerRightRearTireStatus"`
	InnerRightRearTirePressure   interface{}                  `json:"innerRightRearTirePressure"`
	RecommendedFrontTirePressure RecommendedFrontTirePressure `json:"recommendedFrontTirePressure"`
	RecommendedRearTirePressure  RecommendedRearTirePressure  `json:"recommendedRearTirePressure"`
}
type FirmwareUpgInProgress struct {
	Value     bool   `json:"value"`
	Timestamp string `json:"timestamp"`
}
type DeepSleepInProgress struct {
	Value     bool   `json:"value"`
	Timestamp string `json:"timestamp"`
}
type CcsSettings struct {
	Timestamp              string `json:"timestamp"`
	Location               int    `json:"location"`
	VehicleConnectivity    int    `json:"vehicleConnectivity"`
	VehicleData            int    `json:"vehicleData"`
	DrivingCharacteristics int    `json:"drivingCharacteristics"`
	Contacts               int    `json:"contacts"`
}
type OutandAbout struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type DriverWindowPosition struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type PassWindowPosition struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RearDriverWindowPos struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type RearPassWindowPos struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type WindowPosition struct {
	DriverWindowPosition DriverWindowPosition `json:"driverWindowPosition"`
	PassWindowPosition   PassWindowPosition   `json:"passWindowPosition"`
	RearDriverWindowPos  RearDriverWindowPos  `json:"rearDriverWindowPos"`
	RearPassWindowPos    RearPassWindowPos    `json:"rearPassWindowPos"`
}
type RightRearDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type LeftRearDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type DriverDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type PassengerDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type HoodDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type TailgateDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type InnerTailgateDoor struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type DoorStatus struct {
	RightRearDoor     RightRearDoor     `json:"rightRearDoor"`
	LeftRearDoor      LeftRearDoor      `json:"leftRearDoor"`
	DriverDoor        DriverDoor        `json:"driverDoor"`
	PassengerDoor     PassengerDoor     `json:"passengerDoor"`
	HoodDoor          HoodDoor          `json:"hoodDoor"`
	TailgateDoor      TailgateDoor      `json:"tailgateDoor"`
	InnerTailgateDoor InnerTailgateDoor `json:"innerTailgateDoor"`
}
type IgnitionStatus struct {
	Value     string `json:"value"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type DieselSystemStatus struct {
	ExhaustFluidLevel        interface{} `json:"exhaustFluidLevel"`
	FilterSoot               interface{} `json:"filterSoot"`
	UreaRange                interface{} `json:"ureaRange"`
	MetricType               interface{} `json:"metricType"`
	FilterRegenerationStatus interface{} `json:"filterRegenerationStatus"`
}
type VehicleStatus struct {
	Vin                           string                `json:"vin"`
	LockStatus                    LockStatus            `json:"lockStatus"`
	Alarm                         Alarm                 `json:"alarm"`
	PrmtAlarmEvent                PrmtAlarmEvent        `json:"PrmtAlarmEvent"`
	Odometer                      Odometer              `json:"odometer"`
	Fuel                          Fuel                  `json:"fuel"`
	Gps                           Gps                   `json:"gps"`
	RemoteStart                   RemoteStart           `json:"remoteStart"`
	RemoteStartStatus             RemoteStartStatus     `json:"remoteStartStatus"`
	Battery                       Battery               `json:"battery"`
	Oil                           Oil                   `json:"oil"`
	TirePressure                  TirePressure          `json:"tirePressure"`
	Authorization                 string                `json:"authorization"`
	TPMS                          TPMS                  `json:"TPMS"`
	FirmwareUpgInProgress         FirmwareUpgInProgress `json:"firmwareUpgInProgress"`
	DeepSleepInProgress           DeepSleepInProgress   `json:"deepSleepInProgress"`
	CcsSettings                   CcsSettings           `json:"ccsSettings"`
	LastRefresh                   string                `json:"lastRefresh"`
	LastModifiedDate              string                `json:"lastModifiedDate"`
	ServerTime                    string                `json:"serverTime"`
	BatteryFillLevel              interface{}           `json:"batteryFillLevel"`
	ElVehDTE                      interface{}           `json:"elVehDTE"`
	HybridModeStatus              interface{}           `json:"hybridModeStatus"`
	ChargingStatus                interface{}           `json:"chargingStatus"`
	PlugStatus                    interface{}           `json:"plugStatus"`
	ChargeStartTime               interface{}           `json:"chargeStartTime"`
	ChargeEndTime                 interface{}           `json:"chargeEndTime"`
	PreCondStatusDsply            interface{}           `json:"preCondStatusDsply"`
	ChargerPowertype              interface{}           `json:"chargerPowertype"`
	BatteryPerfStatus             interface{}           `json:"batteryPerfStatus"`
	OutandAbout                   OutandAbout           `json:"outandAbout"`
	BatteryChargeStatus           interface{}           `json:"batteryChargeStatus"`
	DcFastChargeData              interface{}           `json:"dcFastChargeData"`
	WindowPosition                WindowPosition        `json:"windowPosition"`
	DoorStatus                    DoorStatus            `json:"doorStatus"`
	IgnitionStatus                IgnitionStatus        `json:"ignitionStatus"`
	BatteryTracLowChargeThreshold interface{}           `json:"batteryTracLowChargeThreshold"`
	BattTracLoSocDDsply           interface{}           `json:"battTracLoSocDDsply"`
	DieselSystemStatus            DieselSystemStatus    `json:"dieselSystemStatus"`
}

func (c *Client) Status() (*VehicleStatus, error) {
	if err := c.acquireToken(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vehicles/v4/%s/status", defaultBaseURL, c.Vin), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lrdt", "01-01-1970 00:00:00")

	var status Status
	if err := c.doRequest(req, &status); err != nil {
		return nil, err
	}

	return &status.VehicleStatus, nil
}
