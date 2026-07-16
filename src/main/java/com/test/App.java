package com.test;

import com.google.common.collect.ImmutableList;
import org.apache.commons.lang3.StringUtils;

/**
 * Minimal Java app for testing javawalker packman dependency resolution.
 */
public class App {
    public static void main(String[] args) {
        ImmutableList<String> items = ImmutableList.of("hello", "world");
        String joined = StringUtils.join(items, ", ");
        System.out.println(joined);
    }
}
