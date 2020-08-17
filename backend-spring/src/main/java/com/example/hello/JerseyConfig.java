package com.example.hello;

import org.glassfish.jersey.server.ResourceConfig;
import org.springframework.context.annotation.Configuration;


@Configuration
public class JerseyConfig extends ResourceConfig {
    public JerseyConfig() {
        register(HelloController.class);
    }
}
