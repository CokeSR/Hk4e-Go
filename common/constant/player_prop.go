package constant

const (
	PLAYER_PROP_EXP                             = 1001  // 角色经验
	PLAYER_PROP_BREAK_LEVEL                     = 1002  // 角色突破等阶
	PLAYER_PROP_SATIATION_VAL                   = 1003  // 角色饱食度
	PLAYER_PROP_SATIATION_PENALTY_TIME          = 1004  // 角色饱食度溢出
	PLAYER_PROP_LEVEL                           = 4001  // 角色等级
	PLAYER_PROP_LAST_CHANGE_AVATAR_TIME         = 10001 // 上一次改变角色的时间 暂不确定
	PLAYER_PROP_MAX_SPRING_VOLUME               = 10002 // 七天神像最大恢复血量 0-8500000
	PLAYER_PROP_CUR_SPRING_VOLUME               = 10003 // 七天神像当前血量 0-PROP_MAX_SPRING_VOLUME
	PLAYER_PROP_IS_SPRING_AUTO_USE              = 10004 // 是否开启靠近自动回血 0 1
	PLAYER_PROP_SPRING_AUTO_USE_PERCENT         = 10005 // 自动回血百分比 0-100
	PLAYER_PROP_IS_FLYABLE                      = 10006 // 禁止使用风之翼 0 1
	PLAYER_PROP_IS_WEATHER_LOCKED               = 10007 // 游戏内天气锁定
	PLAYER_PROP_IS_GAME_TIME_LOCKED             = 10008 // 游戏内时间锁定
	PLAYER_PROP_IS_TRANSFERABLE                 = 10009 // 是否禁止传送 0 1
	PLAYER_PROP_MAX_STAMINA                     = 10010 // 最大体力 0-24000
	PLAYER_PROP_CUR_PERSIST_STAMINA             = 10011 // 当前体力 0-PROP_MAX_STAMINA
	PLAYER_PROP_CUR_TEMPORARY_STAMINA           = 10012 // 当前临时体力 暂不确定
	PLAYER_PROP_PLAYER_LEVEL                    = 10013 // 冒险等级
	PLAYER_PROP_PLAYER_EXP                      = 10014 // 冒险经验
	PLAYER_PROP_PLAYER_HCOIN                    = 10015 // 原石 可以为负数
	PLAYER_PROP_PLAYER_SCOIN                    = 10016 // 摩拉
	PLAYER_PROP_PLAYER_MP_SETTING_TYPE          = 10017 // 多人游戏世界权限 0禁止加入 1直接加入 2需要申请
	PLAYER_PROP_IS_MP_MODE_AVAILABLE            = 10018 // 玩家当前的世界是否可加入 0 1 例如任务中就不可加入
	PLAYER_PROP_PLAYER_WORLD_LEVEL              = 10019 // 世界等级 0-8
	PLAYER_PROP_PLAYER_RESIN                    = 10020 // 树脂 0-2000
	PLAYER_PROP_PLAYER_WAIT_SUB_HCOIN           = 10022 // 暂存的原石 暂不确定
	PLAYER_PROP_PLAYER_WAIT_SUB_SCOIN           = 10023 // 暂存的摩拉 暂不确定
	PLAYER_PROP_IS_ONLY_MP_WITH_PS_PLAYER       = 10024 // 当前玩家多人世界里是否有PS主机玩家 0 1
	PLAYER_PROP_PLAYER_MCOIN                    = 10025 // 创世结晶 可以为负数
	PLAYER_PROP_PLAYER_WAIT_SUB_MCOIN           = 10026 // 暂存的创世结晶 暂不确定
	PLAYER_PROP_PLAYER_LEGENDARY_KEY            = 10027 // 传说任务钥匙
	PLAYER_PROP_IS_HAS_FIRST_SHARE              = 10028 // 是否拥有抽卡结果首次分享奖励 暂不确定
	PLAYER_PROP_PLAYER_FORGE_POINT              = 10029 // 锻造相关
	PLAYER_PROP_CUR_CLIMATE_METER               = 10035 // 天气相关
	PLAYER_PROP_CUR_CLIMATE_TYPE                = 10036 // 天气相关
	PLAYER_PROP_CUR_CLIMATE_AREA_ID             = 10037 // 天气相关
	PLAYER_PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE   = 10038 // 天气相关
	PLAYER_PROP_PLAYER_WORLD_LEVEL_LIMIT        = 10039 // 降低世界等级到此等级 暂不确定
	PLAYER_PROP_PLAYER_WORLD_LEVEL_ADJUST_CD    = 10040 // 降低世界等级的CD
	PLAYER_PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM = 10041 // 传说每日任务数量 暂不确定
	PLAYER_PROP_PLAYER_HOME_COIN                = 10042 // 洞天宝钱
	PLAYER_PROP_PLAYER_WAIT_SUB_HOME_COIN       = 10043 // 暂存的洞天宝钱 暂不确定
)
