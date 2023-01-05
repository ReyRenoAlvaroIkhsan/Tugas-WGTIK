package main

import (
	"fmt"
	"os"
	"os/exec"
)

const T int = 10
const N int = 4
const M int = 4

type kendaraan struct {
	nk         string
	jenis      string
	jam_masuk  string
	jam_keluar string
}

type arrMotor struct {
	data [T]kendaraan
	n    int
}

type arrMobil struct {
	data [N][M]kendaraan
	n    int
}

type struk struct {
	data   kendaraan
	keluar string
}

var area1 arrMotor
var area2 arrMobil

func menu(pilih *int) {
	fmt.Println("SELAMAT DATANG DI PARKIRAN\nSilahkan pilih opsi berikut untuk mengakses fitur-fitur berikutini:\nPilih 1 untuk masukkan data kendaraan\nPilih 2 untuk lihat parkir motor\nPilih 3 untuk lihat parkir mobil\nPilih 4 untuk mengecek penuh atau tidak\nPilih 5 untuk melihat Statistik kendaraan\nPilih 6 untuk melihat presentase okupansi parkir untuk masing-masing area\nPilih 7 untuk mencari kendaraan\nPilih 8 untuk mengurutkan kendaraan berdasarkan nomor kendaraan\nPilih 0 untuk keluar (selesai)")

	fmt.Scan(&*pilih)
}

func kendaraanMasuk() {
	var data kendaraan

	fmt.Print("Masukkan nomor kendaraan : ")
	fmt.Scan(&data.nk)
	fmt.Print("Masukkan jenis kendaraan : ")
	fmt.Scan(&data.jenis)
	fmt.Print("Masukkan waktu masuk : ")
	fmt.Scan(&data.jam_masuk)

	if data.jenis == "motor" || data.jenis == "Motor" {
		if area1.n < T {
			area1.data[area1.n] = data
			area1.n++
			fmt.Println("Data Berhasil Diinput")
		}

		cekArea1()

	} else if data.jenis == "mobil" || data.jenis == "Mobil" {

		var inputed bool = false

		for i := 0; i < N && !inputed; i++ {
			for j := 0; j < M && !inputed; j++ {
				if area2.data[i][j].nk == "" {
					area2.data[i][j] = data

					area2.n++

					inputed = true

					fmt.Println("Data Berhasil Diinput")

					cekArea2()
				}
			}
		}

	} else {
		var inputed bool = false

		for i := 0; i < N && !inputed; i++ {
			for j := 0; j < M-1 && !inputed; j++ {
				if area2.data[i][j].nk == "" && area2.data[i][j+1].nk == "" {
					area2.data[i][j] = data
					area2.data[i][j+1] = data

					area2.n += 2

					inputed = true

					fmt.Println("Data Berhasil Diinput")

					cekArea2()
				}
			}
		}
	}
}

func cekArea1() {
	for i := 0; i < T; i++ {
		if area1.data[i].nk == "" {
			fmt.Print("{}", " ")
		} else {
			fmt.Print(area1.data[i], " ")
		}
	}
}

