package com.example.demo;

import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@Slf4j
@RestController
public class DemoController {

    @PostMapping("/simple/call")
    public Mono<User> save(@RequestBody User user) {
        log.info("Received user to save: {}", user);

        return Mono.just(user);
    }
}
