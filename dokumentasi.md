# Aplikasi Penjualan Makanan Sisa

Aplikasi ini dirancang untuk membantu penjual menawarkan makanan sisa berkualitas dengan harga terjangkau, menghubungkan penjual dan pembeli secara efisien. Aplikasi ini dibangun menggunakan **Golang** dengan arsitektur RESTful API.

---

## 1. Entity Relationship Diagram (ERD)

Berikut adalah ERD untuk aplikasi penjualan makanan sisa:

## 2. High-Level Architecture (HLA)

![HLA Diagram](path_to_hla_image.png)

Komponen utama dari arsitektur aplikasi ini meliputi:

- **Backend API**: Server RESTful yang ditulis dalam Golang untuk menangani proses bisnis dan pengelolaan data.
- **Database**: Penyimpanan data terstruktur (MySQL atau PostgreSQL).
- **External Services** seperti API Pembayaran, API Lokasi, dan AI untuk rekomendasi produk.

---

## 3. Daftar Endpoint API dan Dokumentasi

### a. **Auth Endpoints**

1. **Registrasi Pengguna**

   - **Endpoint**: `POST /api/v1/auth/register`
   - **Body**:
     ```json
     {
       "name": "string",
       "email": "string",
       "password": "string",
       "role": "buyer/seller",
       "address": "string",
       "phone": "string"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "User registered successfully",
       "user_id": "1"
     }
     ```

2. **Login Pengguna**
   - **Endpoint**: `POST /api/v1/auth/login`
   - **Body**:
     ```json
     {
       "email": "johndoe@example.com",
       "password": "password123"
     }
     ```
   - **Response**:
     ```json
     {
       "token": "jwt_token_here"
     }
     ```

### b. **User Endpoints**

1. **Get Profile**
   - **Endpoint**: `GET /api/v1/users/{id}`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Response**:
     ```json
     {
       "id": "1",
       "name": "John Doe",
       "email": "johndoe@example.com",
       "role": "buyer",
       "address": "123 Street",
       "phone": "1234567890"
     }
     ```
2. **Update Profile**

   - **Endpoint**: `PUT /api/v1/users/{id}`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Body**:
     ```json
     {
       "name": "string",
       "email": "string",
       "address": "string",
       "phone": "string"
     }
     ```
   - **Response**:
     ```json
     {
       "id": "1",
       "name": "John Doe",
       "email": "johndoe@example.com",
       "role": "buyer",
       "address": "123 Street",
       "phone": "1234567890"
     }
     ```

3. **Delete User**
   - **Endpoint**: `DELETE /api/v1/users/{id}`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Response**:
     ```json
      204 No Content
     ```

### c. **Store Endpoints**

1. **Tambah Toko**
   - **Endpoint**: `POST /api/v1/stores`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Body**:
     ```json
     {
       "name": "Starbucks",
       "address": "123 Coffee St",
       "phone": "1234567890"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Store created successfully",
       "store_id": "1"
     }
     ```

### d. **Product Endpoints**

1. **Tambah Produk ke Toko Tertentu**
   - **Endpoint**: `POST /api/v1/stores/:storeId/products`
   - **Parameter**:
     - `storeId` (path parameter): ID toko tempat produk akan ditambahkan.
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Body**:
     ```json
     {
       "name": "Sandwich",
       "description": "Chicken sandwich with lettuce",
       "price": 50.0,
       "stock": 10,
       "expiry_date": "2024-12-31",
       "location": "Jakarta"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Product created successfully",
       "product_id": "1"
     }
     ```

### e. **Cart Endpoints**

1. **Tambah ke Keranjang**
   - **Endpoint**: `POST /api/v1/cart`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Body**:
     ```json
     {
       "product_id": 1,
       "quantity": 2
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Product added to cart"
     }
     ```

### f. **Order Endpoints**

1. **Checkout Keranjang**
   - **Endpoint**: `POST /api/v1/orders`
   - **Header**: `Authorization: Bearer jwt_token_here`
   - **Body**:
     ```json
     {
       "payment_method": "credit_card"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Order placed successfully",
       "order_id": "1"
     }
     ```

---

## 4. Cara Instalasi dan Penggunaan

### Prasyarat

- Golang
- Database (MySQL atau PostgreSQL)
- Alat pendukung seperti Postman untuk pengujian API

### Langkah Instalasi

1. Clone repositori:

```
git clone https://github.com/username/repo-name.git
cd repo-name
```

2. Install Depedencies

```
go mod tidy
```

3. Contoh env

```
DATABASE_USER="root"
DATABASE_PASSWORD="password"
DATABASE_HOST="localhost"
DATABASE_PORT="3306"
DATABASE_NAME="food_sale"
JWT_SECRET_KEY="your_jwt_secret"
```

4. Run app

```
go run main.go
```
