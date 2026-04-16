package env

func GetString(key string fallback) string {
  val, err := os.LoadEnv(key)
}
