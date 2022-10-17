package dal

// StudentInfo 学生信息表
type StudentInfo struct {
	Id           int64  `gorm:"primaryKey" json:"-"`
	StudentId    int64  `gorm:"not null; unique; index" json:"id"`
	Account      string `gorm:"not null" json:"account"`
	Name         string `gorm:"not null" json:"name"`
	Gender       string `gorm:"not null; index" json:"gender"`
	Date         string `gorm:"not null" json:"date"`
	Organization string `gorm:"not null; index" json:"organization"`
	Role         string `gorm:"not null" json:"role"`
	State        string `gorm:"not null; index" json:"state"`
}

// FindById 根据 id 查找学生信息
func (s *StudentInfo) FindById(id int64) (StudentInfo, error) {
	var studentInfo StudentInfo
	return studentInfo, DB.First(&studentInfo, id).Error
}

// FindByStudentId 根据 student_id 查找学生信息
func (s *StudentInfo) FindByStudentId(studentId int64) (StudentInfo, error) {
	var studentInfo StudentInfo
	return studentInfo, DB.Where("student_id = ?", studentId).First(&studentInfo).Error
}

// FindAll 获取全部学生信息
func (s *StudentInfo) FindAll() ([]StudentInfo, error) {
	var studentInfos []StudentInfo
	return studentInfos, DB.Find(&studentInfos).Error
}

// FindAllByParams 根据性别、机构、状态筛选用户
func (s *StudentInfo) FindAllByParams(gender, organization, state string) ([]StudentInfo, error) {
	var studentInfos []StudentInfo
	clause := ""
	if gender != "" {
		clause += "gender = '" + gender + "'"
	}
	if organization != "" {
		if clause != "" {
			clause += " AND "
		}
		clause += "organization = '" + organization + "'"
	}
	if state != "" {
		if clause != "" {
			clause += " AND "
		}
		clause += "state = '" + state + "'"
	}
	if clause == "" {
		clause = "true"
	}
	queryStr := "SELECT * FROM student_infos WHERE(" + clause + ");"
	// 有 SQL 注入风险
	return studentInfos, DB.Raw(queryStr).Scan(&studentInfos).Error
}

// Create 添加单个学生信息
func (s *StudentInfo) Create(studentInfo *StudentInfo) (err error) {
	return DB.Create(studentInfo).Error
}

// Update 查找学生信息并对其进行更新
func (s *StudentInfo) Update(studentInfo *StudentInfo) error {
	return DB.Save(studentInfo).Error // 将需要更新的学生信息保存
}

// UpdateNewFields 查找学生信息并根据用户更新信息的每个字段是否是空字符串进行判断是否需要覆盖原有信息并保存到数据库
// TODO 有 Bug，无法唯一确定学生
func (s *StudentInfo) UpdateNewFields(stuUpdate *StudentInfo) error {
	var stuBefore StudentInfo
	if err := DB.Where("student_id = ?", stuUpdate.StudentId).First(&stuBefore).Error; err != nil {
		return err
	}
	if stuUpdate.Account != "" {
		stuBefore.Account = stuUpdate.Account
	}
	if stuUpdate.Name != "" {
		stuBefore.Name = stuUpdate.Name
	}
	if stuUpdate.Gender != "" {
		stuBefore.Gender = stuUpdate.Gender
	}
	if stuUpdate.Date != "" {
		stuBefore.Date = stuUpdate.Date
	}
	if stuUpdate.Organization != "" {
		stuBefore.Organization = stuUpdate.Organization
	}
	if stuUpdate.Role != "" {
		stuBefore.Role = stuUpdate.Role
	}
	if stuUpdate.State != "" {
		stuBefore.State = stuUpdate.State
	}
	return DB.Save(&stuBefore).Error
}

// Delete 删除学生信息
func (s *StudentInfo) Delete(studentInfo *StudentInfo) error {
	return DB.Delete(studentInfo).Error
}
