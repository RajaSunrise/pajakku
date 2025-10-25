package request

type CreateBillingRequest struct {
	KodeAkunPajak    string `json:"kode_akun_pajak" validate:"required,alphanum,len=6"`
	KodeJenisSetoran string `json:"kode_jenis_setoran" validate:"required,numeric,len=3"`
	MasaPajak        int    `json:"masa_pajak" validate:"required,min=1,max=12"`
	TahunPajak       int    `json:"tahun_pajak" validate:"required,min=2020"`
	JumlahSetor      int64  `json:"jumlah_setor" validate:"required,gt=0"`
}
