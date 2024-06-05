import nextPWA from 'next-pwa';

const withPWA = nextPWA({
  dest: 'public',
});

/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['@repo/design-system'],
  output: 'standalone',
};

export default withPWA(nextConfig);
