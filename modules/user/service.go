package user

import (
	"gofiber-boilerplate/modules/app/appmodel"
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/user/userdto"
	"gofiber-boilerplate/modules/user/usermodel"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Init(db db.DbService)
	Insert(user *usermodel.UserModel) (*userdto.UserDTO, error)
	Update(idString string, updateDTO *userdto.UpdateUserDTO) (*userdto.UserDTO, error)
	List(req *appmodel.GetListRequest) (*appmodel.PaginationResponseList, error)
	Detail(idString string) (*userdto.UserDTO, error)
	Delete(idString string) error
}

type userServiceImpl struct {
	db *gorm.DB
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

// impl `UserService` start

func (service *userServiceImpl) Init(db db.DbService) {
	service.db = db.Default()
}

func (service *userServiceImpl) Insert(user *usermodel.UserModel) (*userdto.UserDTO, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	pwdString := string(pwd)
	user.Password = &pwdString
	result := service.db.Create(user)
	dto := userdto.MapUserModelToDTO(user)
	return dto, result.Error
}

func (service *userServiceImpl) Update(idString string, updateDTO *userdto.UpdateUserDTO) (*userdto.UserDTO, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	if updateDTO.Password != nil {
		pwd, err := bcrypt.GenerateFromPassword([]byte(*updateDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		pwdString := string(pwd)
		updateDTO.Password = &pwdString
	}
	user := userdto.UserDTO{ID: id}
	result := service.db.Model(&user).Updates(updateDTO)
	if result.Error != nil {
		return nil, result.Error
	}
	return service.Detail(idString)
}

func (service *userServiceImpl) Detail(idString string) (*userdto.UserDTO, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	var user userdto.UserDTO
	result := service.db.First(&user, id)
	return &user, result.Error
}

func (service *userServiceImpl) Delete(idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	var user userdto.UserDTO
	result := service.db.Delete(&user, id)
	return result.Error
}

func (service *userServiceImpl) List(req *appmodel.GetListRequest) (*appmodel.PaginationResponseList, error) {
	var count int64
	users := []usermodel.ReadonlyUserModel{}
	query := service.db.Model(users)
	if req.Search != "" {
		query.Where("name ILIKE ?", "%"+req.Search+"%")
	}

	query = query.Session(&gorm.Session{})

	var wg sync.WaitGroup
	wg.Add(2)

	// Perform count and find concurrently using goroutines
	errChan := make(chan error, 2)
	go func() {
		defer wg.Done()
		errChan <- query.Count(&count).Error
	}()

	go func() {
		defer wg.Done()
		query = query.Session(&gorm.Session{})
		errChan <- query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&users).Error
	}()

	wg.Wait()

	var err error
	for i := 0; i < 2; i++ {
		select {
		case err = <-errChan:
			if err != nil {
				return nil, err
			}
		default:
		}
	}

	count32 := int(count)

	return &appmodel.PaginationResponseList{
		Pagination: &appmodel.PaginationResponsePagination{
			Page:  &req.Page,
			Size:  &req.Limit,
			Total: &count32,
		},
		Content: users,
	}, nil
}

// impl `UserService` end
