package discoveryGo

import "fmt"

func ExampleAddress() {
	var c Contact
	c.Mobile = "010-0001-0002"
	fmt.Println(c.Telephone.Mobile)
	c.Address.City = "San Francisco"
	c.State = "CA"
	c.Direct = "N/A"
	fmt.Println(c)

	// output:
	// 010-0001-0002
	// {{San Francisco CA} {010-0001-0002 N/A}}
}
