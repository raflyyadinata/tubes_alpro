package main

import "fmt"

type Pakaian struct {
	Nama            string
	Kategori        string
	Warna           string
	Formalitas      int
	TerakhirDipakai int
}

var daftarPakaian []Pakaian

func tambahPakaian() {
	var p Pakaian
	var tambahan string
	fmt.Println("=== Tambah Pakaian ===")
	fmt.Print("Nama: ")
	fmt.Scan(&p.Nama)
	fmt.Print("Kategori (casual/formal/sporty): ")
	fmt.Scan(&p.Kategori)
	fmt.Print("Warna: ")
	fmt.Scan(&p.Warna)
	fmt.Print("Tingkat Formalitas (1-10): ")
	fmt.Scan(&p.Formalitas)
	fmt.Print("Hari terakhir dipakai (1-31): ")
	fmt.Scan(&p.TerakhirDipakai)
	daftarPakaian = append(daftarPakaian, p)
	fmt.Println("Pakaian berhasil ditambahkan!")
	fmt.Println("apakah kamu ingin menambahkan lagi?")
	fmt.Scan(&tambahan)
	if tambahan == "ya" {
		tambahPakaian()
	}
}

func listpakaian() {
	if len(daftarPakaian) == 0 {
		fmt.Println("Belum ada pakaian yang ditambahkan.")
		return
	}

	fmt.Println("=== Daftar Semua Pakaian ===")
	for i, p := range daftarPakaian {
		fmt.Printf("%d. Nama: %s | Kategori: %s | Warna: %s | Formalitas: %d | Terakhir Dipakai: %d\n",
			i+1, p.Nama, p.Kategori, p.Warna, p.Formalitas, p.TerakhirDipakai)
	}
}

func pakaiansequential(kataKunci string) {
	fmt.Println("=== Hasil Pencarian ===")
	ketemu := false
	for _, p := range daftarPakaian {
		if p.Warna == kataKunci || p.Kategori == kataKunci {
			fmt.Println(p)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ditemukan.")
	}
}

func urutkanFormalitas() {
	n := len(daftarPakaian)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarPakaian[j].Formalitas < daftarPakaian[minIdx].Formalitas {
				minIdx = j
			}
		}
		daftarPakaian[i], daftarPakaian[minIdx] = daftarPakaian[minIdx], daftarPakaian[i]
	}
	fmt.Println("=== Diurutkan Berdasarkan Formalitas ===")
	listpakaian()
}

func terakhirdipakai() {
	n := len(daftarPakaian)
	for i := 1; i < n; i++ {
		key := daftarPakaian[i]
		j := i - 1
		for j >= 0 && daftarPakaian[j].TerakhirDipakai > key.TerakhirDipakai {
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = key
	}
	fmt.Println("=== Diurutkan Berdasarkan Terakhir Dipakai ===")
	listpakaian()
}

func urutkanKategori() {
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if daftarPakaian[j].Kategori > daftarPakaian[j+1].Kategori {
				daftarPakaian[j], daftarPakaian[j+1] = daftarPakaian[j+1], daftarPakaian[j]
			}
		}
	}
}

func urutkanWarna() {
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if daftarPakaian[j].Warna > daftarPakaian[j+1].Warna {
				daftarPakaian[j], daftarPakaian[j+1] = daftarPakaian[j+1], daftarPakaian[j]
			}
		}
	}
}

