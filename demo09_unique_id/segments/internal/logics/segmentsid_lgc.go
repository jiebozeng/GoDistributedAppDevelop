package logics

import (
	"GoDistributedAppDevelop/demo09_unique_id/segments/internal/models"
	"GoDistributedAppDevelop/demo09_unique_id/segments/pkg/mysqldb"
	"fmt"
	"github.com/jiebozeng/golangutils/timeutils"
	"gorm.io/gorm"
)

type Segmentsid_lgc struct {
}

// 逻辑实现
func (s *Segmentsid_lgc) GetSegmentsIds(bizType int64) (minId int64, maxId int64, err error) {
	seg := &models.SegmentsId{}
	nowTime := timeutils.GetNowTime()
	//查询
	query := mysqldb.Mysql.Model(seg).
		Where("biz_type = ?", bizType).Limit(1)
	if err = query.Find(seg).Error; err != nil {
		return 0, 0, err
	}
	minId = seg.MaxId
	maxId = seg.MaxId + seg.Step - 1
	// 执行更新操作
	result := mysqldb.Mysql.Model(seg).
		Where("version = ? AND biz_type = ?", seg.Version, bizType).
		Updates(map[string]interface{}{
			"max_id":     gorm.Expr("max_id + step"), //当前max_id+step(步长)
			"version":    gorm.Expr("version + 1"),
			"updated_at": nowTime, // 更新更新时间
		})
	if result.Error != nil {
		return 0, 0, result.Error
	}
	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return 0, 0, fmt.Errorf("获取id段失败, 业务类型或版本不对")
	}
	//返回结果
	return minId, maxId, nil
}
