**About**

Cart is a service written in Golang, it fetches product list from an API and show it in frontend, the user can currently login with google only, the user can add products to cart, update, delete those products  
The service uses Golang **clean architecture**  
Sample Test case of mid layer(usecase) written in src/usecase/cart/cart_test.go


   
**Running the app**  
Setup product and category API in files/config.ini  
Docker Command   `docker build -t cart . && docker run -p 9090:9090 cart`  


Config File - files/config.ini  
Db file - db.sql  
API - src/interfaces/web/api/api.go