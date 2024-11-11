### Jalankan Postgres sebagai container dengan nama container **"my-postgres-joycepatrcia"** dan volume **"my-pg-volume-joycepatrcia"**
dengan code : docker run -d -v my-pg-volume-joycepatrcia:/var/lib/postgresql/data --name my-postgres-joycepatrcia -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
<img width="949" alt="1" src="https://github.com/user-attachments/assets/f766bbb1-dfda-4971-a1f4-c065a2a2a365">
menghasilkan container dan volume:
<img width="959" alt="1 5" src="https://github.com/user-attachments/assets/69b7d794-2721-4156-876c-446b08f0ddd0">

### Cek Connection Postgres
<img width="336" alt="2" src="https://github.com/user-attachments/assets/74ef41b9-7726-4123-a661-28784432690c">

### Buat table baru dengan nama "new_table_joycepatrcia" dan isi 10 record
<img width="448" alt="3" src="https://github.com/user-attachments/assets/2a45042c-e34c-4f00-8c4f-b2d876b1eb2f">
<img width="308" alt="4" src="https://github.com/user-attachments/assets/60e8f600-2e7c-49a6-863d-9ce8a4775df4">

### Stop dan Hapus container postgres **"my-postgres-joycepatrcia"**

### Jalankan ulang postgres sebagai container dengan nama baru "my-postgres-v2-joycepatrcia"
code : docker run -d -v my-pg-volume-joycepatrcia:/var/lib/postgresql/data --name my-postgres-v2-joycepatrcia -e POSTGRES_PASS
WORD=password -p 5432:5432 postgres
<img width="959" alt="5" src="https://github.com/user-attachments/assets/185925e4-aa79-4bf1-965c-a69c29c06aa7">
<img width="959" alt="5 5" src="https://github.com/user-attachments/assets/6a991411-c714-45b7-908b-c81d7f9e3f5f">

### Cek apakah data masih ada
<img width="329" alt="6" src="https://github.com/user-attachments/assets/5a418275-7a28-4741-ad75-29f2267dc52e">
