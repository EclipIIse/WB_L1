package main

type FlightDataAdapter struct {
	flightData *MilesFlightData
}

func (adapter *FlightDataAdapter) SendFlightData() {
	adapter.flightData.ConvertToKilometers()
	println("Flight distance converted from miles into kilometers")
}
