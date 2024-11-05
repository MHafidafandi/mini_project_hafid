# Rangkuman Aplikasi Penjualan Makanan Sisa

## 1. MVP (Minimum Viable Product)

**Tujuan Utama**:
Membuat platform yang memudahkan penjual untuk menawarkan makanan sisa berkualitas dan pembeli untuk membeli dengan harga terjangkau.

**Fitur Utama**:

- **Autentikasi dan Registrasi Pengguna**: Pembeli dan penjual dapat membuat akun dan login.
- **Manajemen Produk**: Penjual dapat menambah, memperbarui, dan menghapus produk makanan sisa.
- **Keranjang Belanja**: Pembeli dapat menambahkan produk ke keranjang dan melakukan checkout.
- **Transaksi**: Pembeli dapat melakukan pembayaran online melalui sistem checkout.
- **Riwayat Pembelian**: Pembeli dan penjual dapat melihat riwayat transaksi masing-masing.

---

## 2. Struktur Database

### Tabel Utama

- **users**

  - `id` (PK)
  - `name` (Nama pengguna)
  - `email`
  - `password`
  - `role` (buyer/seller)
  - `address`
  - `phone`

- **Foods**

  - `id` (PK)
  - `name` (Nama produk)
  - `description` (Deskripsi produk)
  - `price` (Harga produk)
  - `stock` (Stok produk)
  - `expiry_date` (Tanggal kedaluwarsa)
  - `location` (Lokasi produk)
  - `user_id` (FK ke `users`)
  - `store_id` (FK ke `stores`)

- **orders**

  - `id` (PK)
  - `user_id` (FK ke `users`)
  - `total_amount` (Total harga)
  - `status` (Status pesanan)
  - `order_date` (Tanggal pesanan)

- **order_items**
  - `id` (PK)
  - `order_id` (FK ke `orders`)
  - `food_id` (FK ke `foods`)
  - `quantity` (Jumlah produk)
  - `price` (Harga satuan)

### Relasi Antar Tabel

- **Users - Stores**: Satu pengguna (penjual) dapat memiliki banyak toko.
- **Stores - Products**: Satu toko dapat memiliki banyak produk.
- **Users - Orders**: Satu pengguna (pembeli) dapat memiliki banyak transaksi.
- **Orders - Order Items**: Satu transaksi dapat memiliki banyak item produk.

### Manfaat Tabel _Stores_

Dengan adanya tabel _stores_, data toko lebih terorganisir, memungkinkan pembeli untuk melihat informasi lengkap tentang toko penjual, termasuk produk yang mereka tawarkan.

---

## 3. Penerapan AI

**Integrasi Gen AI untuk Meningkatkan Fitur**:

- **Deskripsi Produk Otomatis**: Membantu penjual membuat deskripsi produk menarik dengan memasukkan informasi dasar, di mana AI memberikan saran untuk deskripsi produk.

---

## 4. Penerapan API Eksternal

**API Eksternal yang Dapat Diterapkan**:

1. **Payment Gateway API** (Contoh: Midtrans, Stripe)
   - Untuk memproses pembayaran online dengan aman saat checkout.

---
