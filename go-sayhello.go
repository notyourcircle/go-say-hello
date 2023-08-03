package gosayhello

type Customer struct {
	Name    string
	Age     int
	Address string
}

func SayHello(customer Customer) string {
	return "Hello " + customer.Name + " " + customer.Address + "Sedang cerah ya"
}
