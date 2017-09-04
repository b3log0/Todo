package main

func setCurrentDomain(domain string) {
	client.Set(REDIS_KEY + "." + "current", domain, 0)
	current_domain = domain
}

func getCurrentDomain() string {
	result, err := client.Get(REDIS_KEY + "." + "current").Result()
	if err == nil {
		return result
	} else {
		return getDomain(0)
	}
}

func getDomain(index int64) string {
	result, _ := client.LIndex(REDIS_KEY, index).Result()
	return result
}

func insertDomain(domain string) bool {
	result, _ := client.RPush(REDIS_KEY, domain).Result()
	if result > 0 {
		return true
	} else {
		return false
	}
}

func getDomains() []string {
	result, _ := client.LRange(REDIS_KEY, 0, -1).Result()
	return result
}

func delDomain(domain string) {
	result, _ := client.LRem(REDIS_KEY, 0, domain).Result()
	if result > 0 {
		for _, task := range getAllTasks(domain) {
			delTask(domain, task.key)
		}
	}
}

func setTask(domain string, index string, task string) {
	printError("setTask: " + domain + " " + index + " " + task)
	client.HSet(REDIS_KEY + "." + domain, index, task).Result()
}

func getAllTasks(domain string) []Task {
	result, _ := client.HGetAll(REDIS_KEY + "." + domain).Result()
	return buildTaskList(result)
}

func delTask(domain string, field string) {
	client.HDel(REDIS_KEY + "." + domain, field).Result()
}
