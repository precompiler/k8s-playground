package com.precompiler;

import io.quarkus.runtime.Quarkus;
import io.quarkus.runtime.annotations.QuarkusMain;

/**
 * @author Richard Li
 */
@QuarkusMain
public class Application {
    public static void main(String[] args) {
        Quarkus.run(args);
    }
}
