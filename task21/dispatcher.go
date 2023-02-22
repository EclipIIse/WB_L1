package main

type Dispatcher struct {
}

func (d *Dispatcher) doSomething(flight FlightDataService) {
	flight.SendFlightData()
}
