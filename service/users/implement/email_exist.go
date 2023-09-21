package implement

import (
	"context"
	"errors"
)

func (impl *implementation) CheckEmailExist(ctx context.Context, email *string) (err error) {
	var count int
	filters := []string{
		impl.FilterString.MakeEmail(*email),
		impl.FilterString.MakeStatus("active"),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	count, err = impl.RepoUsers.Count(ctx, filters)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Email is not exist")
	}

	return nil
}
