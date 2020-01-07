
Before running the app
Setup product and category API in files/config.ini
run command by docker   `docker build -t cart . && docker run -p 9090:9090 cart`


The service is written using the clean architecture by Uncle Bob, where each layer can have its own test cases
Sample Test case of mid layer(usecase) written in src/usecase/cart/cart_test.go
