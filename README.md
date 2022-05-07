[![Go](https://github.com/supermarine1377/gopherbank/actions/workflows/go.yml/badge.svg)](https://github.com/supermarine1377/gopherbank/actions/workflows/go.yml)

# エンドポイント

- GET /users
- GET /users/:id
  - 存在しないIDのユーザーのデータは返却できない
- POST /users
  - 名前を入力しないと登録できない
  - 負の預金額は登録できない
- PUT /users/:id
- DELETE /users/:id
  - 存在しないIDのユーザーは削除できない
- POST /transactions