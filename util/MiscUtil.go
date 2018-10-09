package util


func ReachIndexSendTrue(i int, j int, channel chan bool) {
	if i == j {
		channel <- true
	}
}
