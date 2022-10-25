# rest-shoppingcart
Dokumentasi API menggunakan swagger dapat dilihat pada: http://localhost:3000/swagger/index.html

Definisi schema database:
1. User
2. Product
3. Cart
4. Transaksi

Associations:
1. Satu User memiliki satu (Has One) Cart. Cart kosong akan secara otomatis dibuatkan saat melakukan registrasi User
2. Carts memiliki banyak (Many To Many dengan) Products. Products dapat dimiliki oleh banyak Carts
3. Satu User memiliki banyak (Has Many) Transaksi.
4. Transaksis memiliki banyak (Many To Many dengan) Products. Products dapat dimiliki oleh banyak Transaksis

Connect to SQL Server
	dsn := "sqlserver://michaelmaulana:pass123@localhost:1433?database=GolangDB"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
