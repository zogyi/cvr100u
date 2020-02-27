package main

import (
	"fmt"

	"github.com/zogyi/cvr100u/device"
)

func main() {
	conn := device.Connector{IsX64: true}
	if success := conn.Initial(); success {
		if success = conn.Authentication(); success {
			if success = conn.ReadContent(); success {
				if result, _, err := conn.ReadFields(device.ReadName); err == nil {
					fmt.Println(`the card's name is %s`, result)
				} else {
					println(`can't read the name`)
				}
			} else {
				println(`can't read the content`)
			}
		} else {
			println(`can't Authentication, reput the ID card`)
		}
	} else {
		println(`device doesn't work properly`)
	}
}
