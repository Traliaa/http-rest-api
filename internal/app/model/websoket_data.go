package model

type Websoket_data struct {
	Id_device   int
	Device_name string
	Divice_info string
}

type SmartDevice struct {
	IDevice int
	IdIot   int
	Name    string
	Status  bool
}

type AuthDevice struct {
	Username  string
	Email     string
	Processor string
}
