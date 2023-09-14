package main

import (
	"fmt"
)

func main() {
	var bldr = newNotificationBuilder()

	bldr.SetTitle("New Notification")
	bldr.SetSubTitle("New SubTitle")
	bldr.SetMessage("This is a basic notification")
	bldr.SetImage("image.png")
	bldr.SetIcon("icon.png")
	bldr.SetPriority(1)
	bldr.SetNotType("alert")

	notif, err := bldr.Build()

	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Title: %+v\n", notif.title)
		fmt.Printf("Sub Title: %+v\n", notif.subtitle)
		fmt.Printf("Message: %+v\n", notif.message)
		fmt.Printf("Image: %+v\n", notif.image)
		fmt.Printf("Icon: %+v\n", notif.icon)
		fmt.Printf("Priority: %+v\n", notif.priority)
		fmt.Printf("NotType: %+v\n", notif.notType)
	}
}
