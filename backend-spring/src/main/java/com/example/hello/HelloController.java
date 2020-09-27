package com.example.hello;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;


@Component
@Path("/")
public class HelloController {
    private final RedisClient redisClient;
    private final String greetingLabel;

    public HelloController(@Autowired final RedisClient redisClient, @Value("${greetingLabel}") final String greetingLabel) {
        this.redisClient = redisClient;
        this.greetingLabel = greetingLabel;
    }

    @GET
    @Path("{name:.*}")
    @Produces(MediaType.TEXT_PLAIN)
    public String hello(@PathParam("name") String name) {
        final long count = redisClient.incr("backend-spring-counter");
        return "[SPRING] " + greetingLabel + " " + name + " (call #" + count + ")";
    }
}
