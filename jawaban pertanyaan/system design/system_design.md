# System Design: Platform Pembelian Tiket Bioskop Online

## 1. Solusi Teknis

### A. Sistem Pemilihan Kursi & Performa Tinggi (Anti-Double Booking)

1. **Kunci Kursi Sementara (Seat Locking)**:
   - Ketika user memilih kursi (misal `A1`), sistem langsung mengunci kursi tersebut dengan status `locked` selama **10 menit**.
   - User lain yang melihat denah kursi pada waktu bersamaan akan melihat kursi tersebut sudah tidak tersedia.

2. **Proteksi di Level Database**:
   - Pada tabel database dipasang constraint unik: `UNIQUE(showtime_id, seat_number)`.
   - Jika ada 2 user menekan tombol bayar untuk kursi yang sama di milidetik yang sama, database hanya akan menerima transaksi pertama dan otomatis menolak transaksi kedua.

3. **Performa Sistem pada Trafik Tinggi**:
   - Menggunakan **Redis Cache** untuk menyimpan data katalog film, jadwal, dan ketersediaan kursi secara in-memory.
   - Hal ini mencegah database utama kelebihan beban (overload) saat ribuan orang mengakses aplikasi secara bersamaan (misalnya saat presale tiket film besar).

---

### B. Sistem Pencatatan dan Restok Tiket

1. **Status Transaksi & Tiket**:
   - `pending` / `locked`: User sedang proses bayar (berlaku 10 menit).
   - `paid` / `booked`: Pembayaran berhasil dan tiket aktif.
   - `expired` / `cancelled`: Waktu bayar habis atau dibatalkan.
   - `refunded`: Transaksi dibatalkan dan dana dikembalikan.

2. **Mekanisme Restok Otomatis**:
   - Setiap transaksi memiliki `payment_deadline` selama 10 menit.
   - Terdapat **Background Worker / Cron Job** yang mengecek transaksi pending setiap menit.
   - Jika waktu 10 menit habis dan user belum membayar, sistem mengubah status order menjadi `expired` dan menghapus status lock kursi.
   - Kursi otomatis langsung tersedia kembali (**restok**) untuk dipilih pengunjung lain tanpa intervensi manual.

---

### C. Alur Refund & Pembatalan dari Pihak Bioskop

Jika terjadi kendala operasional tak terduga (misal proyektor rusak, studio bermasalah, atau force majeure):

1. **Pembatalan oleh Admin**:
   - Admin bioskop mengubah status jadwal tayang menjadi `cancelled` pada sistem dan memasukkan alasan pembatalan.

2. **Proses Pengembalian Dana Otomatis**:
   - Sistem mendeteksi seluruh transaksi yang berstatus `paid` pada jadwal tersebut.
   - Backend memicu API Payment Gateway untuk mengembalikan dana secara otomatis ke rekening / e-wallet asal pengguna.

3. **Pencatatan & Notifikasi**:
   - Status order dan tiket diperbarui menjadi `refunded` beserta catatan alasan pembatalan (`refund_reason`).
   - Sistem mengirimkan notifikasi (email/WhatsApp) ke pembeli bahwa penayangan dibatalkan dan dana telah berhasil dikembalikan.
