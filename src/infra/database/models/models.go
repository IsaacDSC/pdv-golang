package models

import "gorm.io/gorm"

type Enterprise struct {
	gorm.Model
	ID          string `gorm:"type:varchar(255),primaryKey;default:uuid()"`
	SubDomain   string `gorm:"unique"`
	CompanyName string `gorm:"unique"`
	CNPJ        string `gorm:"unique"`
	Email       string
	Telephone   string
	CreatedAt   string
	UpdatedAt   string
	Clients     []Client `gorm:"many2many:enterprise_client;"`
	Orders      []Order  `gorm:"many2many:enterprise_order;"`
}

type Client struct {
	gorm.Model
	ID           string `gorm:"type:varchar(255),primaryKey;default:uuid()"`
	CPF          string `gorm:"type:varchar(80)"`
	Name         string `gorm:"type:varchar(100)"`
	Email        string `gorm:"unique,type:varchar(120)"`
	Username     string `gorm:"unique,type:varchar(80)"`
	Password     string `gorm:"type:varchar(255)"`
	Telephone    string `gorm:"type:varchar(20)"`
	RecoverToken string `gorm:"type:varchar(255)"`
}

type Delivery struct {
	gorm.Model
	ID         string `gorm:"type:varchar(255),primaryKey;default:uuid()"`
	Country    string `gorm:"type:varchar(30)"`
	State      string `gorm:"type:varchar(30)"`
	Address    string `gorm:"type:varchar(255)"`
	NumberHome int32
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

type Order struct {
	gorm.Model
	ID            string       `gorm:"type:varchar(255),primaryKey;default:uuid()"`
	OrderItems    []OrderItems `gorm:"many2many:order_orderitems;"`
	TypePayment   string       `gorm:"type:varchar(100)"`
	Price         float32
	Discount      float32
	PriceDelivery float32
	TaxDelivery   bool
	Delivery      Delivery
	Client        Client
}

type OrderItems struct {
	gorm.Model
	ID        string `gorm:"primaryKey;default:uuid()"`
	ProductID string `gorm:"type:varchar(255)"`
	Quantity  int32
	Price     float32
}

/*
	SEMANAL
	PIX - DEBITO
	GERAR RELATORIO DE LUCROS OUTRO DATABASE E OUTRA APLICACAO DE FORMA MANUAL
	SALVAR RELATORIO EM UM LUGAR DA EMPRESA MENSAL

	MENSAL
	LER RELATÓRIO MENSAL DA EMPRESA
	CREDITO - IF PIX && DEBITO NOT TRANSFER THEN CALCULATE
	GERAR RELATORIO DE LUCROS OUTRO DATABASE E OUTRA APLICACAO DE FORMA MANUAL
	SALVAR RELATORIO EM UM LUGAR DA EMPRESA MENSAL
	SALVAR FECHAMENTO DO CAIXA MENSAL DA EMPRESA (REPASSE_PLATFORM, RECEBIMENTO, REPASSE_GATWAY)


	VALIDACAO DO CALCULO
	TOTAL MENSAL = BUSCAR TODOS DEPOSITOS MENSAIS E SOMAR
	COMPARAR COM O RELATORIO FINAL DO FECHAMENTO DO MES


*/
