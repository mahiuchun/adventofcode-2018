load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "day09",
    srcs = ["day09.go"],
    data = [":input.txt"],
    deps = [":game"],
)

go_library(
    name = "game",
    srcs = ["game.go"],
    importpath = "game",
)

go_test(
    name = "game_test",
    srcs = ["game_test.go"],
    embed = [":game"],
)

go_binary(
    name = "part2",
    srcs = ["part2.go"],
    data = [":input.txt"],
    deps = [":game"],
)
