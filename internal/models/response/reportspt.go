package response

import "time"

type ReportSPTResponse struct {
	ID            uint      `json:"id"`
	UserProfileID uint      `json:"user_profile_id"`
	JenisSPT      string    `json:"jenis_spt"`
	PeriodePajak  string    `json:"periode_pajak"`
	StatusLaporan string    `json:"status_laporan"`
	TanggalLapor  time.Time `json:"tanggal_lapor"`
	NTTE          string    `json:"ntte"`
	FileBPEPath   string    `json:"file_bpe_path"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
