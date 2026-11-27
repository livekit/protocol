// @ts-check

import { defineBuildConfig } from "unbuild";

export default defineBuildConfig({
  // The main barrel, and the Weaver-generated OTEL semantic conventions as a separate
  // `@livekit/protocol/semconv` entry. semconv.ts has no imports and no runtime cost, so a
  // subpath keeps it out of the protobuf barrel's bundle: importing the constants never pulls
  // in @bufbuild/protobuf. src/gen/semconv.ts is copied from otel/gen/ts/semconv.ts by the
  // generate:semconv script before this runs.
  entries: [
    "src/index",
    { input: "src/gen/semconv", name: "semconv" },
  ],
  declaration: true,
  sourcemap: true,
  rollup: {
    emitCJS: true,
  },
});