func binarySearchKategori(target string) {
	urutkanKategori()
	low := 0
	high := len(daftarPakaian) - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2
		if daftarPakaian[mid].Kategori == target {
			left := mid
			for left >= 0 && daftarPakaian[left].Kategori == target {
				left--
			}
			right := mid + 1
			for right < len(daftarPakaian) && daftarPakaian[right].Kategori == target {
				right++
			}
			for i := left + 1; i < right; i++ {
				fmt.Println("Ditemukan:", daftarPakaian[i])
			}
			ketemu = true
			break
		} else if daftarPakaian[mid].Kategori < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func sequentialSearchWarna(target string) {
	ketemu := false
	fmt.Println("=== Hasil Pencarian Warna ===")
	for _, p := range daftarPakaian {
		if p.Warna == target {
			fmt.Println("Ditemukan:", p)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Warna tidak ditemukan.")
	}
}

func editPakaian() {
	var nama string
	fmt.Print("Masukkan nama pakaian yang ingin diedit: ")
	fmt.Scan(&nama)

	for i := range daftarPakaian {
		if daftarPakaian[i].Nama == nama {
			fmt.Println("Data ditemukan. Silakan isi ulang:")
			fmt.Print("Nama: ")
			fmt.Scan(&daftarPakaian[i].Nama)
			fmt.Print("Kategori: ")
			fmt.Scan(&daftarPakaian[i].Kategori)
			fmt.Print("Warna: ")
			fmt.Scan(&daftarPakaian[i].Warna)
			fmt.Print("Formalitas: ")
			fmt.Scan(&daftarPakaian[i].Formalitas)
			fmt.Print("Terakhir dipakai: ")
			fmt.Scan(&daftarPakaian[i].TerakhirDipakai)
			fmt.Println("Berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Pakaian tidak ditemukan.")
}

func hapusPakaian() {
	var nama string
	fmt.Print("Masukkan nama pakaian yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i := range daftarPakaian {
		if daftarPakaian[i].Nama == nama {
			daftarPakaian = append(daftarPakaian[:i], daftarPakaian[i+1:]...)
			fmt.Println("Pakaian berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pakaian tidak ditemukan.")
}

func rekomendasi(cuaca string, acara string) {
	fmt.Println("=== Rekomendasi Outfit ===")
	for _, p := range daftarPakaian {
		if cuaca == "hujan" && p.Warna == "gelap" {
			fmt.Println("Untuk cuaca hujan:", p)
		} else if acara == "formal" && p.Formalitas >= 7 {
			fmt.Println("Untuk acara formal:", p)
		}
	}
}

func kombinasiOutfit() {
	fmt.Println("=== Kombinasi Outfit Tersedia ===")
	ketemu := false
	for i := 0; i < len(daftarPakaian); i++ {
		for j := i + 1; j < len(daftarPakaian); j++ {
			if daftarPakaian[i].Kategori != daftarPakaian[j].Kategori {
				fmt.Printf("- %s (%s) + %s (%s)\n", daftarPakaian[i].Nama, daftarPakaian[i].Kategori, daftarPakaian[j].Nama, daftarPakaian[j].Kategori)
				ketemu = true
			}
		}
	}
	if !ketemu {
		fmt.Println("Belum ada kombinasi yang bisa dibuat.")
	}
}

func main() {

	daftarPakaian = []Pakaian{
		{"kaos", "casual", "putih", 3, 5},
		{"jasm", "formal", "hitam", 9, 12},
		{"hoodie", "casual", "abu", 5, 2},
		{"celana_pendek", "casual", "putih", 4, 6},
		{"celana_panjang", "formal", "hitam", 8, 4},
		{"kaos", "casual", "hitam", 3, 6},
		{"celana_jeans", "casual", "biru", 7, 20},
	}

	var pilihan, x int
	for {
		fmt.Println("\n=== OOTD Planner ===")
		fmt.Println("1. Tambah Pakaian")
		fmt.Println("2. Tampilkan Semua")
		fmt.Println("3. Edit Pakaian")
		fmt.Println("4. Cari Pakaian")
		fmt.Println("5. Rekomendasi Pakaian")
		fmt.Println("6. Kombinasi Outfit")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPakaian()
		case 2:
			// var x int
			listpakaian()
			fmt.Println("\n=== URUTKAN BERDASARKAN: ===\n1. formalitas \n2. terakhir dipakai\n0. untuk kembali ")
			fmt.Print("Urutkan berdasarkan : ")
			fmt.Scan(&x)
			if x == 1 {
				urutkanFormalitas()
			} else if x == 2 {
				terakhirdipakai()
			}
		case 3:
			// var x int
			listpakaian()
			fmt.Println("\n=== EDIT PAKAIAN ===\n1. edit\n2. hapus pakaian\n0. untuk kembali")
			fmt.Print("Pilih ke : ")
			fmt.Scan(&x)
			if x == 1 {
				editPakaian()
			} else if x == 2 {
				hapusPakaian()
			}
		case 4:
			var a string
			fmt.Println("\n=== CARI PAKAIAN ===\n1. cari berdasarkan warna\n2. cari berdasarkan kategori\n0. untuk kembali")
			fmt.Print("Cari berdasarkan : ")
			fmt.Scan(&x)
			if x == 1 {
				fmt.Print("Masukkan warna yang ingin dicari: ")
				fmt.Scan(&a)
				sequentialSearchWarna(a)
			} else if x == 2 {
				fmt.Print("Masukkan kategori yang ingin dicari: ")
				fmt.Scan(&a)
				binarySearchKategori(a)
			}

		case 5:
			var cuaca, acara string
			fmt.Print("Masukkan cuaca (cerah/hujan): ")
			fmt.Scan(&cuaca)
			fmt.Print("Masukkan acara (santai/formal): ")
			fmt.Scan(&acara)
			rekomendasi(cuaca, acara)

		case 6:
			kombinasiOutfit()

		case 0:
			fmt.Println("terima kasih")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
