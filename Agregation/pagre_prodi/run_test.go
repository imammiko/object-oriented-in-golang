package pagreprodi

import "testing"

func TestRun(t *testing.T) {
	// create new jurusan
	jurusan := NewJurusan("TI", "Teknik Informatika")

	// create new mahasiswa
	mhs1 := Mahasiswa{}
	mhs2 := Mahasiswa{}

	// add mahasiswa to jurusan

	mhs1.SetNim("112233")
	mhs1.SetNama("Budi")

	mhs2.SetNim("223344")
	mhs2.SetNama("Susi")

	jurusan.AddMahasiswa(mhs1)
	jurusan.AddMahasiswa(mhs2)


	// get daftar mahasiswa
	daftarMhs := jurusan.GetDaftarMhs()

	// print daftar mahasiswa
	for _, m := range daftarMhs {
		t.Logf("NIM: %s, Nama: %s", m.GetNim(), m.GetNama())
	}
}