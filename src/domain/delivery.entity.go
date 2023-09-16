package entity

type DeliveryEntity struct {
	ID         string
	Country    string
	State      string
	City       string
	Address    string
	NumberHome string
	Reference  string
	/*
		estado => RJ
		cidade => BM
		endereco => input
		bairro => input
		numero => input
		ponto de referencia => input
	*/
}
