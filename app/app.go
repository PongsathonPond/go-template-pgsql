package app

import (
	"context"
	"fmt"
	"idev-cms-service/app/authen"
	"idev-cms-service/app/category"
	"idev-cms-service/app/content"
	"idev-cms-service/app/menus"
	"idev-cms-service/app/roles"
	"idev-cms-service/app/tokens"
	"idev-cms-service/app/users"
	"idev-cms-service/config"
	"idev-cms-service/middleware"
	"idev-cms-service/repository/mongodb"
	"idev-cms-service/repository/pgsql"
	"idev-cms-service/repository/redis"
	serviceAuthen "idev-cms-service/service/authen/implement"
	serviceCategory "idev-cms-service/service/category/implement"
	serviceContent "idev-cms-service/service/content/implement"
	serviceCron "idev-cms-service/service/cron/implement"
	serviceMenus "idev-cms-service/service/menus/implement"
	serviceRoles "idev-cms-service/service/roles/implement"
	serviceTokens "idev-cms-service/service/tokens/implement"
	serviceUsers "idev-cms-service/service/users/implement"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	serviceValidator "idev-cms-service/service/validator"
)

type App struct {
	mid      middleware.Service
	authen   *authen.Controller
	tokens   *tokens.Controller
	users    *users.Controller
	content  *content.Controller
	category *category.Controller
	roles    *roles.Controller
	menus    *menus.Controller
}

const (
	collUsers      = "users"
	collFlagUsers  = "flag-users"
	collTokens     = "tokens"
	collUserGroups = "user-groups"
	collMenus      = "menus"
	collRoles      = "roles"
	collContent    = "content"
	collCategory   = "category"
)

func newApp(appConfig *config.Config, log logs.Log) *App {
	ctx := context.Background()
	uuid, err := util.NewUUID()
	panicIfErr(log, err)
	datetime := util.NewDateTime(appConfig)
	filter := util.NewFilterString()
	genCode := util.NewGenCode()
	encrypt := util.NewEncrypt()
	perm := util.NewPermission()
	//filters := util.NewFilters()

	// repositories
	redisRepo, err := redis.New(ctx, appConfig)
	panicIfErr(log, err)
	//kafkaRepo, err := kafka.New(setup.ConfigKafka(appConfig))
	//panicIfErr(log, err)

	usersRepoPgsql, err := pgsql.New(ctx, pgsqlConfig(appConfig, collUsers), collUsers)
	panicIfErr(log, err)

	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
	}

	usersRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collUsers))
	panicIfErr(log, err)
	//flagUsersRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collFlagUsers))
	//panicIfErr(log, err)
	userGroupsRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collUserGroups))
	panicIfErr(log, err)
	tokensRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collTokens))
	panicIfErr(log, err)
	rolesRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collRoles))
	panicIfErr(log, err)
	menusRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collMenus))
	panicIfErr(log, err)
	contentRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collContent))
	panicIfErr(log, err)
	categoryRepo, err := mongodb.New(ctx, mongoDBConfig(appConfig, collCategory))
	panicIfErr(log, err)

	// validators
	validator := serviceValidator.New(
		&serviceValidator.ValidatorRepository{
			User: usersRepo,
		},
	)

	// services
	authenService := serviceAuthen.New(&serviceAuthen.AuthenServiceConfig{
		Validator:      validator,
		RepoUsers:      usersRepo,
		RepoUserGroups: userGroupsRepo,
		RepoTokens:     tokensRepo,
		RepoRedis:      redisRepo,
		UUID:           uuid,
		Config:         appConfig,
		DateTime:       datetime,
		FilterString:   filter,
		Encrypt:        encrypt,
		Log:            log,
	})
	tokensService := serviceTokens.New(&serviceTokens.TokensServiceConfig{
		Validator:    validator,
		RepoUsers:    usersRepo,
		RepoTokens:   tokensRepo,
		RepoRedis:    redisRepo,
		Config:       appConfig,
		DateTime:     datetime,
		FilterString: filter,
		Encrypt:      encrypt,
		Log:          log,
	})
	usersService := serviceUsers.New(&serviceUsers.UserServiceConfig{
		Validator:     validator,
		RepoUsers:     usersRepoPgsql,
		RepoTokens:    tokensRepo,
		RepoRoles:     rolesRepo,
		Config:        appConfig,
		ServiceTokens: tokensService,
		UUID:          uuid,
		GenCode:       genCode,
		DateTime:      datetime,
		FilterString:  filter,
		Encrypt:       encrypt,
		Log:           log,
	})

	menuService := serviceMenus.New(&serviceMenus.MenusServiceConfig{
		Validator:    validator,
		Repo:         menusRepo,
		RepoRoles:    rolesRepo,
		UUID:         uuid,
		DateTime:     datetime,
		Perm:         perm,
		FilterString: filter,
		Log:          log,
	})
	roleService := serviceRoles.New(&serviceRoles.RolesServiceConfig{
		Validator:    validator,
		Repo:         rolesRepo,
		MenusRepo:    menusRepo,
		UserService:  usersService,
		ServiceMenus: menuService,
		UUID:         uuid,
		DateTime:     datetime,
		Perm:         perm,
		FilterString: filter,
		Log:          log,
		Config:       appConfig,
	})

	contentService := serviceContent.New(&serviceContent.ContentServiceConfig{
		Validator:     validator,
		RepoContent:   contentRepo,
		RepoUsers:     usersRepo,
		RepoCategory:  categoryRepo,
		RepoRedis:     nil,
		UUID:          uuid,
		DateTime:      datetime,
		FilterString:  filter,
		Log:           log,
		ServiceTokens: tokensService,
	})

	categoryService := serviceCategory.New(&serviceCategory.CategoryServiceConfig{
		Validator:    validator,
		RepoCategory: categoryRepo,
		RepoRedis:    nil,
		UUID:         uuid,
		DateTime:     datetime,
		FilterString: filter,
		Log:          log,
	})

	// middlewares
	midService := middleware.New(&middleware.MiddlewareServiceConfig{
		AuthenService: authenService,
		UserService:   usersService,
		RoleService:   roleService,
	})

	// run cron
	go func() {
		serviceCron.New(&serviceCron.CronServiceConfig{
			RepoTokens:   tokensRepo,
			FilterString: filter,
			DateTime:     datetime,
			Log:          log,
		}).Run()
	}()

	return &App{
		mid:      midService,
		authen:   authen.New(authenService),
		tokens:   tokens.New(tokensService),
		users:    users.New(usersService),
		roles:    roles.New(roleService),
		menus:    menus.New(menuService),
		content:  content.New(contentService),
		category: category.New(categoryService),
	}
}

func mongoDBConfig(appConfig *config.Config, collName string) *mongodb.Config {
	return &mongodb.Config{
		URI:      appConfig.MongoDBEndpoint,
		DBName:   appConfig.MongoDBName,
		CollName: collName,
		Timeout:  appConfig.MongoDBTimeout,
	}
}

func pgsqlConfig(appConfig *config.Config, collName string) *pgsql.Config {
	return &pgsql.Config{
		Host:     appConfig.Host,
		Port:     appConfig.Port,
		User:     appConfig.User,
		Password: appConfig.Password,
		DBName:   collName,
		Timeout:  appConfig.Timeout,
	}
}

func panicIfErr(log logs.Log, err error) {
	if err != nil {
		log.Panic(err)
	}
}