func cekArea2() {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if area2.data[i][j].nk == "" {
				fmt.Print("{}", " ")
			} else {
				fmt.Print(area2.data[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func penuh() {
	var status1, status2 bool = false, false
	for i := 0; i < T && !status1; i++ {
		if area1.data[i].nk == "" {
			status1 = true
		}
	}
	if status1 {
		fmt.Println("Area1 Kosong")
	} else {
		fmt.Println("Area1 Penuh")
	}
	for i := 0; i < N && !status2; i++ {
		for j := 0; j < M && !status2; j++ {
			if area2.data[i][j].nk == "" {
				status2 = true
			}
		}
	}
	if status2 {
		fmt.Println("Area2 Kosong")
	} else {
		fmt.Println("Area2 Penuh")
	}
}

func cari() bool {
	var area1 arrMotor
	var area2 arrMobil
	var nk string
	var i, j int
	var status1, status2 bool

	status1 = false
	status2 = false

	fmt.Print("Berapa nomor kendaraan anda? : ")
	fmt.Scan(&nk)
	for i = 0; i < area1.n; i++ {
		if nk == area1.data[i].nk {
			status1 = status1 || true
		}
	}

	for i = 0; i < area2.n; i++ {
		if nk == area2.data[i][j].nk {
			status2 = status2 || true
		}
	}
	return status1 || status2
}

func nomorKendaraan() {
	var i, j, i1, iMin, iMin1, pass, pass1 int
	var temp, temp1 kendaraan

	pass = 1
	pass1 = 1
	for pass = 1; pass <= area1.n; pass++ {
		iMin = pass - 1
		for i = pass; i <= area1.n; i++ {
			if area1.data[iMin].nk > area1.data[i].nk {
				iMin = i
			}
		}
		temp = area1.data[iMin]
		area1.data[iMin] = area1.data[pass-1]
		area1.data[pass-1] = temp
	}
	for pass1 = 1; pass1 <= area2.n; pass1++ {
		iMin1 = pass1 - 1
		for i = pass1; i <= area2.n; i1++ {
			if area2.data[iMin1][iMin1].nk > area2.data[i1][i1].nk {
				iMin1 = i1
			}
		}
		temp1 = area2.data[iMin1][iMin1]
		area2.data[iMin1][iMin1] = area2.data[pass1-1][pass1-1]
		area2.data[pass1-1][pass1-1] = temp1
	}

	fmt.Println("Data Motor")
	for i := 0; i <= area1.n; i++ {
		fmt.Println(area1.data[i].nk)
	}

	fmt.Println("Data Mobil")
	for i := 0; i <= area2.n; i++ {
		fmt.Println(area2.data[i][j].nk)
	}

}

func harga(jam1, jam2, menit1, menit2 int) int {
	var harga, menittotal, jamtotal int
	menittotal = menit2 - menit1
	jamtotal = jam2 - jam1
	if menittotal > 10 {
		jamtotal++
	}
	harga = jamtotal * 2000
	return harga
}

func statistikKendaraan() {
	var i int = 0
	var j int = 0
	var nMotor, nMobil, nTruk int
	for i < T {
		if area1.data[i].jenis == "motor" || area1.data[i].jenis == "Motor" {
			nMotor++
		}
		i++
	}
	i = 0
	for i < N {
		for j < M {
			if area2.data[i][j].jenis == "mobil" || area2.data[i][j].jenis == "Mobil" {
				nMobil++
			} else if area2.data[i][j].jenis != "" {
				nTruk++
			}
			j++
		}
		i++
	}
	fmt.Println("Total Motor yang parkir: ", nMotor)
	fmt.Println("Total Mobil yang parkir: ", nMobil)
	fmt.Println("Total Truk yang parkir: ", nTruk/2)
	fmt.Println("Total kendaraan yang parkir: ", nMotor+nMobil+(nTruk/2))
	fmt.Println()
}

func statistikOkupansiParkir() {
	var area string
	fmt.Print("Pilih area1/area2 : ")
	fmt.Scan(&area)
	if area == "1" {
		fmt.Println((float64(area1.n)/float64(T))*float64(100), "%")
	} else {
		fmt.Println((float64(area2.n)/float64(N*M))*float64(100), "%")
	}
}

func main() {
	var running bool = true

	area1.n = 0
	area2.n = 0

	for running {

		var pilih int
		menu(&pilih)

		switch pilih {
		case 1:
			// fmt.Print("kendaraan masuk")
			kendaraanMasuk()
		case 2:
			cekArea1()
		case 3:
			cekArea2()
		case 4:
			penuh()
		case 5:
			statistikKendaraan()
		case 6:
			statistikOkupansiParkir()
		case 7:
			cari()
		case 8:
			nomorKendaraan()
		case 0:
			fmt.Println("keluar")
			running = false
		default:
			fmt.Print("Input Nomor yang valid")
			fmt.Println()
			fmt.Println()
		}

		fmt.Scanln()
		fmt.Scanln()

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
