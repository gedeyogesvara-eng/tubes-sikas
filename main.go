package main

import (
	"fmt"
	"time"
)

type Mahasiswa struct {
	Nama         string
	NIM          string
	Totalbayar   int
	StatusLunas  bool
	TanggalBayar string
}

var datamahasiswa []Mahasiswa

const kasperorang = 100000

func tampilkandaftarkas() {
	fmt.Println("Daftar Kas Kelas :")
	fmt.Printf("NIM\t\tNama\t\tTotal Bayar\t\tTanggal\t\tStatus\t\tSisa Tunggakan\n")
	fmt.Println("______________________________________________________________________________________________________")
	for i := 0; i < len(datamahasiswa); i++ {
		m := datamahasiswa[i]
		fmt.Printf("%s\t\t%s\t\t%d\t\t\t%s\t%s\t%d\n", m.NIM, m.Nama, m.Totalbayar, m.TanggalBayar, statuslunas(m.StatusLunas), kasperorang-m.Totalbayar)
	}
}

func statuslunas(lunas bool) string {
	if lunas {
		return "LUNAS"
	} else {
		return "BELUM LUNAS"
	}
}

func TambahMahasiswa() {
	var nim, nama string

	fmt.Println("=== Tambah Mahasiswa ===")

	fmt.Println("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Println("Masukkan Nama: ")
	fmt.Scanln(&nama)

	mhsBaru := Mahasiswa{
		NIM:         nim,
		Nama:        nama,
		Totalbayar:  0,
		StatusLunas: false,
	}

	datamahasiswa = append(datamahasiswa, mhsBaru)

	fmt.Println("Data mahasiswa berhasil ditambahkan")
}

func pembayarankas() {
	var input string
	fmt.Print("Masukkan NIM atau Nama Mahasiswa: ")
	fmt.Scanln(&input)
	idx := -1
	i := 0
	for i < len(datamahasiswa) {
		if datamahasiswa[i].NIM == input || datamahasiswa[i].Nama == input {
			idx = i
		}
		i++
	}
	if idx != -1 {
		fmt.Println("\nData ditemukan!")
		fmt.Println("NIM  :", datamahasiswa[idx].NIM)
		fmt.Println("Nama :", datamahasiswa[idx].Nama)
		var nominal int
		fmt.Print("Masukkan nominal pembayaran: ")
		fmt.Scanln(&nominal)
		tgl := time.Now().Format("02-01-2006")
		datamahasiswa[idx].TanggalBayar = tgl
		datamahasiswa[idx].Totalbayar += nominal
		if datamahasiswa[idx].Totalbayar >= kasperorang {
			datamahasiswa[idx].StatusLunas = true
		}
		fmt.Println("\nPEMBAYARAN BERHASIL")
		fmt.Println("Tanggal Bayar :", datamahasiswa[idx].TanggalBayar)
		fmt.Println("Total Bayar   :", datamahasiswa[idx].Totalbayar)
		if datamahasiswa[idx].StatusLunas {
			fmt.Println("Status : LUNAS")
		} else {
			sisa := kasperorang - datamahasiswa[idx].Totalbayar
			fmt.Println("Status : BELUM LUNAS")
			fmt.Println("Sisa Tunggakan :", sisa)
		}
	} else {
		fmt.Println("\nData mahasiswa tidak ditemukan!")
	}
}

func editData() {
	if len(datamahasiswa) == 0 {
		fmt.Println("\n✗ Belum ada Data Mahasiswa!")
		return
	}
	var inputNIM string
	fmt.Print("\nMasukkan NIM Mahasiswa yang akan diedit: ")
	fmt.Scanln(&inputNIM)
	idx := -1
	for i := 0; i < len(datamahasiswa); i++ {
		if datamahasiswa[i].NIM == inputNIM {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("\n✗ Mahasiswa dengan NIM tidak ditemukan!")
		return
	}
	fmt.Println("\nDATA MAHASISWA")
	fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
	fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
	fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
	fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
	if !datamahasiswa[idx].StatusLunas {
		fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
	}
	fmt.Println("\nPILIH DATA YANG AKAN DIEDIT")
	fmt.Println("1. Edit NIM")
	fmt.Println("2. Edit Nama")
	fmt.Println("3. Edit Total Bayar")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		var NIMbaru string
		fmt.Print("\nMasukkan NIM baru: ")
		fmt.Scanln(&NIMbaru)
		datamahasiswa[idx].NIM = NIMbaru
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ NIM berhasil diubah!")
	} else if pilihan == 2 {
		var namabaru string
		fmt.Print("\nMasukkan Nama Baru: ")
		fmt.Scanln(&namabaru)
		datamahasiswa[idx].Nama = namabaru
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Nama berhasil diubah!")
	} else if pilihan == 3 {
		fmt.Printf("\nTotal Bayar Saat Ini: Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Sisa Tunggakan Saat Ini: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		var totalbaru int
		fmt.Print("Masukkan Total Bayar baru: Rp ")
		fmt.Scanln(&totalbaru)
		if totalbaru < 0 {
			fmt.Println("\n✗ Total bayar tidak boleh negatif!")
			return
		}
		datamahasiswa[idx].Totalbayar = totalbaru
		if datamahasiswa[idx].Totalbayar >= kasperorang {
			datamahasiswa[idx].StatusLunas = true
		} else {
			datamahasiswa[idx].StatusLunas = false
		}
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Total bayar berhasil diubah!")
	} else {
		fmt.Println("\n✗ Pilihan tidak valid! Silakan pilih 1, 2, atau 3.")
		return
	}
}

func main() {
	var pilih int

	for {
		fmt.Println("___________________________________________")
		fmt.Println("       SIKAS APLIKASI KAS MAHASISWA       ")
		fmt.Println("___________________________________________")
		fmt.Println("1. Tampilkan Daftar Kas & Status")
		fmt.Println("2. Input Pembayaran Kas")
		fmt.Println("3. Input Pengeluaran Kelas")
		fmt.Println("4. Cari Data Mahasiswa")
		fmt.Println("5. Urutkan Penunggak Kas")
		fmt.Println("6. Tambah Data Baru")
		fmt.Println("7. Edit Data")
		fmt.Println("8. Menghapus Data")
		fmt.Println("9. Keluar")
		fmt.Println("___________________________________________")
		fmt.Print("Pilih menu (1-9): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Anda memilih 1 : Tampilkan Daftar Kas & Status")
			tampilkandaftarkas()
		case 2:
			fmt.Println("Anda memilih 2 : Input Pembayaran Kas")
			pembayarankas()
		case 3:
			fmt.Println("Anda memilih 3 : Input Pengeluaran Kelas")
		case 4:
			fmt.Println("Anda memilih 4 : Cari Data Mahasiswa")
		case 5:
			fmt.Println("Anda memilih 5 : Urutkan Penunggak Kas")
		case 6:
			fmt.Println("Anda memilih 6 : Tambah Data Baru")
			TambahMahasiswa()
		case 7:
			fmt.Println("Anda memilih 7 : Edit Data")
			editData()
		case 8:
			fmt.Println("Anda memilih 8 : Menghapus Data")
		case 9:
			fmt.Println("Anda memilih 9 : Keluar")
			fmt.Println("Terima kasih telah menggunakan aplikasi SIKAS")
			return
		}
	}
}
