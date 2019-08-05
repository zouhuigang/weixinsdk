package storage

func ExporterMap() (m map[string]interface{}) {
	m = map[string]interface{}{
		"memcache": new(Zmemcache),
		"local":    new(Zmemcache),
		"redis":    new(Zredis),
	}
	return
}
