package db

import "context"

type (
	CreateUserTxParams struct {
		CreateUserParams
		// calls this func after the db query if that returned nil err
		AfterCreate func(user User) error
	}

	CreateUserTxResult struct {
		User User
	}
)

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.User)
	})

	return result, err
}
