package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2" // 注册 Redis 适配器
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

func main() {
	ctx := context.Background()

	fmt.Println("===== Redis 连接测试 =====")

	// 方式1: 使用 gcache 缓存组件（会自动读取配置文件中的 redis 配置）
	fmt.Println("\n[测试1] gcache 缓存组件测试")
	testCache(ctx)

	// 方式2: 直接使用 Redis 客户端
	fmt.Println("\n[测试2] Redis 客户端测试")
	testRedis(ctx)
}

// testCache 测试 gcache 缓存组件
func testCache(ctx context.Context) {
	// 设置缓存，10秒过期
	err := gcache.Set(ctx, "test_key", "hello redis", 10)
	if err != nil {
		fmt.Printf("  ❌ 设置缓存失败: %v\n", err)
		return
	}
	fmt.Println("  ✓ 设置缓存成功: test_key = hello redis")

	// 获取缓存
	val, err := gcache.Get(ctx, "test_key")
	if err != nil {
		fmt.Printf("  ❌ 获取缓存失败: %v\n", err)
		return
	}
	fmt.Printf("  ✓ 获取缓存成功: %v\n", val.String())

	// 删除缓存
	_, err = gcache.Remove(ctx, "test_key")
	if err != nil {
		fmt.Printf("  ❌ 删除缓存失败: %v\n", err)
		return
	}
	fmt.Println("  ✓ 删除缓存成功")
}

// testRedis 测试 Redis 客户端
func testRedis(ctx context.Context) {
	// 获取 Redis 单例
	redis := g.Redis()

	// 执行 SET 命令
	_, err := redis.Do(ctx, "SET", "redis_test_key", "test_value")
	if err != nil {
		fmt.Printf("  ❌ Redis SET 失败: %v\n", err)
		return
	}
	fmt.Println("  ✓ Redis SET 成功: redis_test_key = test_value")

	// 执行 GET 命令
	result, err := redis.Do(ctx, "GET", "redis_test_key")
	if err != nil {
		fmt.Printf("  ❌ Redis GET 失败: %v\n", err)
		return
	}
	fmt.Printf("  ✓ Redis GET 成功: %v\n", gconv.String(result))

	// 执行 DEL 命令
	_, err = redis.Do(ctx, "DEL", "redis_test_key")
	if err != nil {
		fmt.Printf("  ❌ Redis DEL 失败: %v\n", err)
		return
	}
	fmt.Println("  ✓ Redis DEL 成功")

	// 测试连接信息
	fmt.Println("\n[测试3] Redis 连接信息")
	info, err := redis.Do(ctx, "PING")
	if err != nil {
		fmt.Printf("  ❌ Redis PING 失败: %v\n", err)
		return
	}
	fmt.Printf("  ✓ Redis PING 成功: %v\n", gconv.String(info))

	fmt.Println("\n===== 所有测试通过 =====")
}
