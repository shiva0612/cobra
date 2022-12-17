
shoes_in_stock: build
	ecom items shoes --instock=true --number=10 nike puma
shirts_in_stock: build
	ecom items shirts --instock=true --size=l nike adidas
items: build
	ecom items --instock=false nike

build:
	go install

 