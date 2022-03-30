package user

import (
	"reflect"
	"supermarine1377/domain"
	"supermarine1377/interface/db"
	"testing"
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
			var repo = UserRepository{SqlHandler: MockSqlHandler{}}
			if err := repo.Store(tt.args.u); (err != nil) != tt.wantErr {
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
			repo := UserRepository{
				SqlHandler: tt.fields.SqlHandler,
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

type MockSqlHandler struct {
	queryUserResult []domain.User
}

type MockResult struct{}

type MockRow struct{}

func (msh MockSqlHandler) Excute(string, ...interface{}) (db.Result, error) {
	return MockResult{}, nil
}

func (msh MockSqlHandler) Query(s string, args ...interface{}) (db.Row, error) {
	return MockRow{}, nil
}

func (mr MockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (mr MockResult) RowsAffected() (int64, error) {
	return 0, nil
}

func (mr MockRow) Scan(dest ...interface{}) error {
	return nil
}

func (mr MockRow) Next() bool {
	return true
}

func (mr MockRow) Close() error {
	return nil
}
