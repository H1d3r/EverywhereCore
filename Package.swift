// swift-tools-version:5.9
//
// Auto-generated for the v2026.06.10 release by
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
            url: "https://github.com/NodePassProject/EverywhereCore/releases/download/v2026.06.10/EverywhereCore-v2026.06.10.xcframework.zip",
            checksum: "8027d3afeb57630d6b3becf99045e2eb9963cc51f6f0b2920625e4affa182c78"
        ),
    ]
)
