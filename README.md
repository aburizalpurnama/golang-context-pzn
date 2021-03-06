# Belajar Golang Context by. PZN

## Context With Value

### context.WithValue(parent, key, value)

- Digunakan untuk membuat sebuah child context yang terdapat nilai didalamnya

- Nilai yang terdapat pada sebuah context juga akan terdapat pada semua child context turunannya

- Context yang lebih bawah dapat mengambil nilai dari semua parent yang terhubung

## Context With Cancel

### context.WithCancel(parent)

- Digunakan untuk mengirimkan sebuah sinyal pembatalan suatu proses (cancel) pada sebuah goroutine yang sedang berjalan.

- Secara default semua operasi database dan web pasti menggunakan context, sehingga dapat mengirimkan sinyal cancel jika perlu.

- Dapat juga digunakan untuk menghindari adanya goroutine leaks (goroutine yang terus berjalan tanpa guna)

## Context With Timeout

###  context.WithTimeout(parent)

- Mirip seperti context with cancel, bedanya sinyal cancel tidak dijalankan manual melaikan secara otomatis menggunakan batas waktu tertentu. (Durasi)

- Cocok sekali digunakan untuk melakukan query ke database atau http api dengan memberikan batas maksimal (timeout)

- Cancel function harus tetap dijalankan jika suatu proses telah selesai (mengantisipasi jika suatu proses selesai lebih cepat dibandingkan durasi timeout)

## Context With Deadline

### context.WithDeadline(parent, time)

- Mirip seperti context with time out, tapi bedanya jika timeout berdasarkan durasi sedangkan deadline berdasarkan waktu (jam tertentu)