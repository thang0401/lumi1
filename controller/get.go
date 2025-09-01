package controller

import (
	"context"
	"encoding/json"
	models "lumi/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func handleGetData(db *gorm.DB, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Kiểm tra cache
		cacheKey := "f3_data"
		cached, err := redisClient.Get(context.Background(), cacheKey).Result()
		if err == nil {
			c.Data(200, "application/json", []byte(cached))
			return
		}

		// Query DB
		var records []models.Orders
		result := db.Order("ngayKeToanDoiSoatFfm2 DESC").Find(&records)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		// Danh sách headers
		headers := []string{
			"maDonHang", "maTracking", "ngayLenDon", "name", "phone", "add", "city", "state", "zipcode",
			"matHang", "tenMatHang1", "soLuongMatHang1", "tenMatHang2", "soLuongMatHang2", "quaTang",
			"soLuongQuaKem", "giaBan", "loaiTienThanhToan", "tongTienVND", "hinhThucThanhToan",
			"ghiChu", "nhanVienSale", "nhanVienMarketing", "nvVanDon", "ketQuaCheck", "trangThaiGiaoHangNB",
			"lyDo", "donViVanChuyen", "trangThaiThuTien", "ngayHenDayDon", "soTienThucThu", "ngayUpBill",
			"anhBill", "ngayChuyenKhoan", "loaiThanhToan", "tenNguoiChuyenKhoan", "taiKhoanNhan",
			"soTienDaNhan", "tyGiaCuoc", "donViTienDoiSoat", "ngayDoiSoatCuoc", "tienShip", "tyGia",
			"donViTienTeTinhCuoc", "thoiGianLenDon", "ghiChuFFM", "ffmThanhToan", "ngayDoiSoatBill",
			"ngayKeToanDoiSoatFFM1", "soTienVeTKCtyFFM1", "ngayKeToanDoiSoatFFM2", "tienVietDaDoiSoat",
			"ngayKeToanDoiSoatFFM3", "ca", "tongSoTienFFM", "soTienVeTKCtyNgoaiFFM", "soTienDonHangVeTKCty",
			"keToanXacNhanThuTien", "thongTinDon", "thongTinKhachHang", "thongTinNhanSu", "phanHoiTichCuc",
			"phanHoiTieuCuc", "count", "phanLoaiKH", "xacNhanDon", "dienGiai", "tenPage", "giaGoc",
			"mauBacklist", "trangThaiDon", "tienChietKhau", "phiShip", "tienSauShip", "tenLenDon",
			"taoBanIn", "inDon", "canhBaoBlacklistTrungDon", "khuVuc", "phiLuuKho", "team", "maCheck",
			"ghiChuBEE", "danhDau", "ngayDongHang", "trangThaiGiaoHang", "thoiGianGiaoDuKien",
			"phiShipNoiDiaMyUSD", "phiXuLyDongHangUSD", "ghiChuKhac", "ngayDoiSoat", "timeKeToanXacNhan",
			"ghiChuVD", "donViThanhToan", "taiKhoanThanhToan", "updateAt", "trangThaiDVVC", "nhapKho",
			"tongPhiVC", "tienCuoc", "soTienDoiSoat", "maDonHang2", "sstDonMKT", "cskh", "thoiGianCutoff",
		}

		// Format ngày thành MM/dd/yyyy (trừ ngayKeToanDoiSoatFfm2)
		formattedRecords := make([]map[string]interface{}, len(records))
		for i, record := range records {
			formatted := make(map[string]interface{})
			formatted["maDonHang"] = record.MaDonHang
			formatted["maTracking"] = record.MaTracking
			formatted["ngayLenDon"] = record.NgayLenDon.Format("01/02/2006")
			formatted["name"] = record.Name
			formatted["phone"] = record.Phone
			formatted["add"] = record.Add
			formatted["city"] = record.City
			formatted["state"] = record.State
			formatted["zipcode"] = record.Zipcode
			formatted["matHang"] = record.MatHang
			formatted["tenMatHang1"] = record.TenMatHang1
			formatted["soLuongMatHang1"] = record.SoLuongMatHang1
			formatted["tenMatHang2"] = record.TenMatHang2
			formatted["soLuongMatHang2"] = record.SoLuongMatHang2
			formatted["quaTang"] = record.QuaTang
			formatted["soLuongQuaKem"] = record.SoLuongQuaKem
			formatted["giaBan"] = record.GiaBan
			formatted["loaiTienThanhToan"] = record.LoaiTienThanhToan
			formatted["tongTienVND"] = record.TongTienVND
			formatted["hinhThucThanhToan"] = record.HinhThucThanhToan
			formatted["ghiChu"] = record.GhiChu
			formatted["nhanVienSale"] = record.NhanVienSale
			formatted["nhanVienMarketing"] = record.NhanVienMarketing
			formatted["nvVanDon"] = record.NvVanDon
			formatted["ketQuaCheck"] = record.KetQuaCheck
			formatted["trangThaiGiaoHangNB"] = record.TrangThaiGiaoHangNB
			formatted["lyDo"] = record.LyDo
			formatted["donViVanChuyen"] = record.DonViVanChuyen
			formatted["trangThaiThuTien"] = record.TrangThaiThuTien
			formatted["ngayHenDayDon"] = record.NgayHenDayDon.Format("01/02/2006")
			formatted["soTienThucThu"] = record.SoTienThucThu
			formatted["ngayUpBill"] = record.NgayUpBill.Format("01/02/2006")
			formatted["anhBill"] = record.AnhBill
			formatted["ngayChuyenKhoan"] = record.NgayChuyenKhoan.Format("01/02/2006")
			formatted["loaiThanhToan"] = record.LoaiThanhToan
			formatted["tenNguoiChuyenKhoan"] = record.TenNguoiChuyenKhoan
			formatted["taiKhoanNhan"] = record.TaiKhoanNhan
			formatted["soTienDaNhan"] = record.SoTienDaNhan
			formatted["tyGiaCuoc"] = record.TyGiaCuoc
			formatted["donViTienDoiSoat"] = record.DonViTienDoiSoat
			formatted["ngayDoiSoatCuoc"] = record.NgayDoiSoatCuoc.Format("01/02/2006")
			formatted["tienShip"] = record.TienShip
			formatted["tyGia"] = record.TyGia
			formatted["donViTienTeTinhCuoc"] = record.DonViTienTeTinhCuoc
			formatted["thoiGianLenDon"] = record.ThoiGianLenDon.Format("01/02/2006")
			formatted["ghiChuFFM"] = record.GhiChuFFM
			formatted["ffmThanhToan"] = record.FfmThanhToan
			formatted["ngayDoiSoatBill"] = record.NgayDoiSoatBill.Format("01/02/2006")
			formatted["ngayKeToanDoiSoatFFM1"] = record.NgayKeToanDoiSoatFFM1.Format("01/02/2006")
			formatted["soTienVeTKCtyFFM1"] = record.SoTienVeTKCtyFFM1
			formatted["ngayKeToanDoiSoatFFM2"] = record.NgayKeToanDoiSoatFFM2 // Không format
			formatted["tienVietDaDoiSoat"] = record.TienVietDaDoiSoat
			formatted["ngayKeToanDoiSoatFFM3"] = record.NgayKeToanDoiSoatFFM3.Format("01/02/2006")
			formatted["ca"] = record.Ca
			formatted["tongSoTienFFM"] = record.TongSoTienFFM
			formatted["soTienVeTKCtyNgoaiFFM"] = record.SoTienVeTKCtyNgoaiFFM
			formatted["soTienDonHangVeTKCty"] = record.SoTienDonHangVeTKCty
			formatted["keToanXacNhanThuTien"] = record.KeToanXacNhanThuTien
			formatted["thongTinDon"] = record.ThongTinDon
			formatted["thongTinKhachHang"] = record.ThongTinKhachHang
			formatted["thongTinNhanSu"] = record.ThongTinNhanSu
			formatted["phanHoiTichCuc"] = record.PhanHoiTichCuc
			formatted["phanHoiTieuCuc"] = record.PhanHoiTieuCuc
			formatted["count"] = record.Count
			formatted["phanLoaiKH"] = record.PhanLoaiKH
			formatted["xacNhanDon"] = record.XacNhanDon
			formatted["dienGiai"] = record.DienGiai
			formatted["tenPage"] = record.TenPage
			formatted["giaGoc"] = record.GiaGoc
			formatted["mauBacklist"] = record.MauBacklist
			formatted["trangThaiDon"] = record.TrangThaiDon
			formatted["tienChietKhau"] = record.TienChietKhau
			formatted["phiShip"] = record.PhiShip
			formatted["tienSauShip"] = record.TienSauShip
			formatted["tenLenDon"] = record.TenLenDon
			formatted["taoBanIn"] = record.TaoBanIn
			formatted["inDon"] = record.InDon
			formatted["canhBaoBlacklistTrungDon"] = record.CanhBaoBlacklistTrungDon
			formatted["khuVuc"] = record.KhuVuc
			formatted["phiLuuKho"] = record.PhiLuuKho
			formatted["team"] = record.Team
			formatted["maCheck"] = record.MaCheck
			formatted["ghiChuBEE"] = record.GhiChuBEE
			formatted["danhDau"] = record.DanhDau
			formatted["ngayDongHang"] = record.NgayDongHang.Format("01/02/2006")
			formatted["trangThaiGiaoHang"] = record.TrangThaiGiaoHang
			formatted["thoiGianGiaoDuKien"] = record.ThoiGianGiaoDuKien.Format("01/02/2006")
			formatted["phiShipNoiDiaMyUSD"] = record.PhiShipNoiDiaMyUSD
			formatted["phiXuLyDongHangUSD"] = record.PhiXuLyDongHangUSD
			formatted["ghiChuKhac"] = record.GhiChuKhac
			formatted["ngayDoiSoat"] = record.NgayDoiSoat.Format("01/02/2006")
			formatted["timeKeToanXacNhan"] = record.TimeKeToanXacNhan
			formatted["ghiChuVD"] = record.GhiChuVD
			formatted["donViThanhToan"] = record.DonViThanhToan
			formatted["taiKhoanThanhToan"] = record.TaiKhoanThanhToan
			formatted["updateAt"] = record.UpdateAt.Format("01/02/2006")
			formatted["trangThaiDVVC"] = record.TrangThaiDVVC
			formatted["nhapKho"] = record.NhapKho
			formatted["tongPhiVC"] = record.TongPhiVC
			formatted["tienCuoc"] = record.TienCuoc
			formatted["soTienDoiSoat"] = record.SoTienDoiSoat
			formatted["maDonHang2"] = record.MaDonHang2
			formatted["sstDonMKT"] = record.SstDonMKT
			formatted["cskh"] = record.Cskh
			formatted["thoiGianCutoff"] = record.ThoiGianCutoff.Format("01/02/2006")
			formattedRecords[i] = formatted
		}

		response := gin.H{
			"headers": headers,
			"rows":    formattedRecords,
		}

		// Lưu vào cache (TTL: 5 phút)
		jsonData, _ := json.Marshal(response)
		redisClient.Set(context.Background(), cacheKey, jsonData, 5*time.Minute)

		c.JSON(200, response)
	}
}
