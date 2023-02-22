/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

func main() {
	dispatcher := &Dispatcher{}
	kilometers := &KilometersData{}
	dispatcher.doSomething(kilometers)
	miles := &MilesFlightData{}
	kilometersAdapter := &FlightDataAdapter{flightData: miles}
	dispatcher.doSomething(kilometersAdapter)
}
