// swift-tools-version:5.9
//
// Auto-generated for the v2026.09.14 release by
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
            url: "https://github.com/NodePassProject/EverywhereCore/releases/download/v2026.09.14/EverywhereCore-v2026.09.14.xcframework.zip",
            checksum: "8bbfe648cfc788eb84b76815190b7b33ff3e60449a699b07af66d34e34018d9d"
        ),
    ]
)
