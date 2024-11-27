import { Options } from 'tsup';

const defaultOptions: Options = {
  entry: ['src/**/*.ts'],
  format: ['cjs', 'esm'],
  splitting: false,
  sourcemap: true,
  // for the type maps to work, we use tsc's declaration-only command
  dts: false,
  clean: true,
  target: 'node16',
  bundle: false,
  esbuildOptions: (options, context) => {
    if (context.format === 'esm') {
      options.packages = 'external';
    }
  },
};
export default defaultOptions;
