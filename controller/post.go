package controller

// import (
// 	"context"
// 	"fmt"
// 	models "lumi/model"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8"
// 	"gorm.io/gorm"
// )

// func handlePostData(db *gorm.DB, redisClient *redis.Client) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var updates []models.F3
// 		if err := c.ShouldBindJSON(&updates); err != nil {
// 			c.JSON(400, gin.H{"error": "JSON không hợp lệ: " + err.Error()})
// 			return
// 		}

// 		// Parse ngày MM/dd/yyyy trong JSON
// 		dateFields := []string{
// 			"ngayLenDon", "ngayHenDayDon", "ngayUpBill", "ngayChuyenKhoan", "ngayDoiSoatCuoc",
// 			"ngayDoiSoatBill", "ngayKeToanDoiSoatFFM1", "ngayKeToanDoiSoatFFM2",
// 			"ngayKeToanDoiSoatFFM3", "thoiGianLenDon", "thoiGianGiaoDuKien",
// 			"ngayDoiSoat", "thoiGianCutoff", "updateAt",
// 		}

// 		for i, update := range updates {
// 			for _, field := range dateFields {
// 				// Giả sử JSON gửi ngày dạng "MM/dd/yyyy" qua string
// 				val, ok := c.GetPostForm(field)
// 				if ok && val != "" {
// 					parsed, err := time.Parse("01/02/2006", val)
// 					if err == nil {
// 						switch field {
// 						case "ngayLenDon":
// 							updates[i].NgayLenDon = parsed
// 						case "ngayHenDayDon":
// 							updates[i].NgayHenDayDon = parsed
// 						case "ngayUpBill":
// 							updates[i].NgayUpBill = parsed
// 						case "ngayChuyenKhoan":
// 							updates[i].NgayChuyenKhoan = parsed
// 						case "ngayDoiSoatCuoc":
// 							updates[i].NgayDoiSoatCuoc = parsed
// 						case "ngayDoiSoatBill":
// 							updates[i].NgayDoiSoatBill = parsed
// 						case "ngayKeToanDoiSoatFFM1":
// 							updates[i].NgayKeToanDoiSoatFFM1 = parsed
// 						case "ngayKeToanDoiSoatFFM2":
// 							updates[i].NgayKeToanDoiSoatFFM2 = parsed
// 						case "ngayKeToanDoiSoatFFM3":
// 							updates[i].NgayKeToanDoiSoatFFM3 = parsed
// 						case "thoiGianLenDon":
// 							updates[i].ThoiGianLenDon = parsed
// 						case "thoiGianGiaoDuKien":
// 							updates[i].ThoiGianGiaoDuKien = parsed
// 						case "ngayDoiSoat":
// 							updates[i].NgayDoiSoat = parsed
// 						case "thoiGianCutoff":
// 							updates[i].ThoiGianCutoff = parsed
// 						case "updateAt":
// 							updates[i].UpdateAt = parsed
// 						}
// 					}
// 				}
// 			}
// 		}

// 		// Batch update
// 		tx := db.Begin()
// 		for _, update := range updates {
// 			result := tx.Where("maDonHang = ?", update.MaDonHang).Updates(&update)
// 			if result.Error != nil {
// 				tx.Rollback()
// 				c.JSON(500, gin.H{"error": result.Error.Error()})
// 				return
// 			}
// 			if result.RowsAffected == 0 {
// 				// Insert nếu không tìm thấy
// 				result = tx.Create(&update)
// 				if result.Error != nil {
// 					tx.Rollback()
// 					c.JSON(500, gin.H{"error": result.Error.Error()})
// 					return
// 				}
// 			}
// 		}
// 		tx.Commit()

// 		// Xóa cache sau update
// 		redisClient.Del(context.Background(), "f3_data")

// 		c.JSON(200, gin.H{
// 			"message": "Cập nhật thành công " + fmt.Sprint(len(updates)) + " dòng",
// 		})
// 	}
// }
