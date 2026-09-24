// swift-tools-version:5.9
//
// Auto-generated for the v2026.09.24 release by
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
            url: "https://github.com/NodePassProject/EverywhereCore/releases/download/v2026.09.24/EverywhereCore-v2026.09.24.xcframework.zip",
            checksum: "a93064bcae1deaf55b1113a8e4088b4503d278242a51bcc965e68c674ea1ea51"
        ),
    ]
)
