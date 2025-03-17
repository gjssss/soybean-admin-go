package cache

import (
	"context"
	"fmt"

	"github.com/gjssss/soybean-admin-go/global"

	"slices"

	"github.com/gjssss/soybean-admin-go/repositories"
)

var ctx = context.Background()
var apiKeys = make([]string, 0)

type apiCache struct{}

var ApiCache = &apiCache{}

func makeApiKey(path string, method string) string {
	return "api:" + path + ":" + method
}

func (a *apiCache) Refresh() {
	if len(apiKeys) > 0 {
		global.Redis.Del(ctx, apiKeys...)
		apiKeys = make([]string, 0)
	}

	data, _ := repositories.System.Api.GetAllApisRoles()
	for _, v := range data {
		key := makeApiKey(v.Path, v.Method)
		role_ids := make([]any, 0)
		for _, role := range v.Roles {
			role_ids = append(role_ids, fmt.Sprintf("%d", role.ID))
		}
		global.Redis.LPush(ctx, key, role_ids...)
		apiKeys = append(apiKeys, key)
	}
}

func (a *apiCache) Has(path string, method string, role_id uint) bool {
	key := makeApiKey(path, method)
	role_ids, _ := global.Redis.LRange(ctx, key, 0, -1).Result()
	return slices.Contains(role_ids, fmt.Sprintf("%d", role_id))
}
