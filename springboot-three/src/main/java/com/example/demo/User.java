package com.example.demo;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.ToString;

@ToString
public class User {

    @JsonProperty(value = "name")
    private String name;

    @JsonProperty(value = "email")
    private String email;
}
