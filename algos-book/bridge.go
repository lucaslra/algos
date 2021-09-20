package algosbook

//
//import "fmt"
//
//type IEntertainmentDevice interface {
//	buttonFivePressed()
//	buttonSixPressed()
//}
//
//type EntertainmentDevice struct {
//	IEntertainmentDevice
//	deviceState int
//	maxSetting  int
//	volumeLevel int
//}
//
//func (ed *EntertainmentDevice) buttonSevenPressed() {
//	ed.volumeLevel++
//
//	fmt.Printf("Volume raised to %v\n", ed.volumeLevel)
//}
//func (ed *EntertainmentDevice) buttonEightPressed() {
//	ed.volumeLevel--
//
//	fmt.Printf("Volume decreased to %v\n", ed.volumeLevel)
//}
//func (ed *EntertainmentDevice) deviceFeedback() {
//	if ed.deviceState > ed.maxSetting || ed.deviceState < 0 {
//		ed.deviceState = 0
//	}
//
//	fmt.Printf("On %v\n", ed.deviceState)
//}
//
//type TVDevice struct {
//	*EntertainmentDevice
//}
//
//func (td *TVDevice) buttonFivePressed() {
//	fmt.Println("Channel down")
//
//	td.deviceState--
//}
//func (td *TVDevice) buttonSixPressed() {
//	fmt.Println("Channel up")
//
//	td.deviceState++
//}
//
//type RemoteButton struct {
//	*Ente
//}
//
////type RemoteDevice struct {
////	*EntertainmentDevice
////}
////
////func (rd *RemoteDevice) buttonFivePressed(){
////	fmt.Println()
////}
////func (rd *RemoteDevice) buttonSixPressed(){
////	fmt.Println()
////}
//
//func RunBridgeDemo() {
//	var tv EntertainmentDevice = TVDevice{&EntertainmentDevice{deviceState: 0, maxSetting: 10}}
//	tv.buttonSevenPressed()
//	tv.buttonEightPressed()
//
//	fmt.Println(tv.volumeLevel)
//
//	tv.buttonSixPressed()
//	tv.buttonFivePressed()
//}
