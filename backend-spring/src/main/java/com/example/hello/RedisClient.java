package com.example.hello;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;


@Component
public class RedisClient {
    private final RedisTemplate<String, Object> redisTemplate;

    public RedisClient(@Autowired final RedisTemplate<String, Object> redisTemplate) {
        this.redisTemplate = redisTemplate;
    }

    public long incr(final String key) {
        final Long r = redisTemplate.opsForValue().increment(key);
        if (r == null) { // Should not happen unless Redis is unavailable, because "increment" creates a key if not existing
            throw new RuntimeException("Unable to increment Redis key '" + key + "'");
        }
        return r;
    }
}
