// swift-tools-version:5.9
//
// Auto-generated for the v2026.09.10 release by
// .github/workflows/upstream-watch.yml. The `main` branch
// keeps a local `binaryTarget(path:)` variant for in-tree
// development; this variant lives only on the tag.

import PackageDescription

let package = Package(
    name: "EverywhereCore",
    platforms: [
        .iOS(.v15),
        .macOS(.v13),
    ],
    products: [
        .library(name: "EverywhereCore", targets: ["EverywhereCore"]),
    ],
    targets: [
        .binaryTarget(
            name: "EverywhereCore",
            url: "https://github.com/NodePassProject/EverywhereCore/releases/download/v2026.09.10/EverywhereCore-v2026.09.10.xcframework.zip",
            checksum: "efd49ec1f90640b852488a60f7f66256295f2850899e15283ab15cfc05a0a7e8"
        ),
    ]
)
