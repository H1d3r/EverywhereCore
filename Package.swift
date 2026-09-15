// swift-tools-version:5.9
//
// Auto-generated for the v2026.09.15 release by
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
            url: "https://github.com/NodePassProject/EverywhereCore/releases/download/v2026.09.15/EverywhereCore-v2026.09.15.xcframework.zip",
            checksum: "2713e990ec46b71ab8d833e7db3ba976d8d446f3f3038a5d5595e1b6043b91f4"
        ),
    ]
)
