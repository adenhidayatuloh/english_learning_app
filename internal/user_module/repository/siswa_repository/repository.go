package siswarepository

import (
	"english_app/internal/user_module/entity"
	"english_app/pkg/errs"
)

type SiswaRepository interface {
	CreateSiswa(siswa *entity.Siswa) (*entity.Siswa, errs.MessageErr)
	UpdateSiswa(oldSiswa *entity.Siswa, newSiswa *entity.Siswa) (*entity.Siswa, errs.MessageErr)
	GetSiswaByEmail(email string) (*entity.Siswa, errs.MessageErr)
	//GetAllSiswaWithPemeriksaan() ([]entity.Siswa_pemeriksaan, errs.MessageErr)
}
