package user

import (
	"reflect"
	"supermarine1377/domain"
	"supermarine1377/interface/db"
	mock_db "supermarine1377/interface/db/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserRepository_Store(t *testing.T) {
	type fields struct {
		SqlHandler db.SqlHandler
	}
	type args struct {
		u domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1st",
			args: args{
				u: domain.User{
					Name:    "test",
					Balance: 100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			msh := mock_db.NewMockSqlHandler(ctrl)
			var user = domain.User{Name: "test", Balance: 1000}
			var mr = mock_db.NewMockResult(ctrl)
			msh.EXPECT().Excute(gomock.Any()).Return(mr, nil)
			var repo = UserRepository{SqlHandler: msh}
			if err := repo.Store(user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	type fields struct {
		SqlHandler db.SqlHandler
	}
	type args struct {
		u domain.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := UserRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			repo.Update(tt.args.u)
		})
	}
}

func TestUserRepository_FindById(t *testing.T) {
	type fields struct {
		SqlHandler db.SqlHandler
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := UserRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			repo.FindById(tt.args.id)
		})
	}
}

func TestUserRepository_FindAll(t *testing.T) {
	type fields struct {
		SqlHandler db.SqlHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.User
		wantErr bool
	}{
		{
			name:    "1st",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			msh := mock_db.NewMockSqlHandler(ctrl)
			mr := mock_db.NewMockRow(ctrl)
			msh.EXPECT().Query(gomock.Any()).Return(mr, nil)
			mr.EXPECT().Next().Return(false)
			repo := UserRepository{
				SqlHandler: msh,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {
	type fields struct {
		SqlHandler db.SqlHandler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := UserRepository{
				SqlHandler: tt.fields.SqlHandler,
			}
			repo.Delete()
		})
	}
}
