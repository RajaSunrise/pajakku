package request

type ReportSPTRequest struct {
	JenisSPT      string `form:"jenis_spt" validate:"required"`
	PeriodePajak  string `form:"periode_pajak" validate:"required"`
	StatusLaporan string `form:"status_laporan" validate:"required,oneof=Nihil 'Kurang Bayar' 'Lebih Bayar'"`
}
