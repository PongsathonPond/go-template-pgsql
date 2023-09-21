package tasks

import (
	"fmt"
)

func (t *Tasks) ClearTokensExpired() {
	t.Log.Info("Task clear tokens expired")

	filters := []string{
		t.FilterString.MakeRefreshTokenExpiredAtLessThan(t.DateTime.GetUnixNow()),
	}

	cnt, err := t.RepoTokens.Count(t.ctx, filters)
	if err != nil {
		t.Log.Error(err)
		return
	}

	if cnt == 0 {
		t.Log.Info("- Nothing tokens expired to clear")
		return
	}

	if err = t.RepoTokens.DeleteMany(t.ctx, filters); err != nil {
		t.Log.Error(err)
		return
	}

	t.Log.Info(fmt.Sprintf("- Clear tokens expire amount %d token(s)", cnt))
}
